// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
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
func (s *sAdminDept) NameUnique(ctx context.Context, in adminin.DeptNameUniqueInp) (*adminin.DeptNameUniqueModel, error) {

	var res adminin.DeptNameUniqueModel
	isUnique, err := dao.AdminDept.IsUniqueName(ctx, in.Id, in.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.IsUnique = isUnique
	return &res, nil
}

// Delete 删除
func (s *sAdminDept) Delete(ctx context.Context, in adminin.DeptDeleteInp) (err error) {

	var (
		models *entity.AdminDept
	)
	err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Scan(&models)
	if err != nil {
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
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sAdminDept) Edit(ctx context.Context, in adminin.DeptEditInp) (err error) {

	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return err
	}

	uniqueName, err := dao.AdminDept.IsUniqueName(ctx, in.Id, in.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !uniqueName {
		err = gerror.New("名称已存在")
		return err
	}

	in.Pid, in.Level, in.Tree, err = hgorm.GenSubTree(ctx, dao.AdminDept, in.Pid)
	if err != nil {
		return err
	}

	// 修改
	if in.Id > 0 {
		_, err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	_, err = dao.AdminDept.Ctx(ctx).Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

// Status 更新部门状态
func (s *sAdminDept) Status(ctx context.Context, in adminin.DeptStatusInp) (err error) {

	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return err
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return err
	}

	if !validate.InSliceInt(consts.StatusMap, in.Status) {
		err = gerror.New("状态不正确")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	_, err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// MaxSort 最大排序
func (s *sAdminDept) MaxSort(ctx context.Context, in adminin.DeptMaxSortInp) (*adminin.DeptMaxSortModel, error) {
	var res adminin.DeptMaxSortModel

	if in.Id > 0 {
		if err := dao.AdminDept.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	res.Sort = res.Sort + 10

	return &res, nil
}

// View 获取指定字典类型信息
func (s *sAdminDept) View(ctx context.Context, in adminin.DeptViewInp) (res *adminin.DeptViewModel, err error) {

	if err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

// List 获取列表
func (s *sAdminDept) List(ctx context.Context, in adminin.DeptListInp) (list adminin.DeptListModel, err error) {
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
		values, err := dao.AdminDept.Ctx(ctx).Fields("pid").WhereLike("code", "%"+in.Code+"%").Array()
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
		return list, err
	}

	list = gconv.SliceMap(models)
	for k, v := range list {
		list[k]["index"] = v["id"]
		list[k]["key"] = v["id"]
		list[k]["label"] = v["name"]
	}

	return tree.GenTree(list), nil
}

type DeptTree struct {
	entity.AdminDept
	Children []*DeptTree `json:"children"`
}

// getDeptChildIds 将列表转为父子关系列表
func (s *sAdminDept) getDeptChildIds(ctx context.Context, lists []*DeptTree, pid int64) []*DeptTree {

	var (
		count    = len(lists)
		newLists []*DeptTree
	)

	if count == 0 {
		return nil
	}

	for i := 0; i < len(lists); i++ {
		if lists[i].Id > 0 && lists[i].Pid == pid {
			var row *DeptTree
			if err := gconv.Structs(lists[i], &row); err != nil {
				panic(err)
			}
			row.Children = s.getDeptChildIds(ctx, lists, row.Id)
			newLists = append(newLists, row)
		}
	}

	return newLists
}

type DeptListTree struct {
	Id       int64           `json:"id" `
	Key      int64           `json:"key" `
	Pid      int64           `json:"pid"  `
	Label    string          `json:"label"`
	Title    string          `json:"title"`
	Name     string          `json:"name"`
	Type     string          `json:"type"`
	Children []*DeptListTree `json:"children"`
}

// ListTree 获取列表树
func (s *sAdminDept) ListTree(ctx context.Context, in adminin.DeptListTreeInp) (list []*adminin.DeptListTreeModel, err error) {
	var (
		mod      = dao.AdminDept.Ctx(ctx)
		dataList []*entity.AdminDept
		models   []*DeptListTree
	)

	err = mod.Order("id desc").Scan(&dataList)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, err
	}

	_ = gconv.Structs(dataList, &models)

	// 重写树入参
	for i := 0; i < len(models); i++ {
		models[i].Key = models[i].Id
		models[i].Title = models[i].Name
		models[i].Label = models[i].Name
	}

	childIds := s.getDeptTreeChildIds(ctx, models, 0)

	_ = gconv.Structs(childIds, &list)

	return list, nil
}

// getDeptTreeChildIds 将列表转为父子关系列表
func (s *sAdminDept) getDeptTreeChildIds(ctx context.Context, lists []*DeptListTree, pid int64) []*DeptListTree {
	var (
		count    = len(lists)
		newLists []*DeptListTree
	)

	if count == 0 {
		return nil
	}

	for i := 0; i < len(lists); i++ {
		if lists[i].Id > 0 && lists[i].Pid == pid {
			var row *DeptListTree
			if err := gconv.Structs(lists[i], &row); err != nil {
				panic(err)
			}
			row.Children = s.getDeptTreeChildIds(ctx, lists, row.Id)
			newLists = append(newLists, row)
		}
	}

	return newLists
}

// GetName 获取部门名称
func (s *sAdminDept) GetName(ctx context.Context, id int64) (name string, err error) {
	var data entity.AdminDept
	err = dao.AdminDept.Ctx(ctx).
		Where("id", id).
		Fields("name").
		Scan(&data)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return name, err
	}

	return data.Name, nil
}
