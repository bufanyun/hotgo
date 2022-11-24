// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/notice"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
)

var (
	Notice = cNotice{}
)

type cNotice struct{}

// Delete 删除
func (c *cNotice) Delete(ctx context.Context, req *notice.DeleteReq) (res *notice.DeleteRes, err error) {
	var in adminin.NoticeDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.AdminNotice().Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

// Edit 更新
func (c *cNotice) Edit(ctx context.Context, req *notice.EditReq) (res *notice.EditRes, err error) {

	var in adminin.NoticeEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.AdminNotice().Edit(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

// MaxSort 最大排序
func (c *cNotice) MaxSort(ctx context.Context, req *notice.MaxSortReq) (*notice.MaxSortRes, error) {

	data, err := service.AdminNotice().MaxSort(ctx, adminin.NoticeMaxSortInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res notice.MaxSortRes
	res.Sort = data.Sort
	return &res, nil
}

// View 获取指定信息
func (c *cNotice) View(ctx context.Context, req *notice.ViewReq) (*notice.ViewRes, error) {

	data, err := service.AdminNotice().View(ctx, adminin.NoticeViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res notice.ViewRes
	res.NoticeViewModel = data
	return &res, nil
}

// List 查看列表
func (c *cNotice) List(ctx context.Context, req *notice.ListReq) (*notice.ListRes, error) {

	var (
		in  adminin.NoticeListInp
		res notice.ListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := service.AdminNotice().List(ctx, in)
	if err != nil {
		return nil, err
	}

	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage

	return &res, nil
}

// Status 更新部门状态
func (c *cNotice) Status(ctx context.Context, req *notice.StatusReq) (res *notice.StatusRes, err error) {

	var in adminin.NoticeStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.AdminNotice().Status(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}
