package servmsg

import (
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/input/servmsgin"
)

// AuthSummaryReq 授权信息
type AuthSummaryReq struct {
}

// AuthSummaryRes 响应授权信息
type AuthSummaryRes struct {
	tcp.ServerRes
	Data *servmsgin.AuthSummaryModel `json:"data,omitempty" description:"数据集"`
}
