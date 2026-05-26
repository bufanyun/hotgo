// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
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
	LogicWhereComments  = "\n\t// 查询%s\n"
	LogicWhereNoSupport = "\t// TODO 暂不支持生成[ %s ]查询方式，请自行补充此处代码！"
	LogicEditUpdate     = "\tif _, err = s.Model(ctx%s).\n\t\t\tFields(%sin.%sUpdateFields{}).\n\t\t\tWherePri(in.%s).Data(in).Update(); err != nil {\n\t\t\terr = gerror.Wrap(err, \"修改%s失败，请稍后重试！\")\n\t\t}\n\t\treturn"
	LogicEditInsert     = "\tif _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).\n\t\tFields(%sin.%sInsertFields{}).\n\t\tData(in).OmitEmptyData().Insert(); err != nil {\n\t\terr = gerror.Wrap(err, \"新增%s失败，请稍后重试！\")\n\t}"
	LogicEditUnique     = "\t// 验证'%s'唯一\n\tif err = hgorm.IsUnique(ctx, &dao.%s, g.Map{dao.%s.Columns().%s: in.%s}, \"%s已存在\", in.%s); err != nil {\n\t\treturn\n\t}\n"
	LogicSwitchUpdate   = "g.Map{\n\t\tin.Key:                       in.Value,\n%s}"
	LogicStatusUpdate   = "g.Map{\n\t\tdao.%s.Columns().Status:    in.Status,\n%s}"
	LogicDeletedUpdate  = "g.Map{\n%s}"
)

func (l *gCurd) logicTplData(ctx context.Context, in *CurdPreviewInput) (data g.Map, err error) {
	data = make(g.Map)
	data["listWhere"] = l.generateLogicListWhere(ctx, in)
	data["listJoin"] = l.generateLogicListJoin(ctx, in)
	data["listFields"] = l.generateLogicListFields(ctx, in)
	data["listOrder"] = l.generateLogicListOrder(ctx, in)
	data["edit"] = l.generateLogicEdit(ctx, in)
	data["switchFields"] = l.generateLogicSwitchFields(ctx, in)
	data["switchUpdate"] = l.generateLogicSwitchUpdate(ctx, in)
	data["statusUpdate"] = l.generateLogicStatusUpdate(ctx, in)
	data["deletedUpdate"] = l.generateLogicDeletedUpdate(ctx, in)
	return
}

func (l *gCurd) generateLogicDeletedUpdate(ctx context.Context, in *CurdPreviewInput) string {
	isDestroy := false
	var update string
	for _, field := range in.masterFields {
		if field.GoName == "DeletedBy" {
			update += "\t\tdao." + in.In.DaoName + ".Columns().DeletedBy: contexts.GetUserId(ctx),\n"
			isDestroy = true
		}
		if field.GoName == "DeletedAt" {
			update += "\t\tdao." + in.In.DaoName + ".Columns().DeletedAt: gtime.Now(),\n"
		}
	}

	if !isDestroy {
		return ""
	}

	update += "\t"
	return fmt.Sprintf(LogicDeletedUpdate, update)
}

func (l *gCurd) generateLogicStatusUpdate(ctx context.Context, in *CurdPreviewInput) string {
	var update string
	for _, field := range in.masterFields {
		if field.GoName == "UpdatedBy" {
			update += "\t\tdao." + in.In.DaoName + ".Columns().UpdatedBy: contexts.GetUserId(ctx),\n"
		}
	}

	update += "\t"
	return fmt.Sprintf(LogicStatusUpdate, in.In.DaoName, update)
}

func (l *gCurd) generateLogicSwitchUpdate(ctx context.Context, in *CurdPreviewInput) string {
	var update string
	for _, field := range in.masterFields {
		if field.GoName == "UpdatedBy" {
			update += "\t\tdao." + in.In.DaoName + ".Columns().UpdatedBy: contexts.GetUserId(ctx),\n"
		}
	}

	update += "\t"
	return fmt.Sprintf(LogicSwitchUpdate, update)
}

func (l *gCurd) generateLogicSwitchFields(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	if in.options.Step.HasSwitch {
		for _, field := range in.masterFields {
			if field.FormMode == "Switch" {
				buffer.WriteString("\t\tdao." + in.In.DaoName + ".Columns()." + field.GoName + ",\n")
			}
		}
	}
	return buffer.String()
}

