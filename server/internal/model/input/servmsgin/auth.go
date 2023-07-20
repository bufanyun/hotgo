package servmsgin

import "github.com/gogf/gf/v2/os/gtime"

// AuthSummaryModel 授权信息
type AuthSummaryModel struct {
	EndAt  *gtime.Time `json:"end_at" description:"授权过期时间"`
	Online int         `json:"online" description:"在线人数"`
	// 请填充你的授权数据
	// ...
}
