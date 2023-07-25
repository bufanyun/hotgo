// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"hotgo/api/admin/post"
	"hotgo/internal/service"
)

// Post 岗位
var Post = cPost{}

type cPost struct{}

// Delete 删除
func (c *cPost) Delete(ctx context.Context, req *post.DeleteReq) (res *post.DeleteRes, err error) {
	err = service.AdminPost().Delete(ctx, &req.PostDeleteInp)
	return
}

// Edit 修改/新增
func (c *cPost) Edit(ctx context.Context, req *post.EditReq) (res *post.EditRes, err error) {
	err = service.AdminPost().Edit(ctx, &req.PostEditInp)
	return
}

// MaxSort 最大排序
func (c *cPost) MaxSort(ctx context.Context, req *post.MaxSortReq) (res *post.MaxSortRes, err error) {
	res = new(post.MaxSortRes)
	res.PostMaxSortModel, err = service.AdminPost().MaxSort(ctx, &req.PostMaxSortInp)
	return
}

// View 获取指定信息
func (c *cPost) View(ctx context.Context, req *post.ViewReq) (res *post.ViewRes, err error) {
	data, err := service.AdminPost().View(ctx, &req.PostViewInp)
	if err != nil {
		return
	}

	res = new(post.ViewRes)
	res.PostViewModel = data
	return
}

// List 获取列表
func (c *cPost) List(ctx context.Context, req *post.ListReq) (res *post.ListRes, err error) {
	list, totalCount, err := service.AdminPost().List(ctx, &req.PostListInp)
	if err != nil {
		return
	}

	res = new(post.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Status 更新状态
func (c *cPost) Status(ctx context.Context, req *post.StatusReq) (res *post.StatusRes, err error) {
	err = service.AdminPost().Status(ctx, &req.PostStatusInp)
	return
}
