// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// AttachmentMaxSortInp 最大排序
type AttachmentMaxSortInp struct {
	Id int64
}

type AttachmentMaxSortModel struct {
	Sort int
}

// AttachmentEditInp 修改/新增字典数据
type AttachmentEditInp struct {
	entity.SysAttachment
}
type AttachmentEditModel struct{}

// AttachmentDeleteInp 删除字典类型
type AttachmentDeleteInp struct {
	Id interface{}
}
type AttachmentDeleteModel struct{}

// AttachmentViewInp 获取信息
type AttachmentViewInp struct {
	Id int64
}

type AttachmentViewModel struct {
	entity.SysAttachment
}

// AttachmentListInp 获取列表
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

// AttachmentStatusInp 更新状态
type AttachmentStatusInp struct {
	entity.SysAttachment
}
type AttachmentStatusModel struct{}

type UploadFileMeta struct {
	Filename  string
	Size      int64
	Kind      string
	MetaType  string
	NaiveType string
	Ext       string
	Md5       string
}
