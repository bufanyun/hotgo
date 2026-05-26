// Package normaltreedemo
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.15.7
package normaltreedemo

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/utility/tree"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询普通树表列表
type ListReq struct {
	g.Meta `path:"/normalTreeDemo/list" method:"get" tags:"普通树表" summary:"获取普通树表列表"`
	sysin.NormalTreeDemoListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.NormalTreeDemoListModel `json:"list"   dc:"数据列表"`
}

// ViewReq 获取普通树表指定信息
type ViewReq struct {
	g.Meta `path:"/normalTreeDemo/view" method:"get" tags:"普通树表" summary:"获取普通树表指定信息"`
	sysin.NormalTreeDemoViewInp
}

type ViewRes struct {
	*sysin.NormalTreeDemoViewModel
}

// EditReq 修改/新增普通树表
type EditReq struct {
	g.Meta `path:"/normalTreeDemo/edit" method:"post" tags:"普通树表" summary:"修改/新增普通树表"`
	sysin.NormalTreeDemoEditInp
}

type EditRes struct{}

// DeleteReq 删除普通树表
type DeleteReq struct {
	g.Meta `path:"/normalTreeDemo/delete" method:"post" tags:"普通树表" summary:"删除普通树表"`
	sysin.NormalTreeDemoDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取普通树表最大排序
type MaxSortReq struct {
	g.Meta `path:"/normalTreeDemo/maxSort" method:"get" tags:"普通树表" summary:"获取普通树表最大排序"`
	sysin.NormalTreeDemoMaxSortInp
}

type MaxSortRes struct {
	*sysin.NormalTreeDemoMaxSortModel
}

// TreeOptionReq 获取普通树表关系树选项
type TreeOptionReq struct {
	g.Meta `path:"/normalTreeDemo/treeOption" method:"get" tags:"普通树表" summary:"获取普通树表关系树选项"`
}

type TreeOptionRes []tree.Node
