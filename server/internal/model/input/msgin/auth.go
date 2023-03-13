package msgin

import "github.com/gogf/gf/v2/os/gtime"

// AuthSummary 授权摘要
type AuthSummary struct {
	Request
}

// ResponseAuthSummary 响应授权摘要
type ResponseAuthSummary struct {
	Response
	Data *AuthSummaryData `json:"data,omitempty" description:"数据集"`
}

type AuthSummaryData struct {
	EndAt  *gtime.Time `json:"end_at" description:"授权过期时间"`
	Online int         `json:"online" description:"在线人数"`
}
