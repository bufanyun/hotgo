// Package queues
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
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
	jobList = append(jobList, LoginLog)
}

// LoginLog 登录日志
var LoginLog = &qLoginLog{}

type qLoginLog struct{}

// getTopic 主题
func (q *qLoginLog) getTopic() string {
	return consts.QueueLoginLogTopic
}

// handle 处理消息
func (q *qLoginLog) handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	var data entity.SysLoginLog
	if err = json.Unmarshal(mqMsg.Body, &data); err != nil {
		return err
	}
	return service.SysLoginLog().RealWrite(ctx, data)
}
