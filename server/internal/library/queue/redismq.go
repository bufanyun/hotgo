package queue

import (
	"encoding/json"
	"fmt"
	"github.com/bufanyun/pool"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gomodule/redigo/redis"
	"hotgo/internal/consts"
	"hotgo/utility/encrypt"
	"math/rand"
	"time"
)

type RedisMq struct {
	poolName  string
	groupName string
	retry     int
	timeout   int
}

type PoolOption struct {
	InitCap     int
	MaxCap      int
	IdleTimeout int
}

type RedisOption struct {
	Addr    string
	Passwd  string
	DBnum   int
	Timeout int
}

var redisPoolMap map[string]pool.Pool

func init() {
	redisPoolMap = make(map[string]pool.Pool)

}

// SendMsg 按字符串类型生产数据
func (r *RedisMq) SendMsg(topic string, body string) (mqMsg MqMsg, err error) {
	return r.SendByteMsg(topic, []byte(body))
}

// SendByteMsg 生产数据
func (r *RedisMq) SendByteMsg(topic string, body []byte) (mqMsg MqMsg, err error) {
	if r.poolName == "" {
		return mqMsg, gerror.New("RedisMq producer not register")
	}
	if topic == "" {
		return mqMsg, gerror.New("RedisMq topic is empty")
	}

	msgId := getRandMsgId()
	rdx, put, err := getRedis(r.poolName, r.retry)
	defer put()

	if err != nil {
		return mqMsg, gerror.New(fmt.Sprint("queue redis 生产者获取redis实例失败:", err))
	}

	mqMsg = MqMsg{
		RunType:   SendMsg,
		Topic:     topic,
		MsgId:     msgId,
		Body:      body,
		Timestamp: time.Now(),
	}
	mqMsgJson, err := json.Marshal(mqMsg)
	if err != nil {
		return mqMsg, gerror.New(fmt.Sprint("queue redis 生产者解析json消息失败:", err))
	}

	queueName := r.genQueueName(r.groupName, topic)

	_, err = redis.Int64(rdx.Do("LPUSH", queueName, mqMsgJson))
	if err != nil {
		return mqMsg, gerror.New(fmt.Sprint("queue redis 生产者添加消息失败:", err))
	}

	if r.timeout > 0 {
		_, err = rdx.Do("EXPIRE", queueName, r.timeout)
		if err != nil {
			return mqMsg, gerror.New(fmt.Sprint("queue redis 生产者设置过期时间失败:", err))
		}
	}
	return
}

// ListenReceiveMsgDo 消费数据
func (r *RedisMq) ListenReceiveMsgDo(topic string, receiveDo func(mqMsg MqMsg)) (err error) {
	if r.poolName == "" {
		return gerror.New("RedisMq producer not register")
	}
	if topic == "" {
		return gerror.New("RedisMq topic is empty")
	}

	queueName := r.genQueueName(r.groupName, topic)
	go func() {
		for range time.Tick(500 * time.Millisecond) {
			mqMsgList := r.loopReadQueue(queueName)
			for _, mqMsg := range mqMsgList {
				receiveDo(mqMsg)
			}
		}
	}()
	select {}
}

// 生成队列名称
func (r *RedisMq) genQueueName(groupName string, topic string) string {
	return fmt.Sprintf(consts.QueueName+"%s_%s", groupName, topic)
}

func (r *RedisMq) loopReadQueue(queueName string) (mqMsgList []MqMsg) {
	rdx, put, err := getRedis(r.poolName, r.retry)
	defer put()
	if err != nil {
		g.Log().Warningf(ctx, "loopReadQueue getRedis err:%+v", err)
		return
	}

	for {
		infoByte, err := redis.Bytes(rdx.Do("RPOP", queueName))
		if redis.ErrNil == err || len(infoByte) == 0 {
			break
		}
		if err != nil {
			g.Log().Warningf(ctx, "loopReadQueue redis RPOP err:%+v", err)
			break
		}

		var mqMsg MqMsg
		if err = json.Unmarshal(infoByte, &mqMsg); err != nil {
			g.Log().Warningf(ctx, "loopReadQueue Unmarshal err:%+v", err)
			break
		}
		if mqMsg.MsgId != "" {
			mqMsgList = append(mqMsgList, mqMsg)
		}

	}
	return mqMsgList
}

