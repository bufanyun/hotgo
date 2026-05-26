// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"hotgo/utility/tree"
)

type sSysOptionTreeDemo struct{}

func NewSysOptionTreeDemo() *sSysOptionTreeDemo {
	return &sSysOptionTreeDemo{}
}

func init() {
	service.RegisterSysOptionTreeDemo(NewSysOptionTreeDemo())
}

// Model 选项树表ORM模型
func (s *sSysOptionTreeDemo) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysGenTreeDemo.Ctx(ctx), option...)
}

// List 获取选项树表列表
func (s *sSysOptionTreeDemo) List(ctx context.Context, in *sysin.OptionTreeDemoListInp) (list []*sysin.OptionTreeDemoListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.OptionTreeDemoListModel{})

	// 查询标题
	if in.Title != "" {
		mod = mod.WhereLike(dao.SysGenTreeDemo.Columns().Title, "%"+in.Title+"%")
	}

	// 查询上级
	if in.Pid > 0 {
		mod = mod.Where(dao.SysGenTreeDemo.Columns().Pid, in.Pid)
	}

	// 查询测试分类
	if in.CategoryId > 0 {
		mod = mod.Where(dao.SysGenTreeDemo.Columns().CategoryId, in.CategoryId)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.SysGenTreeDemo.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysGenTreeDemo.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 树形列表判断是否需要分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderAsc(dao.SysGenTreeDemo.Columns().Sort).OrderDesc(dao.SysGenTreeDemo.Columns().Id)

	// 操作人摘要信息
	mod = mod.Hook(hook.MemberSummary)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取选项树表列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增选项树表
func (s *sSysOptionTreeDemo) Edit(ctx context.Context, in *sysin.OptionTreeDemoEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		in.Pid, in.Level, in.Tree, err = hgorm.AutoUpdateTree(ctx, &dao.SysGenTreeDemo, in.Id, in.Pid)
		if err != nil {
			return err
		}
		// 修改
		if in.Id > 0 {
			in.UpdatedBy = contexts.GetUserId(ctx)
			if _, err = s.Model(ctx).
				Fields(sysin.OptionTreeDemoUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改选项树表失败，请稍后重试！")
			}
			return
		}

		// 新增
		in.CreatedBy = contexts.GetUserId(ctx)
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.OptionTreeDemoInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增选项树表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除选项树表
func (s *sSysOptionTreeDemo) Delete(ctx context.Context, in *sysin.OptionTreeDemoDeleteInp) (err error) {
	count, err := dao.SysGenTreeDemo.Ctx(ctx).Where(dao.SysGenTreeDemo.Columns().Pid, in.Id).Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if count > 0 {
		return gerror.New("请先删除该选项树表下的所有下级！")
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除选项树表失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取选项树表最大排序
func (s *sSysOptionTreeDemo) MaxSort(ctx context.Context, in *sysin.OptionTreeDemoMaxSortInp) (res *sysin.OptionTreeDemoMaxSortModel, err error) {
	if err = dao.SysGenTreeDemo.Ctx(ctx).Fields(dao.SysGenTreeDemo.Columns().Sort).OrderDesc(dao.SysGenTreeDemo.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取选项树表最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(sysin.OptionTreeDemoMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取选项树表指定信息
func (s *sSysOptionTreeDemo) View(ctx context.Context, in *sysin.OptionTreeDemoViewInp) (res *sysin.OptionTreeDemoViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Hook(hook.MemberSummary).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取选项树表信息，请稍后重试！")
		return
	}
	return
}

// TreeOption 获取选项树表关系树选项
func (s *sSysOptionTreeDemo) TreeOption(ctx context.Context) (nodes []tree.Node, err error) {
	var models []*sysin.OptionTreeDemoTreeOption
	if err = s.Model(ctx).Fields(sysin.OptionTreeDemoTreeOption{}).OrderAsc(dao.SysGenTreeDemo.Columns().Pid).OrderAsc(dao.SysGenTreeDemo.Columns().Sort).OrderDesc(dao.SysGenTreeDemo.Columns().Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取选项树表关系树选项失败！")
		return
	}
	nodes = make([]tree.Node, len(models))
	for i, v := range models {
		nodes[i] = v
	}
	return tree.ListToTree(0, nodes)
}