// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
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
		return
	}

	list = s.treeList(0, models)
	return
}

// Delete 删除
func (s *sSysDictType) Delete(ctx context.Context, in sysin.DictTypeDeleteInp) (err error) {
	var models *entity.SysDictType
	if err = dao.SysDictType.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		return
	}

	if models == nil {
		err = gerror.New("数据不存在或已删除！")
		return
	}

	exist, err := dao.SysDictData.Ctx(ctx).Where("type", models.Type).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !exist.IsEmpty() {
		err = gerror.New("请先删除该字典类型下得所有字典数据！")
		return
	}

	pidExist, err := dao.SysDictType.Ctx(ctx).Where("pid", models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if !pidExist.IsEmpty() {
		err = gerror.New("请先删除该字典类型下得所有子级类型！")
		return
	}

	_, err = dao.SysDictType.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

// Edit 修改/新增
func (s *sSysDictType) Edit(ctx context.Context, in sysin.DictTypeEditInp) (err error) {
	if err = hgorm.IsUnique(ctx, &dao.SysDictType, g.Map{dao.SysDictType.Columns().Name: in.Name}, "名称已存在", in.Id); err != nil {
		return
	}

	// 修改
	if in.Id > 0 {
		if _, err = dao.SysDictType.Ctx(ctx).Fields(sysin.DictTypeUpdateFields{}).WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
		}
		return
	}

	// 新增
	if _, err = dao.SysDictType.Ctx(ctx).Fields(sysin.DictTypeInsertFields{}).Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}
	return
}

// TreeSelect 获取类型关系树选项
func (s *sSysDictType) TreeSelect(ctx context.Context, in sysin.DictTreeSelectInp) (list []*sysin.DictTypeTree, err error) {
	var (
		mod    = dao.SysDictType.Ctx(ctx)
		models []*entity.SysDictType
	)

	if err = mod.Order("pid asc,sort asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
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
