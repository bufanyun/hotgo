// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/tree"
)

type sAdminDept struct{}

func NewAdminDept() *sAdminDept {
	return &sAdminDept{}
}

func init() {
	service.RegisterAdminDept(NewAdminDept())
}

// Delete 删除
func (s *sAdminDept) Delete(ctx context.Context, in adminin.DeptDeleteInp) (err error) {
	var models *entity.AdminDept
	if err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		return err
	}

	if models == nil {
		return gerror.New("数据不存在或已删除！")
	}

	pidExist, err := dao.AdminDept.Ctx(ctx).Where("pid", models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !pidExist.IsEmpty() {
		return gerror.New("请先删除该部门下得所有子级！")
	}

	_, err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

// Edit 修改/新增
func (s *sAdminDept) Edit(ctx context.Context, in adminin.DeptEditInp) (err error) {
	if err = hgorm.IsUnique(ctx, &dao.AdminDept, g.Map{dao.AdminDept.Columns().Name: in.Name}, "名称已存在", in.Id); err != nil {
		return
	}

	if in.Pid, in.Level, in.Tree, err = hgorm.GenSubTree(ctx, &dao.AdminDept, in.Pid); err != nil {
		return
	}

	// 修改
	if in.Id > 0 {
		// 获取父级tree
		var pTree gdb.Value
		pTree, err = dao.AdminDept.Ctx(ctx).Where("id", in.Pid).Fields("tree").Value()
		if err != nil {
			return
		}
		in.Tree = tree.GenLabel(pTree.String(), in.Id)

		err = dao.AdminDept.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 更新数据
			_, err = dao.AdminDept.Ctx(ctx).Fields(adminin.DeptUpdateFields{}).WherePri(in.Id).Data(in).Update()
			if err != nil {
				return err
			}

			// 如果当前部门有子级,更新子级tree关系树
			return updateChildrenTree(ctx, in.Id, in.Level, in.Tree)
		})
		return
	}

	// 新增
	_, err = dao.AdminDept.Ctx(ctx).Fields(adminin.DeptInsertFields{}).Data(in).Insert()
	return
}

func updateChildrenTree(ctx context.Context, _id int64, _level int, _tree string) (err error) {
	var list []*entity.AdminDept
	if err = dao.AdminDept.Ctx(ctx).Where("pid", _id).Scan(&list); err != nil || list == nil {
		return
	}
	for _, child := range list {
		child.Level = _level + 1
		child.Tree = tree.GenLabel(_tree, child.Id)

		if _, err = dao.AdminDept.Ctx(ctx).Where("id", child.Id).Data("level", child.Level, "tree", child.Tree).Update(); err != nil {
			return
		}

		if err = updateChildrenTree(ctx, child.Id, child.Level, child.Tree); err != nil {
			return
		}
	}
	return
}

// Status 更新部门状态
func (s *sAdminDept) Status(ctx context.Context, in adminin.DeptStatusInp) (err error) {
	if _, err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update(); err != nil {
		err = gerror.Wrap(err, "更新部门状态失败！")
	}
	return
}

// MaxSort 最大排序
func (s *sAdminDept) MaxSort(ctx context.Context, in adminin.DeptMaxSortInp) (res *adminin.DeptMaxSortModel, err error) {
	if in.Id > 0 {
		if err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取部门数据异常！")
			return
		}
	}

	if res == nil {
		res = new(adminin.DeptMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(ctx, res.Sort)
	return
}

// View 获取指定部门信息
func (s *sAdminDept) View(ctx context.Context, in adminin.DeptViewInp) (res *adminin.DeptViewModel, err error) {
	if err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取部门信息失败！")
	}
	return
}

// Option 选项
func (s *sAdminDept) Option(ctx context.Context, in adminin.DeptOptionInp) (res *adminin.DeptOptionModel, totalCount int, err error) {
	var (
		mod    = dao.AdminDept.Ctx(ctx)
		models []*entity.AdminDept
		pid    int64 = 0
	)

	// 非超管只获取下级
	if !service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx)) {
		pid = contexts.GetUser(ctx).DeptId
		mod = mod.WhereLike(dao.AdminDept.Columns().Tree, "%"+tree.GetIdLabel(pid)+"%")
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取部门数据行失败！")
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("sort asc,id asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取部门数据失败！")
		return
	}

	res = new(adminin.DeptOptionModel)
	if models != nil {
		res.List = s.treeList(pid, models)
	}
	return
}

// List 获取列表
func (s *sAdminDept) List(ctx context.Context, in adminin.DeptListInp) (res *adminin.DeptListModel, err error) {
	var (
		mod    = dao.AdminDept.Ctx(ctx)
		models []*entity.AdminDept
		ids    []int64
		pids   []int64
	)

	// 部门名称
	if in.Name != "" {
		columns, err := dao.AdminDept.Ctx(ctx).Fields("pid").WhereLike("name", "%"+in.Name+"%").Array()
		if err != nil {
			err = gerror.Wrap(err, "过滤部门列表失败-1！")
			return nil, err
		}

		ds := g.NewVar(columns).Int64s()
		ids = append(ids, ds...)
		pids = append(pids, ds...)
		if len(ids) == 0 {
			return nil, nil
		}
	}

	if in.Code != "" {
		columns, err := dao.AdminDept.Ctx(ctx).Fields("pid").WhereLike("code", "%"+in.Code+"%").Array()
		if err != nil {
			err = gerror.Wrap(err, "过滤部门列表失败-2！")
			return nil, err
		}

		ds := g.NewVar(columns).Int64s()
		ids = append(ids, ds...)
		pids = append(pids, ds...)
		if len(ids) == 0 {
			return nil, nil
		}
	}

	if len(ids) > 0 {
		mod = mod.Wheref(`id in (?) or pid in (?)`, convert.UniqueSlice(ids), convert.UniqueSlice(pids))
	}

	if err = mod.Order("pid asc,sort asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	res = new(adminin.DeptListModel)
	res.List = s.treeList(0, models)
	return
}

// GetName 获取部门名称
func (s *sAdminDept) GetName(ctx context.Context, id int64) (name string, err error) {
	var data *entity.AdminDept
	if err = dao.AdminDept.Ctx(ctx).Where("id", id).Fields("name").Scan(&data); err != nil {
		err = gerror.Wrap(err, "获取部门名称失败！")
		return
	}

	if data == nil {
		err = gerror.Wrap(err, "部门不存在！")
		return
	}
	return data.Name, nil
}

// treeList 树状列表
func (s *sAdminDept) treeList(pid int64, nodes []*entity.AdminDept) (list []*adminin.DeptTree) {
	list = make([]*adminin.DeptTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(adminin.DeptTree)
			item.AdminDept = *v
			item.Label = v.Name
			item.Value = v.Id

			child := s.treeList(v.Id, nodes)
			if len(child) > 0 {
				item.Children = child
			}
			list = append(list, item)
		}
	}
	return
}
