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
	"github.com/olekukonko/tablewriter"
	"hotgo/internal/model/input/sysin"
)

const (
	InputTypeListInp     = 1 // 列表输入
	InputTypeListModel   = 2 // 列表输出
	InputTypeExportModel = 3 // 列表导出
)

func (l *gCurd) inputTplData(ctx context.Context, in *CurdPreviewInput) (data g.Map, err error) {
	data = make(g.Map)
	data["listInpColumns"] = l.generateInputListColumns(ctx, in, InputTypeListInp)
	data["listModelColumns"] = l.generateInputListColumns(ctx, in, InputTypeListModel)
	data["exportModelColumns"] = l.generateInputListColumns(ctx, in, InputTypeExportModel)
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

	default:
		panic("inputType is invalid")
	}

	return result
}
