// Package provinces
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package provinces

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/provinces/list" method:"get" tags:"省市区" summary:"获取省市区列表"`
	sysin.ProvincesListInp
}

type ListRes struct {
	List []*sysin.ProvincesListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取信息
type ViewReq struct {
	g.Meta `path:"/provinces/view" method:"get" tags:"省市区" summary:"获取指定信息"`
	sysin.ProvincesViewInp
}

type ViewRes struct {
	*sysin.ProvincesViewModel
}

// EditReq 修改/新增
type EditReq struct {
	g.Meta `path:"/provinces/edit" method:"post" tags:"省市区" summary:"修改/新增省市区"`
	sysin.ProvincesEditInp
}

type EditRes struct{}

// DeleteReq 删除
type DeleteReq struct {
	g.Meta `path:"/provinces/delete" method:"post" tags:"省市区" summary:"删除省市区"`
	sysin.ProvincesDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	g.Meta `path:"/provinces/maxSort" method:"get" tags:"省市区" summary:"省市区最大排序"`
	sysin.ProvincesMaxSortInp
}

type MaxSortRes struct {
	*sysin.ProvincesMaxSortModel
}

// StatusReq 更新状态
type StatusReq struct {
	g.Meta `path:"/provinces/status" method:"post" tags:"省市区" summary:"更新省市区状态"`
	sysin.ProvincesStatusInp
}

type StatusRes struct{}

// TreeReq 关系树选项列表
type TreeReq struct {
	g.Meta `path:"/provinces/tree" tags:"省市区" method:"get" summary:"省市区关系树选项列表"`
}

type TreeRes struct {
	List []*sysin.ProvincesTree `json:"list"   dc:"数据列表"`
}

// ChildrenListReq 获取省市区下级列表
type ChildrenListReq struct {
	g.Meta `path:"/provinces/childrenList" method:"get" tags:"省市区" summary:"获取省市区下级列表"`
	sysin.ProvincesChildrenListInp
}

type ChildrenListRes struct {
	List []*sysin.ProvincesChildrenListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// UniqueIdReq 地区ID是否唯一
type UniqueIdReq struct {
	g.Meta `path:"/provinces/uniqueId" method:"get" tags:"省市区" summary:"地区ID是否唯一"`
	sysin.ProvincesUniqueIdInp
}

type UniqueIdRes struct {
	*sysin.ProvincesUniqueIdModel
}

// SelectReq 省市区选项
type SelectReq struct {
	g.Meta `path:"/provinces/select" method:"get" tags:"省市区"  summary:"省市区选项"`
	sysin.ProvincesSelectInp
}

type SelectRes struct {
	*sysin.ProvincesSelectModel
}

// CityLabelReq 获取指定城市标签
type CityLabelReq struct {
	g.Meta `path:"/provinces/cityLabel" method:"get" tags:"省市区" summary:"获取指定城市标签" `
	sysin.ProvincesCityLabelInp
}

type CityLabelRes sysin.ProvincesCityLabelModel
