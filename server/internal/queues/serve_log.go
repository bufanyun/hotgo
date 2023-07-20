// Package queues
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package queues

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
)

func init() {
	queue.RegisterConsumer(ServeLog)
}

// ServeLog 登录日志
var ServeLog = &qServeLog{}

type qServeLog struct{}

// GetTopic 主题
func (q *qServeLog) GetTopic() string {
	return consts.QueueServeLogTopic
}

// Handle 处理消息
func (q *qServeLog) Handle(ctx context.Context, mqMsg queue.MqMsg) error {
	var data entity.SysServeLog
	if err := json.Unmarshal(mqMsg.Body, &data); err != nil {
		g.Dump("ServeLog Handle Unmarshal err:%+v", err)
		return nil
	}

	if err := service.SysServeLog().RealWrite(ctx, data); err != nil {
		g.Dump("ServeLog Handle Write err:%+v", err)
		return nil
	}
	return nil
}
