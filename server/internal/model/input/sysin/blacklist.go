// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// BlacklistMaxSortInp 最大排序
type BlacklistMaxSortInp struct {
	Id int64
}

type BlacklistMaxSortModel struct {
	Sort int
}

// BlacklistEditInp 修改/新增字典数据
type BlacklistEditInp struct {
	entity.SysBlacklist
}
type BlacklistEditModel struct{}

// BlacklistDeleteInp 删除字典类型
type BlacklistDeleteInp struct {
	Id interface{}
}
type BlacklistDeleteModel struct{}

// BlacklistViewInp 获取信息
type BlacklistViewInp struct {
	Id int64
}

type BlacklistViewModel struct {
	entity.SysBlacklist
}

// BlacklistListInp 获取列表
type BlacklistListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Ip string
}

type BlacklistListModel struct {
	entity.SysBlacklist
}

// BlacklistStatusInp 更新状态
type BlacklistStatusInp struct {
	entity.SysBlacklist
}
type BlacklistStatusModel struct{}
