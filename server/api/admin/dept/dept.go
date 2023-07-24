// Package dept
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package dept

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/dept/list" method:"get" tags:"部门" summary:"获取部门列表"`
	adminin.DeptListInp
}

type ListRes *adminin.DeptListModel

// ViewReq 获取指定信息
type ViewReq struct {
	g.Meta `path:"/dept/view" method:"get" tags:"部门" summary:"获取指定信息"`
	adminin.DeptViewInp
}

type ViewRes struct {
	*adminin.DeptViewModel
}

// EditReq 修改/新增字典数据
type EditReq struct {
	g.Meta `path:"/dept/edit" method:"post" tags:"部门" summary:"修改/新增部门"`
	adminin.DeptEditInp
}

type EditRes struct{}

// DeleteReq 删除字典类型
type DeleteReq struct {
	g.Meta `path:"/dept/delete" method:"post" tags:"部门" summary:"删除部门"`
	adminin.DeptDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	g.Meta `path:"/dept/maxSort" method:"get" tags:"部门" summary:"部门最大排序"`
	adminin.DeptMaxSortInp
}

type MaxSortRes struct {
	*adminin.DeptMaxSortModel
}

// StatusReq 更新部门状态
type StatusReq struct {
	g.Meta `path:"/dept/status" method:"post" tags:"部门" summary:"更新部门状态"`
	adminin.DeptStatusInp
}

type StatusRes struct{}

// OptionReq 获取部门选项树
type OptionReq struct {
	g.Meta `path:"/dept/option" method:"get" tags:"部门" summary:"获取部门选项树"`
	adminin.DeptOptionInp
}

type OptionRes struct {
	*adminin.DeptOptionModel
	form.PageRes
}
