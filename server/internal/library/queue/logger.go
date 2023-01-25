// Package queue
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package queue

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/utility/charset"
)

const (
	ConsumerLogErrFormat = "消费 [%s] 失败, mqMsgId:%+v, mqMsgData:%+v, err:%+v, stack:%+v"
	ProducerLogErrFormat = "生产 [%s] 失败, data:%+v, err:%+v, stack:%+v"
)

// ConsumerLog 消费日志
func ConsumerLog(ctx context.Context, topic string, mqMsg MqMsg, err error) {
	if err != nil {
		g.Log().Printf(ctx, ConsumerLogErrFormat, topic, mqMsg.MsgId, mqMsg.BodyString(), err, charset.ParseErrStack(err))
	} else {
		g.Log().Print(ctx, "消费 ["+topic+"] 成功", mqMsg.MsgId)
	}
}

// ProducerLog 生产日志
func ProducerLog(ctx context.Context, topic string, data interface{}, err error) {
	if err != nil {
		g.Log().Printf(ctx, ProducerLogErrFormat, topic, gconv.String(data), err, charset.ParseErrStack(err))
	} else {
		g.Log().Print(ctx, "生产 ["+topic+"] 成功", gconv.String(data))
	}
}
