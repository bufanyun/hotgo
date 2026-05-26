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
	"hotgo/internal/library/hgorm/handler"
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

// Model 部门ORM模型
func (s *sAdminDept) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AdminDept.Ctx(ctx), option...)
}

// Delete 删除
func (s *sAdminDept) Delete(ctx context.Context, in *adminin.DeptDeleteInp) (err error) {
	var models *entity.AdminDept
	if err = s.Model(ctx).WherePri(in.Id).Scan(&models); err != nil {
		return err
	}

	if models == nil {
		return gerror.New("数据不存在或已删除！")
	}

	pidExist, err := s.Model(ctx).Where(dao.AdminDept.Columns().Pid, models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !pidExist.IsEmpty() {
		return gerror.New("请先删除该部门下得所有子级！")
	}

	_, err = s.Model(ctx).WherePri(in.Id).Delete()
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
	where := g.Map{
		dao.AdminDept.Columns().Name: in.Name,
		dao.AdminDept.Columns().Code: in.Code,
	}
	// 验证唯一性
	err = s.VerifyUnique(ctx, &adminin.VerifyUniqueInp{Id: in.Id, Where: where})
	if err != nil {
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		in.Pid, in.Level, in.Tree, err = hgorm.AutoUpdateTree(ctx, &dao.AdminDept, in.Id, in.Pid)
		if err != nil {
			return err
		}
		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改部门管理失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增部门管理失败，请稍后重试！")
		}
		return
	})
}

// MaxSort 最大排序
func (s *sAdminDept) MaxSort(ctx context.Context, in *adminin.DeptMaxSortInp) (res *adminin.DeptMaxSortModel, err error) {
	if in.Id > 0 {
		if err = dao.AdminDept.Ctx(ctx).WherePri(in.Id).OrderDesc(dao.AdminDept.Columns().Sort).Scan(&res); err != nil {
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
	if err = dao.AdminDept.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取部门信息失败！")
	}
	return
}

// List 获取列表
func (s *sAdminDept) List(ctx context.Context, in *adminin.DeptListInp) (res *adminin.DeptListModel, err error) {
	res = new(adminin.DeptListModel)

	var (
		mod  = dao.AdminDept.Ctx(ctx)
		cols = dao.AdminDept.Columns()
	)

	// 部门名称
	if in.Name != "" {
		columns, err := dao.AdminDept.Ctx(ctx).Fields(cols.Id).WhereLike(cols.Name, "%"+in.Name+"%").Array()
		if err != nil {
			err = gerror.Wrap(err, "查询部门名称失败！")
			return nil, err
		}

		if len(columns) == 0 {
			return nil, nil
		}
		res.Ids = append(res.Ids, g.NewVar(columns).Int64s()...)
	}

	// 部门编码
	if in.Code != "" {
		columns, err := dao.AdminDept.Ctx(ctx).Fields(cols.Id).WhereLike(cols.Code, "%"+in.Code+"%").Array()
		if err != nil {
			err = gerror.Wrap(err, "查询部门编码失败！")
			return nil, err
		}

		if len(columns) == 0 {
			return nil, nil
		}
		res.Ids = append(res.Ids, g.NewVar(columns).Int64s()...)
	}

	// 负责人
	if in.Leader != "" {
		columns, err := dao.AdminDept.Ctx(ctx).Fields(cols.Id).Where(cols.Leader, in.Leader).Array()
		if err != nil {
			err = gerror.Wrap(err, "查询负责人失败！")
			return nil, err
		}

		if len(columns) == 0 {
			return nil, nil
		}
		res.Ids = append(res.Ids, g.NewVar(columns).Int64s()...)
	}

	// 创建时间
	if len(in.CreatedAt) == 2 {
		columns, err := dao.AdminDept.Ctx(ctx).Fields(cols.Id).WhereBetween(cols.CreatedAt, in.CreatedAt[0], in.CreatedAt[1]).Array()
		if err != nil {
			err = gerror.Wrap(err, "查询创建时间失败！")
			return nil, err
		}

		if len(columns) == 0 {
			return nil, nil
		}
		res.Ids = append(res.Ids, g.NewVar(columns).Int64s()...)
	}

	res.Ids = convert.UniqueSlice(res.Ids)
	if len(res.Ids) > 0 {
		// 找到匹配到的完整上级部门
		columns, err := dao.AdminDept.Ctx(ctx).Fields(cols.Tree).WhereIn(cols.Id, res.Ids).Array()
		if err != nil {
			err = gerror.Wrap(err, "查询部门失败，请稍后重试！")
			return nil, err
		}
		var pids []int64
		for _, tr := range g.NewVar(columns).Strings() {
			pids = append(pids, tree.GetIds(tr)...)
		}
		mod = mod.WhereIn(cols.Id, append(res.Ids, convert.UniqueSlice(pids)...))
	}

	if err = mod.Order("pid asc,sort asc").Scan(&res.List); err != nil {
		err = gerror.Wrap(err, "获取部门列表失败！")
		return
	}
	return
}

// GetName 获取部门名称
func (s *sAdminDept) GetName(ctx context.Context, id int64) (name string, err error) {
	var data *entity.AdminDept
	if err = dao.AdminDept.Ctx(ctx).Where(dao.AdminDept.Columns().Id, id).Fields(dao.AdminDept.Columns().Name).Scan(&data); err != nil {
		err = gerror.Wrap(err, "获取部门名称失败！")
		return
	}

	if data == nil {
		err = gerror.Wrap(err, "部门不存在！")
		return
	}
	return data.Name, nil
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

	if !validate.InSlice(g.NewVar(columns).Int64s(), id) {
		err = gerror.New("部门ID是无效的")
		return
	}
	return
}

// Option 获取当前登录用户可选的部门选项
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

// TreeOption 获取部门关系树选项
func (s *sAdminDept) TreeOption(ctx context.Context) (nodes []tree.Node, err error) {
	var models []*adminin.DeptTreeOption
	if err = s.Model(ctx).Fields(adminin.DeptTreeOption{}).OrderAsc(dao.AdminDept.Columns().Pid).OrderAsc(dao.AdminDept.Columns().Sort).OrderDesc(dao.AdminDept.Columns().Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取部门关系树选项失败！")
		return
	}
	nodes = make([]tree.Node, len(models))
	for i, v := range models {
		nodes[i] = v
	}
	return tree.ListToTree(0, nodes)
}
