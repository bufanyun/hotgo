// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

type sSysDictType struct{}

func NewSysDictType() *sSysDictType {
	return &sSysDictType{}
}

func init() {
	service.RegisterSysDictType(NewSysDictType())
}

// Tree 树
func (s *sSysDictType) Tree(ctx context.Context) (list []*sysin.DictTypeTree, err error) {
	var (
		mod    = dao.SysDictType.Ctx(ctx)
		models []*entity.SysDictType
	)

	if err = mod.Order("sort asc,id asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, err
	}

	list = s.treeList(0, models)
	return
}

// Delete 删除
func (s *sSysDictType) Delete(ctx context.Context, in sysin.DictTypeDeleteInp) error {
	var (
		models *entity.SysDictType
	)
	err := dao.SysDictType.Ctx(ctx).Where("id", in.Id).Scan(&models)
	if err != nil {
		return err
	}

	if models == nil {
		return gerror.New("数据不存在或已删除！")
	}

	exist, err := dao.SysDictData.Ctx(ctx).Where("type", models.Type).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !exist.IsEmpty() {
		return gerror.New("请先删除该字典类型下得所有字典数据！")
	}

	pidExist, err := dao.SysDictType.Ctx(ctx).Where("pid", models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !pidExist.IsEmpty() {
		return gerror.New("请先删除该字典类型下得所有子级类型！")
	}

	_, err = dao.SysDictType.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sSysDictType) Edit(ctx context.Context, in sysin.DictTypeEditInp) (err error) {
	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return err
	}

	uniqueName, err := dao.SysDictType.IsUniqueType(ctx, in.Id, in.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !uniqueName {
		err = gerror.New("名称已存在")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.SysDictType.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()
	_, err = dao.SysDictType.Ctx(ctx).Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

// TreeSelect 获取类型关系树选项
func (s *sSysDictType) TreeSelect(ctx context.Context, in sysin.DictTreeSelectInp) (list []*sysin.DictTypeTree, err error) {
	var (
		mod    = dao.SysDictType.Ctx(ctx)
		models []*entity.SysDictType
	)

	if err = mod.Order("pid asc,sort asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, err
	}

	list = s.treeList(0, models)

	for _, v := range list {
		// 父类一律禁止选中
		if len(v.Children) > 0 {
			v.Disabled = true
			for _, v2 := range v.Children {
				if len(v2.Children) > 0 {
					v2.Disabled = true
				}
			}
		}
	}
	return
}

// treeList 树状列表
func (s *sSysDictType) treeList(pid int64, nodes []*entity.SysDictType) (list []*sysin.DictTypeTree) {
	list = make([]*sysin.DictTypeTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(sysin.DictTypeTree)
			item.SysDictType = *v
			item.Label = v.Name
			item.Value = v.Id
			item.Key = v.Id

			child := s.treeList(v.Id, nodes)
			if len(child) > 0 {
				item.Children = child
			}
			list = append(list, item)
		}
	}
	return
}
