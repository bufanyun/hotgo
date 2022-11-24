// Package excel
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package excel

import (
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	"net/url"
	"reflect"
)

func ExportByStruct(w *ghttp.ResponseWriter, titleList []string, data []interface{}, fileName string, sheetName string) error {

	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName)
	header := make([]string, 0)
	for _, v := range titleList {
		header = append(header, v)
	}

	rowStyleID, _ := f.NewStyle(`{"font":{"color":"#666666","size":13,"family":"arial"},"alignment":{"vertical":"center","horizontal":"center"}}`)
	_ = f.SetSheetRow(sheetName, "A1", &header)
	_ = f.SetRowHeight("Sheet1", 1, 30)
	length := len(titleList)
	headStyle := letter(length)
	var lastRow string
	var widthRow string
	for k, v := range headStyle {
		if k == length-1 {
			lastRow = fmt.Sprintf("%s1", v)
			widthRow = v
		}
	}
	if err := f.SetColWidth(sheetName, "A", widthRow, 30); err != nil {
		return err
	}

	rowNum := 1
	for _, v := range data {
		t := reflect.TypeOf(v)
		value := reflect.ValueOf(v)
		row := make([]interface{}, 0)
		for l := 0; l < t.NumField(); l++ {
			val := value.Field(l).Interface()
			row = append(row, val)
		}
		rowNum++
		err := f.SetSheetRow(sheetName, "A"+gconv.String(rowNum), &row)
		_ = f.SetCellStyle(sheetName, fmt.Sprintf("A%d", rowNum), fmt.Sprintf("%s", lastRow), rowStyleID)
		if err != nil {
			return err
		}
	}
	disposition := fmt.Sprintf("attachment; filename=%s.xlsx", url.QueryEscape(fileName))

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", disposition)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	err := f.Write(w)
	if err != nil {
		return err
	}

	return nil
}

// Letter 遍历a-z
func letter(length int) []string {
	var str []string
	for i := 0; i < length; i++ {
		str = append(str, string(rune('A'+i)))
	}
	return str
}
