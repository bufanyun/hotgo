## 消息队列

目录

- 缓存驱动
- 上下文（待写）
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

### 上下文
```go
// 待写
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