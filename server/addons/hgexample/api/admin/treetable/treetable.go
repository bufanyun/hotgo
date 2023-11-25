// Package treetable
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package treetable

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/internal/model/input/form"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/treeTable/list" method:"get" tags:"表格" summary:"获取表格列表"`
	sysin.TreeTableListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.TreeTableListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出列表
type ExportReq struct {
	g.Meta `path:"/treeTable/export" method:"get" tags:"表格" summary:"导出表格列表"`
	sysin.TableListInp
}

type ExportRes struct{}

// ViewReq 获取信息
type ViewReq struct {
	g.Meta `path:"/treeTable/view" method:"get" tags:"表格" summary:"获取指定信息"`
	sysin.TableViewInp
}

type ViewRes struct {
	*sysin.TableViewModel
}

// EditReq 修改/新增
type EditReq struct {
	g.Meta `path:"/treeTable/edit" method:"post" tags:"表格" summary:"修改/新增表格"`
	sysin.TableEditInp
}

type EditRes struct{}

// DeleteReq 删除
type DeleteReq struct {
	g.Meta `path:"/treeTable/delete" method:"post" tags:"表格" summary:"删除表格"`
	sysin.TableDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	g.Meta `path:"/treeTable/maxSort" method:"get" tags:"表格" summary:"表格最大排序"`
	sysin.TableMaxSortInp
}

type MaxSortRes struct {
	*sysin.TableMaxSortModel
}

// StatusReq 更新状态
type StatusReq struct {
	g.Meta `path:"/treeTable/status" method:"post" tags:"表格" summary:"更新表格状态"`
	sysin.TableStatusInp
}

type StatusRes struct{}

// SwitchReq 更新开关状态
type SwitchReq struct {
	g.Meta `path:"/treeTable/switch" method:"post" tags:"表格" summary:"更新表格状态"`
	sysin.TableSwitchInp
}

type SwitchRes struct{}

// SelectReq 树形选项
type SelectReq struct {
	g.Meta `path:"/treeTable/select" method:"get" tags:"表格" summary:"树形选项"`
}

type SelectRes struct {
	List []*sysin.TableTree `json:"list"   dc:"数据列表"`
}
