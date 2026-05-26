// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package views

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/dict"
	"hotgo/internal/model/input/sysin"
	"hotgo/utility/convert"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type StateItem struct {
	Name         string
	DefaultValue interface{}
	Dc           string
	DataType     string // 新增字段：存储字段类型信息
}

func (l *gCurd) webModelTplData(ctx context.Context, in *CurdPreviewInput) (data g.Map, err error) {
	data = make(g.Map)
	data["stateItems"] = l.generateWebModelStateItems(ctx, in)
	data["rules"] = l.generateWebModelRules(ctx, in)
	data["formSchema"] = l.generateWebModelFormSchema(ctx, in)
	if data["columns"], err = l.generateWebModelColumns(ctx, in); err != nil {
		return nil, err
	}

	// 根据表单生成情况，按需导包
	data["import"], data["const"] = l.generateWebModelImport(ctx, in)
	return
}

func (l *gCurd) generateWebModelImport(ctx context.Context, in *CurdPreviewInput) (string, string) {
	importBuffer := bytes.NewBuffer(nil)
	constBuffer := bytes.NewBuffer(nil)

	// 导入基础组件
	if len(in.options.Step.ImportModel.NaiveUI) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.options.Step.ImportModel.NaiveUI) + " from 'naive-ui';\n")
	}

	importBuffer.WriteString("import { cloneDeep } from 'lodash-es';\n")

	// 导入表单搜索
	if in.options.Step.HasSearchForm {
		importBuffer.WriteString("import { FormSchema } from '@/components/Form';\n")
	}

	// 导入工具类
	if len(in.options.Step.ImportModel.UtilsIs) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.options.Step.ImportModel.UtilsIs) + " from '@/utils/is';\n")
	}

	if len(in.options.Step.ImportModel.UtilsUrl) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.options.Step.ImportModel.UtilsUrl) + " from '@/utils/urlUtils';\n")
	}

	if len(in.options.Step.ImportModel.UtilsDate) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.options.Step.ImportModel.UtilsDate) + " from '@/utils/dateUtil';\n")
	}

	if in.options.Step.HasRulesValidator {
		importBuffer.WriteString("import { validate } from '@/utils/validateUtil';\n")
	}

	if len(in.options.Step.ImportModel.UtilsHotGo) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.options.Step.ImportModel.UtilsHotGo) + " from '@/utils/hotgo';\n")
	}

	if len(in.options.Step.ImportModel.UtilsIndex) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.options.Step.ImportModel.UtilsIndex) + " from '@/utils';\n")
	}

	// 导入api
	var importApiMethod []string
	if in.options.Step.HasSwitch {
		importApiMethod = append(importApiMethod, "Switch")
	}
	if in.options.Step.IsTreeTable {
		importApiMethod = append(importApiMethod, "TreeOption")
	}
	if len(importApiMethod) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(importApiMethod) + " from '" + in.options.ImportWebApi + "';\n")
	}

	// 导入字典选项
	if in.options.DictOps.Has {
		importBuffer.WriteString("import { useDictStore } from '@/store/modules/dict';\n")
		constBuffer.WriteString("const dict = useDictStore();\n")
	}

	if in.options.Step.HasSwitch {
		importBuffer.WriteString("import { usePermission } from '@/hooks/web/usePermission';\n")
		constBuffer.WriteString("const { hasPermission } = usePermission();\n")
		constBuffer.WriteString("const $message = window['$message'];\n")
	}

	return importBuffer.String(), constBuffer.String()
}

func (l *gCurd) generateWebModelStateItems(ctx context.Context, in *CurdPreviewInput) (items []*StateItem) {
	for _, field := range in.masterFields {
		dataType := field.DataType
		if dataType == "" && field.SqlType != "" {
			if parts := strings.SplitN(field.SqlType, "(", 2); len(parts) > 0 {
				dataType = parts[0]
			}
		}

		isStr := isStringType(field.TsType, dataType)
		isArray := strings.HasPrefix(dataType, "_") || strings.Contains(field.SqlType, "[]")

		var value = field.DefaultValue
		if value == nil {
			if isArray {
				value = "null"
			} else if isStr {
				value = ""
			} else {
				value = "null"
			}
		}
		if value == "" {
			if isArray || !isStr {
				value = "null"
			}
		} else if valueStr, ok := value.(string); ok && isStr {
			value = valueStr
		}

		// 选项组件默认值调整
		if gconv.Int(value) == 0 && IsSelectFormMode(field.FormMode) {
			value = "null"
		}

		if field.Name == "status" {
			value = 1
		}
		if field.FormMode == FormModeSwitch {
			value = 2
		}
		if field.FormMode == FormModeInputDynamic {
			value = "[]"
		}
		items = append(items, &StateItem{
			Name:         field.TsName,
			DefaultValue: value,
			DataType:     dataType,
			Dc:           field.Dc,
		})

		// 查询用户摘要
		if field.IsList && in.options.Step.HasHookMemberSummary && IsMemberSummaryField(field.Name) {
			items = append(items, &StateItem{
				Name:         field.TsName + "Summa?: null | MemberSumma",
				DefaultValue: "null",
				DataType:     dataType,
				Dc:           field.Dc + "摘要信息",
			})
		}
	}
	return
}

