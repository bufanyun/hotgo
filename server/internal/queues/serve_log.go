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
	jobList = append(jobList, ServeLog)
}

// ServeLog 登录日志
var ServeLog = &qServeLog{}

type qServeLog struct{}

// getTopic 主题
func (q *qServeLog) getTopic() string {
	return consts.QueueServeLogTopic
}

// handle 处理消息
func (q *qServeLog) handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	var data entity.SysServeLog
	if err = json.Unmarshal(mqMsg.Body, &data); err != nil {
		return err
	}
	return service.SysServeLog().RealWrite(ctx, data)
}
