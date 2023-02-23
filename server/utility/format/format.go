// Package format
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package format

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
)

// RoundInt64 四舍五入
func RoundInt64(x float64) int64 {
	return int64(math.Floor(x + 0/5))
}

// Round2String 四舍五入保留小数，默认2位
func Round2String(value float64, args ...interface{}) (v string) {
	var places int32 = 2
	if len(args) > 0 {
		places = args[0].(int32)
	}

	return decimal.NewFromFloat(value).Round(places).String()
}

// Round2Float64 四舍五入保留小数，默认2位
func Round2Float64(value float64, args ...interface{}) (v float64) {
	var places int32 = 2
	if len(args) > 0 {
		places = args[0].(int32)
	}

	return decimal.NewFromFloat(value).Round(places).InexactFloat64()
}

// FileSize 字节的单位转换 保留两位小数
func FileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else {
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}