func (l *gCurd) generateWebModelDictOptions(ctx context.Context, in *CurdPreviewInput) error {
	type DictType struct {
		Id   int64  `json:"id"`
		Type string `json:"type"`
	}

	var (
		dictTypeIds         []int64
		dictTypeList        []*DictType
		builtinDictTypeIds  []int64
		builtinDictTypeList []*DictType
	)

	for _, field := range in.masterFields {
		if field.DictType > 0 {
			dictTypeIds = append(dictTypeIds, field.DictType)
		}

		if field.DictType < 0 {
			builtinDictTypeIds = append(builtinDictTypeIds, field.DictType)
		}
	}

	dictTypeIds = convert.UniqueSlice(dictTypeIds)
	builtinDictTypeIds = convert.UniqueSlice(builtinDictTypeIds)

	if len(dictTypeIds) == 0 && len(builtinDictTypeIds) == 0 {
		return nil
	}

	if len(dictTypeIds) > 0 {
		err := dao.SysDictType.Ctx(ctx).
			Fields(dao.SysDictType.Columns().Id, dao.SysDictType.Columns().Type).
			WhereIn(dao.SysDictType.Columns().Id, dictTypeIds).
			Scan(&dictTypeList)
		if err != nil {
			return err
		}
	}

	if len(builtinDictTypeIds) > 0 {
		for _, id := range builtinDictTypeIds {
			typ, err := dict.GetTypeById(ctx, id)
			if err != nil && !errors.Is(err, dict.NotExistKeyError) {
				return err
			}
			if len(typ) > 0 {
				row := new(DictType)
				row.Id = id
				row.Type = typ
				builtinDictTypeList = append(builtinDictTypeList, row)
			}
		}
	}

	if len(dictTypeList) == 0 && len(builtinDictTypeList) == 0 {
		return nil
	}

	if len(builtinDictTypeList) > 0 {
		dictTypeList = append(dictTypeList, builtinDictTypeList...)
	}

	in.options.DictOps.Has = true

	for _, v := range dictTypeList {
		// 字段映射字典
		for _, field := range in.masterFields {
			if field.DictType != 0 && v.Id == field.DictType {
				in.options.dictMap[field.TsName] = v.Type
				in.options.DictOps.Schemas = append(in.options.DictOps.Schemas, &OptionsSchemasField{
					Field: field.TsName,
					Type:  v.Type,
				})
			}
		}
		in.options.DictOps.Types = append(in.options.DictOps.Types, v.Type)
	}
	return nil
}

