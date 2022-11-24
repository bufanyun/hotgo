// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// ProvincesMaxSortInp 最大排序
type ProvincesMaxSortInp struct {
	Id int64
}

type ProvincesMaxSortModel struct {
	Sort int
}

// ProvincesEditInp 修改/新增字典数据
type ProvincesEditInp struct {
	entity.SysProvinces
}
type ProvincesEditModel struct{}

// ProvincesDeleteInp 删除字典类型
type ProvincesDeleteInp struct {
	Id interface{}
}
type ProvincesDeleteModel struct{}

// ProvincesViewInp 获取信息
type ProvincesViewInp struct {
	Id int64
}

type ProvincesViewModel struct {
	entity.SysProvinces
}

// ProvincesListInp 获取列表
type ProvincesListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Title   string
	Content string
}

type ProvincesListModel struct {
	entity.SysProvinces
}

// ProvincesStatusInp 更新状态
type ProvincesStatusInp struct {
	entity.SysProvinces
}
type ProvincesStatusModel struct{}
