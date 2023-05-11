// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/post"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
)

// Post 岗位
var Post = cPost{}

type cPost struct{}

// Delete 删除
func (c *cPost) Delete(ctx context.Context, req *post.DeleteReq) (res *post.DeleteRes, err error) {
	var in adminin.PostDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminPost().Delete(ctx, in)
	return
}

// Edit 修改/新增
func (c *cPost) Edit(ctx context.Context, req *post.EditReq) (res *post.EditRes, err error) {
	var in adminin.PostEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminPost().Edit(ctx, in)
	return
}

// MaxSort 最大排序
func (c *cPost) MaxSort(ctx context.Context, req *post.MaxSortReq) (res *post.MaxSortRes, err error) {
	data, err := service.AdminPost().MaxSort(ctx, adminin.PostMaxSortInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(post.MaxSortRes)
	res.Sort = data.Sort
	return
}

// NameUnique 名称是否唯一
func (c *cPost) NameUnique(ctx context.Context, req *post.NameUniqueReq) (res *post.NameUniqueRes, err error) {
	data, err := service.AdminPost().NameUnique(ctx, adminin.PostNameUniqueInp{Id: req.Id, Name: req.Name})
	if err != nil {
		return
	}

	res = new(post.NameUniqueRes)
	res.IsUnique = data.IsUnique
	return
}

// CodeUnique 编码是否唯一
func (c *cPost) CodeUnique(ctx context.Context, req *post.CodeUniqueReq) (res *post.CodeUniqueRes, err error) {
	data, err := service.AdminPost().CodeUnique(ctx, adminin.PostCodeUniqueInp{Id: req.Id, Code: req.Code})
	if err != nil {
		return nil, err
	}

	res = new(post.CodeUniqueRes)
	res.IsUnique = data.IsUnique
	return
}

// View 获取指定信息
func (c *cPost) View(ctx context.Context, req *post.ViewReq) (res *post.ViewRes, err error) {
	data, err := service.AdminPost().View(ctx, adminin.PostViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(post.ViewRes)
	res.PostViewModel = data
	return
}

// List 获取列表
func (c *cPost) List(ctx context.Context, req *post.ListReq) (res *post.ListRes, err error) {
	var in adminin.PostListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.AdminPost().List(ctx, in)
	if err != nil {
		return
	}

	res = new(post.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Status 更新状态
func (c *cPost) Status(ctx context.Context, req *post.StatusReq) (res *post.StatusRes, err error) {
	var in adminin.PostStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminPost().Status(ctx, in)
	return
}
