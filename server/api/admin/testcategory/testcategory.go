// Package testcategory
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.15.7
package testcategory

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询测试分类列表
type ListReq struct {
	g.Meta `path:"/testCategory/list" method:"get" tags:"测试分类" summary:"获取测试分类列表"`
	sysin.TestCategoryListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.TestCategoryListModel `json:"list"   dc:"数据列表"`
}

// ViewReq 获取测试分类指定信息
type ViewReq struct {
	g.Meta `path:"/testCategory/view" method:"get" tags:"测试分类" summary:"获取测试分类指定信息"`
	sysin.TestCategoryViewInp
}

type ViewRes struct {
	*sysin.TestCategoryViewModel
}

// EditReq 修改/新增测试分类
type EditReq struct {
	g.Meta `path:"/testCategory/edit" method:"post" tags:"测试分类" summary:"修改/新增测试分类"`
	sysin.TestCategoryEditInp
}

type EditRes struct{}

// DeleteReq 删除测试分类
type DeleteReq struct {
	g.Meta `path:"/testCategory/delete" method:"post" tags:"测试分类" summary:"删除测试分类"`
	sysin.TestCategoryDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取测试分类最大排序
type MaxSortReq struct {
	g.Meta `path:"/testCategory/maxSort" method:"get" tags:"测试分类" summary:"获取测试分类最大排序"`
	sysin.TestCategoryMaxSortInp
}

type MaxSortRes struct {
	*sysin.TestCategoryMaxSortModel
}

// StatusReq 更新测试分类状态
type StatusReq struct {
	g.Meta `path:"/testCategory/status" method:"post" tags:"测试分类" summary:"更新测试分类状态"`
	sysin.TestCategoryStatusInp
}

type StatusRes struct{}
