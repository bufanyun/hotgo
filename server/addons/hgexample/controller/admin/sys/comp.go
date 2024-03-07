// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"github.com/xuri/excelize/v2"
	"hotgo/addons/hgexample/api/admin/comp"
)

var (
	Comp = cComp{}
)

type cComp struct{}

// ImportExcel 导入Excel
func (c *cComp) ImportExcel(ctx context.Context, req *comp.ImportExcelReq) (res *comp.ImportExcelRes, err error) {
	file, err := req.File.Open()
	defer file.Close()
	if err != nil {
		return
	}

	excel, err := excelize.OpenReader(file)
	if err != nil {
		return
	}
	defer excel.Close()

	res = new(comp.ImportExcelRes)
	sheetList := excel.GetSheetList()
	for _, sheet := range sheetList {
		item := new(comp.ImportExcelSheet)
		item.Sheet = sheet
		item.Rows, err = excel.GetRows(sheet)
		if err != nil {
			return nil, err
		}
		res.Sheets = append(res.Sheets, item)
	}
	return
}
