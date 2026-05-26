// Package queue
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package queue

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

const (
	ConsumerLogErrFormat = "消费 [%s] 失败, body:%+v, err:%+v"
	ProducerLogErrFormat = "生产 [%s] 失败, body:%+v, err:%+v"
)

func Logger() *glog.Logger {
	return g.Log("queue")
}

// ConsumerLog 消费日志
func ConsumerLog(ctx context.Context, topic string, mqMsg MqMsg, err error) {
	if err != nil {
		Logger().Errorf(ctx, ConsumerLogErrFormat, topic, string(mqMsg.Body), err)
	}
}

// ProducerLog 生产日志
func ProducerLog(ctx context.Context, topic string, mqMsg MqMsg, err error) {
	if err != nil {
		Logger().Infof(ctx, ProducerLogErrFormat, topic, string(mqMsg.Body), err)
	}
}
