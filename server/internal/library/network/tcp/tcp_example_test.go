package tcp_test

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"hotgo/internal/library/network/tcp"
	"testing"
	"time"
)

var T *testing.T // 声明一个全局的 *testing.T 变量

type TestMsgReq struct {
	Name string `json:"name"`
}

type TestMsgRes struct {
	tcp.ServerRes
}

type TestRPCMsgReq struct {
	Name string `json:"name"`
}

type TestRPCMsgRes struct {
	tcp.ServerRes
}

func onTestMsg(ctx context.Context, req *TestMsgReq) {
	fmt.Printf("服务器收到消息 ==> onTestMsg:%+v\n", req)
	conn := tcp.ConnFromCtx(ctx)
	gtest.C(T, func(t *gtest.T) {
		t.AssertNE(conn, nil)
	})

	res := new(TestMsgRes)
	res.Message = fmt.Sprintf("你的名字：%v", req.Name)
	conn.Send(ctx, res)
}

func onResponseTestMsg(ctx context.Context, req *TestMsgRes) {
	fmt.Printf("客户端收到响应消息 ==> TestMsgRes:%+v\n", req)
	err := req.GetError()
	gtest.C(T, func(t *gtest.T) {
		t.AssertNil(err)
	})
}

func onTestRPCMsg(ctx context.Context, req *TestRPCMsgReq) (res *TestRPCMsgRes, err error) {
	fmt.Printf("服务器收到消息 ==> onTestRPCMsg:%+v\n", req)
	res = new(TestRPCMsgRes)
	res.Message = fmt.Sprintf("你的名字：%v", req.Name)
	return
}

func startTCPServer() {
	serv := tcp.NewServer(&tcp.ServerConfig{
		Name: "hotgo",
		Addr: ":8002",
	})

	// 注册路由
	serv.RegisterRouter(
		onTestMsg,
	)

	// 注册RPC路由
	serv.RegisterRPCRouter(
		onTestRPCMsg,
	)

	// 服务监听
	err := serv.Listen()
	gtest.C(T, func(t *gtest.T) {
		t.AssertNil(err)
	})
}

// 一个基本的消息收发
func TestSendMsg(t *testing.T) {
	T = t
	go startTCPServer()

	ctx := gctx.New()
	client := tcp.NewClient(&tcp.ClientConfig{
		Addr: "127.0.0.1:8002",
	})

	// 注册路由
	client.RegisterRouter(
		onResponseTestMsg,
	)

	go func() {
		err := client.Start()
		gtest.C(T, func(t *gtest.T) {
			t.AssertNil(err)
		})
	}()

	// 确保服务都启动完成
	time.Sleep(time.Second * 1)

	// 拿到客户端的连接
	conn := client.Conn()
	gtest.C(T, func(t *gtest.T) {
		t.AssertNE(conn, nil)
	})

	// 向服务器发送tcp消息，不会阻塞程序执行
	err := conn.Send(ctx, &TestMsgReq{Name: "Tom"})
	gtest.C(T, func(t *gtest.T) {
		t.AssertNil(err)
	})

	// 向服务器发送rpc消息，会等待服务器响应结果，直到拿到结果或响应超时才会继续
	var res TestRPCMsgRes
	if err = conn.RequestScan(ctx, &TestRPCMsgReq{Name: "Tony"}, &res); err != nil {
		gtest.C(T, func(t *gtest.T) {
			t.AssertNil(err)
		})
	}

	fmt.Printf("客户端收到RPC消息响应 ==> TestRPCMsgRes:%+v\n", res)
	time.Sleep(time.Second * 1)
}
