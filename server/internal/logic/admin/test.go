// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/validate"
)

type sAdminTest struct{}

func NewAdminTest() *sAdminTest {
	return &sAdminTest{}
}

func init() {
	service.RegisterAdminTest(NewAdminTest())
}

// Model Orm模型
func (s *sAdminTest) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Test.Ctx(ctx), option...)
}

// List 获取列表
func (s *sAdminTest) List(ctx context.Context, in adminin.TestListInp) (list []*adminin.TestListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	if in.Title != "" {
		mod = mod.WhereLike(dao.Test.Columns().Title, "%"+in.Title+"%")
	}

	if in.Content != "" {
		mod = mod.WhereLike(dao.Test.Columns().Content, "%"+in.Content+"%")
	}

	if in.Status > 0 {
		mod = mod.Where(dao.Test.Columns().Status, in.Status)
	}

	if in.Switch > 0 {
		mod = mod.Where(dao.Test.Columns().Switch, in.Switch)
	}

	if len(in.Price) > 0 {
		if in.Price[0] > float64(0) && in.Price[1] > float64(0) {
			mod = mod.WhereBetween(dao.Test.Columns().Price, in.Price[0], in.Price[1])
		} else if in.Price[0] > float64(0) && in.Price[1] == float64(0) {
			mod = mod.WhereGTE(dao.Test.Columns().Price, in.Price[0])
		} else if in.Price[0] == float64(0) && in.Price[1] > float64(0) {
			mod = mod.WhereLTE(dao.Test.Columns().Price, in.Price[1])
		}
	}

	if in.ActivityAt != nil {
		mod = mod.Where(dao.Test.Columns().ActivityAt, in.ActivityAt)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.Test.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if !in.Flag.IsNil() {
		mod = mod.Where(fmt.Sprintf(`JSON_CONTAINS(%s,'%v')`, dao.Test.Columns().Flag, in.Flag))
	}

	if !in.Hobby.IsNil() {
		mod = mod.Where(fmt.Sprintf(`JSON_CONTAINS(%s,'%v')`, dao.Test.Columns().Hobby, in.Hobby))
	}

	//// 关联表testCategory
	//mod = mod.LeftJoin(hgorm.GenJoinOnRelation(
	//	dao.Test.Table(), dao.Test.Columns().CategoryId, // 主表表名,关联条件
	//	dao.TestCategory.Table(), "testCategory", dao.TestCategory.Columns().Id, // 关联表表名,别名,关联条件
	//)...)
	//
	//mod = mod.Where(`testCategory.`+dao.TestCategory.Columns().Name, "微信公众号")

	totalCount, err = mod.Clone().Count(1)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if totalCount == 0 {
		return
	}

	////关联表select
	//fields, err := hgorm.GenJoinSelect(ctx, adminin.TestListModel{}, dao.Test, []*hgorm.Join{
	//	{Dao: dao.TestCategory, Alias: "testCategory"},
	//	//{Dao: dao.TestCategory, Alias: "testCategory"},
	//})

	fields, err := hgorm.GenSelect(ctx, adminin.TestListModel{}, dao.Test)
	if err != nil {
		return
	}

	if err = mod.Fields(fields).Handler(handler.FilterAuth, handler.ForceCache).Page(in.Page, in.PerPage).OrderAsc(dao.Test.Columns().Sort).OrderDesc(dao.Test.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return
}

// Export 导出
func (s *sAdminTest) Export(ctx context.Context, in adminin.TestListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(adminin.TestExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "测试导出-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []adminin.TestExportModel
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
func (s *sAdminTest) Edit(ctx context.Context, in adminin.TestEditInp) (err error) {
	if err = hgorm.IsUnique(ctx, dao.Test, g.Map{dao.Test.Columns().Qq: in.Qq}, "QQ号码已存在，请换一个", in.Id); err != nil {
		return
	}

	// 修改
	if in.Id > 0 {
		in.UpdatedBy = contexts.GetUserId(ctx)
		_, err = s.Model(ctx).Where(dao.Test.Columns().Id, in.Id).Data(in).Update()
		return
	}

	// 新增
	in.CreatedBy = contexts.GetUserId(ctx)
	_, err = s.Model(ctx, &handler.Option{FilterAuth: false}).Data(in).Insert()
	return
}

// Delete 删除
func (s *sAdminTest) Delete(ctx context.Context, in adminin.TestDeleteInp) (err error) {
	_, err = s.Model(ctx).Where(dao.Test.Columns().Id, in.Id).Delete()
	return
}

// Status 更新状态
func (s *sAdminTest) Status(ctx context.Context, in adminin.TestStatusInp) (err error) {
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
	_, err = s.Model(ctx).Where(dao.Test.Columns().Id, in.Id).Data(g.Map{
		dao.Test.Columns().Status:    in.Status,
		dao.Test.Columns().UpdatedBy: contexts.GetUserId(ctx),
	}).Update()
	return
}

// Switch 更新开关状态
func (s *sAdminTest) Switch(ctx context.Context, in adminin.TestSwitchInp) (err error) {
	var fields = []string{
		dao.Test.Columns().Switch,
		// ...
	}

	if !validate.InSliceString(fields, in.Key) {
		err = gerror.New("开关键名不在白名单")
		return
	}

	// 修改
	_, err = s.Model(ctx).Where(dao.Test.Columns().Id, in.Id).Data(g.Map{
		in.Key:                       in.Value,
		dao.Test.Columns().UpdatedBy: contexts.GetUserId(ctx),
	}).Update()
	return
}

// MaxSort 最大排序
func (s *sAdminTest) MaxSort(ctx context.Context, in adminin.TestMaxSortInp) (res *adminin.TestMaxSortModel, err error) {
	if err = dao.Test.Ctx(ctx).Fields(dao.Test.Columns().Sort).OrderDesc(dao.Test.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if res == nil {
		res = new(adminin.TestMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(ctx, res.Sort)
	return
}

// View 获取指定字典类型信息
func (s *sAdminTest) View(ctx context.Context, in adminin.TestViewInp) (res *adminin.TestViewModel, err error) {
	err = s.Model(ctx).Where(dao.Test.Columns().Id, in.Id).Scan(&res)
	return
}
