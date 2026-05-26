// Package websocket
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package websocket

import "github.com/gogf/gf/v2/frame/g"

// WRequest 输入对象
type WRequest struct {
	Event string `json:"event"` // 事件名称
	Data  g.Map  `json:"data"`  // 数据
}

// WResponse 输出对象
type WResponse struct {
	Event     string      `json:"event"`              // 事件名称
	Data      interface{} `json:"data,omitempty"`     // 数据
	Code      int         `json:"code"`               // 状态码
	ErrorMsg  string      `json:"errorMsg,omitempty"` // 错误消息
	Timestamp int64       `json:"timestamp"`          // 服务器时间
}

type TagWResponse struct {
	Tag       string
	WResponse *WResponse
}

type UserWResponse struct {
	UserID    int64
	WResponse *WResponse
}

type ClientWResponse struct {
	ID        string
	WResponse *WResponse
}

// EventHandler 消息处理器
type EventHandler func(client *Client, req *WRequest)

type EventHandlers map[string]EventHandler