func RegisterRedisMqProducerMust(connOpt RedisOption, poolOpt PoolOption, groupName string, retry int) (client MqProducer) {
	var err error
	client, err = RegisterRedisMq(connOpt, poolOpt, groupName, retry)
	if err != nil {
		g.Log().Fatal(ctx, "RegisterRedisMqProducerMust err:%+v", err)
		return
	}
	return client
}

// RegisterRedisMqConsumerMust 注册消费者
func RegisterRedisMqConsumerMust(connOpt RedisOption, poolOpt PoolOption, groupName string) (client MqConsumer) {
	var err error
	client, err = RegisterRedisMq(connOpt, poolOpt, groupName, 0)
	if err != nil {
		g.Log().Fatal(ctx, "RegisterRedisMqConsumerMust err:%+v", err)
		return
	}
	return client
}

// RegisterRedisMq 注册redis实例
func RegisterRedisMq(connOpt RedisOption, poolOpt PoolOption, groupName string, retry int) (mqIns *RedisMq, err error) {
	poolName, err := registerRedis(connOpt.Addr, connOpt.Passwd, connOpt.DBnum, poolOpt)
	if err != nil {
		return
	}

	if retry <= 0 {
		retry = 0
	}

	mqIns = &RedisMq{
		poolName:  poolName,
		groupName: groupName,
		retry:     retry,
		timeout:   connOpt.Timeout,
	}
	return mqIns, nil
}

// RegisterRedis 注册一个redis配置
func registerRedis(host, pass string, dbNum int, opt PoolOption) (poolName string, err error) {
	poolName = encrypt.Md5ToString(fmt.Sprintf("%s-%s-%d", host, pass, dbNum))
	if _, ok := redisPoolMap[poolName]; ok {
		return poolName, nil
	}

	connRedis := func() (interface{}, error) {
		conn, err := redis.Dial("tcp", host)
		if err != nil {
			return nil, err
		}
		if pass != "" {
			if _, err := conn.Do("AUTH", pass); err != nil {
				return nil, err
			}
		}
		if dbNum > 0 {
			if _, err := conn.Do("SELECT", dbNum); err != nil {
				return nil, err
			}
		}
		return conn, err
	}

	// closeRedis 关闭连接
	closeRedis := func(v interface{}) error {
		return v.(redis.Conn).Close()
	}

	// pingRedis 检测连接连通性
	pingRedis := func(v interface{}) error {
		conn := v.(redis.Conn)
		val, err := redis.String(conn.Do("PING"))
		if err != nil {
			return err
		}
		if val != "PONG" {
			return gerror.New("queue redis ping is error ping => " + val)
		}

		return nil
	}

	p, err := pool.NewChannelPool(&pool.Config{
		InitialCap:  opt.InitCap,
		MaxCap:      opt.MaxCap,
		Factory:     connRedis,
		Close:       closeRedis,
		Ping:        pingRedis,
		IdleTimeout: time.Duration(opt.IdleTimeout) * time.Second,
	})

	if err != nil {
		return poolName, err
	}

	mutex.Lock()
	defer mutex.Unlock()
	redisPoolMap[poolName] = p
	return poolName, nil
}

//  getRedis 获取一个redis db连接
func getRedis(poolName string, retry int) (db redis.Conn, put func(), err error) {
	put = func() {}
	if _, ok := redisPoolMap[poolName]; ok == false {
		return nil, put, gerror.New("db connect is nil")
	}
	redisPool := redisPoolMap[poolName]

	conn, err := redisPool.Get()
	for i := 0; i < retry; i++ {
		if err == nil {
			break
		}
		conn, err = redisPool.Get()
		time.Sleep(time.Second)
	}

	if err != nil {
		return nil, put, err
	}
	put = func() {
		redisPool.Put(conn)
	}

	db = conn.(redis.Conn)
	return db, put, nil
}

func getRandMsgId() (msgId string) {
	rand.Seed(time.Now().UnixNano())
	radium := rand.Intn(999) + 1
	timeCode := time.Now().UnixNano()

	msgId = fmt.Sprintf("%d%.4d", timeCode, radium)
	return msgId
}
