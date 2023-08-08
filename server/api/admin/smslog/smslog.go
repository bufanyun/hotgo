// Package smslog
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package smslog

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/smsLog/list" method:"get" tags:"短信记录" summary:"获取短信记录列表"`
	sysin.SmsLogListInp
}

type ListRes struct {
	List []*sysin.SmsLogListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取指定信息
type ViewReq struct {
	g.Meta `path:"/smsLog/view" method:"get" tags:"短信记录" summary:"获取指定信息"`
	sysin.SmsLogViewInp
}

type ViewRes struct {
	*sysin.SmsLogViewModel
}

// EditReq 修改/新增数据
type EditReq struct {
	g.Meta `path:"/smsLog/edit" method:"post" tags:"短信记录" summary:"修改/新增短信记录"`
	sysin.SmsLogEditInp
}

type EditRes struct{}

// DeleteReq 删除
type DeleteReq struct {
	g.Meta `path:"/smsLog/delete" method:"post" tags:"短信记录" summary:"删除短信记录"`
	sysin.SmsLogDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新状态
type StatusReq struct {
	g.Meta `path:"/smsLog/status" method:"post" tags:"短信记录" summary:"更新短信记录状态"`
	sysin.SmsLogStatusInp
}

type StatusRes struct{}
