package queue

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

const (
	queueStatusPending = 0
	queueStatusDead    = 1
	queueTable         = "hg_sys_queue_message"
)

// PgsqlProducerMq 把消息写到 PostgreSQL。
type PgsqlProducerMq struct {
	db gdb.DB
}

// PgsqlConsumerMq 单进程消费 PostgreSQL 队列。
type PgsqlConsumerMq struct {
	db        gdb.DB
	maxRetry  int
	retryBase time.Duration
	stop      context.Context
}

// RegisterPgsqlMqProducer 注册把消息写入 PostgreSQL 的生产者。
func RegisterPgsqlMqProducer(conf *PgsqlConf) (MqProducer, error) {
	return &PgsqlProducerMq{db: g.DB()}, nil
}

// RegisterPgsqlMqConsumer 注册 PostgreSQL 队列消费者。成功才删除，失败退避，超过次数进入死信。
func RegisterPgsqlMqConsumer(conf *PgsqlConf) (MqConsumer, error) {
	maxRetry := defaultDiskRetry
	base := defaultDiskRetryBase
	if conf != nil && conf.Retry > 0 {
		maxRetry = conf.Retry
	}
	if conf != nil && conf.RetryInterval > 0 {
		base = time.Duration(conf.RetryInterval) * time.Second
	}
	db := g.DB()
	startPgsqlListen(db)
	return &PgsqlConsumerMq{db: db, maxRetry: maxRetry, retryBase: base}, nil
}

// SendMsg 发送即时消息。
func (p *PgsqlProducerMq) SendMsg(topic string, body string) (MqMsg, error) {
	return p.SendByteMsg(topic, []byte(body))
}

// SendByteMsg 发送即时字节消息。
func (p *PgsqlProducerMq) SendByteMsg(topic string, body []byte) (MqMsg, error) {
	return p.insert(topic, body, time.Now())
}

// SendDelayMsg 写入延迟消息，到点后才会被消费。
func (p *PgsqlProducerMq) SendDelayMsg(topic string, body string, delaySecond int64) (MqMsg, error) {
	if delaySecond < 0 {
		delaySecond = 0
	}
	return p.insert(topic, []byte(body), time.Now().Add(time.Duration(delaySecond)*time.Second))
}

// insert 把消息插入 PostgreSQL 队列表。executeAt 为可执行时间。
func (p *PgsqlProducerMq) insert(topic string, body []byte, executeAt time.Time) (MqMsg, error) {
	if topic == "" {
		return MqMsg{}, gerror.New("pgsql queue topic is empty")
	}
	mqMsg := MqMsg{
		RunType:   SendMsg,
		Topic:     topic,
		MsgId:     getRandMsgId(),
		Body:      body,
		Timestamp: time.Now(),
		ExecuteAt: executeAt,
	}
	raw, err := json.Marshal(mqMsg)
	if err != nil {
		return mqMsg, err
	}
	_, err = p.db.Exec(context.Background(),
		"INSERT INTO "+queueTable+"(topic, msg_id, body, status, execute_at, retry, created_at) VALUES(?,?,?,?,?,?,?)",
		topic, mqMsg.MsgId, string(raw), queueStatusPending, unixAtLeast(executeAt), 0, time.Now().Unix(),
	)
	if err == nil {
		publishQueue(topic)
		notifyPgsql(context.Background(), p.db, topic)
	}
	return mqMsg, err
}

// ListenReceiveMsgDo 循环消费。新消息由 NOTIFY 唤醒，到点才再查库。处理成功才删除。
func (q *PgsqlConsumerMq) ListenReceiveMsgDo(topic string, receiveDo func(mqMsg MqMsg) error) error {
	if topic == "" {
		return gerror.New("pgsql queue topic is empty")
	}
	ctx := q.stop
	release := func() {}
	if ctx == nil {
		ctx, release = beginConsumer()
	}
	defer release()
	for {
		if ctx.Err() != nil {
			return nil
		}
		row, err := q.nextPending(ctx, topic)
		if err != nil {
			return err
		}
		if row == nil {
			if !waitQueue(ctx, topic, time.Now().Add(queueIdleRecheck)) {
				return nil
			}
			continue
		}
		if row.executeAt > time.Now().Unix() {
			if !waitQueue(ctx, topic, time.Unix(row.executeAt, 0)) {
				return nil
			}
			continue
		}
		var mqMsg MqMsg
		if unmarshalErr := json.Unmarshal([]byte(row.body), &mqMsg); unmarshalErr != nil || mqMsg.MsgId == "" {
			_ = q.markDead(ctx, row.id, row.retry)
			continue
		}
		if handleErr := receiveDo(mqMsg); handleErr == nil {
			if _, err = q.db.Exec(ctx, "DELETE FROM "+queueTable+" WHERE id=?", row.id); err != nil {
				return err
			}
			continue
		}
		row.retry++
		if row.retry > q.maxRetry {
			_ = q.markDead(ctx, row.id, row.retry)
			continue
		}
		next := time.Now().Add(backoff(q.retryBase, row.retry))
		_, err = q.db.Exec(ctx, "UPDATE "+queueTable+" SET retry=?, execute_at=? WHERE id=?", row.retry, unixAtLeast(next), row.id)
		if err != nil {
			return err
		}
	}
}

type queueRow struct {
	id        int64
	body      string
	retry     int
	executeAt int64
}

// nextPending 取下一条待处理消息，含尚未到点的延迟消息。没有则返回空。
func (q *PgsqlConsumerMq) nextPending(ctx context.Context, topic string) (*queueRow, error) {
	one, err := q.db.GetOne(ctx,
		"SELECT id, body, retry, execute_at FROM "+queueTable+" WHERE topic=? AND status=? ORDER BY execute_at ASC, id ASC LIMIT 1",
		topic, queueStatusPending,
	)
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, nil
	}
	return &queueRow{
		id:        one["id"].Int64(),
		body:      one["body"].String(),
		retry:     one["retry"].Int(),
		executeAt: one["execute_at"].Int64(),
	}, nil
}

// markDead 把消息标成死信，避免堵住同主题后续消息。
func (q *PgsqlConsumerMq) markDead(ctx context.Context, id int64, retry int) error {
	_, err := q.db.Exec(ctx, "UPDATE "+queueTable+" SET status=?, retry=? WHERE id=?", queueStatusDead, retry, id)
	return err
}

// unixAtLeast 向上取整到秒。字段只存 Unix 秒，直接截断会让延迟消息提前到期。
func unixAtLeast(t time.Time) int64 {
	sec := t.Unix()
	if t.After(time.Unix(sec, 0)) {
		sec++
	}
	return sec
}

// backoff 按已重试次数指数增加等待时间，最长 5 分钟。
func backoff(base time.Duration, retry int) time.Duration {
	if base <= 0 {
		base = defaultDiskRetryBase
	}
	d := base
	for i := 1; i < retry; i++ {
		d *= 2
		if d > 5*time.Minute {
			return 5 * time.Minute
		}
	}
	return d
}
