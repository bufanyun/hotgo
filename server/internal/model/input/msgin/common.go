package msgin

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/utility/encrypt"
)

type RpcMsg struct {
	AppId     string `json:"appID" v:"0" example:"10001"                               description:"应用ID"`
	TraceID   string `json:"traceID" v:"0" example:"d0bb93048bc5c9164cdee845dcb7f820"  description:"链路ID"`
	Timestamp int64  `json:"timestamp" example:"1640966400"                            description:"服务器时间戳"`
	Sign      string `json:"sign" example:"d0bb93048bc5c9164cdee845dcb7f820"           description:"签名"`
}

func (i *RpcMsg) SetSign(appId, secretKey string) *RpcMsg {
	i.AppId = appId
	i.Timestamp = gtime.Timestamp()
	i.Sign = i.GetSign(secretKey)
	return i
}

func (i *RpcMsg) GetSign(secretKey string) string {
	return encrypt.Md5ToString(fmt.Sprintf("%v%v%v%v", i.AppId, i.TraceID, i.Timestamp, secretKey))
}

func (i *RpcMsg) GetTraceID() string {
	return i.TraceID
}

func (i *RpcMsg) SetTraceID(traceID string) {
	i.TraceID = traceID
}

type Response struct {
	RpcMsg
	Code    int    `json:"code" example:"2000"                    description:"状态码"`
	Message string `json:"message,omitempty" example:"操作成功"     description:"提示消息"`
}

// PkgResponse 打包响应消息
func (m *Response) PkgResponse() {
	m.SetCode()
	// ...
}

// SetCode 设置状态码
func (m *Response) SetCode(code ...int) {
	if len(code) > 0 {
		m.Code = code[0]
		return
	}

	// 默认值，转为成功的状态码
	if m.Code == 0 {
		m.Code = consts.TCPMsgCodeSuccess
	}
}

// GetError 获取响应中的错误
func (m *Response) GetError() (err error) {
	if m.Code != consts.TCPMsgCodeSuccess {
		if m.Message == "" {
			m.Message = "操作失败"
		}
		err = gerror.New(m.Message)
	}
	return
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
	RpcMsg
	Group string `json:"group"   description:"分组"`
	Name  string `json:"name"    description:"名称"`
}

// ResponseServerLogin 响应服务登录
type ResponseServerLogin struct {
	Response
}

// ServerOffline 服务离线
type ServerOffline struct {
	RpcMsg
}

// ResponseServerOffline 响应服务离线
type ResponseServerOffline struct {
	Response
}
