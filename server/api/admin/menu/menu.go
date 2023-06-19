// Package menu
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package menu

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/adminin"
)

// EditReq 修改/新增菜单
type EditReq struct {
	g.Meta `path:"/menu/edit" method:"post" tags:"菜单" summary:"修改/新增菜单"`
	adminin.MenuEditInp
}

type EditRes struct{}

// DeleteReq 删除菜单
type DeleteReq struct {
	g.Meta `path:"/menu/delete" method:"post" tags:"菜单" summary:"删除菜单"`
	adminin.MenuDeleteInp
}

type DeleteRes struct{}

// ListReq 获取菜单列表
type ListReq struct {
	g.Meta `path:"/menu/list" method:"get" tags:"菜单" summary:"获取菜单列表"`
	adminin.MenuListInp
}

type ListRes struct {
	List []map[string]interface{} `json:"list"   dc:"数据列表"`
}
