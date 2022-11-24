// Package provinces
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package provinces

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询列表
type ListReq struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Title   string `json:"title"`
	Content string `json:"content"`
	g.Meta  `path:"/provinces/list" method:"get" tags:"省市区" summary:"获取省市区列表"`
}

type ListRes struct {
	List []*sysin.ProvincesListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取信息
type ViewReq struct {
	Id     int64 `json:"id" v:"required#省市区ID不能为空" dc:"省市区ID"`
	g.Meta `path:"/provinces/view" method:"get" tags:"省市区" summary:"获取指定信息"`
}
type ViewRes struct {
	*sysin.ProvincesViewModel
}

// EditReq 修改/新增
type EditReq struct {
	entity.SysProvinces
	g.Meta `path:"/provinces/edit" method:"post" tags:"省市区" summary:"修改/新增省市区"`
}
type EditRes struct{}

// DeleteReq 删除
type DeleteReq struct {
	Id     interface{} `json:"id" v:"required#省市区ID不能为空" dc:"省市区ID"`
	g.Meta `path:"/provinces/delete" method:"post" tags:"省市区" summary:"删除省市区"`
}
type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	Id     int64 `json:"id" dc:"省市区ID"`
	g.Meta `path:"/provinces/max_sort" method:"get" tags:"省市区" summary:"省市区最大排序"`
}
type MaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

// StatusReq 更新状态
type StatusReq struct {
	entity.SysProvinces
	g.Meta `path:"/provinces/status" method:"post" tags:"省市区" summary:"更新省市区状态"`
}
type StatusRes struct{}
