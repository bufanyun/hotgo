## 定时任务

目录

- 实现接口
- 一个例子
- 更多

> 在实际的项目开发中，定时任务几乎成为不可或缺的一部分。HotGo为定时任务提供一个方便的后台操作界面，让您能够轻松地进行在线启停、修改和立即执行等操作。这样的设计可以极大地改善您在使用定时任务过程中的体验，让整个过程更加顺畅、高效。


### 实现接口
- 为了提供高度的扩展性，定时任务在设计上采用了接口化的思路。只需要实现以下接口，您就可以在任何地方注册和使用定时任务功能，从而实现更大的灵活性和可扩展性。

```go
// Cron 定时任务接口
type Cron interface {
    // GetName 获取任务名称
    GetName() string
    // Execute 执行任务
	Execute(ctx context.Context, parser *Parser) (err error)
}
```


### 一个例子

定时任务的文件结构可以根据具体需要进行调整，以下是一个常见的参考结构：

- 文件路径：server/internal/crons/test.go

```go 
package crons

import (
	"context"
	"hotgo/internal/library/cron"
	"time"
)

func init() {
	cron.Register(Test)
}

// Test 测试任务（无参数）
var Test = &cTest{name: "test"}

type cTest struct {
	name string
}

func (c *cTest) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *cTest) Execute(ctx context.Context, parser *cron.Parser) (err error) {
	parser.Logger.Infof(ctx, "cron test Execute:%v", time.Now()) // 记录任务调度日志
	return
}
```

继续在后台系统设置-定时任务-添加任务，填写的任务名称需要和上面的名称保持一致，再进行简单的策略配置以后，一个后台可控的定时任务就添加好了！


### 更多

定时任务源码路径：server/internal/library/cron/cron.go

更多文档请参考：https://goframe.org/pages/viewpage.action?pageId=1114187