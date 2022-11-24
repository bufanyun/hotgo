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
	"hotgo/internal/consts"
)

// ConsumerLog 消费日志
func ConsumerLog(ctx context.Context, topic string, mqMsg MqMsg, err error) {
	if err != nil {
		g.Log(consts.QueueLogPath).Error(ctx, "消费 ["+topic+"] 失败", mqMsg, err)
	} else {
		g.Log(consts.QueueLogPath).Debug(ctx, "消费 ["+topic+"] 成功", mqMsg.MsgId)
	}
}

// ProducerLog 生产日志
func ProducerLog(ctx context.Context, topic string, data interface{}, err error) {
	if err != nil {
		g.Log(consts.QueueLogPath).Error(ctx, "生产 ["+topic+"] 失败", gconv.String(data))
	} else {
		g.Log(consts.QueueLogPath).Debug(ctx, "生产 ["+topic+"] 成功", gconv.String(data))
	}
}

// FatalLog 致命日志
func FatalLog(ctx context.Context, text string, err error) {
	g.Log(consts.QueueLogPath).Fatal(ctx, text+":", err)
}

// Log 通用日志
func Log(ctx context.Context, text string) {
	g.Log(consts.QueueLogPath).Debug(ctx, text)
}
