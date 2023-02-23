## 模块辅助说明

目录

- 模块结构
- 获取模块信息
- 插件路由规则


#### 模块结构
```go
// Skeleton 模块骨架
type Skeleton struct {
	Label       string `json:"label"`       // 标识
	Name        string `json:"name"`        // 名称
	Group       int    `json:"group"`       // 分组
	Logo        string `json:"logo"`        // logo
	Brief       string `json:"brief"`       // 简介
	Description string `json:"description"` // 详细描述
	Author      string `json:"author"`      // 作者
	Version     string `json:"version"`     // 版本号
	RootPath    string `json:"rootPath"`    // 根路径
}

func (s *Skeleton) GetModule() Module {
	return GetModule(s.Name)
}

// Module 插件模块
type Module interface {
	Init(ctx context.Context)                                 // 初始化
	InitRouter(ctx context.Context, group *ghttp.RouterGroup) // 初始化并注册路由
	Ctx() context.Context                                     // 上下文
	GetSkeleton() *Skeleton                                   // 架子
	Install(ctx context.Context) error                        // 安装模块
	Upgrade(ctx context.Context) error                        // 更新模块
	UnInstall(ctx context.Context) error                      // 卸载模块
}
```

#### 获取模块信息

- 在插件模块内
```go
package main

import (
	"fmt"
	"hotgo/addons/hgexample/global"
)

func test()  {
	fmt.Printf("当前插件模块是：%+v", global.GetSkeleton())
}
```

- 在插件模块外
```go
package main

import (
	"context"
	"fmt"
	"hotgo/internal/library/addons"
	"hotgo/internal/library/contexts"
)

func test(ctx context.Context)  {
	fmt.Printf("当前是否为插件请求：%v", contexts.IsAddonRequest(ctx))
	if contexts.IsAddonRequest(ctx) {
		fmt.Printf("当前插件名称：%v", contexts.GetAddonName(ctx))
		fmt.Printf("当前插件信息：%v", addons.GetModule(contexts.GetAddonName(ctx)))
	}
}
```

- 更多辅助方法请参考：\server\internal\library\addons

#### 插件路由规则
- 如果你不喜欢现在的路由风格，可以自行调整。修改位置在：\server\internal\library\addons\addons.go的RouterPrefix方法。 
- 调整后如web前端页面中有之前的路由风格也需同步修改。

```go
package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
)

// RouterPrefix 路由前缀
// 最终效果：/应用名称/插件模块名称/xxx/xxx
func RouterPrefix(ctx context.Context, app, name string) string {
	var prefix = "/"
	switch app {
	case consts.AppAdmin:
		prefix = g.Cfg().MustGet(ctx, "router.admin.prefix", "/admin").String()
	case consts.AppApi:
		prefix = g.Cfg().MustGet(ctx, "router.api.prefix", "/api").String()
	case consts.AppHome:
		prefix = g.Cfg().MustGet(ctx, "router.home.prefix", "/home").String()
	case consts.AppWebSocket:
		prefix = g.Cfg().MustGet(ctx, "router.ws.prefix", "/socket").String()
	}

	return prefix + "/" + name
}
```