func (l *gCurd) generateLogicEdit(ctx context.Context, in *CurdPreviewInput) g.Map {
	var (
		data             = make(g.Map)
		updateBuffer     = bytes.NewBuffer(nil)
		insertBuffer     = bytes.NewBuffer(nil)
		uniqueBuffer     = bytes.NewBuffer(nil)
		validationBuffer = bytes.NewBuffer(nil)
	)

	for _, field := range in.masterFields {
		if field.GoName == "UpdatedBy" {
			updateBuffer.WriteString("\t\tin.UpdatedBy = contexts.GetUserId(ctx)\n")
		}

		if field.GoName == "CreatedBy" {
			insertBuffer.WriteString("\tin.CreatedBy = contexts.GetUserId(ctx)\n")
		}

		if field.Unique {
			uniqueBuffer.WriteString(fmt.Sprintf(LogicEditUnique, field.GoName, in.In.DaoName, in.In.DaoName, field.GoName, field.GoName, field.Dc, in.pk.GoName))
		}

		// 添加 YAML 格式验证
		if field.IsEdit && field.FormRole == FormRoleYaml {
			validationBuffer.WriteString(fmt.Sprintf(EditInpValidatorYaml, field.GoName, field.Dc))
		}
	}

	notFilterAuth := ""
	if gstr.InArray(in.options.ColumnOps, "notFilterAuth") {
		notFilterAuth = ", &handler.Option{FilterAuth: false}"
	}

	updateBuffer.WriteString(fmt.Sprintf(LogicEditUpdate, notFilterAuth, in.options.TemplateGroup, in.In.VarName, in.pk.GoName, in.In.TableComment))
	insertBuffer.WriteString(fmt.Sprintf(LogicEditInsert, in.options.TemplateGroup, in.In.VarName, in.In.TableComment))

	data["update"] = updateBuffer.String()
	data["insert"] = insertBuffer.String()
	data["unique"] = uniqueBuffer.String()
	data["validation"] = validationBuffer.String()
	return data
}

func (l *gCurd) generateLogicListOrder(ctx context.Context, in *CurdPreviewInput) string {
	statement := ""
	if hasEffectiveJoins(in.options.Join) {
		statement = "dao." + in.In.DaoName + ".Table() + \".\" +"
	}
	buffer := bytes.NewBuffer(nil)
	if in.options.Step.HasMaxSort {
		buffer.WriteString("OrderAsc(" + statement + "dao." + in.In.DaoName + ".Columns().Sort).")
	}
	buffer.WriteString("OrderDesc(" + statement + "dao." + in.In.DaoName + ".Columns()." + in.pk.GoName + ")")
	return buffer.String()
}

func (l *gCurd) generateLogicListJoin(ctx context.Context, in *CurdPreviewInput) (link string) {
	connector := `"="`
	if hasEffectiveJoins(in.options.Join) {
		linkBuffer := bytes.NewBuffer(nil)
		for _, join := range in.options.Join {
			if isEffectiveJoin(join) {
				linkBuffer.WriteString("\tmod = mod." + consts.GenCodesJoinLinkMap[join.LinkMode] + "OnFields(dao." + join.DaoName + ".Table(), dao." + in.In.DaoName + ".Columns()." + gstr.CaseCamel(join.MasterField) + "," + connector + ", dao." + join.DaoName + ".Columns()." + gstr.CaseCamel(join.Field) + ")\n")
			}
		}
		link = linkBuffer.String()
	}
	return
}

func (l *gCurd) generateLogicListFields(ctx context.Context, in *CurdPreviewInput) (fields string) {
	selectBuffer := bytes.NewBuffer(nil)
	if hasEffectiveJoins(in.options.Join) {
		selectBuffer.WriteString("mod = mod.FieldsPrefix(dao." + in.In.DaoName + ".Table(), " + in.options.TemplateGroup + "in." + in.In.VarName + "ListModel{})\n")
		for _, join := range in.options.Join {
			if isEffectiveJoin(join) {
				selectBuffer.WriteString("mod = mod.Fields(hgorm.JoinFields(ctx, " + in.options.TemplateGroup + "in." + in.In.VarName + "ListModel{}, &dao." + join.DaoName + ", \"" + join.Alias + "\"))\n")
			}
		}
		fields = selectBuffer.String()
	} else {
		fields = fmt.Sprintf("mod = mod.Fields(%sin.%sListModel{})", in.options.TemplateGroup, in.In.VarName)
	}
	return
}

