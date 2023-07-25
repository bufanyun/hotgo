## TCP服务器

目录

- 配置文件
- 一个基本的消息收发例子
- 注册路由
- 拦截器
- 服务认证
- 更多

> HotGo基于GF框架的TCP服务器组件，提供了一个简单而灵活的方式快速搭建基于TCP的服务应用。集成了许多常用功能，如长连接、服务认证、路由分发、RPC消息、拦截器和数据绑定等，大大简化和规范了服务器开发流程。

### 配置文件
- 配置文件：server/manifest/config/config.yaml

```yaml
tcp:
  # 服务器
  server:
    address: ":8099"
  # 客户端
  client:
    # 定时任务
    cron:
      group: "cron"                                                 # 分组名称
      name: "cron1"                                                 # 客户端名称
      address: "127.0.0.1:8099"                                     # 服务器地址
      appId: "1002"                                                 # 应用名称
      secretKey: "hotgo"                                            # 密钥
    # 系统授权
    auth:
      group: "auth"                                                 # 分组名称
      name: "auth1"                                                 # 客户端名称
      address: "127.0.0.1:8099"                                     # 服务器地址
      appId: "mengshuai"                                            # 应用名称
      secretKey: "123456"                                           # 密钥

```
- 可以看到，除了服务器配置外，还有两个客户端配置`cron` 和`auth`
- `cron`是HotGo内置的定时任务服务，和http服务通过RPC通讯以实现和后台交互，使其可以独立、集群部署。
- `auth`可以为第三方平台提供授权服务。如果你需要他，可以将它部署在第三方程序中，在重要的位置进行授权验证。

### 一个基本的消息收发测试用例

- 文件路径：server/internal/library/network/tcp/tcp_example_test.go

```go
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

```


### 注册路由

- 从上面的例子可以看到，不管是普通TCP消息和RPC消息的请求/响应结构体都采用类似GF框架的规范路由的结构，请求`XxxRes`/响应`XxxRes`的格式，是不是很亲切？


### 拦截器

- 不管是服务端还是客户端，在初始化时都可以注册多个拦截器来满足更多场景的服务开发，下面是一个使用例子：

```go
package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/library/network/tcp"
)

func main()  {
	serv = tcp.NewServer(&tcp.ServerConfig{
		Name: "hotgo",
		Addr: ":8002",
	})
	
	// 注册拦截器
	// 执行顺序是从前到后，即Interceptor -> Interceptor2 -> Interceptor3。如果中间有任意一个抛出错误，则会中断后续处理
	serv.RegisterInterceptor(Interceptor, Interceptor2, Interceptor3)

	// 服务监听
	if err := serv.Listen(); err != nil {
		if !serv.IsClose() {
			g.Log().Warningf(ctx, "TCPServer Listen err:%v", err)
		}
	}
}

func  Interceptor(ctx context.Context, msg *tcp.Message) (err error) {
	// 可以在拦截器中通过上下文拿到连接
	conn := tcp.ConnFromCtx(ctx)

	// 拿到原始请求消息
	g.Dump(msg)
	
	// 如果想要中断后续处理只需返回一个错误即可，但注意两种情况
	// tcp消息：如果你还想对该消息进行回复应在拦截器中进行处理，例如：conn.Send(ctx, 回复消息内容)
	// rpc消息：返回一个错误后系统会将错误自动回复到rpc响应中，无需单独处理
	return
}

func  Interceptor2(ctx context.Context, msg *tcp.Message) (err error) {
	// ...
	return
}

func  Interceptor3(ctx context.Context, msg *tcp.Message) (err error) {
	// ...
	return
}

```


### 服务认证

- 一般情况下，建议客户端连接到服务器时都通过`授权许可证`的方式进行登录认证，当初始化客户端配置认证数据时，连接成功后会自动进行登录认证。

```go
	// 创建客户端配置
	clientConfig := &tcp.ClientConfig{
		Addr:          "127.0.0.1:8002",
		AutoReconnect: true,
		// 认证数据
		// 认证数据可以在后台-系统监控-在线服务-许可证列表中添加，同一个授权支持多个服务使用，但多个服务不能使用相同的名称进行连接
		Auth: &tcp.AuthMeta{
			Name:      "服务名称",
			Group:     "服务分组",
			AppId:     "APPID",
			SecretKey: "SecretKey",
		},
	}

	// 初始化客户端
	client = tcp.NewClient(clientConfig)
```


### 更多

TCP服务器源码路径：server/internal/library/network/tcp

更多文档请参考：https://goframe.org/pages/viewpage.action?pageId=1114625
