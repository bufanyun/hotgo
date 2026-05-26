// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package views

import (
	"bytes"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func (l *gCurd) webIndexTplData(ctx context.Context, in *CurdPreviewInput) (g.Map, error) {
	var (
		data              = make(g.Map)
		importBuffer      = bytes.NewBuffer(nil)
		importVueMethod   = []string{"h", "reactive", "ref", "computed"}
		importApiMethod   = []string{"List"}
		importModelMethod = []string{"columns", "schemas", "State"}
		importUtilsMethod = []string{"adaTableScrollX"}
		importIcons       []string
		actionWidth       int64 = 72
	)

	// 添加
	if in.options.Step.HasAdd {
		importIcons = append(importIcons, "PlusOutlined")
	}

	// 编辑
	if in.options.Step.HasEdit {
		in.options.Step.ActionColumnWidth += actionWidth
		if in.options.Step.IsTreeTable && !in.options.Step.IsOptionTreeTable {
			in.options.Step.ActionColumnWidth += actionWidth
		}
		if in.options.Step.IsOptionTreeTable {
			importIcons = append(importIcons, "EditOutlined")
		}
	}

	// 导出
	if in.options.Step.HasExport {
		importIcons = append(importIcons, "ExportOutlined")
		importApiMethod = append(importApiMethod, "Export")
	}

	// 删除
	if in.options.Step.HasDel {
		importApiMethod = append(importApiMethod, "Delete")
		in.options.Step.ActionColumnWidth += actionWidth
	}

	// 批量删除
	if in.options.Step.HasBatchDel {
		importIcons = append(importIcons, "DeleteOutlined")
		importApiMethod = append(importApiMethod, "Delete")
	}

	// 修改状态
	if in.options.Step.HasStatus {
		importApiMethod = append(importApiMethod, "Status")
		in.options.Step.ActionColumnWidth += actionWidth
	}

	// 更多
	// 查看详情
	if in.options.Step.HasView {
		in.options.Step.ActionColumnWidth += actionWidth
	}

	// 展开树
	if in.options.Step.IsTreeTable {
		importIcons = append(importIcons, "AlignLeftOutlined")
	}

	// 存在字典数据选项
	if in.options.DictOps.Has {
		importVueMethod = append(importVueMethod, "onMounted")
		importModelMethod = append(importModelMethod, "loadOptions")
	}

	// 普通树表
	if in.options.Step.IsTreeTable && !in.options.Step.IsOptionTreeTable {
		importUtilsMethod = append(importUtilsMethod, "convertListToTree")
	}

	// 选项式树表
	if in.options.Step.IsOptionTreeTable {
		importVueMethod = append(importVueMethod, []string{"onMounted", "unref"}...)
		importIcons = append(importIcons, []string{"FormOutlined", "SearchOutlined"}...)
		importApiMethod = append(importApiMethod, "TreeOption")
		importUtilsMethod = append(importUtilsMethod, "getTreeKeys")
		importModelMethod = append(importModelMethod, []string{"loadTreeOption", "treeOption", "State"}...)
	}

	// 操作按钮宽度最小值
	if in.options.Step.ActionColumnWidth > 0 && in.options.Step.ActionColumnWidth < actionWidth*2 {
		in.options.Step.ActionColumnWidth = 100
	}

	// 导入基础包
	importBuffer.WriteString("  import { useDialog, useMessage } from 'naive-ui';\n")
	importBuffer.WriteString("  import { BasicTable, TableAction } from '@/components/Table';\n")
	importBuffer.WriteString("  import { BasicForm, useForm } from '@/components/Form/index';\n")
	importBuffer.WriteString("  import { usePermission } from '@/hooks/web/usePermission';\n")

	// 导入字典
	if in.options.DictOps.Has {
		importBuffer.WriteString("  import { useDictStore } from '@/store/modules/dict';\n")
	}

	// 导入api
	importBuffer.WriteString("  import " + ImportWebMethod(importApiMethod) + " from '" + in.options.ImportWebApi + "';\n")

	// 导入icons
	if len(importIcons) > 0 {
		importBuffer.WriteString("  import " + ImportWebMethod(importIcons) + " from '@vicons/antd';\n")
	}

	// 导入model
	if in.options.Step.IsTreeTable {
		importModelMethod = append(importModelMethod, "newState")
	}
	importBuffer.WriteString("  import " + ImportWebMethod(importModelMethod) + " from './model';\n")

	// 导入utils
	if len(importUtilsMethod) > 0 {
		importBuffer.WriteString("  import " + ImportWebMethod(importUtilsMethod) + " from '@/utils/hotgo';\n")
	}

	// 导入edit组件
	if in.options.Step.HasEdit {
		importBuffer.WriteString("  import Edit from './edit.vue';\n")
	}

	// 导入view组件
	if in.options.Step.HasView {
		importBuffer.WriteString("  import View from './view.vue';\n")
	}

	// 没有需要查询的字段则隐藏搜索表单
	isSearchForm := false
	for _, field := range in.masterFields {
		if field.IsQuery {
			isSearchForm = true
			break
		}
	}
	if !isSearchForm {
		if len(in.options.Join) > 0 {
		LoopOut:
			for _, v := range in.options.Join {
				for _, column := range v.Columns {
					if column.IsQuery {
						isSearchForm = true
						break LoopOut
					}
				}
			}
		}
	}
	data["isSearchForm"] = isSearchForm
	data["import"] = importBuffer.String()
	return data, nil
}