func (l *gCurd) generateLogicListWhere(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)

	// 主表
	l.generateLogicListWhereEach(buffer, in, in.masterFields, in.In.DaoName, "")

	// 关联表
	if hasEffectiveJoins(in.options.Join) {
		for _, v := range in.options.Join {
			if isEffectiveJoin(v) {
				l.generateLogicListWhereEach(buffer, in, v.Columns, v.DaoName, v.Alias)
			}
		}
	}
	return buffer.String()
}

func (l *gCurd) generateLogicListWhereEach(buffer *bytes.Buffer, in *CurdPreviewInput, fields []*sysin.GenCodesColumnListModel, daoName string, alias string) {
	isLink := false
	if alias != "" {
		alias = `"` + alias + `."+`
		isLink = true
	}

	tablePrefix := ""
	wherePrefix := "Where"
	if isLink {
		wherePrefix = "WherePrefix"
		tablePrefix = "dao." + daoName + ".Table(), "
	}

	for _, field := range fields {
		isQuery := false
		// 树表查询上级
		if in.options.Step.IsTreeTable && IsPidName(field.Name) {
			isQuery = true
			field.QueryWhere = WhereModeEq
		}

		if (!field.IsQuery && !isQuery) || field.QueryWhere == "" {
			continue
		}

		buffer.WriteString(fmt.Sprintf(LogicWhereComments, field.Dc))

		var (
			linkMode   string
			whereTag   string
			columnName string
		)

		// 查询用户摘要
		if field.IsQuery && in.options.Step.HasQueryMemberSummary && IsMemberSummaryField(field.Name) {
			servicePackName := "service"
			if in.options.Step.IsAddon {
				servicePackName = "isc"
			}
			buffer.WriteString(fmt.Sprintf("if in.%v != \"\" {\n\t\t\t\tids, err := %v.AdminMember().GetIdsByKeyword(ctx, in.%v)\n\t\t\t\tif err != nil {\n\t\t\t\t\treturn nil, 0, err\n\t\t\t\t}\n\t\t\t\tmod = mod.WhereIn(dao.%v.Columns().%v, ids)\n\t\t\t}\n", field.GoName, servicePackName, field.GoName, in.In.DaoName, field.GoName))
			continue
		}

		if IsNumberType(field.GoType) {
			linkMode = `in.` + field.GoName + ` > 0`
		} else if field.GoType == GoTypeBool {
			linkMode = `true` // bool 类型始终可以查询
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

		// 如果是关联表重新转换字段
		columnName = field.GoName
		if isLink {
			columnName = gstr.CaseCamel(field.Name)
		}

		switch field.QueryWhere {
		case WhereModeEq:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeNeq:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "Not(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeGt:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "GT(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeGte:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "GTE(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeLt:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "LT(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeLte:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "LTE(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeIn:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "In(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeNotIn:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "NotIn(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeBetween:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "Between(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + "[0], in." + field.GoName + "[1])\n\t}"
		case WhereModeNotBetween:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "NotBetween(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + "[0], in." + field.GoName + "[1])\n\t}"
		case WhereModeLike:
			fieldValue := "in." + field.GoName
			if field.GoType != GoTypeString && field.GoType != GoTypeBytes {
				fieldValue = "gconv.String(in." + field.GoName + ")"
			}
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "Like(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", " + fieldValue + ")\n\t}"
		case WhereModeLikeAll:
			val := `"%"+in.` + field.GoName + `+"%"`
			if field.GoType != GoTypeString && field.GoType != GoTypeBytes {
				val = `"%"+gconv.String(in.` + field.GoName + `)+"%"`
			}
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "Like(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", " + val + ")\n\t}"
		case WhereModeNotLike:
			fieldValue := "in." + field.GoName
			if field.GoType != GoTypeString && field.GoType != GoTypeBytes {
				fieldValue = "gconv.String(in." + field.GoName + ")"
			}
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "NotLike(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", " + fieldValue + ")\n\t}"
		case WhereModeJsonContains:
			val := tablePrefix + `"JSON_CONTAINS("+dao.` + daoName + `.Columns().` + columnName + `+",?)", in.` + field.GoName
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "(" + val + ")\n\t}"

		default:
			buffer.WriteString(fmt.Sprintf(LogicWhereNoSupport, field.QueryWhere))
		}

		buffer.WriteString(whereTag + "\n")
	}
}
