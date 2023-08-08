// Package queue
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/utility/encrypt"
	"math/rand"
	"strconv"
	"time"
)

type RedisMq struct {
	poolName  string
	groupName string
	timeout   int64
}

type RedisOption struct {
	Timeout int64
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

	mqMsg = MqMsg{
		RunType:   SendMsg,
		Topic:     topic,
		MsgId:     getRandMsgId(),
		Body:      body,
		Timestamp: time.Now(),
	}

	data, err := json.Marshal(mqMsg)
	if err != nil {
		return
	}

	key := r.genKey(r.groupName, topic)
	if _, err = g.Redis().Do(ctx, "LPUSH", key, data); err != nil {
		return
	}

	if r.timeout > 0 {
		if _, err = g.Redis().Do(ctx, "EXPIRE", key, r.timeout); err != nil {
			return
		}
	}

	return
}

func (r *RedisMq) SendDelayMsg(topic string, body string, delaySecond int64) (mqMsg MqMsg, err error) {
	if delaySecond < 1 {
		return r.SendMsg(topic, body)
	}

	if r.poolName == "" {
		err = gerror.New("SendDelayMsg RedisMq not register")
		return
	}

	if topic == "" {
		err = gerror.New("SendDelayMsg RedisMq topic is empty")
		return
	}

	mqMsg = MqMsg{
		RunType:   SendMsg,
		Topic:     topic,
		MsgId:     getRandMsgId(),
		Body:      []byte(body),
		Timestamp: time.Now(),
	}

	data, err := json.Marshal(mqMsg)
	if err != nil {
		return
	}

	var (
		conn         = g.Redis()
		key          = r.genKey(r.groupName, "delay:"+topic)
		expireSecond = time.Now().Unix() + delaySecond
		timePiece    = fmt.Sprintf("%s:%d", key, expireSecond)
		z            = gredis.ZAddMember{Score: float64(expireSecond), Member: timePiece}
	)

	if _, err = conn.ZAdd(ctx, key, &gredis.ZAddOption{}, z); err != nil {
		return
	}

	if _, err = conn.RPush(ctx, timePiece, data); err != nil {
		return
	}

	// consumer will also delete the item
	if r.timeout > 0 {
		_, _ = conn.Expire(ctx, timePiece, r.timeout+delaySecond)
		_, _ = conn.Expire(ctx, key, r.timeout)
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

	var (
		key  = r.genKey(r.groupName, topic)
		key2 = r.genKey(r.groupName, "delay:"+topic)
	)

	go func() {
		for range time.Tick(300 * time.Millisecond) {
			mqMsgList := r.loopReadQueue(key)
			for _, mqMsg := range mqMsgList {
				receiveDo(mqMsg)
			}
		}
	}()

	go func() {
		mqMsgCh, errCh := r.loopReadDelayQueue(key2)
		for mqMsg := range mqMsgCh {
			receiveDo(mqMsg)
		}
		for err = range errCh {
			if err != nil && err != context.Canceled && err != context.DeadlineExceeded {
				Logger().Infof(ctx, "ListenReceiveMsgDo Delay topic:%v, err:%+v", topic, err)
			}
		}
	}()

	select {}
}

// 生成队列key
func (r *RedisMq) genKey(groupName string, topic string) string {
	return fmt.Sprintf("queue:%s_%s", groupName, topic)
}

func (r *RedisMq) loopReadQueue(key string) (mqMsgList []MqMsg) {
	conn := g.Redis()
	for {
		data, err := conn.Do(ctx, "RPOP", key)
		if err != nil {
			Logger().Warningf(ctx, "loopReadQueue redis RPOP err:%+v", err)
			break
		}

		if data.IsEmpty() {
			break
		}

		var mqMsg MqMsg
		if err = data.Scan(&mqMsg); err != nil {
			Logger().Warningf(ctx, "loopReadQueue Scan err:%+v", err)
			break
		}

		if mqMsg.MsgId != "" {
			mqMsgList = append(mqMsgList, mqMsg)
		}
	}
	return mqMsgList
}

func RegisterRedisMqProducer(connOpt RedisOption, groupName string) (client MqProducer) {
	return RegisterRedisMq(connOpt, groupName)
}

// RegisterRedisMqConsumer 注册消费者
func RegisterRedisMqConsumer(connOpt RedisOption, groupName string) (client MqConsumer) {
	return RegisterRedisMq(connOpt, groupName)
}

// RegisterRedisMq 注册redis实例
func RegisterRedisMq(connOpt RedisOption, groupName string) *RedisMq {
	return &RedisMq{
		poolName:  encrypt.Md5ToString(fmt.Sprintf("%s-%d", groupName, time.Now().UnixNano())),
		groupName: groupName,
		timeout:   connOpt.Timeout,
	}
}

func getRandMsgId() string {
	rand.NewSource(time.Now().UnixNano())
	radium := rand.Intn(999) + 1
	timeCode := time.Now().UnixNano()
	return fmt.Sprintf("%d%.4d", timeCode, radium)
}

func (r *RedisMq) loopReadDelayQueue(key string) (resCh chan MqMsg, errCh chan error) {
	resCh = make(chan MqMsg)
	errCh = make(chan error, 1)

	go func() {
		defer close(resCh)
		defer close(errCh)

		conn := g.Redis()
		for {
			now := time.Now().Unix()
			do, err := conn.Do(ctx, "zrangebyscore", key, "0", strconv.FormatInt(now, 10), "limit", 0, 1)
			if err != nil {
				return
			}

			val := do.Strings()
			if len(val) == 0 {
				select {
				case <-ctx.Done():
					errCh <- ctx.Err()
					return
				case <-time.After(time.Second):
					continue
				}
			}
			for _, listK := range val {
				for {
					pop, err := conn.LPop(ctx, listK)
					if err != nil {
						errCh <- err
						return
					} else if pop.IsEmpty() {
						_, _ = conn.ZRem(ctx, key, listK)
						_, _ = conn.Del(ctx, listK)
						break
					} else {
						var mqMsg MqMsg
						if err = pop.Scan(&mqMsg); err != nil {
							g.Log().Warningf(ctx, "loopReadDelayQueue Scan err:%+v", err)
							break
						}

						if mqMsg.MsgId == "" {
							continue
						}

						select {
						case resCh <- mqMsg:
						case <-ctx.Done():
							errCh <- ctx.Err()
							return
						}
					}
				}
			}
		}
	}()
	return resCh, errCh
}
