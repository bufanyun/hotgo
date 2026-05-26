// Package queues
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package queues

import (
	"context"
	"encoding/json"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
)

func init() {
	queue.RegisterConsumer(SysLog)
}

// SysLog 系统日志
var SysLog = &qSysLog{}

type qSysLog struct{}

// GetTopic 主题
func (q *qSysLog) GetTopic() string {
	return consts.QueueLogTopic
}

// Handle 处理消息
func (q *qSysLog) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	var data entity.SysLog
	if err = json.Unmarshal(mqMsg.Body, &data); err != nil {
		return err
	}
	return service.SysLog().RealWrite(ctx, data)
}
