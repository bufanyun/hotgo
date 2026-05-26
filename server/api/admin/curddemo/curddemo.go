// Package curddemo
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package curddemo

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询CURD列表列表
type ListReq struct {
	g.Meta `path:"/curdDemo/list" method:"get" tags:"CURD列表" summary:"获取CURD列表列表"`
	sysin.CurdDemoListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.CurdDemoListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出CURD列表列表
type ExportReq struct {
	g.Meta `path:"/curdDemo/export" method:"get" tags:"CURD列表" summary:"导出CURD列表列表"`
	sysin.CurdDemoListInp
}

type ExportRes struct{}

// ViewReq 获取CURD列表指定信息
type ViewReq struct {
	g.Meta `path:"/curdDemo/view" method:"get" tags:"CURD列表" summary:"获取CURD列表指定信息"`
	sysin.CurdDemoViewInp
}

type ViewRes struct {
	*sysin.CurdDemoViewModel
}

// EditReq 修改/新增CURD列表
type EditReq struct {
	g.Meta `path:"/curdDemo/edit" method:"post" tags:"CURD列表" summary:"修改/新增CURD列表"`
	sysin.CurdDemoEditInp
}

type EditRes struct{}

// DeleteReq 删除CURD列表
type DeleteReq struct {
	g.Meta `path:"/curdDemo/delete" method:"post" tags:"CURD列表" summary:"删除CURD列表"`
	sysin.CurdDemoDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取CURD列表最大排序
type MaxSortReq struct {
	g.Meta `path:"/curdDemo/maxSort" method:"get" tags:"CURD列表" summary:"获取CURD列表最大排序"`
	sysin.CurdDemoMaxSortInp
}

type MaxSortRes struct {
	*sysin.CurdDemoMaxSortModel
}

// StatusReq 更新CURD列表状态
type StatusReq struct {
	g.Meta `path:"/curdDemo/status" method:"post" tags:"CURD列表" summary:"更新CURD列表状态"`
	sysin.CurdDemoStatusInp
}

type StatusRes struct{}

// SwitchReq 更新CURD列表开关状态
type SwitchReq struct {
	g.Meta `path:"/curdDemo/switch" method:"post" tags:"CURD列表" summary:"更新CURD列表状态"`
	sysin.CurdDemoSwitchInp
}

type SwitchRes struct{}