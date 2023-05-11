// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
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
	"hotgo/utility/validate"
)

type sAdminDept struct{}

func NewAdminDept() *sAdminDept {
	return &sAdminDept{}
}

func init() {
	service.RegisterAdminDept(NewAdminDept())
}

// NameUnique 菜单名称是否唯一
func (s *sAdminDept) NameUnique(ctx context.Context, in adminin.DeptNameUniqueInp) (res *adminin.DeptNameUniqueModel, err error) {
	isUnique, err := dao.AdminDept.IsUniqueName(ctx, in.Id, in.Name)
	if err != nil {
		return
	}

	res = new(adminin.DeptNameUniqueModel)
	res.IsUnique = isUnique
	return
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
	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return
	}

	uniqueName, err := dao.AdminDept.IsUniqueName(ctx, in.Id, in.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	if !uniqueName {
		err = gerror.New("名称已存在")
		return
	}

	if in.Pid, in.Level, in.Tree, err = hgorm.GenSubTree(ctx, dao.AdminDept, in.Pid); err != nil {
		return
	}

	// 修改
	if in.Id > 0 {
		_, err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		return
	}

	// 新增
	_, err = dao.AdminDept.Ctx(ctx).Data(in).Insert()
	return
}

// Status 更新部门状态
func (s *sAdminDept) Status(ctx context.Context, in adminin.DeptStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSliceInt(consts.StatusMap, in.Status) {
		err = gerror.New("状态不正确")
		return
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	_, err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	return
}

// MaxSort 最大排序
func (s *sAdminDept) MaxSort(ctx context.Context, in adminin.DeptMaxSortInp) (res *adminin.DeptMaxSortModel, err error) {
	if in.Id > 0 {
		if err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
	}

	if res == nil {
		res = new(adminin.DeptMaxSortModel)
	}
	res.Sort = form.DefaultMaxSort(ctx, res.Sort)
	return
}

// View 获取指定字典类型信息
func (s *sAdminDept) View(ctx context.Context, in adminin.DeptViewInp) (res *adminin.DeptViewModel, err error) {
	err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Scan(&res)
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
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("sort asc,id asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	res = new(adminin.DeptOptionModel)
	res.List = s.treeList(pid, models)
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
		values, err := dao.AdminDept.Ctx(ctx).Fields("pid").WhereLike("name", "%"+in.Name+"%").Array()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
		for i := 0; i < len(values); i++ {
			ids = append(ids, values[i].Int64())
			pids = append(pids, values[i].Int64())
		}

		if len(ids) == 0 {
			return nil, nil
		}
	}

	if in.Code != "" {
		values, err := dao.AdminDept.Ctx(ctx).Fields("pid").
			WhereLike("code", "%"+in.Code+"%").Array()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
		for i := 0; i < len(values); i++ {
			ids = append(ids, values[i].Int64())
			pids = append(pids, values[i].Int64())
		}

		if len(ids) == 0 {
			return nil, nil
		}
	}

	if len(ids) > 0 {
		ids = convert.UniqueSliceInt64(ids)
		pids = convert.UniqueSliceInt64(pids)
		mod = mod.Wheref(`id in (?) or pid in (?)`, ids, pids)
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
	var data entity.AdminDept
	err = dao.AdminDept.Ctx(ctx).Where("id", id).Fields("name").Scan(&data)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return name, err
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
