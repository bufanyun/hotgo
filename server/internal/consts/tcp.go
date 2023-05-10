package consts

import "time"

const (
	TCPMsgCodeSuccess = 2000 // 成功的状态码
)

// 定时任务
const (
	TCPCronHeartbeatVerify = "tcpHeartbeatVerify"
	TCPCronHeartbeat       = "tcpHeartbeat"
	TCPCronAuthVerify      = "tcpAuthVerify"
)

// 认证分组
const (
	TCPClientGroupCron  = "cron"  // 定时任务
	TCPClientGroupQueue = "queue" // 消息队列
	TCPClientGroupAuth  = "auth"  // 服务授权
)

const (
	TCPRpcTimeout = time.Second * 10 // rpc通讯超时时间， 默认10s
)
