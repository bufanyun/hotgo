// Package convert
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package convert

// UniqueSliceInt64 切片去重
func UniqueSliceInt64(languages []int64) []int64 {
	result := make([]int64, 0, len(languages))
	temp := map[int64]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok { //如果字典中找不到元素，ok=false，!ok为true，就往切片中append元素。
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
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

// InSliceInt 元素是否存在于切片中
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
