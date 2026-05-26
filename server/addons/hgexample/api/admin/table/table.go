// Package table
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package table

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/internal/model/input/form"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/table/list" method:"get" tags:"表格例子" summary:"获取表格列表"`
	sysin.TableListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.TableListModel `json:"list" dc:"数据列表"`
}

// ExportReq 导出列表
type ExportReq struct {
	g.Meta `path:"/table/export" method:"get" tags:"表格例子" summary:"导出表格列表"`
	sysin.TableListInp
}

type ExportRes struct{}

// ViewReq 获取信息
type ViewReq struct {
	g.Meta `path:"/table/view" method:"get" tags:"表格例子" summary:"获取指定信息"`
	sysin.TableViewInp
}

type ViewRes struct {
	*sysin.TableViewModel
}

// EditReq 修改/新增
type EditReq struct {
	g.Meta `path:"/table/edit" method:"post" tags:"表格例子" summary:"修改/新增表格"`
	sysin.TableEditInp
}

type EditRes struct{}

// DeleteReq 删除
type DeleteReq struct {
	g.Meta `path:"/table/delete" method:"post" tags:"表格例子" summary:"删除表格"`
	sysin.TableDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	g.Meta `path:"/table/maxSort" method:"get" tags:"表格例子" summary:"表格最大排序"`
	sysin.TableMaxSortInp
}

type MaxSortRes struct {
	*sysin.TableMaxSortModel
}

// StatusReq 更新状态
type StatusReq struct {
	g.Meta `path:"/table/status" method:"post" tags:"表格例子" summary:"更新表格状态"`
	sysin.TableStatusInp
}

type StatusRes struct{}

// SwitchReq 更新开关状态
type SwitchReq struct {
	g.Meta `path:"/table/switch" method:"post" tags:"表格例子" summary:"更新表格状态"`
	sysin.TableSwitchInp
}

type SwitchRes struct{}
