// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// AttachmentDeleteInp 删除附件
type AttachmentDeleteInp struct {
	Id interface{}
}
type AttachmentDeleteModel struct{}

// AttachmentViewInp 获取附件信息
type AttachmentViewInp struct {
	Id int64
}

type AttachmentViewModel struct {
	entity.SysAttachment
}

// AttachmentListInp 获取附件列表
type AttachmentListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	MemberId int64
	Drive    string
}

type AttachmentListModel struct {
	entity.SysAttachment
	SizeFormat string `json:"sizeFormat"      description:"长度"`
}
