package msgin

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/utility/encrypt"
)

type Request struct {
	AppId     string `json:"appID" v:"0" example:"d0bb93048bc5c9164cdee845dcb7f820" description:"应用ID"`
	TraceID   string `json:"traceID" v:"0" example:"d0bb93048bc5c9164cdee845dcb7f820" description:"链路ID"`
	Timestamp int64  `json:"timestamp" example:"1640966400" description:"服务器时间戳"`
	Sign      string `json:"sign" example:"d0bb93048bc5c9164cdee845dcb7f820" description:"签名"`
}

func (i *Request) SetSign(traceID, appId, secretKey string) {
	i.AppId = appId
	i.TraceID = traceID
	i.Timestamp = gtime.Timestamp()
	i.Sign = i.GetSign(secretKey)
}

func (i *Request) GetSign(secretKey string) string {
	return encrypt.Md5ToString(fmt.Sprintf("%s%s%s%s", i.AppId, i.TraceID, i.Timestamp, secretKey))
}

type Response struct {
	Code    int    `json:"code" example:"0" description:"状态码"`
	Message string `json:"message,omitempty" example:"操作成功" description:"提示消息"`
	//Data    interface{} `json:"data,omitempty" description:"数据集"`
}

// ServerHeartbeat 心跳
type ServerHeartbeat struct {
}

// ResponseServerHeartbeat 响应心跳
type ResponseServerHeartbeat struct {
	Response
}

// ServerLogin 服务登录
type ServerLogin struct {
	Request
	Group string
	Name  string
}

// ResponseServerLogin 响应服务登录
type ResponseServerLogin struct {
	Response
}

// ServerOffline 服务离线
type ServerOffline struct {
}

// ResponseServerOffline 响应服务离线
type ResponseServerOffline struct {
	Response
}
