// Package post
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package post

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// EditReq 修改/新增岗位
type EditReq struct {
	g.Meta `path:"/post/edit" method:"post" tags:"岗位" summary:"修改/新增岗位"`
	entity.AdminPost
}
type EditRes struct{}

// DeleteReq 删除岗位
type DeleteReq struct {
	g.Meta `path:"/post/delete" method:"post" tags:"岗位" summary:"删除岗位"`
	Id     interface{} `json:"id" v:"required#岗位ID不能为空" description:"岗位ID"`
}
type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	g.Meta `path:"/post/max_sort" method:"get" tags:"岗位" summary:"岗位最大排序"`
	Id     int64 `json:"id" description:"岗位ID"`
}
type MaxSortRes struct {
	Sort int `json:"sort" description:"排序"`
}

// ListReq 获取列表
type ListReq struct {
	g.Meta `path:"/post/list" method:"get" tags:"岗位" summary:"获取岗位列表"`
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name string `json:"name"   description:"岗位名称"`
	Code string `json:"code"   description:"岗位编码"`
}

type ListRes struct {
	List []*adminin.PostListModel `json:"list"   description:"数据列表"`
	form.PageRes
}

// ViewReq 获取指定信息
type ViewReq struct {
	g.Meta `path:"/post/view" method:"get" tags:"岗位" summary:"获取指定信息"`
	Id     string `json:"id" v:"required#岗位ID不能为空" description:"岗位ID"`
}
type ViewRes struct {
	*adminin.PostViewModel
}

// CodeUniqueReq 编码是否唯一
type CodeUniqueReq struct {
	g.Meta `path:"/post/code_unique" method:"get" tags:"岗位" summary:"岗位编码是否唯一"`
	Code   string `json:"code" v:"required#岗位编码不能为空"  description:"岗位编码"`
	Id     int64  `json:"id" description:"岗位ID"`
}
type CodeUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}

// NameUniqueReq 名称是否唯一
type NameUniqueReq struct {
	g.Meta `path:"/post/name_unique" method:"get" tags:"岗位" summary:"岗位名称是否唯一"`
	Name   string `json:"name" v:"required#岗位名称不能为空"  description:"岗位名称"`
	Id     int64  `json:"id" description:"岗位ID"`
}
type NameUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}

// StatusReq 更新状态
type StatusReq struct {
	entity.AdminPost
	g.Meta `path:"/post/status" method:"post" tags:"岗位" summary:"更新岗位状态"`
}
type StatusRes struct{}
