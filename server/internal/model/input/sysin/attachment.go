// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"
)

// AttachmentDeleteInp 删除附件
type AttachmentDeleteInp struct {
	Id interface{} `json:"id" v:"required#附件ID不能为空" dc:"附件ID"`
}

type AttachmentDeleteModel struct{}

// AttachmentViewInp 获取附件信息
type AttachmentViewInp struct {
	Id int64 `json:"id" v:"required#附件ID不能为空" dc:"附件ID"`
}

type AttachmentViewModel struct {
	entity.SysAttachment
}

// AttachmentClearKindInp 清空上传类型
type AttachmentClearKindInp struct {
	Kind string `json:"kind" v:"required#上传类型不能为空" dc:"上传类型"`
}

func (in *AttachmentClearKindInp) Filter(ctx context.Context) (err error) {
	if !validate.InSlice(storager.KindSlice, in.Kind) {
		err = gerror.New("上传类型是无效的")
		return
	}
	return
}

// AttachmentListInp 获取附件列表
type AttachmentListInp struct {
	form.PageReq
	form.StatusReq
	MemberId  int64         `json:"member_id"  dc:"用户ID"`
	Name      string        `json:"name"       dc:"文件名称"`
	Drive     string        `json:"drive"      dc:"驱动"`
	Kind      string        `json:"kind"       dc:"上传类型"`
	UpdatedAt []*gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

type AttachmentListModel struct {
	entity.SysAttachment
	SizeFormat string `json:"sizeFormat"      dc:"大小"`
}

// AttachmentChooserListInp 获取附件列表
type AttachmentChooserListInp struct {
	form.PageReq
	Drive     string  `json:"drive"      dc:"驱动"`
	Kind      string  `json:"kind"       dc:"上传类型"`
	UpdatedAt []int64 `json:"updatedAt"  dc:"更新时间"`
}

type AttachmentChooserListModel struct {
	entity.SysAttachment
	SizeFormat string `json:"sizeFormat"      dc:"大小"`
}

// AttachmentClearInp 清空分类
type AttachmentClearInp struct {
	Kind string `json:"kind"       dc:"上传类型"`
}
