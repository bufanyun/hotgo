// Package convert
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package convert

// UniqueSlice 切片去重
func UniqueSlice[K comparable](languages []K) []K {
	result := make([]K, 0, len(languages))
	temp := map[K]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func Remove(sl []interface{}, f func(v1 interface{}) bool) []interface{} {
	for k, v := range sl {
		if f(v) {
			sl[k] = sl[len(sl)-1]
			sl = sl[:len(sl)-1]
			return sl
		}
	}
	return sl
}

func RemoveSlice[K comparable](src []K, sub K) []K {
	for k, v := range src {
		if v == sub {
			copy(src[k:], src[k+1:])
			return src[:len(src)-1]
		}
	}
	return src
}

// DifferenceSlice 比较两个切片，返回他们的差集
// slice1 := []int{1, 2, 3, 4, 5}
// slice2 := []int{4, 5, 6, 7, 8}
// fmt.Println(Difference(slice1, slice2)) // Output: [1 2 3]
func DifferenceSlice[T comparable](s1, s2 []T) []T {
	m := make(map[T]bool)
	for _, item := range s1 {
		m[item] = true
	}

	var diff []T
	for _, item := range s2 {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}
