// Package optiontreedemo
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package optiontreedemo

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/utility/tree"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询选项树表列表
type ListReq struct {
	g.Meta `path:"/optionTreeDemo/list" method:"get" tags:"选项树表" summary:"获取选项树表列表"`
	sysin.OptionTreeDemoListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.OptionTreeDemoListModel `json:"list"   dc:"数据列表"`
}

// ViewReq 获取选项树表指定信息
type ViewReq struct {
	g.Meta `path:"/optionTreeDemo/view" method:"get" tags:"选项树表" summary:"获取选项树表指定信息"`
	sysin.OptionTreeDemoViewInp
}

type ViewRes struct {
	*sysin.OptionTreeDemoViewModel
}

// EditReq 修改/新增选项树表
type EditReq struct {
	g.Meta `path:"/optionTreeDemo/edit" method:"post" tags:"选项树表" summary:"修改/新增选项树表"`
	sysin.OptionTreeDemoEditInp
}

type EditRes struct{}

// DeleteReq 删除选项树表
type DeleteReq struct {
	g.Meta `path:"/optionTreeDemo/delete" method:"post" tags:"选项树表" summary:"删除选项树表"`
	sysin.OptionTreeDemoDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取选项树表最大排序
type MaxSortReq struct {
	g.Meta `path:"/optionTreeDemo/maxSort" method:"get" tags:"选项树表" summary:"获取选项树表最大排序"`
	sysin.OptionTreeDemoMaxSortInp
}

type MaxSortRes struct {
	*sysin.OptionTreeDemoMaxSortModel
}

// TreeOptionReq 获取选项树表关系树选项
type TreeOptionReq struct {
	g.Meta `path:"/optionTreeDemo/treeOption" method:"get" tags:"选项树表" summary:"获取选项树表关系树选项"`
}

type TreeOptionRes []tree.Node