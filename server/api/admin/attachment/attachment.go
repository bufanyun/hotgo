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
	g.Meta `path:"/attachment/list" method:"get" tags:"附件" summary:"获取附件列表"`
	sysin.AttachmentListInp
}

type ListRes struct {
	List []*sysin.AttachmentListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取附件信息
type ViewReq struct {
	g.Meta `path:"/attachment/view" method:"get" tags:"附件" summary:"获取指定附件信息"`
	sysin.AttachmentViewInp
}

type ViewRes struct {
	*sysin.AttachmentViewModel
}

// DeleteReq 删除附件
type DeleteReq struct {
	g.Meta `path:"/attachment/delete" method:"post" tags:"附件" summary:"删除附件"`
	sysin.AttachmentDeleteInp
}

type DeleteRes struct{}

// ClearKindReq 清空上传类型
type ClearKindReq struct {
	g.Meta `path:"/attachment/clearKind" method:"post" tags:"附件" summary:"清空上传类型"`
	sysin.AttachmentClearKindInp
}

type ClearKindRes struct{}

// ChooserOptionReq 获取选择器选项
type ChooserOptionReq struct {
	g.Meta `path:"/attachment/chooserOption" method:"get" tags:"附件" summary:"获取选择器选项"`
}

type ChooserOptionRes struct {
	Drive sysin.DataSelectModel `json:"drive" dc:"驱动"`
	Kind  []KindSelect          `json:"kind"  dc:"上传类型"`
}

type KindSelect struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Tag   string `json:"listClass"`
	Label string `json:"label"`
	Icon  string `json:"icon"`
}
