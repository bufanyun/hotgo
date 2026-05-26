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
)

func (l *gCurd) webEditTplData(ctx context.Context, in *CurdPreviewInput) (data g.Map, err error) {
	data = make(g.Map)
	data["formItem"] = l.generateWebEditFormItem(ctx, in)
	data["script"] = l.generateWebEditScript(ctx, in)
	return
}

func (l *gCurd) generateWebEditFormItem(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	for _, field := range in.masterFields {
		if !field.IsEdit {
			continue
		}

		if IsIndexPK(field.Index) {
			continue
		}

		var (
			defaultComponent = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <n-input\n                    placeholder=\"请输入%s\"\n                    v-model:value=\"formValue.%s\"\n                  />\n                </n-form-item>", field.Dc, field.TsName, field.Dc, field.TsName)
			component        string
		)

		if in.options.Step.IsTreeTable && IsPidName(field.Name) {
			field.FormMode = FormModePidTreeSelect
		}

		switch field.FormMode {
		case FormModeInput:
			component = defaultComponent

		case FormModeInputNumber:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <n-input-number\n                    placeholder=\"请输入%s\"\n                    v-model:value=\"formValue.%s\"\n                  />\n                </n-form-item>", field.Dc, field.TsName, field.Dc, field.TsName)

		case FormModeInputTextarea:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <n-input\n                    type=\"textarea\"\n                    placeholder=\"%s\"\n                    v-model:value=\"formValue.%s\"\n                  />\n                </n-form-item>", field.Dc, field.TsName, field.Dc, field.TsName)

		case FormModeInputEditor:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <Editor style=\"height: 450px\" id=\"%s\" v-model:value=\"formValue.%s\" />\n                </n-form-item>", field.Dc, field.TsName, field.TsName, field.TsName)

		case FormModeInputYaml:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <YamlEditor\n                    ref=\"%sYamlRef\"\n                    v-model:value=\"formValue.%s\"\n                    :height=\"400\"\n                    placeholder=\"请输入%s\"\n                    validate-on-blur\n                    show-stats\n                    prevent-invalid-submit\n                    @error=\"handleYamlError\"\n                    @valid=\"handleYamlValid\"\n                    @validation-change=\"(isValid) => handleYamlValidationChange(isValid, '%s')\"\n                  />\n                </n-form-item>", field.Dc, field.TsName, field.TsName, field.TsName, field.Dc, field.TsName)

		case FormModeInputDynamic:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <n-dynamic-input\n                    v-model:value=\"formValue.%s\"\n                    preset=\"pair\"\n                    key-placeholder=\"键名\"\n                    value-placeholder=\"键值\"\n                  />\n                </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeDate:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <DatePicker v-model:formValue=\"formValue.%s\" type=\"date\" />\n                </n-form-item>", field.Dc, field.TsName, field.TsName)

		// case FormModeDateRange:  // 必须要有两个字段，后面优化下

		case FormModeTime:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <DatePicker v-model:formValue=\"formValue.%s\" type=\"datetime\" />\n                </n-form-item>", field.Dc, field.TsName, field.TsName)

		// case FormModeTimeRange: // 必须要有两个字段，后面优化下

		case FormModeRadio:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <n-radio-group v-model:value=\"formValue.%s\" name=\"%s\">\n                    <n-radio-button\n                      v-for=\"%s in dict.getOptionUnRef('%s')\"\n                      :key=\"%s.value\"\n                      :value=\"%s.value\"\n                      :label=\"%s.label\"\n                    />\n                  </n-radio-group>\n                </n-form-item>", field.Dc, field.TsName, field.TsName, field.TsName, field.TsName, in.options.dictMap[field.TsName], field.TsName, field.TsName, field.TsName)

		case FormModeCheckbox:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <n-checkbox-group v-model:value=\"formValue.%s\">\n                    <n-space>\n                      <n-checkbox\n                        v-for=\"item in dict.getOptionUnRef('%s')\"\n                        :key=\"item.value\"\n                        :value=\"item.value\"\n                        :label=\"item.label\"\n                      />\n                    </n-space>\n                  </n-checkbox-group>\n                </n-form-item>", field.Dc, field.TsName, field.TsName, in.options.dictMap[field.TsName])

		case FormModeSelect:
			if in.options.dictMap[field.TsName] != nil {
				component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <n-select v-model:value=\"formValue.%s\" :options=\"dict.getOptionUnRef('%s')\" />\n                </n-form-item>", field.Dc, field.TsName, field.TsName, in.options.dictMap[field.TsName])
			} else {
				component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <n-select v-model:value=\"formValue.%s\" options=\"\" />\n                </n-form-item>", field.Dc, field.TsName, field.TsName)
			}

		case FormModeSelectMultiple:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <n-select multiple v-model:value=\"formValue.%s\" :options=\"dict.getOptionUnRef('%s')\" />\n                </n-form-item>", field.Dc, field.TsName, field.TsName, in.options.dictMap[field.TsName])

		case FormModeUploadImage:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <UploadImage :maxNumber=\"1\" v-model:value=\"formValue.%s\" />\n                </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeUploadImages:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <UploadImage :maxNumber=\"10\" v-model:value=\"formValue.%s\" />\n                </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeUploadFile:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <UploadFile :maxNumber=\"1\" v-model:value=\"formValue.%s\" />\n                </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeUploadFiles:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <UploadFile :maxNumber=\"10\" v-model:value=\"formValue.%s\" />\n                </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeSwitch:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <n-switch\n                    :unchecked-value=\"2\"\n                    :checked-value=\"1\"\n                    v-model:value=\"formValue.%s\"\n                  />\n                </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModeRate:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <n-rate allow-half :default-value=\"formValue.%s\" :on-update:value=\"update%s\" />\n                </n-form-item>", field.Dc, field.TsName, field.TsName, field.GoName)

		case FormModeCitySelector:
			component = fmt.Sprintf("                <n-form-item label=\"%s\" path=\"%s\">\n                  <CitySelector v-model:value=\"formValue.%s\" />\n                </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FormModePidTreeSelect:
			component = fmt.Sprintf(`                <n-form-item label="%v" path="pid">
                  <n-tree-select
                    :options="treeOption"
                    v-model:value="formValue.pid"
                    key-field="%v"
                    label-field="%v"
                    clearable
                    filterable
                    default-expand-all
                    show-path
                  />
                </n-form-item>`, field.Dc, in.pk.TsName, in.options.Tree.TitleField.TsName)

		case FormModeTreeSelect:
			component = fmt.Sprintf(`                <n-form-item label="%v" path="%v">
                  <n-tree-select
                    placeholder="请选择%v"
                    v-model:value="formValue.%v"
                    :options="[{ label: 'AA', key: 1, children: [{ label: 'BB', key: 2 }] }]"
                    clearable
                    filterable
                    default-expand-all
                  />
                </n-form-item>`, field.Dc, field.TsName, field.Dc, field.TsName)

		case FormModeCascader:
			component = fmt.Sprintf(`                <n-form-item label="%v" path="%v">
                  <n-cascader
                    placeholder="请选择%v"
                    v-model:value="formValue.%v"
                    :options="[{ label: 'AA', value: 1, children: [{ label: 'BB', value: 2 }] }]"
                    clearable
                    filterable
                  />
                </n-form-item>`, field.Dc, field.TsName, field.Dc, field.TsName)

		default:
			component = defaultComponent
		}

		buffer.WriteString(fmt.Sprintf("<n-gi span=\"%v\">\n%v\n              </n-gi>\n\n", field.FormGridSpan, component))
	}
	return buffer.String()
}