func (l *gCurd) generateWebModelRules(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const rules = {\n")
	for _, field := range in.masterFields {
		if !field.IsEdit || (!field.Required && (field.FormRole == "" || field.FormRole == FormRoleNone)) {
			continue
		}

		in.options.Step.HasRules = true
		if field.FormRole == "" || field.FormRole == FormRoleNone || field.FormRole == "required" {
			buffer.WriteString(fmt.Sprintf("  %s: {\n    required: %v,\n    trigger: ['blur', 'input'],\n    type: '%s',\n    message: '请输入%s',\n  },\n", field.TsName, field.Required, field.TsType, field.Dc))
		} else {
			in.options.Step.HasRulesValidator = true
			buffer.WriteString(fmt.Sprintf("  %s: {\n    required: %v,\n    trigger: ['blur', 'input'],\n    type: '%s',\n    validator: validate.%v,\n  },\n", field.TsName, field.Required, field.TsType, field.FormRole))
		}
	}
	buffer.WriteString("};\n")
	return buffer.String()
}

func (l *gCurd) generateWebModelFormSchema(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const schemas = ref<FormSchema[]>([\n")

	// 主表
	l.generateWebModelFormSchemaEach(buffer, in.masterFields, in)

	// 关联表
	if len(in.options.Join) > 0 {
		for _, v := range in.options.Join {
			if !isEffectiveJoin(v) {
				continue
			}
			l.generateWebModelFormSchemaEach(buffer, v.Columns, in)
		}
	}

	buffer.WriteString("]);\n")
	return buffer.String()
}

// generateWebDictOption 生产字典选项
func (l *gCurd) generateWebDictOption(typ any) string {
	return fmt.Sprintf(`dict.getOption('%v')`, typ)
}

func (l *gCurd) generateWebModelFormSchemaEach(buffer *bytes.Buffer, fields []*sysin.GenCodesColumnListModel, in *CurdPreviewInput) {
	for _, field := range fields {
		if !field.IsQuery {
			continue
		}
		in.options.Step.HasSearchForm = true

		// 查询用户摘要
		if field.IsQuery && in.options.Step.HasQueryMemberSummary && IsMemberSummaryField(field.Name) {
			buffer.WriteString(fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      placeholder: '请输入ID|用户名|姓名|手机号',\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NInput", field.Dc))
			continue
		}

		var (
			defaultComponent = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      placeholder: '请输入%s',\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NInput", field.Dc, field.Dc)
			component        string
		)

		// 这里根据编辑表单组件来进行推断，如果没有则使用默认input，这可能会导致和查询条件所需参数不符的情况
		switch field.FormMode {
		case FormModeInput, FormModeInputTextarea, FormModeInputEditor, FormModeInputYaml:
			component = defaultComponent

		case FormModeInputNumber:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      placeholder: '请输入%s',\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NInputNumber", field.Dc, field.Dc)

		case FormModeDate:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NDatePicker", field.Dc, "date", "defShortcuts()")
			in.options.Step.ImportModel.UtilsDate = append(in.options.Step.ImportModel.UtilsDate, "defShortcuts")

		case FormModeDateRange:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NDatePicker", field.Dc, "daterange", "defRangeShortcuts()")
			in.options.Step.ImportModel.UtilsDate = append(in.options.Step.ImportModel.UtilsDate, "defRangeShortcuts")

		case FormModeTime:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NDatePicker", field.Dc, "datetime", "defShortcuts()")
			in.options.Step.ImportModel.UtilsDate = append(in.options.Step.ImportModel.UtilsDate, "defShortcuts")

		case FormModeTimeRange:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NDatePicker", field.Dc, "datetimerange", "defRangeShortcuts()")
			in.options.Step.ImportModel.UtilsDate = append(in.options.Step.ImportModel.UtilsDate, "defRangeShortcuts")

		case FormModeSwitch:
			fallthrough
		case FormModeRadio:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    giProps: {\n      //span: 24,\n    },\n    componentProps: {\n      options: %v,\n      onUpdateChecked: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NRadioGroup", field.Dc, l.generateWebDictOption(in.options.dictMap[field.TsName]))

		case FormModeCheckbox:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    giProps: {\n      span: 1,\n    },\n    componentProps: {\n      placeholder: '请选择%s',\n      options: %v,\n      onUpdateChecked: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NCheckbox", field.Dc, field.Dc, l.generateWebDictOption(in.options.dictMap[field.TsName]))

		case FormModeSelect:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    defaultValue: null,\n    componentProps: {\n      placeholder: '请选择%s',\n      options: %v,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NSelect", field.Dc, field.Dc, l.generateWebDictOption(in.options.dictMap[field.TsName]))

		case FormModeSelectMultiple:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    defaultValue: null,\n    componentProps: {\n      multiple: true,\n      placeholder: '请选择%s',\n      options: %v,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NSelect", field.Dc, field.Dc, l.generateWebDictOption(in.options.dictMap[field.TsName]))

		default:
			component = defaultComponent
		}

		buffer.WriteString(component)
	}
}

func (l *gCurd) generateWebModelColumns(ctx context.Context, in *CurdPreviewInput) (string, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const columns = [\n")

	// 主表
	if err := l.generateWebModelColumnsEach(buffer, in, in.masterFields); err != nil {
		return "", err
	}

	// 关联表
	if len(in.options.Join) > 0 {
		for _, v := range in.options.Join {
			if !isEffectiveJoin(v) {
				continue
			}
			if err := l.generateWebModelColumnsEach(buffer, in, v.Columns); err != nil {
				return "", err
			}
		}
	}

	buffer.WriteString("];\n")
	return buffer.String(), nil
}

func (l *gCurd) generateWebModelColumnsEach(buffer *bytes.Buffer, in *CurdPreviewInput, fields []*sysin.GenCodesColumnListModel) (err error) {
	for _, field := range fields {
		if !field.IsList {
			continue
		}
		var (
			defaultComponent = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n  },\n", field.Dc, field.TsName, field.Align, field.Width)
			component        string
		)

		// 查询用户摘要
		if in.options.Step.HasHookMemberSummary && IsMemberSummaryField(field.Name) {
			buffer.WriteString(fmt.Sprintf("  {\n    title: '%v',\n    key: '%v',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return renderPopoverMemberSumma(row.%vSumma);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName))
			in.options.Step.ImportModel.UtilsIndex = append(in.options.Step.ImportModel.UtilsIndex, []string{"renderPopoverMemberSumma", "MemberSumma"}...)
			continue
		}

		// 这里根据编辑表单组件来进行推断，如果没有则使用默认input，这可能会导致和查询条件所需参数不符的情况
		switch field.FormMode {
		case FormModeDate:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return formatToDate(row.%s);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName)
			in.options.Step.ImportModel.UtilsDate = append(in.options.Step.ImportModel.UtilsDate, "formatToDate")

		case FormModeRadio:
			fallthrough
		case FormModeSelect:
			if g.IsEmpty(in.options.dictMap[field.TsName]) {
				err = gerror.Newf("设置单选下拉框选项时，必须选择字典类型，字段名称:%v", field.Name)
				return
			}
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return renderOptionTag('%v', row.%v);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, in.options.dictMap[field.TsName], field.TsName)
			in.options.Step.ImportModel.UtilsIndex = append(in.options.Step.ImportModel.UtilsIndex, "renderOptionTag")

		case FormModeSelectMultiple:
			if g.IsEmpty(in.options.dictMap[field.TsName]) {
				err = gerror.Newf("设置多选下拉框选项时，必须选择字典类型，字段名称:%v", field.Name)
				return
			}
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      if (isNullObject(row.%s) || !Array.isArray(row.%s)) {\n        return ``;\n      }\n      return row.%s.map((tagKey) => {\n        return renderOptionTag('%s', tagKey)\n      });\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName, field.TsName, field.TsName, in.options.dictMap[field.TsName])
			in.options.Step.ImportModel.NaiveUI = append(in.options.Step.ImportModel.NaiveUI, "NTag")
			in.options.Step.ImportModel.UtilsIs = append(in.options.Step.ImportModel.UtilsIs, "isNullObject")

		case FormModeUploadImage:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return renderImage(row.%v);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName)
			in.options.Step.ImportModel.UtilsIndex = append(in.options.Step.ImportModel.UtilsIndex, "renderImage")

		case FormModeUploadImages:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return renderImageGroup(row.%v);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName)
			in.options.Step.ImportModel.UtilsIndex = append(in.options.Step.ImportModel.UtilsIndex, "renderImageGroup")

		case FormModeUploadFile:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return renderFile(row.%v);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName)
			in.options.Step.ImportModel.UtilsIndex = append(in.options.Step.ImportModel.UtilsIndex, "renderFile")

		case FormModeUploadFiles:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return renderFileGroup(row.%v);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName)
			in.options.Step.ImportModel.UtilsIndex = append(in.options.Step.ImportModel.UtilsIndex, "renderFileGroup")

		case FormModeSwitch:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return h(%s, {\n        value: row.%s === 1,\n        checked: '开启',\n        unchecked: '关闭',\n        disabled: !hasPermission(['%s']),\n        onUpdateValue: function (e) {\n          console.log('onUpdateValue e:' + JSON.stringify(e));\n          row.%s = e ? 1 : 2;\n          Switch({ %s: row.%s, key: '%s', value: row.%s }).then((_res) => {\n            $message.success('操作成功');\n          });\n        },\n      });\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, "NSwitch", field.TsName, "/"+in.options.ApiPrefix+"/switch", field.TsName, in.pk.TsName, in.pk.TsName, convert.CamelCaseToUnderline(field.TsName), field.TsName)
			in.options.Step.ImportModel.NaiveUI = append(in.options.Step.ImportModel.NaiveUI, "NSwitch")

		case FormModeRate:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return h(%s, {\n        allowHalf: true,\n        readonly: true,\n        defaultValue: row.%s,\n      });\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, "NRate", field.TsName)
			in.options.Step.ImportModel.NaiveUI = append(in.options.Step.ImportModel.NaiveUI, "NRate")

		default:
			component = defaultComponent
		}

		buffer.WriteString(component)
	}
	return
}

// isStringType 判断字段是否为字符串类型
func isStringType(tsType, dataType string) bool {
	// 根据 TypeScript 类型判断
	if tsType == "string" {
		return true
	}

	// 根据数据库数据类型判断
	stringTypes := []string{
		// mysql
		"varchar", "char", "text", "longtext", "mediumtext", "tinytext",
		"nvarchar", "nchar", "ntext",
		"string", "enum", "set",
		// pgsql
		"bpchar", "date", "time", "timetz", "timestamp", "timestamptz", "interval",
		"uuid", "bytea", "inet", "cidr", "macaddr", "macaddr8", "bit", "varbit", "json", "jsonb",
		"point", "line", "lseg", "box", "path", "polygon", "circle", "money",
	}

	for _, t := range stringTypes {
		if strings.EqualFold(dataType, t) {
			return true
		}
	}
	return false
}
