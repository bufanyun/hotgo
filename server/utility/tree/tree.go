// Package tree
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package tree

import (
	"github.com/gogf/gf/v2/util/gconv"
)

var pidName = "pid"

// GenTree 生成关系树 参考：https://blog.csdn.net/weixin_51546892/article/details/122876793
func GenTree(menus []map[string]interface{}) (realMenu []map[string]interface{}) {
	if len(menus) < 1 {
		return nil
	}

	minPid := GetMinPid(menus)
	formatMenu := make(map[int]map[string]interface{})
	for _, m := range menus {
		formatMenu[gconv.Int(m["id"])] = m
		if gconv.Int(m[pidName]) == minPid {
			realMenu = append(realMenu, m) // 需要返回的顶级菜单
		}
	}
	// 得益于都是地址操作,可以直接往父级塞
	for _, m := range formatMenu {
		if formatMenu[gconv.Int(m[pidName])] == nil {
			continue
		}
		if formatMenu[gconv.Int(m[pidName])]["children"] == nil {
			formatMenu[gconv.Int(m[pidName])]["children"] = []map[string]interface{}{}
		}

		formatMenu[gconv.Int(m[pidName])]["children"] = append(formatMenu[gconv.Int(m[pidName])]["children"].([]map[string]interface{}), m)
	}
	return
}

func GetMinPid(menus []map[string]interface{}) int {
	index := -1
	for _, m := range menus {
		if index == -1 {
			index = gconv.Int(m[pidName])
			continue
		}
		if gconv.Int(m[pidName]) < index {
			index = gconv.Int(m[pidName])
		}
	}

	if index == -1 {
		return 0
	}
	return index
}
