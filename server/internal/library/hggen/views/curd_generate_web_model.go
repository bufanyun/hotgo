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
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/sysin"
	"hotgo/utility/convert"
)

const (
	ModelLoadOptionsTemplate = "async function loadOptions() {\n  options.value = await Dicts({\n    types: [\n  %v   ],\n  });\n  for (const item of schemas.value) {\n    switch (item.field) {\n%v     }\n  }\n}\n\nawait loadOptions();"
)

func (l *gCurd) webModelTplData(ctx context.Context, in *CurdPreviewInput) (data g.Map, err error) {
	data = make(g.Map)
	data["state"] = l.generateWebModelState(ctx, in)
	data["defaultState"] = l.generateWebModelDefaultState(ctx, in)
	data["rules"] = l.generateWebModelRules(ctx, in)
	data["formSchema"] = l.generateWebModelFormSchema(ctx, in)
	if data["columns"], err = l.generateWebModelColumns(ctx, in); err != nil {
		return nil, err
	}
	return
}

func (l *gCurd) generateWebModelState(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export interface State {\n")
	for _, field := range in.masterFields {
		buffer.WriteString(fmt.Sprintf("  %s: %s;\n", field.TsName, field.TsType))
	}
	buffer.WriteString("}")

	return buffer.String()
}

func (l *gCurd) generateWebModelDefaultState(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const defaultState = {\n")
	for _, field := range in.masterFields {
		var value = field.DefaultValue
		if value == nil {
			value = "null"
		}
		if value == "" {
			value = "''"
		}
		buffer.WriteString(fmt.Sprintf("  %s: %v,\n", field.TsName, value))
	}
	buffer.WriteString("};")

	return buffer.String()
}

func (l *gCurd) generateWebModelDictOptions(ctx context.Context, in *CurdPreviewInput) (g.Map, error) {
	type DictType struct {
		Id   int64  `json:"id"`
		Type string `json:"type"`
	}

	var (
		options      = make(g.Map)
		dictTypeIds  []int64
		dictTypeList []*DictType
	)

	for _, field := range in.masterFields {
		if field.DictType > 0 {
			dictTypeIds = append(dictTypeIds, field.DictType)
		}
	}

	dictTypeIds = convert.UniqueSliceInt64(dictTypeIds)
	if len(dictTypeIds) == 0 {
		options["has"] = false
		return options, nil
	}

	err := g.Model("sys_dict_type").Ctx(ctx).
		Fields("id", "type").
		WhereIn("id", dictTypeIds).
		Scan(&dictTypeList)
	if err != nil {
		return nil, err
	}

	if len(dictTypeList) == 0 {
		options["has"] = false
		return options, nil
	}

	options["has"] = true

	var (
		awaitLoadOptions  string
		switchLoadOptions string
	)

	constOptionsBuffer := bytes.NewBuffer(nil)
	constOptionsBuffer.WriteString("export const options = ref<Options>({\n")

	for _, v := range dictTypeList {
		// 字段映射字典
		for _, field := range in.masterFields {
			if field.DictType > 0 && v.Id == field.DictType {
				in.options.dictMap[field.TsName] = v.Type
				switchLoadOptions = fmt.Sprintf("%s      case '%s':\n        item.componentProps.options = options.value.%s;\n        break;\n", switchLoadOptions, field.TsName, v.Type)
			}
		}

		awaitLoadOptions = fmt.Sprintf("%s    '%s',\n", awaitLoadOptions, v.Type)
		constOptionsBuffer.WriteString("  " + v.Type + ": [],\n")
	}

	constOptionsBuffer.WriteString("});\n")

	loadOptionsBuffer := bytes.NewBuffer(nil)
	loadOptionsBuffer.WriteString(fmt.Sprintf(ModelLoadOptionsTemplate, awaitLoadOptions, switchLoadOptions))

	options["const"] = constOptionsBuffer.String()
	options["load"] = loadOptionsBuffer.String()

	return options, nil
}

func (l *gCurd) generateWebModelRules(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const rules = {\n")
	for _, field := range in.masterFields {
		if !field.IsEdit || (!field.Required && (field.FormRole == "" || field.FormRole == FormRoleNone)) {
			continue
		}

		if field.FormRole == "" || field.FormRole == FormRoleNone {
			buffer.WriteString(fmt.Sprintf("  %s: {\n    required: %v,\n    trigger: ['blur', 'input'],\n    type: '%s',\n    message: '请输入%s',\n  },\n", field.TsName, field.Required, field.TsType, field.Dc))
		} else {
			buffer.WriteString(fmt.Sprintf("  %s: {\n    required: %v,\n    trigger: ['blur', 'input'],\n    type: '%s',\n    message: '请输入%s',\n    validator: validate.%v,\n  },\n", field.TsName, field.Required, field.TsType, field.Dc, field.FormRole))
		}
	}
	buffer.WriteString("};\n")
	return buffer.String()
}

func (l *gCurd) generateWebModelFormSchema(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const schemas = ref<FormSchema[]>([\n")

	// 主表
	l.generateWebModelFormSchemaEach(buffer, in.masterFields)

	// 关联表
	if len(in.options.Join) > 0 {
		for _, v := range in.options.Join {
			if !isEffectiveJoin(v) {
				continue
			}
			l.generateWebModelFormSchemaEach(buffer, v.Columns)
		}
	}

	buffer.WriteString("]);\n")
	return buffer.String()
}

func (l *gCurd) generateWebModelFormSchemaEach(buffer *bytes.Buffer, fields []*sysin.GenCodesColumnListModel) {
	for _, field := range fields {
		if !field.IsQuery {
			continue
		}

		var (
			defaultComponent = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      placeholder: '请输入%s',\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NInput", field.Dc, field.Dc)
			component        string
		)

		// 这里根据编辑表单组件来进行推断，如果没有则使用默认input，这可能会导致和查询条件所需参数不符的情况
		switch field.FormMode {
		case FormModeInput, FormModeInputTextarea, FormModeInputEditor:
			component = defaultComponent

		case FormModeInputNumber:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      placeholder: '请输入%s',\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NInputNumber", field.Dc, field.Dc)

		case FormModeDate:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NDatePicker", field.Dc, "date", "defShortcuts()")

		case FormModeDateRange:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NDatePicker", field.Dc, "daterange", "defRangeShortcuts()")

		case FormModeTime:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NDatePicker", field.Dc, "datetime", "defShortcuts()")

		case FormModeTimeRange:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NDatePicker", field.Dc, "datetimerange", "defRangeShortcuts()")

		case FormModeRadio:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    giProps: {\n      //span: 24,\n    },\n    componentProps: {\n      options: [],\n      onUpdateChecked: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NRadioGroup", field.Dc)

		case FormModeCheckbox:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    giProps: {\n      span: 1,\n    },\n    componentProps: {\n      placeholder: '请选择%s',\n      options: [],\n      onUpdateChecked: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NCheckbox", field.Dc, field.Dc)

		case FormModeSelect:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    defaultValue: null,\n    componentProps: {\n      placeholder: '请选择%s',\n      options: [],\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NSelect", field.Dc, field.Dc)

		case FormModeSelectMultiple:
			component = fmt.Sprintf("  {\n    field: '%s',\n    component: '%s',\n    label: '%s',\n    defaultValue: null,\n    componentProps: {\n      multiple: true,\n      placeholder: '请选择%s',\n      options: [],\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "NSelect", field.Dc, field.Dc)

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
			defaultComponent = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n  },\n", field.Dc, field.TsName)
			component        string
		)

		// 这里根据编辑表单组件来进行推断，如果没有则使用默认input，这可能会导致和查询条件所需参数不符的情况
		switch field.FormMode {
		case FormModeDate:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    render(row) {\n      return formatToDate(row.%s);\n    },\n  },\n", field.Dc, field.TsName, field.TsName)

		case FormModeSelect:
			if g.IsEmpty(in.options.dictMap[field.TsName]) {
				err = gerror.Newf("设置单选下拉框选项时，必须选择字典类型，字段名称:%v", field.Name)
				return
			}
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    render(row) {\n      if (isNullObject(row.%s)) {\n        return ``;\n      }\n      return h(\n        NTag,\n        {\n          style: {\n            marginRight: '6px',\n          },\n          type: getOptionTag(options.value.%s, row.%s),\n          bordered: false,\n        },\n        {\n          default: () => getOptionLabel(options.value.%s, row.%s),\n        }\n      );\n    },\n  },\n", field.Dc, field.TsName, field.TsName, in.options.dictMap[field.TsName], field.TsName, in.options.dictMap[field.TsName], field.TsName)

		case FormModeSelectMultiple:
			if g.IsEmpty(in.options.dictMap[field.TsName]) {
				err = gerror.Newf("设置多选下拉框选项时，必须选择字典类型，字段名称:%v", field.Name)
				return
			}
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    render(row) {\n      if (isNullObject(row.%s) || !isArray(row.%s)) {\n        return ``;\n      }\n      return row.%s.map((tagKey) => {\n        return h(\n          NTag,\n          {\n            style: {\n              marginRight: '6px',\n            },\n            type: getOptionTag(options.value.%s, tagKey),\n            bordered: false,\n          },\n          {\n            default: () => getOptionLabel(options.value.%s, tagKey),\n          }\n        );\n      });\n    },\n  },\n", field.Dc, field.TsName, field.TsName, field.TsName, field.TsName, in.options.dictMap[field.TsName], in.options.dictMap[field.TsName])

		case FormModeUploadImage:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    render(row) {\n      return h(%s, {\n        width: 32,\n        height: 32,\n        src: row.%s,\n        onError: errorImg,\n        style: {\n          width: '32px',\n          height: '32px',\n          'max-width': '100%%',\n          'max-height': '100%%',\n        },\n      });\n    },\n  },\n", field.Dc, field.TsName, "NImage", field.TsName)

		case FormModeUploadImages:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    render(row) {\n      if (isNullObject(row.%s)) {\n        return ``;\n      }\n      return row.%s.map((image) => {\n        return h(%s, {\n          width: 32,\n          height: 32,\n          src: image,\n        onError: errorImg,\n          style: {\n            width: '32px',\n            height: '32px',\n            'max-width': '100%%',\n            'max-height': '100%%',\n            'margin-left': '2px',\n          },\n        });\n      });\n    },\n  },\n", field.Dc, field.TsName, field.TsName, field.TsName, "NImage")

		case FormModeUploadFile:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    render(row) {\n      if (row.%s === '') {\n        return ``;\n      }\n      return h(\n        %s,\n        {\n          size: 'small',\n        },\n        {\n          default: () => getFileExt(row.%s),\n        }\n      );\n    },\n  },\n", field.Dc, field.TsName, field.TsName, "NAvatar", field.TsName)

		case FormModeUploadFiles:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    render(row) {\n      if (isNullObject(row.%s)) {\n        return ``;\n      }\n      return row.%s.map((attachfile) => {\n        return h(\n          %s,\n          {\n            size: 'small',\n            style: {\n              'margin-left': '2px',\n            },\n          },\n          {\n            default: () => getFileExt(attachfile),\n          }\n        );\n      });\n    },\n  },\n", field.Dc, field.TsName, field.TsName, field.TsName, "NAvatar")

		case FormModeSwitch:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    width: 100,\n    render(row) {\n      return h(%s, {\n        value: row.%s === 1,\n        checked: '开启',\n        unchecked: '关闭',\n        disabled: !hasPermission(['%s']),\n        onUpdateValue: function (e) {\n          console.log('onUpdateValue e:' + JSON.stringify(e));\n          row.%s = e ? 1 : 2;\n          Switch({ %s: row.%s, key: '%s', value: row.%s }).then((_res) => {\n            $message.success('操作成功');\n          });\n        },\n      });\n    },\n  },\n", field.Dc, field.TsName, "NSwitch", field.TsName, "/"+in.options.ApiPrefix+"/switch", field.TsName, in.pk.TsName, in.pk.TsName, field.TsName, field.TsName)

		case FormModeRate:
			component = fmt.Sprintf("  {\n    title: '%s',\n    key: '%s',\n    render(row) {\n      return h(%s, {\n        allowHalf: true,\n        readonly: true,\n        defaultValue: row.%s,\n      });\n    },\n  },\n", field.Dc, field.TsName, "NRate", field.TsName)

		default:
			component = defaultComponent
		}

		buffer.WriteString(component)
	}

	return
}
