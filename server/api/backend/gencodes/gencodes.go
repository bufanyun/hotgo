// Package hggen
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package gencodes

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta  `path:"/genCodes/list" method:"get" tags:"生成代码" summary:"获取生成代码列表"`
	sysin.GenCodesListInp
}

type ListRes struct {
	List []*sysin.GenCodesListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取指定信息
type ViewReq struct {
	Id     int64 `json:"id" v:"required#生成代码ID不能为空" dc:"生成代码ID"`
	g.Meta `path:"/genCodes/view" method:"get" tags:"生成代码" summary:"获取指定信息"`
}
type ViewRes struct {
	*sysin.GenCodesViewModel
}

// EditReq 修改/新增数据
type EditReq struct {
	entity.SysGenCodes
	g.Meta `path:"/genCodes/edit" method:"post" tags:"生成代码" summary:"修改/新增生成代码"`
}
type EditRes struct {
	*sysin.GenCodesEditModel
}

// DeleteReq 删除
type DeleteReq struct {
	Id     interface{} `json:"id" v:"required#生成代码ID不能为空" dc:"生成代码ID"`
	g.Meta `path:"/genCodes/delete" method:"post" tags:"生成代码" summary:"删除生成代码"`
}
type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	Id     int64 `json:"id" dc:"生成代码ID"`
	g.Meta `path:"/genCodes/max_sort" method:"get" tags:"生成代码" summary:"生成代码最大排序"`
}
type MaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

// StatusReq 更新状态
type StatusReq struct {
	entity.SysGenCodes
	g.Meta `path:"/genCodes/status" method:"post" tags:"生成代码" summary:"更新生成代码状态"`
}
type StatusRes struct{}

type SelectsReq struct {
	g.Meta `path:"/genCodes/selects" method:"get" tags:"生成代码" summary:"生成入口选项"`
}
type SelectsRes struct {
	*sysin.GenCodesSelectsModel
}

type TableSelectReq struct {
	g.Meta `path:"/genCodes/tableSelect" method:"get" tags:"生成代码" summary:"数据库表选项"`
	sysin.GenCodesTableSelectInp
}
type TableSelectRes []*sysin.GenCodesTableSelectModel

type ColumnSelectReq struct {
	g.Meta `path:"/genCodes/columnSelect" method:"get" tags:"生成代码" summary:"表字段选项"`
	sysin.GenCodesColumnSelectInp
}
type ColumnSelectRes []*sysin.GenCodesColumnSelectModel

type ColumnListReq struct {
	g.Meta `path:"/genCodes/columnList" method:"get" tags:"生成代码" summary:"表字段列表"`
	sysin.GenCodesColumnListInp
}
type ColumnListRes []*sysin.GenCodesColumnListModel

// PreviewReq 生成预览
type PreviewReq struct {
	g.Meta `path:"/genCodes/preview" method:"post" tags:"生成代码" summary:"生成预览"`
	sysin.GenCodesPreviewInp
}
type PreviewRes struct {
	*sysin.GenCodesPreviewModel
}

// BuildReq 提交生成
type BuildReq struct {
	g.Meta `path:"/genCodes/build" method:"post" tags:"生成代码" summary:"提交生成"`
	sysin.GenCodesPreviewInp
}
type BuildRes struct {
}
