## 控制台

目录

- 启动所有服务
- HTTP服务
- 消息队列
- 定时任务
- 常用工具
- Makefile

### 启动所有服务
- 仅推荐在开发期间快速调试使用，线上实际部署时建议将各个服务分开部署，这样重新部署某个服务时无需全部重启。

```shell
# 默认
go run main.go

# 通过热编译启动
gf run main.go
```

### HTTP服务
- 启动HTTP服务，包含websocket。
```shell
# 默认
go run main.go http

# 通过热编译启动
gf run main.go --args "http"
```

### 消息队列
- 启动消息队列的消费者。

```shell
# 默认
go run main.go queue

# 通过热编译启动
gf run main.go --args "queue"
```

### 定时任务（暂未拆分，目前随HTTP服务启动）
- 启动系统中统一注册的定时任务。

```shell
# 默认
go run main.go cron

# 通过热编译启动
gf run main.go --args "cron"
```


### 常用工具
- 释放casbin权限，用于清理无效的权限设置。
```shell
go run main.go tools -m=casbin -a1=refresh
```


### Makefile
- 通过make提供一些快捷命令
```shell
# 一键编译，打包前后端代码到可执行文件
make build

# 更多请查看 /server/Makefile文件
```