## 消息队列

目录

- 配置文件
- 实现接口
- 一个例子
- 控制台
- 自定义队列驱动

> 系统默认的队列驱动为disk(磁盘队列)，目前已支持：disk、redis、rocketmq、kafka等多种驱动。请自行选择适合你的驱动使用。

### 配置文件
- 配置文件：server/manifest/config/config.yaml

```yaml
#消息队列
queue:
  switch: true                                        # 队列开关，可选：true|false，默认为true
  driver: "disk"                                      # 队列驱动，可选：redis|rocketmq|kafka，默认为disk
  retry: 2                                            # 重试次数，仅rocketmq|redis支持
  groupName: "hotgo"                                  # mq群组名称
  #磁盘队列
  disk:
    path: "./storage/diskqueue"                       # 数据存放路径
    batchSize: 100                                    # 每100条消息同步一次，batchSize和batchTime满足其一就会同步一次
    batchTime: 1                                      # 每1秒消息同步一次
    segmentSize: 10485760                             # 每个topic分片数据文件最大字节，默认10M
    segmentLimit: 3000                                # 每个topic最大分片数据文件数量，超出部分将会丢弃
  redis:
    address: "127.0.0.1:6379"                         # redis服务地址，默认为127.0.0.1:6379
    db: 2                                             # 指定redis库
    pass: ""                                          # redis密码
    timeout: 0                                        # 队列超时时间(s) ，0为永不超时，当队列一直没有被消费到达超时时间则队列会被销毁
  rocketmq:
    address: "127.0.0.1:9876"                         # brocker地址+端口
    logLevel: "all"                                   # 系统日志级别，可选：all|close|debug|info|warn|error|fatal
  kafka:
    address: "127.0.0.1:9092"                         # kafka地址+端口
    version: "2.0.0.0"                                # kafka专属配置，默认2.0.0.0
    randClient: true                                  # 开启随机生成clientID，可以实现启动多实例同时一起消费相同topic，加速消费能力的特性，默认为true
    multiConsumer: true                               # 是否支持创建多个消费者

```

### 实现接口
- 为了提供高度的扩展性，消费队列在设计上采用了接口化的思路。只需要实现以下接口，您就可以在任何地方注册和使用消费队列消费功能，从而实现更大的灵活性和可扩展性。

```go
// Consumer 消费者接口，实现该接口即可加入到消费队列中
type Consumer interface {
    GetTopic() string                                    // 获取消费主题
    Handle(ctx context.Context, mqMsg MqMsg) (err error) // 处理消息的方法
}
```


### 一个例子

每个被发送到队列的消息应该被定义为一个单独的文件结构。

例如，如果您需要异步记录系统日志，内容大致如下：

- 文件路径：server/internal/queues/sys_log.go

```go 
package queues

import (
	"context"
	"encoding/json"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
)

func init() {
	queue.RegisterConsumer(SysLog)
}

// SysLog 系统日志
var SysLog = &qSysLog{}

type qSysLog struct{}

// GetTopic 主题
func (q *qSysLog) GetTopic() string {
	return consts.QueueLogTopic
}

// Handle 处理消息
func (q *qSysLog) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	var data entity.SysLog
	if err = json.Unmarshal(mqMsg.Body, &data); err != nil {
		return err
	}
	return service.SysLog().RealWrite(ctx, data)
}

```

下面是将消息添加到队列的方式，大概内容如下:

```go
package main

import (
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
)

func test()  {
	data := &entity.SysLog{
		//...
    }
	if err := queue.Push(consts.QueueLogTopic, data); err != nil {
		fmt.Printf("queue.Push err:%+v", err)
	}
}

```

延迟队列，目前只有redis驱动支持:

```go
package main

import (
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
)

func test()  {
	data := &entity.SysLog{
		//...
    }
	
	// 延迟10秒
	if err := queue.SendDelayMsg(consts.QueueLogTopic, data, 10); err != nil {
		fmt.Printf("queue.Push err:%+v", err)
	}
}

```

### 控制台

控制台用于处理队列消息，即消费者。

相关命令请参考： [控制台](sys-console.md)


### 自定义队列驱动

只需实现消息队列的生成者和消费者接口，并加入到初始化中进行相应调用即可。

- 接口片段：server/internal/library/queue/init.go

```go
package queue

import (
	"time"
)

type MqMsg struct {
	RunType   int       `json:"run_type"`
	Topic     string    `json:"topic"`
	MsgId     string    `json:"msg_id"`
	Offset    int64     `json:"offset"`
	Partition int32     `json:"partition"`
	Timestamp time.Time `json:"timestamp"`
	Body      []byte    `json:"body"`
}


type MqProducer interface {
	SendMsg(topic string, body string) (mqMsg MqMsg, err error)
	SendByteMsg(topic string, body []byte) (mqMsg MqMsg, err error)
}

type MqConsumer interface {
	ListenReceiveMsgDo(topic string, receiveDo func(mqMsg MqMsg)) (err error)
}

```

将实现过接口（MqProducer、MqConsumer）的实例方法分别加入到NewProducer、NewConsumer中进行相应调用即可。

