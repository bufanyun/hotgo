// Package attachment
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package attachment

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询列表
type ListReq struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	MemberId int64  `json:"member_id"`
	Drive    string `json:"drive"`
	g.Meta   `path:"/attachment/list" method:"get" tags:"附件" summary:"获取附件列表"`
}

type ListRes struct {
	List []*sysin.AttachmentListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取信息
type ViewReq struct {
	Id     int64 `json:"id" v:"required#附件ID不能为空" dc:"附件ID"`
	g.Meta `path:"/attachment/view" method:"get" tags:"附件" summary:"获取指定信息"`
}
type ViewRes struct {
	*sysin.AttachmentViewModel
}

// EditReq 修改/新增
type EditReq struct {
	entity.SysAttachment
	g.Meta `path:"/attachment/edit" method:"post" tags:"附件" summary:"修改/新增附件"`
}
type EditRes struct{}

// DeleteReq 删除
type DeleteReq struct {
	Id     interface{} `json:"id" v:"required#附件ID不能为空" dc:"附件ID"`
	g.Meta `path:"/attachment/delete" method:"post" tags:"附件" summary:"删除附件"`
}
type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	Id     int64 `json:"id" dc:"附件ID"`
	g.Meta `path:"/attachment/maxSort" method:"get" tags:"附件" summary:"附件最大排序"`
}
type MaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

// StatusReq 更新状态
type StatusReq struct {
	entity.SysAttachment
	g.Meta `path:"/attachment/status" method:"post" tags:"附件" summary:"更新附件状态"`
}
type StatusRes struct{}
