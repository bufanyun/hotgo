// Package attachment
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package attachment

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询附件列表
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

// ViewReq 获取附件信息
type ViewReq struct {
	Id     int64 `json:"id" v:"required#附件ID不能为空" dc:"附件ID"`
	g.Meta `path:"/attachment/view" method:"get" tags:"附件" summary:"获取指定附件信息"`
}
type ViewRes struct {
	*sysin.AttachmentViewModel
}

// DeleteReq 删除附件
type DeleteReq struct {
	Id     interface{} `json:"id" v:"required#附件ID不能为空" dc:"附件ID"`
	g.Meta `path:"/attachment/delete" method:"post" tags:"附件" summary:"删除附件"`
}
type DeleteRes struct{}
