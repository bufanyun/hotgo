// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/dict"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var (
	DictData = cDictData{}
)

type cDictData struct{}

// Delete 删除
func (c *cDictData) Delete(ctx context.Context, req *dict.DataDeleteReq) (res *dict.DataDeleteRes, err error) {
	var in sysin.DictDataDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.SysDictData().Delete(ctx, in)
	return
}

// Edit 更新
func (c *cDictData) Edit(ctx context.Context, req *dict.DataEditReq) (res *dict.DataEditRes, err error) {
	var in sysin.DictDataEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	err = service.SysDictData().Edit(ctx, in)
	return
}

// List 查看列表
func (c *cDictData) List(ctx context.Context, req *dict.DataListReq) (res *dict.DataListRes, err error) {
	var in sysin.DictDataListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.SysDictData().List(ctx, in)
	if err != nil {
		return
	}

	res = new(dict.DataListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Select 指定选项
func (c *cDictData) Select(ctx context.Context, req *dict.DataSelectReq) (res dict.DataSelectRes, err error) {
	var in sysin.DataSelectInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, err := service.SysDictData().Select(ctx, in)
	if err != nil {
		return
	}

	res = dict.DataSelectRes(list)
	return
}

// Selects 多个选项
func (c *cDictData) Selects(ctx context.Context, req *dict.DataSelectsReq) (res dict.DataSelectsRes, err error) {
	res = make(dict.DataSelectsRes)
	for _, v := range req.Types {
		option, err := service.SysDictData().Select(ctx, sysin.DataSelectInp{Type: v})
		if err != nil {
			return nil, err
		}
		res[v] = option
	}

	return
}
