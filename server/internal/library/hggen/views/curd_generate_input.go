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
	"strings"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/olekukonko/tablewriter"
)

const (
	InputTypeListInp          = 1 // 列表输入
	InputTypeListModel        = 2 // 列表输出
	InputTypeExportModel      = 3 // 列表导出
	InputTypeEditInpValidator = 4 // 添加&编辑验证器
	InputTypeUpdateFields     = 5 // 编辑修改过滤字段
	InputTypeInsertFields     = 6 // 编辑新增过滤字段
	EditInpValidatorGenerally = "if err := g.Validator().Rules(\"%s\").Data(in.%s).Messages(\"%s\").Run(ctx); err != nil {\n\t\treturn err.Current()\n\t}\n"
)

func (l *gCurd) inputTplData(ctx context.Context, in *CurdPreviewInput) (data g.Map, err error) {
	data = make(g.Map)
	data["listInpColumns"] = l.generateInputListColumns(ctx, in, InputTypeListInp)
	data["listModelColumns"] = l.generateInputListColumns(ctx, in, InputTypeListModel)
	data["exportModelColumns"] = l.generateInputListColumns(ctx, in, InputTypeExportModel)
	data["editInpValidator"] = l.generateInputListColumns(ctx, in, InputTypeEditInpValidator)
	data["updateFieldsColumns"] = l.generateInputListColumns(ctx, in, InputTypeUpdateFields)
	data["insertFieldsColumns"] = l.generateInputListColumns(ctx, in, InputTypeInsertFields)
	return
}

func (l *gCurd) generateInputListColumns(ctx context.Context, in *CurdPreviewInput, inputType int) string {
	buffer := bytes.NewBuffer(nil)
	index := 0
	array := make([][]string, 1000)
	// 主表
	for _, field := range in.masterFields {
		row := l.generateStructFieldDefinition(field, inputType)
		if row == nil {
			continue
		}
		array[index] = row
		index++
	}

	// 关联表
	if len(in.options.Join) > 0 {
		for _, v := range in.options.Join {
			if !isEffectiveJoin(v) {
				continue
			}
			for _, field := range v.Columns {
				row := l.generateStructFieldDefinition(field, inputType)
				if row != nil {
					array[index] = row
					index++
				}
			}
		}
	}

	tw := tablewriter.NewWriter(buffer)
	tw.SetBorder(false)
	tw.SetRowLine(false)
	tw.SetAutoWrapText(false)
	tw.SetColumnSeparator("")
	tw.AppendBulk(array)
	tw.Render()
	stContent := buffer.String()
	// Let's do this hack of table writer for indent!
	stContent = gstr.Replace(stContent, "  #", "")
	stContent = gstr.Replace(stContent, "` ", "`")
	stContent = gstr.Replace(stContent, "``", "")
	stContent = removeEndWrap(stContent)

	buffer.Reset()
	buffer.WriteString(stContent)
	return buffer.String()
}

// generateStructFieldForModel generates and returns the attribute definition for specified field.
func (l *gCurd) generateStructFieldDefinition(field *sysin.GenCodesColumnListModel, inputType int) []string {
	var (
		tagKey         = "`"
		result         = []string{"    #" + field.GoName}
		descriptionTag = gstr.Replace(formatComment(field.Dc), `"`, `\"`)
	)

	switch inputType {
	case InputTypeListInp:
		if !field.IsQuery {
			return nil
		}

		if field.QueryWhere == WhereModeBetween {
			result = append(result, " #[]"+field.GoType)
		} else {
			result = append(result, " #"+field.GoType)
		}
		result = append(result, " #"+fmt.Sprintf(tagKey+`json:"%s"`, field.TsName))
		result = append(result, " #"+fmt.Sprintf(`dc:"%s"`+tagKey, descriptionTag))

	case InputTypeListModel:
		if !field.IsList {
			return nil
		}

		result = append(result, " #"+field.GoType)
		result = append(result, " #"+fmt.Sprintf(tagKey+`json:"%s"`, field.TsName))
		result = append(result, " #"+fmt.Sprintf(`dc:"%s"`+tagKey, descriptionTag))
	case InputTypeExportModel:
		if !field.IsExport {
			return nil
		}

		result = append(result, " #"+field.GoType)
		result = append(result, " #"+fmt.Sprintf(tagKey+`json:"%s"`, field.TsName))
		result = append(result, " #"+fmt.Sprintf(`dc:"%s"`+tagKey, descriptionTag))
	case InputTypeEditInpValidator:
		if !field.IsEdit {
			return nil
		}
		if !field.Required && (field.FormRole == "none" || field.FormRole == "") {
			return nil
		}
		rule := "// 验证" + field.Dc + "\n"
		if field.Required && (field.FormRole == FormRoleNone || field.FormRole == "") {
			field.FormRole = "required"
		}
		if err, s := makeValidatorFunc(field); err != nil {
			return nil
		} else {
			rule += s
		}
		result = []string{rule}
	case InputTypeUpdateFields:
		if !field.IsEdit && field.GoName != "UpdatedBy" {
			return nil
		}

		result = append(result, " #"+field.GoType)
		result = append(result, " #"+fmt.Sprintf(tagKey+`json:"%s"`, field.TsName))
		result = append(result, " #"+fmt.Sprintf(`dc:"%s"`+tagKey, descriptionTag))
	case InputTypeInsertFields:
		if !field.IsEdit && field.GoName != "CreatedBy" {
			return nil
		}

		result = append(result, " #"+field.GoType)
		result = append(result, " #"+fmt.Sprintf(tagKey+`json:"%s"`, field.TsName))
		result = append(result, " #"+fmt.Sprintf(`dc:"%s"`+tagKey, descriptionTag))
	default:
		panic("inputType is invalid")
	}
	return result
}

