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
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	DictType = cDictType{}
)

type cDictType struct{}

// Tree 树
func (c *cDictType) Tree(ctx context.Context, req *dict.TypeTreeReq) (*dict.TypeTreeRes, error) {
	var (
		res dict.TypeTreeRes
		err error
	)
	res.List, err = service.SysDictType().Tree(ctx)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Delete 删除
func (c *cDictType) Delete(ctx context.Context, req *dict.TypeDeleteReq) (res *dict.TypeDeleteRes, err error) {
	var in sysin.DictTypeDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysDictType().Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

// Edit 更新
func (c *cDictType) Edit(ctx context.Context, req *dict.TypeEditReq) (res *dict.TypeEditRes, err error) {

	var in sysin.DictTypeEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.SysDictType().Edit(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

// Select 选项
func (c *cDictType) Select(ctx context.Context, req *dict.TypeSelectReq) (res *dict.TypeSelectRes, err error) {
	list, err := service.SysDictType().Select(ctx, sysin.DictTypeSelectInp{})
	if err != nil {
		return nil, err
	}
	res = (*dict.TypeSelectRes)(&list)

	return res, nil
}
