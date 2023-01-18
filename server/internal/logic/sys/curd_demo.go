// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.1.0
// @AutoGenerate Date 2023-01-18 15:19:42
//
package sys

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysCurdDemo struct{}

func NewSysCurdDemo() *sSysCurdDemo {
	return &sSysCurdDemo{}
}

func init() {
	service.RegisterSysCurdDemo(NewSysCurdDemo())
}

// Model 生成演示Orm模型
func (s *sSysCurdDemo) Model(ctx context.Context) *gdb.Model {
	return dao.SysGenCurdDemo.Ctx(ctx)
}

// List 获取生成演示列表
func (s *sSysCurdDemo) List(ctx context.Context, in sysin.CurdDemoListInp) (list []*sysin.CurdDemoListModel, totalCount int, err error) {
	mod := dao.SysGenCurdDemo.Ctx(ctx)

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.SysGenCurdDemo.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.SysGenCurdDemo.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysGenCurdDemo.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 查询分类名称
	if in.TestCategoryName != "" {
		mod = mod.WhereLike(dao.TestCategory.Columns().Name, in.TestCategoryName)
	}

	// 关联表testCategory
	mod = mod.LeftJoin(hgorm.GenJoinOnRelation(
		dao.SysGenCurdDemo.Table(), dao.SysGenCurdDemo.Columns().CategoryId, // 主表表名,关联条件
		dao.TestCategory.Table(), "testCategory", dao.TestCategory.Columns().Id, // 关联表表名,别名,关联条件
	)...)

	totalCount, err = mod.Clone().Count(1)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	if totalCount == 0 {
		return list, totalCount, nil
	}

	//关联表select
	fields, err := hgorm.GenJoinSelect(ctx, sysin.CurdDemoListModel{}, dao.SysGenCurdDemo, []*hgorm.Join{
		{Dao: dao.TestCategory, Alias: "testCategory"},
	})
	if err = mod.Fields(fields).Handler(hgorm.HandlerFilterAuth).Page(in.Page, in.PerPage).OrderAsc(dao.SysGenCurdDemo.Columns().Sort).OrderDesc(dao.SysGenCurdDemo.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	return list, totalCount, err
}

// Export 导出生成演示
func (s *sSysCurdDemo) Export(ctx context.Context, in sysin.CurdDemoListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return err
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.CurdDemoExportModel{})
	if err != nil {
		return err
	}

	var (
		fileName  = "导出生成演示-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.CurdDemoExportModel
	)

	err = gconv.Scan(list, &exports)
	if err != nil {
		return err
	}
	if err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName); err != nil {
		return
	}
	return
}

// Edit 修改/新增生成演示
func (s *sSysCurdDemo) Edit(ctx context.Context, in sysin.CurdDemoEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		in.UpdatedBy = contexts.GetUserId(ctx)
		_, err = dao.SysGenCurdDemo.Ctx(ctx).
			FieldsEx(
				dao.SysGenCurdDemo.Columns().Id,
				dao.SysGenCurdDemo.Columns().CreatedBy,
				dao.SysGenCurdDemo.Columns().CreatedAt,
				dao.SysGenCurdDemo.Columns().DeletedAt,
			).
			Where(dao.SysGenCurdDemo.Columns().Id, in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}
		return nil
	}

	// 新增
	in.CreatedBy = contexts.GetUserId(ctx)
	_, err = dao.SysGenCurdDemo.Ctx(ctx).
		FieldsEx(
			dao.SysGenCurdDemo.Columns().Id,
			dao.SysGenCurdDemo.Columns().UpdatedBy,
			dao.SysGenCurdDemo.Columns().DeletedAt,
		).
		Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return
}

// Delete 删除生成演示
func (s *sSysCurdDemo) Delete(ctx context.Context, in sysin.CurdDemoDeleteInp) (err error) {
	_, err = dao.SysGenCurdDemo.Ctx(ctx).Where(dao.SysGenCurdDemo.Columns().Id, in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// MaxSort 获取生成演示最大排序
func (s *sSysCurdDemo) MaxSort(ctx context.Context, in sysin.CurdDemoMaxSortInp) (res *sysin.CurdDemoMaxSortModel, err error) {
	if err = dao.SysGenCurdDemo.Ctx(ctx).Fields(dao.SysGenCurdDemo.Columns().Sort).OrderDesc(dao.SysGenCurdDemo.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.Sort = res.Sort + g.Cfg().MustGet(ctx, "hotgo.admin.maxSortIncrement").Int()
	return res, nil
}

// View 获取生成演示指定信息
func (s *sSysCurdDemo) View(ctx context.Context, in sysin.CurdDemoViewInp) (res *sysin.CurdDemoViewModel, err error) {
	if err = dao.SysGenCurdDemo.Ctx(ctx).Where(dao.SysGenCurdDemo.Columns().Id, in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

// Status 更新生成演示状态
func (s *sSysCurdDemo) Status(ctx context.Context, in sysin.CurdDemoStatusInp) (err error) {
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

	_, err = dao.SysGenCurdDemo.Ctx(ctx).Where(dao.SysGenCurdDemo.Columns().Id, in.Id).Data(dao.SysGenCurdDemo.Columns().Status, in.Status).Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Switch 更新生成演示开关
func (s *sSysCurdDemo) Switch(ctx context.Context, in sysin.CurdDemoSwitchInp) (err error) {
	var fields = []string{
		dao.SysGenCurdDemo.Columns().Switch,

		// ...
	}

	if !validate.InSliceString(fields, in.Key) {
		err = gerror.New("开关键名不在白名单")
		return err
	}

	_, err = dao.SysGenCurdDemo.Ctx(ctx).Where(dao.SysGenCurdDemo.Columns().Id, in.Id).Data(in.Key, in.Value).Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}
