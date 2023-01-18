// Package test
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package test

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/test/list" method:"get" tags:"测试" summary:"获取测试列表"`
	adminin.TestListInp
}

type ListRes struct {
	form.PageRes
	List []*adminin.TestListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出列表
type ExportReq struct {
	g.Meta `path:"/test/export" method:"get" tags:"测试" summary:"导出测试列表"`
	adminin.TestListInp
}

type ExportRes struct{}

// ViewReq 获取信息
type ViewReq struct {
	g.Meta `path:"/test/view" method:"get" tags:"测试" summary:"获取指定信息"`
	adminin.TestViewInp
}
type ViewRes struct {
	*adminin.TestViewModel
}

// EditReq 修改/新增
type EditReq struct {
	g.Meta `path:"/test/edit" method:"post" tags:"测试" summary:"修改/新增测试"`
	adminin.TestEditInp
}
type EditRes struct{}

// DeleteReq 删除
type DeleteReq struct {
	g.Meta `path:"/test/delete" method:"post" tags:"测试" summary:"删除测试"`
	adminin.TestDeleteInp
}
type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	g.Meta `path:"/test/maxSort" method:"get" tags:"测试" summary:"测试最大排序"`
}
type MaxSortRes struct {
	*adminin.TestMaxSortModel
}

// StatusReq 更新状态
type StatusReq struct {
	g.Meta `path:"/test/status" method:"post" tags:"测试" summary:"更新测试状态"`
	adminin.TestStatusInp
}
type StatusRes struct{}

// SwitchReq 更新开关状态
type SwitchReq struct {
	g.Meta `path:"/test/switch" method:"post" tags:"测试" summary:"更新测试状态"`
	adminin.TestSwitchInp
}
type SwitchRes struct{}
