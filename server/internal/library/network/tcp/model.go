package tcp

// 定时任务
const (
	cronHeartbeatVerify = "tcpHeartbeatVerify"
	cronHeartbeat       = "tcpHeartbeat"
)

// 认证分组
const (
	ClientGroupCron  = "cron"  // 定时任务
	ClientGroupQueue = "queue" // 消息队列
	ClientGroupAuth  = "auth"  // 服务授权
)

// AuthMeta 认证元数据
type AuthMeta struct {
	Group     string `json:"group"`
	Name      string `json:"name"`
	AppId     string `json:"appId"`
	SecretKey string `json:"secretKey"`
}

// CallbackEvent 回调事件
type CallbackEvent func()
