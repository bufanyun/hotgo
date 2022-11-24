// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/dict"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	DictData = cDictData{}
)

type cDictData struct{}

// Delete 删除
func (c *cDictData) Delete(ctx context.Context, req *dict.DataDeleteReq) (res *dict.DataDeleteRes, err error) {
	var in sysin.DictDataDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysDictData().Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

// Edit 更新
func (c *cDictData) Edit(ctx context.Context, req *dict.DataEditReq) (res *dict.DataEditRes, err error) {

	var in sysin.DictDataEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysDictData().Edit(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

// List 查看列表
func (c *cDictData) List(ctx context.Context, req *dict.DataListReq) (*dict.DataListRes, error) {

	var (
		in  sysin.DictDataListInp
		res dict.DataListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := service.SysDictData().List(ctx, in)
	if err != nil {
		return nil, err
	}

	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage

	return &res, nil
}
