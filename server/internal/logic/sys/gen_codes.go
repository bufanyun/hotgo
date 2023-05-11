// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hggen"
	"hotgo/internal/model"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

type sSysGenCodes struct{}

func NewSysGenCodes() *sSysGenCodes {
	return &sSysGenCodes{}
}

func init() {
	service.RegisterSysGenCodes(NewSysGenCodes())
}

// Delete 删除
func (s *sSysGenCodes) Delete(ctx context.Context, in sysin.GenCodesDeleteInp) (err error) {
	_, err = dao.SysGenCodes.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

// Edit 修改/新增
func (s *sSysGenCodes) Edit(ctx context.Context, in sysin.GenCodesEditInp) (res *sysin.GenCodesEditModel, err error) {
	if in.GenType == 0 {
		err = gerror.New("生成类型不能为空")
		return
	}
	if in.VarName == "" {
		err = gerror.New("实体名称不能为空")
		return
	}

	if in.GenType == consts.GenCodesTypeCurd {
		var temp *model.GenerateAppCrudTemplate
		cfg := fmt.Sprintf("hggen.application.crud.templates.%v", in.GenTemplate)
		if err = g.Cfg().MustGet(ctx, cfg).Scan(&temp); err != nil {
			return
		}

		if temp == nil {
			err = gerror.Newf("选择的模板不存在:%v", cfg)
			return
		}

		if temp.IsAddon && in.AddonName == "" {
			err = gerror.New("插件模板必须选择一个有效的插件")
			return
		}
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.SysGenCodes.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
		return &sysin.GenCodesEditModel{SysGenCodes: in.SysGenCodes}, nil
	}

	// 新增
	if in.Options.IsNil() {
		in.Options = gjson.New(consts.NilJsonToString)
	}

	in.MasterColumns = gjson.New("[]")
	in.Status = consts.GenCodesStatusWait
	in.CreatedAt = gtime.Now()
	id, err := dao.SysGenCodes.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	in.Id = id
	return &sysin.GenCodesEditModel{SysGenCodes: in.SysGenCodes}, nil
}

// Status 更新部门状态
func (s *sSysGenCodes) Status(ctx context.Context, in sysin.GenCodesStatusInp) (err error) {
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

	_, err = dao.SysGenCodes.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	return
}

// MaxSort 最大排序
func (s *sSysGenCodes) MaxSort(ctx context.Context, in sysin.GenCodesMaxSortInp) (res *sysin.GenCodesMaxSortModel, err error) {
	if in.Id > 0 {
		if err = dao.SysGenCodes.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	if res == nil {
		res = new(sysin.GenCodesMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(ctx, res.Sort)
	return
}

// View 获取指定字典类型信息
func (s *sSysGenCodes) View(ctx context.Context, in sysin.GenCodesViewInp) (res *sysin.GenCodesViewModel, err error) {
	err = dao.SysGenCodes.Ctx(ctx).Where("id", in.Id).Scan(&res)
	return
}

// List 获取列表
func (s *sSysGenCodes) List(ctx context.Context, in sysin.GenCodesListInp) (list []*sysin.GenCodesListModel, totalCount int, err error) {
	mod := dao.SysGenCodes.Ctx(ctx)

	if in.GenType > 0 {
		mod = mod.Where("gen_type", in.GenType)
	}

	if in.VarName != "" {
		mod = mod.Where("var_name", in.VarName)
	}

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	totalCount, err = mod.Count()
	if err != nil || totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	typeSelect, err := hggen.GenTypeSelect(ctx)
	if err != nil {
		return
	}

	getTpGroup := func(row *sysin.GenCodesListModel) string {
		if row == nil {
			return ""
		}
		for _, v := range typeSelect {
			if v.Value == int(row.GenType) {
				for index, template := range v.Templates {
					if index == row.GenTemplate {
						return template.Label
					}
				}
			}
		}

		return ""
	}

	if len(list) > 0 {
		for _, v := range list {
			v.GenTemplateGroup = getTpGroup(v)
		}
	}

	return
}

// Selects 选项
func (s *sSysGenCodes) Selects(ctx context.Context, in sysin.GenCodesSelectsInp) (res *sysin.GenCodesSelectsModel, err error) {
	return hggen.TableSelects(ctx, in)
}

// TableSelect 表选项
func (s *sSysGenCodes) TableSelect(ctx context.Context, in sysin.GenCodesTableSelectInp) (res []*sysin.GenCodesTableSelectModel, err error) {
	var (
		sql           = "SELECT TABLE_NAME as value, TABLE_COMMENT as label FROM information_schema.`TABLES` WHERE TABLE_SCHEMA = '%s'"
		config        = g.DB(in.Name).GetConfig()
		disableTables = g.Cfg().MustGet(ctx, "hggen.disableTables").Strings()
		lists         []*sysin.GenCodesTableSelectModel
	)

	if err = g.DB(in.Name).Ctx(ctx).Raw(fmt.Sprintf(sql, config.Name)).Scan(&lists); err != nil {
		return
	}

	patternStr := `addon_(\w+)_`
	repStr := ``

	for _, v := range lists {
		if gstr.InArray(disableTables, v.Value) {
			continue
		}

		newValue := v.Value
		if config.Prefix != "" {
			newValue = gstr.SubStrFromEx(v.Value, config.Prefix)
		}
		if newValue == "" {
			err = gerror.Newf("表名[%v]前缀必须和配置中的前缀设置[%v] 保持一致", v.Value, config.Prefix)
			return
		}

		// 如果是插件模块，则移除掉插件表前缀
		bt, err := gregex.Replace(patternStr, []byte(repStr), []byte(newValue))
		if err != nil {
			err = gerror.Newf("表名[%v] gregex.Replace err:%v", v.Value, err.Error())
			break
		}

		row := new(sysin.GenCodesTableSelectModel)
		row = v
		row.DefTableComment = v.Label
		row.DaoName = gstr.CaseCamel(newValue)
		row.DefVarName = gstr.CaseCamel(string(bt))
		row.DefAlias = gstr.CaseCamelLower(newValue)
		row.Name = fmt.Sprintf("%s (%s)", v.Value, v.Label)
		row.Label = row.Name

		res = append(res, row)
	}

	return
}

// ColumnSelect 表字段选项
func (s *sSysGenCodes) ColumnSelect(ctx context.Context, in sysin.GenCodesColumnSelectInp) (res []*sysin.GenCodesColumnSelectModel, err error) {
	var (
		sql    = "select COLUMN_NAME as value,COLUMN_COMMENT as label from information_schema.COLUMNS where TABLE_SCHEMA = '%s' and TABLE_NAME = '%s'"
		config = g.DB(in.Name).GetConfig()
	)

	if err = g.DB(in.Name).Ctx(ctx).Raw(fmt.Sprintf(sql, config.Name, in.Table)).Scan(&res); err != nil {
		return
	}

	if len(res) == 0 {
		err = gerror.Newf("table does not exist:%v", in.Table)
		return
	}

	for k, v := range res {
		res[k].Name = fmt.Sprintf("%s (%s)", v.Value, v.Label)
		res[k].Label = res[k].Name
	}

	return
}

// ColumnList 表字段列表
func (s *sSysGenCodes) ColumnList(ctx context.Context, in sysin.GenCodesColumnListInp) (res []*sysin.GenCodesColumnListModel, err error) {
	return hggen.TableColumns(ctx, in)
}

// Preview 生成预览
func (s *sSysGenCodes) Preview(ctx context.Context, in sysin.GenCodesPreviewInp) (res *sysin.GenCodesPreviewModel, err error) {
	return hggen.Preview(ctx, in)
}

// Build 提交生成
func (s *sSysGenCodes) Build(ctx context.Context, in sysin.GenCodesBuildInp) (err error) {
	// 先保存配置
	if _, err = s.Edit(ctx, sysin.GenCodesEditInp{SysGenCodes: in.SysGenCodes}); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if err = s.Status(ctx, sysin.GenCodesStatusInp{Id: in.Id, Status: consts.GenCodesStatusOk}); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if err = hggen.Build(ctx, in); err != nil {
		_ = s.Status(ctx, sysin.GenCodesStatusInp{Id: in.Id, Status: consts.GenCodesStatusFail})
		return err
	}

	return
}
