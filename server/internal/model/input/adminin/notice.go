// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminin

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// NoticeMaxSortInp 最大排序
type NoticeMaxSortInp struct {
	Id int64
}

type NoticeMaxSortModel struct {
	Sort int
}

// NoticeEditInp 修改/新增字典数据
type NoticeEditInp struct {
	entity.AdminNotice
	Receiver string `json:"receiver"  description:"接收者"`
	Reader   string `json:"reader"    description:"已读人"`
}
type NoticeEditModel struct{}

// NoticeDeleteInp 删除字典类型
type NoticeDeleteInp struct {
	Id interface{}
}
type NoticeDeleteModel struct{}

// NoticeViewInp 获取信息
type NoticeViewInp struct {
	Id int64
}

type NoticeViewModel struct {
	entity.AdminNotice
}

// NoticeListInp 获取列表
type NoticeListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Title   string
	Content string
}

type NoticeListModel struct {
	entity.AdminNotice
	ReceiveNum int `json:"receiveNum"`
}

// NoticeStatusInp 更新状态
type NoticeStatusInp struct {
	entity.AdminNotice
}
type NoticeStatusModel struct{}
