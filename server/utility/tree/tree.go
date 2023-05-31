// Package tree
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tree

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

const (
	treeDefaultId       = "id"       // 树ID
	treeDefaultPid      = "pid"      // 上级树ID
	treeDefaultChildren = "children" // 下级属性
	treeBeginCut        = "tr_"      // 树标识开头
	treeEndCut          = " "        // 树标识结尾
)

// GenOption 生成选项
type GenOption struct {
	IdField       string
	PidField      string
	ChildrenField string
}

// GenLabel 生成关系树标识
func GenLabel(basic string, appendId int64) string {
	return fmt.Sprintf("%v%v%v%v", basic, treeBeginCut, appendId, treeEndCut)
}

// GetIdLabel 获取指定Id的树标签
func GetIdLabel(Id int64) string {
	return fmt.Sprintf("%v%v%v", treeBeginCut, Id, treeEndCut)
}

// GetIds 获取上级ID集合
func GetIds(tree string) (ids []int64) {
	idsStr := strings.Split(tree, treeEndCut)
	if len(idsStr) == 0 {
		return
	}

	for _, v := range idsStr {
		newId := gconv.Int64(strings.ReplaceAll(v, treeBeginCut, ""))
		if newId > 0 {
			ids = append(ids, newId)
		}
	}
	return
}

// GenTree 生成关系树
func GenTree(menus []map[string]interface{}) (realMenu []map[string]interface{}) {
	return GenTreeWithField(menus, GenOption{
		IdField:       treeDefaultId,
		PidField:      treeDefaultPid,
		ChildrenField: treeDefaultChildren,
	})
}

// GenTreeWithField 生成关系树 自定义生成属性
func GenTreeWithField(menus []map[string]interface{}, op GenOption) (realMenu []map[string]interface{}) {
	if len(menus) < 1 {
		return
	}

	minPid := GetMinPid(menus, op.PidField)
	formatMenu := make(map[int]map[string]interface{})
	for _, m := range menus {
		formatMenu[gconv.Int(m[op.IdField])] = m
		if gconv.Int(m[op.PidField]) == minPid {
			realMenu = append(realMenu, m) // 需要返回的顶级菜单
		}
	}
	// 得益于都是地址操作,可以直接往父级塞
	for _, m := range formatMenu {
		if formatMenu[gconv.Int(m[op.PidField])] == nil {
			continue
		}
		if formatMenu[gconv.Int(m[op.PidField])][op.ChildrenField] == nil {
			formatMenu[gconv.Int(m[op.PidField])][op.ChildrenField] = []map[string]interface{}{}
		}

		formatMenu[gconv.Int(m[op.PidField])][op.ChildrenField] = append(formatMenu[gconv.Int(m[op.PidField])][op.ChildrenField].([]map[string]interface{}), m)
	}
	return
}

func GetMinPid(menus []map[string]interface{}, pidField string) int {
	index := -1
	for _, m := range menus {
		if index == -1 {
			index = gconv.Int(m[pidField])
			continue
		}
		if gconv.Int(m[pidField]) < index {
			index = gconv.Int(m[pidField])
		}
	}

	if index == -1 {
		return 0
	}
	return index
}
