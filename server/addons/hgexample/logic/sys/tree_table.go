// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/addons/hgexample/service"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
)

type sSysTreeTable struct{}

func NewSysTreeTable() *sSysTreeTable {
	return &sSysTreeTable{}
}

func init() {
	service.RegisterSysTreeTable(NewSysTreeTable())
}

// Model Orm模型
func (s *sSysTreeTable) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AddonHgexampleTable.Ctx(ctx), option...)
}

// List 获取列表
func (s *sSysTreeTable) List(ctx context.Context, in *sysin.TreeTableListInp) (list []*sysin.TreeTableListModel, totalCount int, err error) {
	mod := s.Model(ctx)
	cols := dao.AddonHgexampleTable.Columns()

	if in.Title != "" {
		mod = mod.WhereLike(cols.Title, "%"+in.Title+"%")
	}

	if in.Content != "" {
		mod = mod.WhereLike(cols.Content, "%"+in.Content+"%")
	}

	if in.Status > 0 {
		mod = mod.Where(cols.Status, in.Status)
	}

	if in.Switch > 0 {
		mod = mod.Where(cols.Switch, in.Switch)
	}

	if in.Pid > 0 {
		mod = mod.Where(cols.Pid, in.Pid)
	}

	if len(in.Price) > 0 {
		if in.Price[0] > 0 && in.Price[1] > 0 {
			mod = mod.WhereBetween(cols.Price, in.Price[0], in.Price[1])
		} else if in.Price[0] > 0 && in.Price[1] == 0 {
			mod = mod.WhereGTE(cols.Price, in.Price[0])
		} else if in.Price[0] == 0 && in.Price[1] > 0 {
			mod = mod.WhereLTE(cols.Price, in.Price[1])
		}
	}

	if in.ActivityAt != nil {
		mod = mod.Where(cols.ActivityAt, in.ActivityAt)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(cols.CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if !in.Flag.IsNil() {
		mod = mod.Where(fmt.Sprintf(`JSON_CONTAINS(%s,'%v')`, cols.Flag, in.Flag))
	}

	if !in.Hobby.IsNil() {
		mod = mod.Where(fmt.Sprintf(`JSON_CONTAINS(%s,'%v')`, cols.Hobby, in.Hobby))
	}

	totalCount, err = mod.Clone().Count(1)
	if err != nil {
		err = gerror.Wrap(err, "获取表格数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.TableListModel{}).Page(in.Page, in.PerPage).Handler(handler.Sorter(in)).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取表格列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增
func (s *sSysTreeTable) Edit(ctx context.Context, in *sysin.TableEditInp) (err error) {
	cols := dao.AddonHgexampleTable.Columns()
	if err = hgorm.IsUnique(ctx, &dao.AddonHgexampleTable, g.Map{cols.Qq: in.Qq}, "QQ号码已存在，请换一个", in.Id); err != nil {
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		in.Pid, in.Level, in.Tree, err = hgorm.AutoUpdateTree(ctx, &dao.AddonHgexampleTable, in.Id, in.Pid)
		if err != nil {
			return err
		}

		if in.Id > 0 {
			in.UpdatedBy = contexts.GetUserId(ctx)
			if _, err = s.Model(ctx).WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改表格失败，请稍后重试！")
				return err
			}
		} else {
			in.CreatedBy = contexts.GetUserId(ctx)
			if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).Data(in).Insert(); err != nil {
				err = gerror.Wrap(err, "新增表格失败，请稍后重试！")
				return err
			}
		}
		return
	})
}

// Delete 删除
func (s *sSysTreeTable) Delete(ctx context.Context, in *sysin.TableDeleteInp) (err error) {
	count, err := dao.AdminMenu.Ctx(ctx).Where("pid", in.Id).Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if count > 0 {
		return gerror.New("请先删除该表格下的所有下级！")
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除表格失败，请稍后重试！")
		return
	}
	return
}

// Select 关系树选项列表
func (s *sSysTreeTable) Select(ctx context.Context) (list []*sysin.TableTree, err error) {
	var models []*entity.AddonHgexampleTable
	if err = dao.AddonHgexampleTable.Ctx(ctx).Order("pid asc,id asc,sort asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取关系树选项列表失败！")
		return
	}

	list = s.treeList(0, models)
	return
}

// treeList 树状列表
func (s *sSysTreeTable) treeList(pid int64, nodes []*entity.AddonHgexampleTable) (list []*sysin.TableTree) {
	list = make([]*sysin.TableTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(sysin.TableTree)
			item.AddonHgexampleTable = *v
			item.Label = v.Title
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
