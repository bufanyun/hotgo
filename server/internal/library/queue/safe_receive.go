package queue

import (
	"runtime/debug"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 一条消息 panic 不能带崩进程，按消费失败处理
func safeReceiveDo(receiveDo func(MqMsg) error, msg MqMsg) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = gerror.Newf("queue 消费回调 panic: %v", r)
			Logger().Errorf(ctx, "%+v topic:%s msgId:%s\n%s", err, msg.Topic, msg.MsgId, debug.Stack())
		}
	}()
	return receiveDo(msg)
}
