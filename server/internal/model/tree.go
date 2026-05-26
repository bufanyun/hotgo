// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package model

import "hotgo/internal/model/entity"

// DefaultTree 默认树表字段
type DefaultTree struct {
	Id    int64  `json:"id"             description:"ID"`
	Pid   int64  `json:"pid"            description:"父ID"`
	Level int    `json:"level"          description:"关系树等级"`
	Tree  string `json:"tree"           description:"关系树"`
}

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
