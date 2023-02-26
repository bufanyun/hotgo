## 消息队列

目录

- 缓存驱动
- 请求上下文（待写）
- JWT（待写）
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
```go
// 待写
```

### 地理定位
```go
// 待写
```

### 通知
```go
// 待写
```