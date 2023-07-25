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
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/addons/hgexample/service"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/validate"
)

type sSysTable struct{}

func NewSysTable() *sSysTable {
	return &sSysTable{}
}

func init() {
	service.RegisterSysTable(NewSysTable())
}

// Model Orm模型
func (s *sSysTable) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AddonHgexampleTable.Ctx(ctx), option...)
}

// List 获取列表
func (s *sSysTable) List(ctx context.Context, in *sysin.TableListInp) (list []*sysin.TableListModel, totalCount int, err error) {
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

	if err = mod.Fields(sysin.TableListModel{}).Page(in.Page, in.PerPage).OrderAsc(cols.Sort).OrderDesc(cols.Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取表格列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出
func (s *sSysTable) Export(ctx context.Context, in *sysin.TableListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.TableExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "表格例子导出-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.TableExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	if err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName); err != nil {
		return
	}
	return
}

// Edit 修改/新增
func (s *sSysTable) Edit(ctx context.Context, in *sysin.TableEditInp) (err error) {
	cols := dao.AddonHgexampleTable.Columns()
	if err = hgorm.IsUnique(ctx, &dao.AddonHgexampleTable, g.Map{cols.Qq: in.Qq}, "QQ号码已存在，请换一个", in.Id); err != nil {
		return
	}

	// 修改
	if in.Id > 0 {
		in.UpdatedBy = contexts.GetUserId(ctx)
		if _, err = s.Model(ctx).WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改表格失败，请稍后重试！")
			return
		}
		return
	}

	// 新增
	in.CreatedBy = contexts.GetUserId(ctx)
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增表格失败，请稍后重试！")
		return
	}
	return
}

// Delete 删除
func (s *sSysTable) Delete(ctx context.Context, in *sysin.TableDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除表格失败，请稍后重试！")
		return
	}
	return
}

// Status 更新状态
func (s *sSysTable) Status(ctx context.Context, in *sysin.TableStatusInp) (err error) {
	update := g.Map{
		dao.AddonHgexampleTable.Columns().Status:    in.Status,
		dao.AddonHgexampleTable.Columns().UpdatedBy: contexts.GetUserId(ctx),
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, "更新表格状态失败，请稍后重试！")
		return
	}
	return
}

// Switch 更新开关状态
func (s *sSysTable) Switch(ctx context.Context, in *sysin.TableSwitchInp) (err error) {
	var fields = []string{
		dao.AddonHgexampleTable.Columns().Switch,
		// ...
	}

	if !validate.InSlice(fields, in.Key) {
		err = gerror.New("开关键名不在白名单")
		return
	}

	update := g.Map{
		in.Key: in.Value,
		dao.AddonHgexampleTable.Columns().UpdatedBy: contexts.GetUserId(ctx),
	}

	if _, err = s.Model(ctx).Where(dao.AddonHgexampleTable.Columns().Id, in.Id).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, "更新表格开关失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 最大排序
func (s *sSysTable) MaxSort(ctx context.Context, in *sysin.TableMaxSortInp) (res *sysin.TableMaxSortModel, err error) {
	dx := dao.AddonHgexampleTable
	if err = dx.Ctx(ctx).Fields(dx.Columns().Sort).OrderDesc(dx.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取表格最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(sysin.TableMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取指定信息
func (s *sSysTable) View(ctx context.Context, in *sysin.TableViewInp) (res *sysin.TableViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取生成演示信息，请稍后重试！")
		return
	}
	return
}
