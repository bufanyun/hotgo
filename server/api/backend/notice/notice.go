// Package notice
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package notice

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// ListReq 查询列表
type ListReq struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Title   string `json:"title"`
	Content string `json:"content"`
	g.Meta  `path:"/notice/list" method:"get" tags:"公告" summary:"获取公告列表"`
}

type ListRes struct {
	List []*adminin.NoticeListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取指定信息
type ViewReq struct {
	Id     int64 `json:"id" v:"required#公告ID不能为空" dc:"公告ID"`
	g.Meta `path:"/notice/view" method:"get" tags:"公告" summary:"获取指定信息"`
}
type ViewRes struct {
	*adminin.NoticeViewModel
}

// EditReq 修改/新增字典数据
type EditReq struct {
	entity.AdminNotice
	g.Meta `path:"/notice/edit" method:"post" tags:"公告" summary:"修改/新增公告"`
}
type EditRes struct{}

// DeleteReq 删除字典类型
type DeleteReq struct {
	Id     interface{} `json:"id" v:"required#公告ID不能为空" dc:"公告ID"`
	g.Meta `path:"/notice/delete" method:"post" tags:"公告" summary:"删除公告"`
}
type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	Id     int64 `json:"id" dc:"公告ID"`
	g.Meta `path:"/notice/max_sort" method:"get" tags:"公告" summary:"公告最大排序"`
}
type MaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

// StatusReq 更新公告状态
type StatusReq struct {
	entity.AdminNotice
	g.Meta `path:"/notice/status" method:"post" tags:"公告" summary:"更新公告状态"`
}
type StatusRes struct{}
