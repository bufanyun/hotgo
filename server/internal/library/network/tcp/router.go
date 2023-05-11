package tcp

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/util/gconv"
)

var GoPool = grpool.New(100)

type RouterHandler func(ctx context.Context, args ...interface{})

// Message 路由消息
type Message struct {
	Router string      `json:"router"`
	Data   interface{} `json:"data"`
}

func SendPkg(conn *gtcp.Conn, message *Message) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return conn.SendPkg(b)
}

func RecvPkg(conn *gtcp.Conn) (*Message, error) {
	if data, err := conn.RecvPkg(); err != nil {
		return nil, err
	} else {
		var msg = new(Message)
		if err = gconv.Scan(data, &msg); err != nil {
			return nil, gerror.Newf("invalid package structure: %s", err.Error())
		}
		if msg.Router == "" {
			return nil, gerror.Newf("message is not routed: %+v", msg)
		}
		return msg, err
	}
}

// MsgPkg 打包消息
func MsgPkg(data interface{}, auth *AuthMeta, traceID string) string {
	// 打包签名
	msg := PkgSign(data, auth.AppId, auth.SecretKey, traceID)

	// 打包响应消息
	PkgResponse(data)

	if msg == nil {
		return ""
	}

	return msg.TraceID
}

// doHandleRouterMsg 处理路由消息
func doHandleRouterMsg(fun RouterHandler, ctx context.Context, cancel context.CancelFunc, args ...interface{}) {
	GoPool.Add(ctx, func(ctx context.Context) {
		fun(ctx, args...)
		cancel()
	})
}
