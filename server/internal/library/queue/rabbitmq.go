// Package queue

package queue

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/utility/simple"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	rabbitDefaultPrefetch    = 10               // 默认预取数量
	rabbitConfirmTimeout     = 5 * time.Second  // 生产者等待broker确认超时时间
	rabbitReconnectInterval  = 3 * time.Second  // 消费者断线重连间隔
	rabbitDelayQueueExpireMs = int64(60 * 1000) // 延迟队列在最后一条消息过期后的保留时间
)

type RabbitMq struct {
	groupName string
	channel   *amqp.Channel       // 生产者通道，使用确认模式
	declared  map[string]struct{} // 已声明的主题队列
	mutex     sync.Mutex
}

type RabbitManager struct {
	conn   *amqp.Connection
	mutex  sync.Mutex
	wg     sync.WaitGroup // 运行中的消费者
	ctx    context.Context
	cancel context.CancelFunc
}

var rabbitManager = newRabbitManager()

func init() {
	setRabbitCloseEvent()
}

func newRabbitManager() *RabbitManager {
	m := &RabbitManager{}
	m.ctx, m.cancel = context.WithCancel(context.Background())
	return m
}

func setRabbitCloseEvent() {
	simple.Event().Register(consts.EventServerClose, func(ctx context.Context, args ...interface{}) {
		rabbitManager.mutex.Lock()
		rabbitManager.cancel()
		rabbitManager.mutex.Unlock()

		// 等待消费中的消息处理完成
		rabbitManager.wg.Wait()

		rabbitManager.mutex.Lock()
		defer rabbitManager.mutex.Unlock()
		if rabbitManager.conn != nil && !rabbitManager.conn.IsClosed() {
			if err := rabbitManager.conn.Close(); err != nil {
				Logger().Warningf(ctx, "rabbitmq connection close err:%v", err)
				return
			}
			Logger().Debug(ctx, "rabbitmq connection close...")
		}
	})
}

func GetRabbitManager() *RabbitManager {
	return rabbitManager
}

// getConn 获取连接，连接断开时自动重连
func (m *RabbitManager) getConn() (*amqp.Connection, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.ctx.Err() != nil {
		return nil, gerror.New("rabbitmq manager is closed")
	}

	if m.conn != nil && !m.conn.IsClosed() {
		return m.conn, nil
	}

	conn, err := amqp.Dial(config.Rabbitmq.Address)
	if err != nil {
		return nil, gerror.Wrap(err, "连接rabbitmq失败，请检查`queue.rabbitmq.address`配置是否有误")
	}
	m.conn = conn
	return m.conn, nil
}

// openChannel 打开一个新通道并声明交换机
func (m *RabbitManager) openChannel() (ch *amqp.Channel, err error) {
	conn, err := m.getConn()
	if err != nil {
		return
	}

	if ch, err = conn.Channel(); err != nil {
		return
	}

	if err = ch.ExchangeDeclare(rabbitExchange(), amqp.ExchangeDirect, true, false, false, false, nil); err != nil {
		_ = ch.Close()
		return nil, err
	}
	return
}

// SendMsg 按字符串类型生产数据
func (r *RabbitMq) SendMsg(topic string, body string) (mqMsg MqMsg, err error) {
	return r.SendByteMsg(topic, []byte(body))
}

// SendByteMsg 生产数据
func (r *RabbitMq) SendByteMsg(topic string, body []byte) (mqMsg MqMsg, err error) {
	if topic == "" {
		return mqMsg, gerror.New("rabbitmq topic is empty")
	}

	mqMsg = MqMsg{
		RunType:   SendMsg,
		Topic:     topic,
		MsgId:     getRandMsgId(),
		Body:      body,
		Timestamp: time.Now(),
	}

	err = r.withChannel(func(ch *amqp.Channel) error {
		if err := r.declareTopic(ch, topic); err != nil {
			return err
		}
		return r.confirmPublish(ch, rabbitExchange(), topic, mqMsg)
	})
	return
}

