// Package validate
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package validate

import (
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// 包含判断

// InSliceExistStr 判断字符或切片字符是否存在指定字符
func InSliceExistStr(elems interface{}, search string) bool {
	switch elems.(type) {
	case []string:
		elem := gconv.Strings(elems)
		for i := 0; i < len(elem); i++ {
			if gconv.String(elem[i]) == search {
				return true
			}
		}
	default:
		return gconv.String(elems) == search
	}

	return false
}

// InSliceInt64 元素是否存在于切片中
func InSliceInt64(slice []int64, key int64) bool {
	if len(slice) == 0 {
		return false
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == key {
			return true
		}
	}
	return false
}

func InSliceInt(slice []int, key int) bool {
	if len(slice) == 0 {
		return false
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == key {
			return true
		}
	}
	return false
}

func InSliceString(slice []string, key string) bool {
	return gstr.InArray(slice, key)
}
