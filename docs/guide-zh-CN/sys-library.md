## 功能扩展库

目录

- 缓存驱动
- 请求上下文
- JWT
- 地理定位（待写）
- 通知（待写）


### 缓存驱动

> 系统默认的缓存驱动为file，目前已支持：memory|redis|file等多种驱动。请自行选择适合你的驱动使用。

- 配置文件：server/manifest/config/config.yaml

```yaml
#缓存
cache:
  adapter: "file"                    # 缓存驱动方式，支持：memory|redis|file，不填默认memory
  fileDir: "./storage/cache"         # 文件缓存路径，adapter=file时必填
```

#### 使用方式
```go
package main

import (
	"hotgo/internal/library/cache"
	"github.com/gogf/gf/v2/os/gctx"
)

func  test() {
	ctx := gctx.New()

	// 添加/修改
	cache.Instance().Set(ctx, "qwe", 123, 0)

	// 查询
	cache.Instance().Get(ctx, "qwe")

	// 删除
	cache.Instance().Remove(ctx, "qwe")
	
	// 更多方法请参考：https://goframe.org/pages/viewpage.action?pageId=27755640
}

```

### 请求上下文

- 主要用于在处理HTTP和websocket请求时通过中间件将用户、应用、插件等信息绑定到上下文中，方便在做业务处理时用到这些信息

```go
package admin

import (
	"fmt"
	"context"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/addons"
)


func test(ctx context.Context) {
	// 获取当前请求的所有上下文变量
	var ctxModel = contexts.Get(ctx)
	fmt.Printf("当前请求的所有上下文变量：%+v\n", ctxModel)

	// 获取当前请求的应用模块
	var module = contexts.GetModule(ctx)
	fmt.Printf("当前请求的应用：%+v\n", module)
	
	// 获取当前请求的用户信息
	var member = contexts.GetUser(ctx)
	fmt.Printf("当前访问用户信息：%+v\n", member)
	
	// 获取当前请求的插件模块
	fmt.Printf("当前是否为插件请求：%v", contexts.IsAddonRequest(ctx))
	if contexts.IsAddonRequest(ctx) {
		fmt.Printf("当前插件名称：%v", contexts.GetAddonName(ctx))
		fmt.Printf("当前插件信息：%v", addons.GetModule(contexts.GetAddonName(ctx)))
	}
}

```

### JWT

- 基于jwt+缓存驱动实现的用户登录令牌功能，支持自动续约，解决了jwt服务端无法退出问题和jwt令牌无法主动失效问题

#### 配置示例
```yaml
# 登录令牌
token:
  secretKey: "hotgo123"                  # 令牌加密秘钥，考虑安全问题生产环境中请修改默认值
  expires: 604800                        # 令牌有效期，单位：秒。默认7天
  autoRefresh: true                      # 是否开启自动刷新过期时间， false|true 默认为true
  refreshInterval: 86400                 # 刷新间隔，单位：秒。必须小于expires，否则无法触发。默认1天内只允许刷新一次
  maxRefreshTimes: 30                    # 最大允许刷新次数，-1不限制。默认30次
  multiLogin: true                       # 是否允许多端登录， false|true 默认为true

```

```go
package admin

import (
	"fmt"
	"context"
	"hotgo/internal/library/token"
	"hotgo/internal/model"
)


func test(ctx context.Context) {
	// 登录
	user := &model.Identity{
		Id:       mb.Id,
		Pid:      mb.Pid,
		DeptId:   mb.DeptId,
		RoleId:   ro.Id,
		RoleKey:  ro.Key,
		Username: mb.Username,
		RealName: mb.RealName,
		Avatar:   mb.Avatar,
		Email:    mb.Email,
		Mobile:   mb.Mobile,
		App:      consts.AppAdmin,
		LoginAt:  gtime.Now(),
	}

	loginToken, expires, err := token.Login(ctx, user)
	if err != nil {
		return nil, err
	}
	
	// gf请求对象
	r := *ghttp.Request
	
	// 获取登录用户信息
	user, err := token.ParseLoginUser(r)
	if err != nil {
		return
	}
	
	// 注销登录
	err = token.Logout(r)
}

```

### 地理定位
```go
// 待写
```

### 通知
```go
// 待写
```