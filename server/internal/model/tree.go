// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package model

import "hotgo/internal/model/entity"

// TreeMenu 菜单树
type TreeMenu struct {
	entity.AdminMenu
	Children []*TreeMenu `json:"children"`
}

// LabelTreeMenu 菜单kv树
type LabelTreeMenu struct {
	entity.AdminMenu
	Key      int64            `json:"key"       description:"键名"`
	Label    string           `json:"label"       description:"键标签"`
	Children []*LabelTreeMenu `json:"children"`
}
