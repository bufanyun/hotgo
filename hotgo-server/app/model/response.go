//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package model

//  HTTP响应
type Response struct {
	Code      int         `json:"code" example:"0" description:"状态码"`
	Message   string      `json:"message" example:"操作成功" description:"提示消息"`
	Data      interface{} `json:"data" description:"数据集"`
	Error     interface{} `json:"error" description:"错误信息"`
	Timestamp int64       `json:"timestamp" example:"1640966400" description:"服务器时间戳"`
	ReqId     string      `json:"req_id" v:"0" example:"d0bb93048bc5c9164cdee845dcb7f820" description:"唯一请求ID"`
}
