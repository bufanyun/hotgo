// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"net"
	"sync/atomic"
)

// AuthMeta 认证元数据
type AuthMeta struct {
	Name      string      `json:"name"`      // 客户端名称，当同一个应用ID有多个客户端时请使用不同的名称区分。比如cron1,cron2
	Extra     g.Map       `json:"extra"`     // 自定义数据，可以传递一些额外的自定义数据
	Group     string      `json:"group"`     // 客户端分组
	AppId     string      `json:"appId"`     // 应用ID
	SecretKey string      `json:"secretKey"` // 应用秘钥
	EndAt     *gtime.Time `json:"endAt"`     // 授权过期时间
	Routes    []string    `json:"routes"`    // 授权路由
}

// Conn tcp连接
type Conn struct {
	CID       int64      // 连接ID
	Conn      *gtcp.Conn // 连接对象
	Auth      *AuthMeta  // 认证元数据
	Heartbeat int64      // 心跳
	FirstTime int64      // 首次连接时间

	writeChan chan []byte  // 发数据
	closeFlag *gtype.Bool  // 关闭标签
	logger    *glog.Logger // 日志处理器
	msgParser *MsgParser   // 消息处理器
}

var idCounter int64

func NewConn(conn *gtcp.Conn, logger *glog.Logger, msgParser *MsgParser) *Conn {
	tcpConn := new(Conn)
	tcpConn.CID = atomic.AddInt64(&idCounter, 1)
	tcpConn.Conn = conn
	tcpConn.Heartbeat = gtime.Timestamp()
	tcpConn.FirstTime = gtime.Timestamp()

	tcpConn.writeChan = make(chan []byte, 1000)
	tcpConn.closeFlag = gtype.NewBool(false)
	tcpConn.logger = logger
	tcpConn.msgParser = msgParser

	go func() {
		for b := range tcpConn.writeChan {
			if b == nil {
				break
			}
			if err := conn.SendPkg(b); err != nil {
				break
			}
		}
	}()
	return tcpConn
}

func (c *Conn) Run() error {
	for {
		data, err := c.Conn.RecvPkg()
		if err != nil {
			return gerror.NewCodef(gcode.CodeInvalidRequest, "read packet err:%+v conn closed", err)
		}

		if c.closeFlag.Val() {
			return nil
		}

		msg, err := c.msgParser.Encoding(data)
		if err != nil {
			return gerror.NewCodef(gcode.CodeInternalError, "message encoding err:%+v conn closed", err)
		}

		ctx, err := c.bindContext(msg)
		if err != nil {
			return gerror.NewCodef(gcode.CodeInternalError, "bindContext err:%+v message: %+v", err, msg)
		}

		if err = c.msgParser.handleInterceptor(ctx, msg); err != nil {
			c.logger.Warning(ctx, gerror.Wrap(err, "interceptor authentication failed"))
			continue
		}

		if err = c.msgParser.handleRouterMsg(ctx, msg); err != nil {
			return err
		}
	}
}

// RemoteAddr returns the remote network address, if known.
func (c *Conn) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// LocalAddr returns the local network address, if known.
func (c *Conn) LocalAddr() net.Addr {
	return c.Conn.LocalAddr()
}

// Write
func (c *Conn) Write(b []byte) {
	if !c.closeFlag.Val() {
		c.writeChan <- b
	}
}

// Send 发送消息
func (c *Conn) Send(ctx context.Context, data interface{}) error {
	if c.closeFlag.Val() {
		return gerror.New("conn is closed")
	}
	b, err := c.msgParser.Decoding(ctx, data, "")
	if err != nil {
		return err
	}
	c.Write(b)
	return nil
}

func (c *Conn) Close() {
	if c.closeFlag.Val() {
		return
	}
	c.closeFlag.Set(true)
	close(c.writeChan)
	c.Conn.Close()
}

// Request 发送消息并等待响应结果
func (c *Conn) Request(ctx context.Context, data interface{}) (interface{}, error) {
	if c.closeFlag.Val() {
		return nil, gerror.New("conn is closed")
	}
	msgId := grand.S(16)
	b, err := c.msgParser.Decoding(ctx, data, msgId)
	if err != nil {
		return nil, err
	}
	return c.msgParser.rpc.Request(ctx, msgId, func() {
		c.Write(b)
	})
}

// RequestScan 发送消息并等待响应结果，将结果保存在response中
func (c *Conn) RequestScan(ctx context.Context, data, response interface{}) error {
	body, err := c.Request(ctx, data)
	if err != nil {
		return err
	}
	return gvar.New(body).Scan(response)
}

// bindContext 将用户身份绑定到上下文
func (c *Conn) bindContext(msg *Message) (ctx context.Context, err error) {
	ctx = initCtx(gctx.New(), &Context{
		Conn: c,
	})
	return SetCtxTraceID(ctx, msg.TraceId)
}
