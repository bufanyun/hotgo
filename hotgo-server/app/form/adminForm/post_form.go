//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

//  修改/新增字典数据
type PostEditReq struct {
	g.Meta `path:"/post/edit" method:"post" tags:"岗位" summary:"修改/新增岗位"`
	entity.AdminPost
}
type PostEditRes struct{}

//  删除字典类型
type PostDeleteReq struct {
	g.Meta `path:"/post/delete" method:"post" tags:"岗位" summary:"删除岗位"`
	Id     interface{} `json:"id" v:"required#岗位ID不能为空" description:"岗位ID"`
}
type PostDeleteRes struct{}

//  最大排序
type PostMaxSortReq struct {
	g.Meta `path:"/post/max_sort" method:"get" tags:"岗位" summary:"岗位最大排序"`
	Id     int64 `json:"id" description:"岗位ID"`
}
type PostMaxSortRes struct {
	Sort int `json:"sort" description:"排序"`
}

//  获取列表
type PostListReq struct {
	g.Meta `path:"/post/list" method:"get" tags:"岗位" summary:"获取岗位列表"`
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name string `json:"name"   description:"岗位名称"`
	Code string `json:"code"   description:"岗位编码"`
}

type PostListRes struct {
	List []*input.AdminPostListModel `json:"list"   description:"数据列表"`
	form.PageRes
}

//  获取指定信息
type PostViewReq struct {
	g.Meta `path:"/post/view" method:"get" tags:"岗位" summary:"获取指定信息"`
	Id     string `json:"id" v:"required#岗位ID不能为空" description:"岗位ID"`
}
type PostViewRes struct {
	*input.AdminPostViewModel
}

//  编码是否唯一
type PostCodeUniqueReq struct {
	g.Meta `path:"/post/code_unique" method:"get" tags:"岗位" summary:"岗位编码是否唯一"`
	Code   string `json:"code" v:"required#岗位编码不能为空"  description:"岗位编码"`
	Id     int64  `json:"id" description:"岗位ID"`
}
type PostCodeUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}

//  名称是否唯一
type PostNameUniqueReq struct {
	g.Meta `path:"/post/name_unique" method:"get" tags:"岗位" summary:"岗位名称是否唯一"`
	Name   string `json:"name" v:"required#岗位名称不能为空"  description:"岗位名称"`
	Id     int64  `json:"id" description:"岗位ID"`
}
type PostNameUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}