func (l *gCurd) generateWebEditScript(ctx context.Context, in *CurdPreviewInput) g.Map {
	var (
		data         = make(g.Map)
		importBuffer = bytes.NewBuffer(nil)
		setupBuffer  = bytes.NewBuffer(nil)
		hasYamlField = false
	)

	// 导入字典
	if in.options.DictOps.Has {
		importBuffer.WriteString("  import { useDictStore } from '@/store/modules/dict';\n")
	}

	// 导入api
	var importApiMethod = []string{"Edit", "View"}
	if in.options.Step.HasMaxSort {
		importApiMethod = append(importApiMethod, "MaxSort")
	}
	importBuffer.WriteString("  import " + ImportWebMethod(importApiMethod) + " from '" + in.options.ImportWebApi + "';\n")

	// 导入model
	var importModelMethod = []string{"State", "newState"}
	if in.options.Step.IsTreeTable {
		importModelMethod = append(importModelMethod, []string{"treeOption", "loadTreeOption"}...)
	}

	if in.options.Step.HasRules {
		importModelMethod = append(importModelMethod, "rules")
	}
	importBuffer.WriteString("  import " + ImportWebMethod(importModelMethod) + " from './model';\n")

	for _, field := range in.masterFields {
		if !field.IsEdit {
			continue
		}
		switch field.FormMode {
		case FormModeDate, FormModeDateRange, FormModeTime, FormModeTimeRange:
			if !gstr.Contains(importBuffer.String(), `import DatePicker`) {
				importBuffer.WriteString("  import DatePicker from '@/components/DatePicker/datePicker.vue';\n")
			}
		case FormModeInputEditor:
			if !gstr.Contains(importBuffer.String(), `import Editor`) {
				importBuffer.WriteString("  import Editor from '@/components/Editor/editor.vue';\n")
			}
		case FormModeInputYaml:
			if !gstr.Contains(importBuffer.String(), `import YamlEditor`) {
				importBuffer.WriteString("  import YamlEditor from '@/components/YamlEditor/index.vue';\n")
			}
			hasYamlField = true
		case FormModeUploadImage, FormModeUploadImages:
			if !gstr.Contains(importBuffer.String(), `import UploadImage`) {
				importBuffer.WriteString("  import UploadImage from '@/components/Upload/uploadImage.vue';\n")
			}
		case FormModeUploadFile, FormModeUploadFiles:
			if !gstr.Contains(importBuffer.String(), `import UploadFile`) {
				importBuffer.WriteString("  import UploadFile from '@/components/Upload/uploadFile.vue';\n")
			}
		case FormModeRate:
			setupBuffer.WriteString(fmt.Sprintf("  function update%s(num) {\n    formValue.value.%s = num;\n  }\n", field.GoName, field.TsName))
		case FormModeCitySelector:
			if !gstr.Contains(importBuffer.String(), `import CitySelector`) {
				importBuffer.WriteString("  import CitySelector from '@/components/CitySelector/citySelector.vue';\n")
			}
		}
	}

	// 根据是否有 YAML 字段添加相应的验证逻辑
	if hasYamlField {
		setupBuffer.WriteString("  const yamlValidationStates = ref(new Map());\n  const isFormValid = computed(() => {\n    for (const [fieldName, isValid] of yamlValidationStates.value) {\n      if (!isValid) return false;\n    }\n    return true;\n  });\n  const handleYamlError = (error) => {\n    console.error('YAML 验证错误:', error);\n  };\n  const handleYamlValid = (content) => {\n    console.log('YAML 验证通过:', content);\n  };\n  const handleYamlValidationChange = (isValid, fieldName) => {\n    yamlValidationStates.value.set(fieldName, isValid);\n  };\n")
	} else {
		setupBuffer.WriteString("  const isFormValid = ref(true);\n")
	}

	data["import"] = importBuffer.String()
	data["setup"] = setupBuffer.String()
	return data
}
