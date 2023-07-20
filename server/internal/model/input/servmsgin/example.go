package servmsgin

import "github.com/gogf/gf/v2/os/gtime"

// ExampleHelloModel 授权信息
type ExampleHelloModel struct {
	Desc      string      `json:"desc"  description:"描述信息"`
	Timestamp *gtime.Time `json:"timestamp"  description:"服务器时间"`
}
