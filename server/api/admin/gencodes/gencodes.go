// Package gencodes
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package gencodes

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/genCodes/list" method:"get" tags:"生成代码" summary:"获取生成代码列表"`
	sysin.GenCodesListInp
}

type ListRes struct {
	List []*sysin.GenCodesListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取指定信息
type ViewReq struct {
	g.Meta `path:"/genCodes/view" method:"get" tags:"生成代码" summary:"获取指定信息"`
	sysin.GenCodesViewInp
}

type ViewRes struct {
	*sysin.GenCodesViewModel
}

// EditReq 修改/新增数据
type EditReq struct {
	g.Meta `path:"/genCodes/edit" method:"post" tags:"生成代码" summary:"修改/新增生成代码"`
	sysin.GenCodesEditInp
}

type EditRes struct {
	*sysin.GenCodesEditModel
}

// DeleteReq 删除
type DeleteReq struct {
	g.Meta `path:"/genCodes/delete" method:"post" tags:"生成代码" summary:"删除生成代码"`
	sysin.GenCodesDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	g.Meta `path:"/genCodes/maxSort" method:"get" tags:"生成代码" summary:"生成代码最大排序"`
	sysin.GenCodesMaxSortInp
}

type MaxSortRes struct {
	*sysin.GenCodesMaxSortModel
}

// StatusReq 更新状态
type StatusReq struct {
	g.Meta `path:"/genCodes/status" method:"post" tags:"生成代码" summary:"更新生成代码状态"`
	sysin.GenCodesStatusInp
}

type StatusRes struct{}

type SelectsReq struct {
	g.Meta `path:"/genCodes/selects" method:"get" tags:"生成代码" summary:"生成入口选项"`
	sysin.GenCodesSelectsInp
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
	sysin.GenCodesBuildInp
}

type BuildRes struct {
}