// SendDelayMsg 生产延迟消息，delaySecond 为延迟秒数
// 基于消息TTL+死信交换机实现，无需安装延迟插件
func (r *RabbitMq) SendDelayMsg(topic string, body string, delaySecond int64) (mqMsg MqMsg, err error) {
	if delaySecond < 1 {
		return r.SendMsg(topic, body)
	}

	if topic == "" {
		return mqMsg, gerror.New("rabbitmq topic is empty")
	}

	mqMsg = MqMsg{
		RunType:   SendMsg,
		Topic:     topic,
		MsgId:     getRandMsgId(),
		Body:      []byte(body),
		Timestamp: time.Now(),
	}

	err = r.withChannel(func(ch *amqp.Channel) error {
		// 先确保目标队列存在，否则消息到期后无处投递
		if err := r.declareTopic(ch, topic); err != nil {
			return err
		}

		// 同一主题同一延迟时长共用一个延迟队列，避免队头阻塞
		// 每次发送都重新声明以刷新x-expires计时，保证最后一条消息到期前队列不会被删除
		var (
			delayQueue = rabbitDelayQueueName(r.groupName, topic, delaySecond)
			ttl        = delaySecond * 1000
		)
		if _, err := ch.QueueDeclare(delayQueue, true, false, false, false, amqp.Table{
			"x-message-ttl":             ttl,
			"x-dead-letter-exchange":    rabbitExchange(),
			"x-dead-letter-routing-key": topic,
			"x-expires":                 ttl + rabbitDelayQueueExpireMs,
		}); err != nil {
			return err
		}
		return r.confirmPublish(ch, "", delayQueue, mqMsg)
	})
	return
}

// withChannel 使用生产者通道执行操作，通道断开时自动重建并重试一次
func (r *RabbitMq) withChannel(fn func(ch *amqp.Channel) error) (err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := 0; i < 2; i++ {
		if r.channel == nil || r.channel.IsClosed() {
			ch, err := rabbitManager.openChannel()
			if err != nil {
				return err
			}
			if err = ch.Confirm(false); err != nil {
				_ = ch.Close()
				return err
			}
			r.channel = ch
			r.declared = make(map[string]struct{})
		}

		if err = fn(r.channel); err == nil || !r.channel.IsClosed() {
			return
		}
	}
	return
}

// declareTopic 声明主题队列
func (r *RabbitMq) declareTopic(ch *amqp.Channel, topic string) error {
	if _, ok := r.declared[topic]; ok {
		return nil
	}
	if err := declareRabbitTopic(ch, r.groupName, topic); err != nil {
		return err
	}
	r.declared[topic] = struct{}{}
	return nil
}

// confirmPublish 发布消息并等待broker确认
func (r *RabbitMq) confirmPublish(ch *amqp.Channel, exchange, key string, mqMsg MqMsg) error {
	sendCtx, cancel := context.WithTimeout(context.Background(), rabbitConfirmTimeout)
	defer cancel()

	dc, err := ch.PublishWithDeferredConfirmWithContext(sendCtx, exchange, key, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		MessageId:    mqMsg.MsgId,
		Timestamp:    mqMsg.Timestamp,
		Body:         mqMsg.Body,
	})
	if err != nil {
		return err
	}

	ack, err := dc.WaitContext(sendCtx)
	if err != nil {
		return err
	}
	if !ack {
		return gerror.Newf("rabbitmq broker nack msg, topic:%v", mqMsg.Topic)
	}
	return nil
}