func makeValidatorFunc(field *sysin.GenCodesColumnListModel) (err error, rule string) {
	if field.FormRole == "required" {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "required", field.GoName, field.Dc+"不能为空")
	} else if field.FormRole == FormRoleIp {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "ip", field.GoName, field.Dc+"必须为IPV4或IPV6")
	} else if field.FormRole == FormRolePercentage {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "min:0|max:100", field.GoName, field.Dc+"必须0-100之间")
	} else if field.FormRole == FormRoleTel {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "telephone", field.GoName, field.Dc+"不是座机号码")
	} else if field.FormRole == FormRolePhone {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "phone", field.GoName, field.Dc+"不是手机号码")
	} else if field.FormRole == FormRoleQq {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "qq", field.GoName, field.Dc+"不是QQ号码")
	} else if field.FormRole == FormRoleEmail {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "email", field.GoName, field.Dc+"不是邮箱地址")
	} else if field.FormRole == FormRoleIdCard {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "resident-id", field.GoName, field.Dc+"不是身份证号码")
	} else if field.FormRole == FormRoleNum {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "min:1", field.GoName, field.Dc+"必须大于0")
	} else if field.FormRole == FormRoleBankCard {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "bank-card", field.GoName, field.Dc+"不是银行卡号")
	} else if field.FormRole == FormRoleWeibo {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "regex:^[0-9a-zA-Z\\u4e00-\\u9fa5_-]*$", field.GoName, field.Dc+"不是微博号")
	} else if field.FormRole == FormRoleUserName {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "regex:^[0-9a-zA-Z]{6,16}$", field.GoName, field.Dc+"必须为6-16位由字母和数字组成")
	} else if field.FormRole == FormRoleAccount {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "regex:^[\\w_\\d]{6,16}$", field.GoName, field.Dc+"必须为6-16位由字母、数字或下划线组成")
	} else if field.FormRole == FormRolePassword {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "regex:^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,18}$", field.GoName, field.Dc+"必须包含6-18为字母和数字")
	} else if field.FormRole == FormRoleAmount {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "regex:(^[0-9]{1,10}$)|(^[0-9]{1,10}[\\\\.]{1}[0-9]{1,2}$)", field.GoName, field.Dc+"最多允许输入10位整数及2位小数")
	} else {
		err = gerror.New("not support")
	}

	// 生成验证字典
	if field.DictType > 0 {
		var (
			ctx         = context.Background()
			valueType   gdb.Value
			resultValue gdb.Result
		)
		valueType, err = dao.SysDictType.Ctx(ctx).Fields(dao.SysDictType.Columns().Type).Where(dao.SysDictType.Columns().Id, field.DictType).Value()
		if err != nil {
			return
		}
		resultValue, err = dao.SysDictData.Ctx(context.Background()).Fields(dao.SysDictData.Columns().Value).Where(dao.SysDictData.Columns().Type, valueType.String()).All()
		if err != nil {
			return
		}
		if resultValue.Len() > 0 {
			names := make([]string, 0)
			for _, item := range resultValue {
				names = append(names, item["value"].String())
			}
			dictRule := "in:" + strings.Join(names, ",")
			rule += fmt.Sprintf(EditInpValidatorGenerally, dictRule, field.GoName, field.Dc+"值不正确")
			err = nil
		}
	}
	return
}
