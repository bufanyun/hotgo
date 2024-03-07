## WebSocket服务器

目录

- 一个基本的消息收发例子
- 常用方法
- HTTP接口
- 其他

> hotgo提供了一个WebSocket服务器，随`HTTP服务`启停。集成了许多常用功能，如JWT身份认证、路由消息处理器、一对一消息/群组消息/广播消息、在线用户管理、心跳保持等，大大简化和规范了WebSocket服务器的开发流程。
- [Websocket客户端](sys-websocket-client.md)

###  一个基本的消息收发例子
- 这是一个基本的消息接收并进行处理的简单例子

#### 1.消息处理接口
- 消息处理在设计上采用了接口化的思路。只需要实现以下接口，即可进行WebSocket消息注册
- 文件路径：server/internal/websocket/model.go
```go
package websocket

// EventHandler 消息处理器
type EventHandler func(client *Client, req *WRequest)
```

#### 2.定义消息处理方法
- 以下是功能案例中的一个简单演示，实现了消息处理接口，并将收到的消息原样发送给客户端
- 文件路径：server/addons/hgexample/controller/websocket/handler/index.go
```go
package handler

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/websocket"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

// TestMessage 测试消息
func (c *cIndex) TestMessage(client *websocket.Client, req *websocket.WRequest) {
	g.Log().Infof(client.Context(), "收到客户端测试消息:%v", gjson.New(req).String())
	// 将收到的消息原样发送给客户端
	websocket.SendSuccess(client, req.Event, req.Data) 
}
```

#### 3.注册消息
- 定义消息处理方法后，需要将其注册到WebSocket消息处理器，一般放在对应应用模块的`router/websocket.go`下即可
- 文件路径：server/addons/hgexample/router/websocket.go
```go
package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/addons/hgexample/controller/websocket"
	"hotgo/addons/hgexample/controller/websocket/handler"
	ws "hotgo/internal/websocket"
)

// WebSocket ws路由配置
func WebSocket(ctx context.Context, group *ghttp.RouterGroup) {
	// 注册消息路由
	ws.RegisterMsg(ws.EventHandlers{
		"admin/addons/hgexample/testMessage": handler.Index.TestMessage, // 测试消息
	})
	
	// 这里"admin/addons/hgexample/testMessage"代表的是一个消息处理ID，可以自定义。建议的格式是和HTTP接口格式保持一致，这样还可以便于对用户请求的消息进行权限验证
	// 客户端连接后，向WebSocket服务器发送event为"admin/addons/hgexample/testMessage"的消息时，会调用TestMessage方法
}
```

- 到此，你已了解了WebSocket消息接收并进行处理的基本流程


### 常用方法
- websocket服务器还提供了一些常用的方法，下面只对部分进行说明
```go
func test() {
	websocket.SendToAll()      // 发送全部客户端
	websocket.SendToClientID() // 发送单个客户端
	websocket.SendToUser()     // 发送单个用户
	websocket.SendToTag()      // 发送某个标签、群组

    client := websocket.Manager().GetClient(id)         // 通过连接ID获取客户端连接
    client := websocket.Manager().GetUserClient(userId) // 通过用户ID获取客户端连接，因为用户是可多端登录的，这里返回的是一个切片
    
    websocket.SendSuccess(client, "admin/addons/hgexample/testMessage", "消息内容")           // 向指定客户端发送一条成功的消息
    websocket.SendError(client, "admin/addons/hgexample/testMessage", gerror.New("错误内容")) // 向指定客户端发送一条失败的消息

}
```


### HTTP接口
- 你还可以通过http接口方式调用WebSocket发送消息
- 参考文件：server/internal/controller/websocket/send.go


### 其他
- WebSocket被连接时需验证用户认证中间件，所以用户必须登录成功后才能连接成功
- 参考文件：server/internal/logic/middleware/weboscket_auth.go
```go
package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/library/response"
	"hotgo/utility/simple"
)

// WebSocketAuth websocket鉴权中间件
func (s *sMiddleware) WebSocketAuth(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		path = gstr.Replace(r.URL.Path, simple.RouterPrefix(ctx, consts.AppWebSocket), "", 1)
	)

	// 不需要验证登录的路由地址
	if s.IsExceptLogin(ctx, consts.AppWebSocket, path) {
		r.Middleware.Next()
		return
	}

	// 将用户信息传递到上下文中
	if err := s.DeliverUserContext(r); err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
		return
	}

	r.Middleware.Next()
}
```

- 如果您不要求用户进行登录即可使用 WebSocket，那么需要对身份验证中间件进行修改。然而，这样做会降低连接的安全性，并且无法应用于需要确定用户身份的情景，因此并不建议采取这种策略
