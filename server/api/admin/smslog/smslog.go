package smslog

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
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
	Id     int64 `json:"id" v:"required#短信记录ID不能为空" dc:"短信记录ID"`
	g.Meta `path:"/smsLog/view" method:"get" tags:"短信记录" summary:"获取指定信息"`
}
type ViewRes struct {
	*sysin.SmsLogViewModel
}

// EditReq 修改/新增数据
type EditReq struct {
	entity.SysSmsLog
	g.Meta `path:"/smsLog/edit" method:"post" tags:"短信记录" summary:"修改/新增短信记录"`
}
type EditRes struct{}

// DeleteReq 删除
type DeleteReq struct {
	Id     interface{} `json:"id" v:"required#短信记录ID不能为空" dc:"短信记录ID"`
	g.Meta `path:"/smsLog/delete" method:"post" tags:"短信记录" summary:"删除短信记录"`
}
type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	Id     int64 `json:"id" dc:"短信记录ID"`
	g.Meta `path:"/smsLog/maxSort" method:"get" tags:"短信记录" summary:"短信记录最大排序"`
}
type MaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

// StatusReq 更新状态
type StatusReq struct {
	entity.SysSmsLog
	g.Meta `path:"/smsLog/status" method:"post" tags:"短信记录" summary:"更新短信记录状态"`
}
type StatusRes struct{}

// SendTestReq 更新状态
type SendTestReq struct {
	entity.SysSmsLog
	g.Meta `path:"/smsLog/sendTest" method:"post" tags:"短信记录" summary:"发送测试短信"`
}
type SendTestRes struct{}
