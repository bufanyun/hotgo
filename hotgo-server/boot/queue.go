//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package boot

import (
	"context"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/factory/queue"
	"github.com/bufanyun/hotgo/app/service/sysService"
)

//
//  @Title  消息队列监听
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//
func QueueListen(ctx context.Context) {

	consumer, err := queue.InstanceConsumer()
	if err != nil {
		queue.FatalLog(ctx, "InstanceConsumer异常", err)
		return
	}

	// 全局日志
	if listenErr := consumer.ListenReceiveMsgDo(consts.QueueLogTopic, func(mqMsg queue.MqMsg) {

		// 自定义消费回调
		err := sysService.Log.QueueJob(ctx, mqMsg)

		// 记录消费日志
		queue.ConsumerLog(ctx, consts.QueueLogTopic, mqMsg, err)
	}); listenErr != nil {
		queue.FatalLog(ctx, "主题："+consts.QueueLogTopic+" 监听失败", listenErr)
	}

}
