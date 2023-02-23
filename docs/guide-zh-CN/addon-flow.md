## 模块开发流程

目录

- 创建新插件
- 开发
- 调用主模块服务接口
- 访问路径
- 数据迁移



### 创建新插件

1、HotGo 后台进入 开发工具->插件管理->找到创建新插件，根据引导进行创建即可。

> 创建成功后会在 根目录的 addons 目录下生成插件文件

2、创建插件完毕重启服务端后，插件管理中会出现你新创建的插件信息。操作栏有几个按钮，在此进行说明
- 安装：会自动执行 server/xxx插件/main.go 文件中的Install方法，方法中的具体逻辑默认为空，可以根据实际情况自行配置。如生成后台菜单、生成插件配置表初始化数据、迁移home页面、web项目文件等。
```
// Install 安装模块
func (m *module) Install(ctx context.Context) (err error) {
	// ...
	return
}
```

- 更新：会自动执行 server/xxx插件/main.go 文件中的Upgrade方法，方法中的具体逻辑默认为空，可以根据实际情况自行配置。
```
// Upgrade 更新模块
func (m *module) Upgrade(ctx context.Context) (err error) {
	// ...
	return
}
```

- 卸载：会自动执行 server/xxx插件/main.go 文件中的UnInstall方法，方法中的具体逻辑默认为空，可以根据实际情况自行配置。如会清除所有的数据表和已安装的信息等。
```
// UnInstall 卸载模块
func (m *module) UnInstall(ctx context.Context) (err error) {
	// ...
	return
}
```



### 开发

完全可以根据HotGo正常的开发流程去开发对应的API、控制器、业务逻辑、插件内的应用

### 调用主模块服务接口

这里推荐的方式是在插件input层新建一个结构，继承主模块中的input结构。这样做的目的是为了服务与服务之间的输入/输出关系解耦，便于参数扩展和避免插件模块下使用`gf gen service`时出现`import cycle not allowed`。

一个简单的例子：
> 假设hgexample插件模块要通过主模块的服务接口更新插件配置

文件：\server\addons\hgexample\model\input\sysin\config.go
```go
package sysin

import (
	"hotgo/internal/model/input/sysin"
)

// UpdateConfigInp 更新指定配置
type UpdateConfigInp struct {
	sysin.UpdateAddonsConfigInp
}

```

插件模块业务逻辑：\server\addons\hgexample\logic\sys\config.go
```go
package sys

import (
	"context"
	"hotgo/addons/hgexample/global"
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/addons/hgexample/service"
	isc "hotgo/internal/service"
)

type sSysConfig struct{}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

// UpdateConfigByGroup 更新指定分组的配置
func (s *sSysConfig) UpdateConfigByGroup(ctx context.Context, in sysin.UpdateConfigInp) error {
	in.UpdateAddonsConfigInp.AddonName = global.GetSkeleton().Name
	return isc.SysAddonsConfig().UpdateConfigByGroup(ctx, in.UpdateAddonsConfigInp)
}

```

主模块input：\server\internal\model\input\sysin\addons_config.go
```go
package sysin

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UpdateAddonsConfigInp 更新指定插件的配置
type UpdateAddonsConfigInp struct {
	AddonName string `json:"addonName"`
	Group     string `json:"group"`
	List      g.Map  `json:"list"`
}

```

主模块业务逻辑：\server\internal\logic\sys\addons_config.go
```go
package sys

import (
	"context"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

type sSysAddonsConfig struct{}

func NewSysAddonsConfig() *sSysAddonsConfig {
	return &sSysAddonsConfig{}
}

func init() {
	service.RegisterSysAddonsConfig(NewSysAddonsConfig())
}

// UpdateConfigByGroup 更新指定分组的配置
func (s *sSysAddonsConfig) UpdateConfigByGroup(ctx context.Context, in sysin.UpdateAddonsConfigInp) error {
	// ...
	return nil
}

```




### 访问路径

#### 后台插件访问路径

```
// IP+端口或域名/admin/插件名称/API路径
如：127.0.0.1:8000/admin/hgexample/index/test
```

对应控制器路径：`server/addons/hgexample/controller/admin/sys/index.go`

#### 前端API插件访问路径

```
// IP+端口或域名/api/插件名称/API路径
如：127.0.0.1:8000/api/hgexample/index/test
```

对应控制器路径：`server/addons/hgexample/controller/api/index.go`

#### 前台页面插件访问路径

```
// IP+端口或域名/home/插件名称/API路径
如：127.0.0.1:8000/home/hgexample/index/test
```

对应控制器路径：`server/addons/hgexample/controller/home/index.go`


#### Websocket插件访问路径

```
// IP+端口或域名/socket/插件名称/API路径
如：127.0.0.1:8000/socket/hgexample/index/test
```

对应控制器路径：`server/addons/hgexample/controller/socket/index.go`


### 数据迁移

可以将数据迁移逻辑写进server/xxx插件/main.go 文件中的Install方法中，并遵循系统规范进行数据安装
