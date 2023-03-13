package tcpclient

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/input/msgin"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

// tcp授权
type sTCPAuth struct {
	client *tcp.Client
}

func init() {
	service.RegisterTCPAuth(newTCPAuth())
}

func newTCPAuth() *sTCPAuth {
	return &sTCPAuth{}
}

// Start 启动服务
func (s *sTCPAuth) Start(ctx context.Context) {
	g.Log().Debug(ctx, "TCPAuth start..")
	simple.SafeGo(ctx, func(ctx context.Context) {
		client, err := tcp.NewClient(&tcp.ClientConfig{
			Addr: "127.0.0.1:8099",
			Auth: &tcp.AuthMeta{
				Group:     "auth",
				Name:      "auth1",
				AppId:     "mengshuai",
				SecretKey: "123456",
			},
			LoginEvent: s.loginEvent,
			CloseEvent: s.closeEvent,
		})
		if err != nil {
			g.Log().Infof(ctx, "TCPAuth NewClient fail：%+v", err)
			return
		}

		s.client = client

		err = s.client.RegisterRouter(map[string]tcp.RouterHandler{
			"ResponseAuthSummary": s.onResponseAuthSummary, // 获取授权信息
		})

		if err != nil {
			g.Log().Infof(ctx, "TCPAuth RegisterRouter fail：%+v", err)
			return
		}

		if err = s.client.Start(); err != nil {
			g.Log().Infof(ctx, "TCPAuth Start fail：%+v", err)
			return
		}
	})
}

// Stop 关闭服务
func (s *sTCPAuth) Stop(ctx context.Context) {
	if s.client != nil {
		s.client.Stop()
		g.Log().Debug(ctx, "TCPAuth stop..")
	}
}

func (s *sTCPAuth) loginEvent() {
	// 登录成功后立即请求一次授权信息
	s.client.Write(&msgin.AuthSummary{})

	// 定时检查授权
	gcron.Add(s.client.Ctx, "@every 1200s", func(ctx context.Context) {
		if !s.client.IsLogin {
			g.Log().Infof(ctx, "TCPAuthVerify client is not logged in, skipped")
			return
		}
		s.client.Write(&msgin.AuthSummary{})
	}, "TCPAuthVerify")
}

func (s *sTCPAuth) closeEvent() {
	// 关闭连接后，删除定时检查授权
	gcron.Remove("TCPAuthVerify")
}

func (s *sTCPAuth) onResponseAuthSummary(args ...interface{}) {
	var in *msgin.ResponseAuthSummary
	if err := gconv.Scan(args[0], &in); err != nil {
		s.client.Logger.Infof(s.client.Ctx, "ResponseAuthSummary message Scan failed:%+v, args:%+v", err, args[0])
		return
	}
	s.client.Logger.Infof(s.client.Ctx, "onResponseAuthSummary in:%+v", *in)

	// 授权异常
	if in.Code != gcode.CodeOK.Code() {
		s.client.Logger.Infof(s.client.Ctx, "onResponseAuthSummary authorization verification failed:%+v", in.Message)
		s.client.Destroy()
		return
	}

	// 授权通过
	// 后续可以做一些操作...
}
