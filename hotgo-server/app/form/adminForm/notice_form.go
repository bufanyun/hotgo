package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

//  名称是否唯一
type NoticeNameUniqueReq struct {
	Title  string `json:"name" v:"required#公告名称不能为空"  description:"公告名称"`
	Id     int64  `json:"id" description:"公告ID"`
	g.Meta `path:"/notice/name_unique" method:"get" tags:"公告" summary:"公告名称是否唯一"`
}
type NoticeNameUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}

//  查询列表
type NoticeListReq struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name   string `json:"name"   description:"公告名称"`
	g.Meta `path:"/notice/list" method:"get" tags:"公告" summary:"获取公告列表"`
}

type NoticeListRes struct {
	List []*input.AdminNoticeListModel `json:"list"   description:"数据列表"`
	form.PageRes
}

//  获取指定信息
type NoticeViewReq struct {
	Id     string `json:"id" v:"required#公告ID不能为空" description:"公告ID"`
	g.Meta `path:"/notice/view" method:"get" tags:"公告" summary:"获取指定信息"`
}
type NoticeViewRes struct {
	*input.AdminNoticeViewModel
}

//  修改/新增
type NoticeEditReq struct {
	entity.AdminNotice
	g.Meta `path:"/notice/edit" method:"post" tags:"公告" summary:"修改/新增公告"`
}
type NoticeEditRes struct{}

//  删除
type NoticeDeleteReq struct {
	Id     interface{} `json:"id" v:"required#公告ID不能为空" description:"公告ID"`
	g.Meta `path:"/notice/delete" method:"post" tags:"公告" summary:"删除公告"`
}
type NoticeDeleteRes struct{}

//  最大排序
type NoticeMaxSortReq struct {
	Id     int64 `json:"id" description:"公告ID"`
	g.Meta `path:"/notice/max_sort" method:"get" tags:"公告" summary:"公告最大排序"`
}
type NoticeMaxSortRes struct {
	Sort int `json:"sort" description:"排序"`
}
