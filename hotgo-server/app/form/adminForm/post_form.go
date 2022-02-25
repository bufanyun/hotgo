package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

//  修改/新增字典数据
type PostEditReq struct {
	entity.AdminPost
	g.Meta `path:"/post/edit" method:"post" tags:"岗位" summary:"修改/新增岗位"`
}
type PostEditRes struct{}

//  删除字典类型
type PostDeleteReq struct {
	Id     interface{} `json:"id" v:"required#岗位ID不能为空" description:"岗位ID"`
	g.Meta `path:"/post/delete" method:"post" tags:"岗位" summary:"删除岗位"`
}
type PostDeleteRes struct{}

//  最大排序
type PostMaxSortReq struct {
	Id     int64 `json:"id" description:"岗位ID"`
	g.Meta `path:"/post/max_sort" method:"get" tags:"岗位" summary:"岗位最大排序"`
}
type PostMaxSortRes struct {
	Sort int `json:"sort" description:"排序"`
}

//  获取列表
type PostListReq struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name   string `json:"name"   description:"岗位名称"`
	Code   string `json:"code"   description:"岗位编码"`
	g.Meta `path:"/post/list" method:"get" tags:"岗位" summary:"获取岗位列表"`
}

type PostListRes struct {
	List []*input.AdminPostListModel `json:"list"   description:"数据列表"`
	form.PageRes
}

//  获取指定信息
type PostViewReq struct {
	Id     string `json:"id" v:"required#岗位ID不能为空" description:"岗位ID"`
	g.Meta `path:"/post/view" method:"get" tags:"岗位" summary:"获取指定信息"`
}
type PostViewRes struct {
	*input.AdminPostViewModel
}

//  编码是否唯一
type PostCodeUniqueReq struct {
	Code   string `json:"code" v:"required#岗位编码不能为空"  description:"岗位编码"`
	Id     int64  `json:"id" description:"岗位ID"`
	g.Meta `path:"/post/code_unique" method:"get" tags:"岗位" summary:"岗位编码是否唯一"`
}
type PostCodeUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}

//  名称是否唯一
type PostNameUniqueReq struct {
	Name   string `json:"name" v:"required#岗位名称不能为空"  description:"岗位名称"`
	Id     int64  `json:"id" description:"岗位ID"`
	g.Meta `path:"/post/name_unique" method:"get" tags:"岗位" summary:"岗位名称是否唯一"`
}
type PostNameUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}
