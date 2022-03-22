//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminForm

import (
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

//  名称是否唯一
type DeptNameUniqueReq struct {
	Name   string `json:"name" v:"required#部门名称不能为空"  dc:"部门名称"`
	Id     int64  `json:"id" dc:"部门ID"`
	g.Meta `path:"/dept/name_unique" method:"get" tags:"部门" summary:"部门名称是否唯一"`
}
type DeptNameUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

//  查询列表树
type DeptListTreeReq struct {
	Id     int64 `json:"id" dc:"部门ID"`
	g.Meta `path:"/dept/list_tree" method:"get" tags:"部门" summary:"获取部门列表树"`
}

type DeptListTreeRes []*input.AdminDeptListTreeModel

//  查询列表
type DeptListReq struct {
	Name   string `json:"name" dc:"部门名称"`
	g.Meta `path:"/dept/list" method:"get" tags:"部门" summary:"获取部门列表"`
}

type DeptListRes []*input.AdminDeptListModel

//  获取指定信息
type DeptViewReq struct {
	Id     int64 `json:"id" v:"required#部门ID不能为空" dc:"部门ID"`
	g.Meta `path:"/dept/view" method:"get" tags:"部门" summary:"获取指定信息"`
}
type DeptViewRes struct {
	*input.AdminDeptViewModel
}

//  修改/新增字典数据
type DeptEditReq struct {
	entity.AdminDept
	g.Meta `path:"/dept/edit" method:"post" tags:"部门" summary:"修改/新增部门"`
}
type DeptEditRes struct{}

//  删除字典类型
type DeptDeleteReq struct {
	Id     interface{} `json:"id" v:"required#部门ID不能为空" dc:"部门ID"`
	g.Meta `path:"/dept/delete" method:"post" tags:"部门" summary:"删除部门"`
}
type DeptDeleteRes struct{}

//  最大排序
type DeptMaxSortReq struct {
	Id     int64 `json:"id" dc:"部门ID"`
	g.Meta `path:"/dept/max_sort" method:"get" tags:"部门" summary:"部门最大排序"`
}
type DeptMaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}