// ListenReceiveMsgDo 消费数据
func (r *RabbitMq) ListenReceiveMsgDo(topic string, receiveDo func(mqMsg MqMsg) error) (err error) {
	if topic == "" {
		return gerror.New("rabbitmq topic is empty")
	}

	ch, deliveries, err := r.subscribe(topic)
	if err != nil {
		return
	}

	rabbitManager.mutex.Lock()
	if rabbitManager.ctx.Err() != nil {
		rabbitManager.mutex.Unlock()
		_ = ch.Close()
		return gerror.New("rabbitmq manager is closed")
	}
	rabbitManager.wg.Add(1)
	rabbitManager.mutex.Unlock()

	go func() {
		defer rabbitManager.wg.Done()
		for {
			r.consume(ch, deliveries, topic, receiveDo)

			// 通道关闭后，非退出状态则自动重连
			for {
				select {
				case <-rabbitManager.ctx.Done():
					return
				case <-time.After(rabbitReconnectInterval):
				}

				var subErr error
				if ch, deliveries, subErr = r.subscribe(topic); subErr == nil {
					Logger().Debugf(ctx, "rabbitmq consumer reconnected, topic:%v", topic)
					break
				}
				Logger().Warningf(ctx, "rabbitmq consumer reconnect failed, topic:%v, err:%+v", topic, subErr)
			}
		}
	}()

	Logger().Debugf(ctx, "rabbitmq consumer up and running, topic:%v", topic)
	return
}

// subscribe 订阅主题队列
func (r *RabbitMq) subscribe(topic string) (ch *amqp.Channel, deliveries <-chan amqp.Delivery, err error) {
	if ch, err = rabbitManager.openChannel(); err != nil {
		return
	}

	defer func() {
		if err != nil {
			_ = ch.Close()
		}
	}()

	prefetch := config.Rabbitmq.Prefetch
	if prefetch <= 0 {
		prefetch = rabbitDefaultPrefetch
	}

	if err = ch.Qos(prefetch, 0, false); err != nil {
		return
	}

	if err = declareRabbitTopic(ch, r.groupName, topic); err != nil {
		return
	}

	deliveries, err = ch.ConsumeWithContext(rabbitManager.ctx, rabbitQueueName(r.groupName, topic), "", false, false, false, false, nil)
	return
}

// consume 处理消息直到通道关闭，处理完成后再确认，保证至少消费一次
func (r *RabbitMq) consume(ch *amqp.Channel, deliveries <-chan amqp.Delivery, topic string, receiveDo func(mqMsg MqMsg) error) {
	defer func() {
		_ = ch.Close()
	}()

	for d := range deliveries {
		_ = receiveDo(MqMsg{
			RunType:   ReceiveMsg,
			Topic:     topic,
			MsgId:     d.MessageId,
			Timestamp: d.Timestamp,
			Body:      d.Body,
		})

		if err := d.Ack(false); err != nil {
			Logger().Warningf(ctx, "rabbitmq ack err, topic:%v, msgId:%v, err:%+v", topic, d.MessageId, err)
		}
	}
}

// RegisterRabbitMqProducer 注册rabbitmq生产者
func RegisterRabbitMqProducer(groupName string) (client MqProducer, err error) {
	mqIns := &RabbitMq{groupName: groupName}

	// 检查一下连通性
	if err = mqIns.withChannel(func(ch *amqp.Channel) error { return nil }); err != nil {
		return nil, err
	}
	return mqIns, nil
}

// RegisterRabbitMqConsumer 注册rabbitmq消费者
func RegisterRabbitMqConsumer(groupName string) (client MqConsumer, err error) {
	// 检查一下连通性
	if _, err = rabbitManager.getConn(); err != nil {
		return nil, err
	}
	return &RabbitMq{groupName: groupName}, nil
}

// declareRabbitTopic 声明主题队列并绑定到交换机
func declareRabbitTopic(ch *amqp.Channel, groupName, topic string) error {
	name := rabbitQueueName(groupName, topic)
	if _, err := ch.QueueDeclare(name, true, false, false, false, nil); err != nil {
		return err
	}
	return ch.QueueBind(name, topic, rabbitExchange(), false, nil)
}

// rabbitExchange 交换机名称，未配置时使用groupName
func rabbitExchange() string {
	if config.Rabbitmq.Exchange != "" {
		return config.Rabbitmq.Exchange
	}
	return config.GroupName
}

// rabbitQueueName 主题队列名称
func rabbitQueueName(groupName, topic string) string {
	return fmt.Sprintf("%s.%s", groupName, topic)
}

// rabbitDelayQueueName 延迟队列名称
func rabbitDelayQueueName(groupName, topic string, delaySecond int64) string {
	return fmt.Sprintf("%s.%s.delay.%ds", groupName, topic, delaySecond)
}
