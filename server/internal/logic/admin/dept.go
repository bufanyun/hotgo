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
	"hotgo/utility/validate"
)

type sAdminDept struct{}

func NewAdminDept() *sAdminDept {
	return &sAdminDept{}
}

func init() {
	service.RegisterAdminDept(NewAdminDept())
}

// Delete 删除
func (s *sAdminDept) Delete(ctx context.Context, in *adminin.DeptDeleteInp) (err error) {
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

// VerifyUnique 验证部门唯一属性
func (s *sAdminDept) VerifyUnique(ctx context.Context, in *adminin.VerifyUniqueInp) (err error) {
	if in.Where == nil {
		return
	}

	cols := dao.AdminDept.Columns()
	msgMap := g.MapStrStr{
		cols.Name: "部门名称已存在，请换一个",
		cols.Code: "部门编码已存在，请换一个",
	}

	for k, v := range in.Where {
		if v == "" {
			continue
		}
		message, ok := msgMap[k]
		if !ok {
			err = gerror.Newf("字段 [ %v ] 未配置唯一属性验证", k)
			return
		}
		if err = hgorm.IsUnique(ctx, &dao.AdminDept, g.Map{k: v}, message, in.Id); err != nil {
			return
		}
	}
	return
}

// Edit 修改/新增
func (s *sAdminDept) Edit(ctx context.Context, in *adminin.DeptEditInp) (err error) {
	// 验证唯一性
	err = s.VerifyUnique(ctx, &adminin.VerifyUniqueInp{
		Id: in.Id,
		Where: g.Map{
			dao.AdminDept.Columns().Name: in.Name,
			dao.AdminDept.Columns().Code: in.Code,
		},
	})
	if err != nil {
		return
	}

	// 生成下级关系树
	if in.Pid, in.Level, in.Tree, err = hgorm.GenSubTree(ctx, &dao.AdminDept, in.Pid); err != nil {
		return
	}

	// 修改
	if in.Id > 0 {
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
		child.Tree = tree.GenLabel(_tree, child.Pid)

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
func (s *sAdminDept) Status(ctx context.Context, in *adminin.DeptStatusInp) (err error) {
	if _, err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update(); err != nil {
		err = gerror.Wrap(err, "更新部门状态失败！")
	}
	return
}

// MaxSort 最大排序
func (s *sAdminDept) MaxSort(ctx context.Context, in *adminin.DeptMaxSortInp) (res *adminin.DeptMaxSortModel, err error) {
	if in.Id > 0 {
		if err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取部门数据异常！")
			return
		}
	}

	if res == nil {
		res = new(adminin.DeptMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取指定部门信息
func (s *sAdminDept) View(ctx context.Context, in *adminin.DeptViewInp) (res *adminin.DeptViewModel, err error) {
	if err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取部门信息失败！")
	}
	return
}

// Option 选项
func (s *sAdminDept) Option(ctx context.Context, in *adminin.DeptOptionInp) (res *adminin.DeptOptionModel, totalCount int, err error) {
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
func (s *sAdminDept) List(ctx context.Context, in *adminin.DeptListInp) (res *adminin.DeptListModel, err error) {
	var (
		mod    = dao.AdminDept.Ctx(ctx)
		cols   = dao.AdminDept.Columns()
		models []*entity.AdminDept
		ids    []int64
		pids   []int64
	)

	appends := func(columns []gdb.Value) {
		ds := g.NewVar(columns).Int64s()
		ids = append(ids, ds...)
		pids = append(pids, ds...)
	}

	// 部门名称
	if in.Name != "" {
		columns, err := dao.AdminDept.Ctx(ctx).Fields(cols.Pid).WhereLike(cols.Name, "%"+in.Name+"%").Array()
		if err != nil {
			err = gerror.Wrap(err, "查询部门名称失败！")
			return nil, err
		}

		if len(columns) == 0 {
			return nil, nil
		}
		appends(columns)
	}

	if in.Code != "" {
		columns, err := dao.AdminDept.Ctx(ctx).Fields(cols.Pid).WhereLike(cols.Code, "%"+in.Code+"%").Array()
		if err != nil {
			err = gerror.Wrap(err, "查询部门编码失败！")
			return nil, err
		}

		if len(columns) == 0 {
			return nil, nil
		}
		appends(columns)
	}

	if in.Leader != "" {
		columns, err := dao.AdminDept.Ctx(ctx).Fields(cols.Pid).Where(cols.Leader, in.Leader).Array()
		if err != nil {
			err = gerror.Wrap(err, "查询负责人失败！")
			return nil, err
		}

		if len(columns) == 0 {
			return nil, nil
		}
		appends(columns)
	}

	if len(in.CreatedAt) == 2 {
		columns, err := dao.AdminDept.Ctx(ctx).Fields(cols.Pid).WhereBetween(cols.CreatedAt, in.CreatedAt[0], in.CreatedAt[1]).Array()
		if err != nil {
			err = gerror.Wrap(err, "查询创建时间失败！")
			return nil, err
		}

		if len(columns) == 0 {
			return nil, nil
		}
		appends(columns)
	}

	if len(ids) > 0 {
		mod = mod.Wheref(`id in (?) or pid in (?)`, convert.UniqueSlice(ids), convert.UniqueSlice(pids))
	}

	if err = mod.Order("pid asc,sort asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取部门列表失败！")
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

// VerifyDeptId 验证部门ID
func (s *sAdminDept) VerifyDeptId(ctx context.Context, id int64) (err error) {
	var (
		pid int64 = 0
		mb        = contexts.GetUser(ctx)
		mod       = dao.AdminDept.Ctx(ctx).Fields(dao.AdminDept.Columns().Id)
	)

	if mb == nil {
		err = gerror.New("用户信息获取失败！")
		return
	}

	// 非超管只获取下级
	if !service.AdminMember().VerifySuperId(ctx, mb.Id) {
		pid = mb.DeptId
		mod = mod.WhereNot(dao.AdminDept.Columns().Id, pid).WhereLike(dao.AdminDept.Columns().Tree, "%"+tree.GetIdLabel(pid)+"%")
	}

	columns, err := mod.Array()
	if err != nil {
		return err
	}

	ids := g.NewVar(columns).Int64s()
	if !validate.InSlice(ids, id) {
		err = gerror.New("部门ID是无效的")
		return
	}
	return
}
