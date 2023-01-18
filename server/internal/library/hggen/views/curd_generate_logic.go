// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package views

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/sysin"
)

const (
	LogicWhereComments      = "\n\t// 查询%s\n"
	LogicWhereNoSupport     = "\t// TODO 暂不支持生成[ %s ]查询方式，请自行补充此处代码！"
	LogicListSimpleSelect   = "\tfields, err := hgorm.GenSelect(ctx, sysin.%sListModel{}, dao.%s)\n\tif err != nil {\n\t\treturn nil, 0, err\n\t}"
	LogicListJoinSelect     = "\t//关联表select\n\tfields, err := hgorm.GenJoinSelect(ctx, %sin.%sListModel{}, dao.%s, []*hgorm.Join{\n%v\t})"
	LogicListJoinOnRelation = "\t// 关联表%s\n\tmod = mod.%s(hgorm.GenJoinOnRelation(\n\t\tdao.%s.Table(), dao.%s.Columns().%s, // 主表表名,关联条件\n\t\tdao.%s.Table(), \"%s\", dao.%s.Columns().%s, // 关联表表名,别名,关联条件\n\t)...)\n\n"
	LogicEditUpdate         = "\t\t_, err = dao.%s.Ctx(ctx).\n\t\t\tFieldsEx(\n%s\t\t\t).\n\t\t\tWhere(dao.%s.Columns().%s, in.%s).Data(in).Update()\n\t\tif err != nil {\n\t\t\terr = gerror.Wrap(err, consts.ErrorORM)\n\t\t\treturn err\n\t\t}\n\t\treturn nil"
	LogicEditInsert         = "\t_, err = dao.%s.Ctx(ctx).\n\t\tFieldsEx(\n%s\t\t).\n\t\tData(in).Insert()\n\tif err != nil {\n\t\terr = gerror.Wrap(err, consts.ErrorORM)\n\t\treturn err\n\t}"
)

func (l *gCurd) logicTplData(ctx context.Context, in *CurdPreviewInput) (data g.Map, err error) {
	data = make(g.Map)
	data["listWhere"] = l.generateLogicListWhere(ctx, in)
	data["listJoin"] = l.generateLogicListJoin(ctx, in)
	data["listOrder"] = l.generateLogicListOrder(ctx, in)
	data["edit"] = l.generateLogicEdit(ctx, in)
	data["switchFields"] = l.generateLogicSwitchFields(ctx, in)
	return
}

func (l *gCurd) generateLogicSwitchFields(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	if in.options.Step.HasSwitch {
		buffer.WriteString("\t\tdao." + in.In.DaoName + ".Columns().Switch,\n")
	}
	return buffer.String()
}

func (l *gCurd) generateLogicEdit(ctx context.Context, in *CurdPreviewInput) g.Map {
	var (
		data           = make(g.Map)
		updateFieldsEx = ""
		updateBuffer   = bytes.NewBuffer(nil)
		insertFieldsEx = ""
		insertBuffer   = bytes.NewBuffer(nil)
	)

	for _, field := range in.masterFields {
		if field.GoName == "UpdatedBy" {
			updateBuffer.WriteString("\t\tin.UpdatedBy = contexts.GetUserId(ctx)\n")
		}

		if field.GoName == "CreatedBy" {
			insertBuffer.WriteString("\tin.CreatedBy = contexts.GetUserId(ctx)\n")
		}

		if field.Index == consts.GenCodesIndexPK || field.GoName == "CreatedAt" || field.GoName == "CreatedBy" || field.GoName == "DeletedAt" {
			updateFieldsEx = updateFieldsEx + "\t\t\t\tdao." + in.In.DaoName + ".Columns()." + field.GoName + ",\n"
		}

		if field.Index == consts.GenCodesIndexPK || field.GoName == "UpdatedBy" || field.GoName == "DeletedAt" {
			insertFieldsEx = insertFieldsEx + "\t\t\t\tdao." + in.In.DaoName + ".Columns()." + field.GoName + ",\n"
		}
	}

	updateBuffer.WriteString(fmt.Sprintf(LogicEditUpdate, in.In.DaoName, updateFieldsEx, in.In.DaoName, in.pk.GoName, in.pk.GoName))
	insertBuffer.WriteString(fmt.Sprintf(LogicEditInsert, in.In.DaoName, insertFieldsEx))

	data["update"] = updateBuffer.String()
	data["insert"] = insertBuffer.String()
	return data
}

