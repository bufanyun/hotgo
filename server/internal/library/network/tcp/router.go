// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/util/gconv"
)

// RouterHandler 路由消息处理器
type RouterHandler func(ctx context.Context, args ...interface{})

// Message 路由消息
type Message struct {
	Router string      `json:"router"`
	Data   interface{} `json:"data"`
}

// SendPkg 打包发送的数据包
func SendPkg(conn *gtcp.Conn, message *Message) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return conn.SendPkg(b)
}

// RecvPkg 解包
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
