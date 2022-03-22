//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

//  名称是否唯一
type NoticeNameUniqueReq struct {
	g.Meta `path:"/notice/name_unique" method:"get" tags:"公告" summary:"公告名称是否唯一"`
	Title  string `json:"name" v:"required#公告名称不能为空"  dc:"公告名称"`
	Id     int64  `json:"id" dc:"公告ID"`
}
type NoticeNameUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

//  查询列表
type NoticeListReq struct {
	g.Meta `path:"/notice/list" method:"get" tags:"公告" summary:"获取公告列表"`
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name string `json:"name"   dc:"公告名称"`
}

type NoticeListRes struct {
	List []*input.AdminNoticeListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

//  获取指定信息
type NoticeViewReq struct {
	g.Meta `path:"/notice/view" method:"get" tags:"公告" summary:"获取指定信息"`
	Id     string `json:"id" v:"required#公告ID不能为空" dc:"公告ID"`
}
type NoticeViewRes struct {
	*input.AdminNoticeViewModel
}

//  修改/新增
type NoticeEditReq struct {
	g.Meta `path:"/notice/edit" method:"post" tags:"公告" summary:"修改/新增公告"`
	entity.AdminNotice
}
type NoticeEditRes struct{}

//  删除
type NoticeDeleteReq struct {
	g.Meta `path:"/notice/delete" method:"post" tags:"公告" summary:"删除公告"`
	Id     interface{} `json:"id" v:"required#公告ID不能为空" dc:"公告ID"`
}
type NoticeDeleteRes struct{}

//  最大排序
type NoticeMaxSortReq struct {
	g.Meta `path:"/notice/max_sort" method:"get" tags:"公告" summary:"公告最大排序"`
	Id     int64 `json:"id" dc:"公告ID"`
}
type NoticeMaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}