func (l *gCurd) generateLogicListOrder(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	if in.options.Step.HasMaxSort {
		buffer.WriteString("OrderAsc(dao." + in.In.DaoName + ".Columns().Sort).")
	}
	buffer.WriteString("OrderDesc(dao." + in.In.DaoName + ".Columns()." + in.pk.GoName + ")")
	return buffer.String()
}

func (l *gCurd) generateLogicListJoin(ctx context.Context, in *CurdPreviewInput) g.Map {
	var data = make(g.Map)
	data["link"] = ""
	if hasEffectiveJoins(in.options.Join) {
		var (
			selectBuffer   = bytes.NewBuffer(nil)
			linkBuffer     = bytes.NewBuffer(nil)
			joinSelectRows string
		)

		for _, join := range in.options.Join {
			if isEffectiveJoin(join) {
				joinSelectRows = joinSelectRows + fmt.Sprintf("\t\t{Dao: dao.%s, Alias: \"%s\"},\n", join.DaoName, join.Alias)
				linkBuffer.WriteString(fmt.Sprintf(LogicListJoinOnRelation, join.Alias, consts.GenCodesJoinLinkMap[join.LinkMode], in.In.DaoName, in.In.DaoName, gstr.CaseCamel(join.MasterField), join.DaoName, join.Alias, join.DaoName, gstr.CaseCamel(join.Field)))
			}
		}

		selectBuffer.WriteString(fmt.Sprintf(LogicListJoinSelect, in.options.TemplateGroup, in.In.VarName, in.In.DaoName, joinSelectRows))

		data["select"] = selectBuffer.String()
		data["link"] = linkBuffer.String()

	} else {
		data["select"] = fmt.Sprintf(LogicListSimpleSelect, in.In.VarName, in.In.DaoName)
	}

	return data
}

func (l *gCurd) generateLogicListWhere(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)

	// 主表
	l.generateLogicListWhereEach(buffer, in.masterFields, in.In.DaoName, "")

	// 关联表
	if hasEffectiveJoins(in.options.Join) {
		for _, v := range in.options.Join {
			if isEffectiveJoin(v) {
				l.generateLogicListWhereEach(buffer, v.Columns, v.DaoName, v.Alias)
			}
		}
	}

	return buffer.String()
}

func (l *gCurd) generateLogicListWhereEach(buffer *bytes.Buffer, fields []*sysin.GenCodesColumnListModel, daoName string, alias string) {
	isLink := false
	if alias != "" {
		alias = `"` + alias + `."+`
		isLink = true
	}
	for _, field := range fields {
		if !field.IsQuery || field.QueryWhere == "" {
			continue
		}

		var (
			linkMode   string
			whereTag   string
			columnName string
		)

		if IsNumberType(field.GoType) {
			linkMode = `in.` + field.GoName + ` > 0`
		} else if field.GoType == GoTypeGTime {
			linkMode = `in.` + field.GoName + ` != nil`
		} else if field.GoType == GoTypeJson {
			linkMode = `!in.` + field.GoName + `.IsNil()`
		} else {
			linkMode = `in.` + field.GoName + ` != ""`
		}

		if field.QueryWhere == WhereModeBetween || field.QueryWhere == WhereModeNotBetween {
			linkMode = `len(in.` + field.GoName + `) == 2`
		}

		buffer.WriteString(fmt.Sprintf(LogicWhereComments, field.Dc))

		// 如果是关联表重新转换字段
		columnName = field.GoName
		if isLink {
			columnName = gstr.CaseCamel(field.Name)
		}

		switch field.QueryWhere {
		case WhereModeEq:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.Where(" + alias + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeNeq:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereNot(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeGt:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereGT(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeGte:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereGTE(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeLt:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereLT(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeLte:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereLTE(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeIn:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereIn(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeNotIn:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereNotIn(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeBetween:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereBetween(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + "[0], in." + field.GoName + "[1])\n\t}"
		case WhereModeNotBetween:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereNotBetween(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + "[0], in." + field.GoName + "[1])\n\t}"
		case WhereModeLike:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereLike(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeLikeAll:
			val := `"%"+in.` + field.GoName + `+"%"`
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereLike(dao." + daoName + ".Columns()." + columnName + ", " + val + ")\n\t}"
		case WhereModeNotLike:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereNotLike(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeJsonContains:
			val := "fmt.Sprintf(`JSON_CONTAINS(%s,'%v')`, dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")"
			whereTag = "\tif in." + field.GoName + linkMode + " {\n\t\tmod = mod.Where(" + val + ")\n\t}"

		default:
			buffer.WriteString(fmt.Sprintf(LogicWhereNoSupport, field.QueryWhere))
		}

		buffer.WriteString(whereTag + "\n")
	}
}
