// Package excel
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package excel

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model"
	"net/url"
	"reflect"
	"time"
	"unicode"
)

var (
	// 单元格表头
	char = []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	// 默认行样式
	defaultRowStyle = &excelize.Style{
		Font: &excelize.Font{
			Color:  "#666666",
			Size:   13,
			Family: "Arial",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}
)

// ExportByStructs 导出切片结构体到excel表格
func ExportByStructs(ctx context.Context, tags []string, list interface{}, fileName string, sheetName string) (err error) {
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName)
	_ = f.SetRowHeight("Sheet1", 1, 30)

	rowStyleID, _ := f.NewStyle(defaultRowStyle)
	if err != nil {
		return
	}
	_ = f.SetSheetRow(sheetName, "A1", &tags)

	var (
		length    = len(tags)
		headStyle = letter(length)
		lastRow   string
		widthRow  string
	)

	for k, v := range headStyle {
		if k == length-1 {
			lastRow = fmt.Sprintf("%s1", v)
			widthRow = v
		}
	}

	if err = f.SetColWidth(sheetName, "A", widthRow, 30); err != nil {
		return err
	}

	var rowNum = 1
	for _, v := range gconv.Interfaces(list) {
		t := reflect.TypeOf(v)
		value := reflect.ValueOf(v)
		row := make([]interface{}, 0)
		for l := 0; l < t.NumField(); l++ {
			val := value.Field(l).Interface()
			row = append(row, val)
		}
		rowNum++
		if err = f.SetSheetRow(sheetName, "A"+gconv.String(rowNum), &row); err != nil {
			return
		}
		if err = f.SetCellStyle(sheetName, fmt.Sprintf("A%d", rowNum), lastRow, rowStyleID); err != nil {
			return
		}
	}

	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		err = gerror.New("ctx not http request")
		return
	}

	writer := r.Response.Writer
	writer.Header().Set("Content-Type", "application/octet-stream")
	writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", url.QueryEscape(fileName)))
	writer.Header().Set("Content-Transfer-Encoding", "binary")
	writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	if err = f.Write(writer); err != nil {
		return
	}

	// 加入到上下文
	contexts.SetResponse(ctx, &model.Response{
		Code:      gcode.CodeOK.Code(),
		Message:   "export successfully!",
		Timestamp: time.Now().Unix(),
		TraceID:   gctx.CtxId(ctx),
	})
	return
}

// ExportByStructsStream 按批拉取行并写入，避免一次性把全部数据放进内存。
// fetchRows 返回空切片时结束。
func ExportByStructsStream(ctx context.Context, tags []string, fileName string, sheetName string, batchSize int, fetchRows func(offset, limit int) ([][]interface{}, error)) (err error) {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		return gerror.New("ctx not http request")
	}
	if batchSize <= 0 {
		batchSize = 500
	}

	f := excelize.NewFile()
	defer f.Close()
	f.SetSheetName("Sheet1", sheetName)

	sw, err := f.NewStreamWriter(sheetName)
	if err != nil {
		return gerror.Wrap(err, "创建StreamWriter失败")
	}

	headerStyleID, err := f.NewStyle(streamHeaderStyle)
	if err != nil {
		return gerror.Wrap(err, "创建表头样式失败")
	}
	rowStyleID, err := f.NewStyle(streamRowStyle)
	if err != nil {
		return gerror.Wrap(err, "创建数据行样式失败")
	}
	altStyleID, err := f.NewStyle(streamAltStyle)
	if err != nil {
		return gerror.Wrap(err, "创建交替行样式失败")
	}

	length := len(tags)
	if length > 0 {
		if err = sw.SetColWidth(1, length, 30); err != nil {
			return gerror.Wrap(err, "设置列宽失败")
		}
	}

	headerRow := make([]interface{}, len(tags))
	for i, tag := range tags {
		headerRow[i] = excelize.Cell{Value: tag, StyleID: headerStyleID}
	}
	if err = sw.SetRow("A1", headerRow, excelize.RowOpts{Height: 30}); err != nil {
		return gerror.Wrap(err, "写入表头失败")
	}

	rowIdx := 2
	offset := 0
	for {
		batch, fetchErr := fetchRows(offset, batchSize)
		if fetchErr != nil {
			return gerror.Wrap(fetchErr, "获取导出数据失败")
		}
		if len(batch) == 0 {
			break
		}
		for _, cols := range batch {
			styleID := rowStyleID
			if (rowIdx-2)%2 != 0 {
				styleID = altStyleID
			}
			cells := make([]interface{}, len(cols))
			for i, v := range cols {
				cells[i] = excelize.Cell{Value: v, StyleID: styleID}
			}
			cell, _ := excelize.CoordinatesToCellName(1, rowIdx)
			if err = sw.SetRow(cell, cells, excelize.RowOpts{Height: 24}); err != nil {
				return gerror.Wrap(err, "写入数据行失败")
			}
			rowIdx++
		}
		offset += batchSize
		if len(batch) < batchSize {
			break
		}
	}
	if err = sw.Flush(); err != nil {
		return gerror.Wrap(err, "Flush失败")
	}

	writer := r.Response.Writer
	writer.Header().Set("Content-Type", "application/octet-stream")
	writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", url.QueryEscape(fileName)))
	writer.Header().Set("Content-Transfer-Encoding", "binary")
	writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	if err = f.Write(writer); err != nil {
		return gerror.Wrap(err, "写入Excel文件失败")
	}
	contexts.SetResponse(ctx, &model.Response{
		Code:      gcode.CodeOK.Code(),
		Message:   "export successfully!",
		Timestamp: time.Now().Unix(),
		TraceID:   gctx.CtxId(ctx),
	})
	return
}

var (
	streamHeaderStyle = &excelize.Style{
		Font:      &excelize.Font{Bold: true, Color: "#FFFFFF", Size: 12, Family: "Calibri"},
		Alignment: &excelize.Alignment{Horizontal: "left", Vertical: "center"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#3367D6"}, Pattern: 1},
	}
	streamRowStyle = &excelize.Style{
		Font:      &excelize.Font{Color: "#333333", Size: 11, Family: "Calibri"},
		Alignment: &excelize.Alignment{Horizontal: "left", Vertical: "center"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#FFFFFF"}, Pattern: 1},
	}
	streamAltStyle = &excelize.Style{
		Font:      &excelize.Font{Color: "#333333", Size: 11, Family: "Calibri"},
		Alignment: &excelize.Alignment{Horizontal: "left", Vertical: "center"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#F5F7FA"}, Pattern: 1},
	}
)

// letter 生成完整的表头
func letter(length int) []string {
	var str []string
	for i := 1; i <= length; i++ {
		str = append(str, numToChars(i))
	}
	return str
}

// numToChars 将数字转换为具体的表格表头名称
func numToChars(num int) string {
	var cols string
	v := num
	for v > 0 {
		k := v % 26
		if k == 0 {
			k = 26
		}
		v = (v - k) / 26
		cols = char[k] + cols
	}
	return cols
}

// NextLetter 传入一个字母，获取下一个字母
func NextLetter(input string) string {
	if len(input) == 0 {
		return ""
	}
	upperInput := unicode.ToUpper(rune(input[0]))
	if upperInput >= 'A' && upperInput < 'Z' {
		return string(upperInput + 1)
	}
	return "A"
}
