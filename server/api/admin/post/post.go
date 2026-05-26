// Package post
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package post

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// EditReq 修改/新增岗位
type EditReq struct {
	g.Meta `path:"/post/edit" method:"post" tags:"岗位" summary:"修改/新增岗位"`
	adminin.PostEditInp
}

type EditRes struct{}

// DeleteReq 删除岗位
type DeleteReq struct {
	g.Meta `path:"/post/delete" method:"post" tags:"岗位" summary:"删除岗位"`
	adminin.PostDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	g.Meta `path:"/post/maxSort" method:"get" tags:"岗位" summary:"岗位最大排序"`
	adminin.PostMaxSortInp
}

type MaxSortRes struct {
	*adminin.PostMaxSortModel
}

// ListReq 获取列表
type ListReq struct {
	g.Meta `path:"/post/list" method:"get" tags:"岗位" summary:"获取岗位列表"`
	adminin.PostListInp
}

type ListRes struct {
	List []*adminin.PostListModel `json:"list"   description:"数据列表"`
	form.PageRes
}

// ViewReq 获取指定信息
type ViewReq struct {
	g.Meta `path:"/post/view" method:"get" tags:"岗位" summary:"获取指定信息"`
	adminin.PostViewInp
}

type ViewRes struct {
	*adminin.PostViewModel
}

// StatusReq 更新状态
type StatusReq struct {
	g.Meta `path:"/post/status" method:"post" tags:"岗位" summary:"更新岗位状态"`
	adminin.PostStatusInp
}

type StatusRes struct{}
