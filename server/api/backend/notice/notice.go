// Package notice
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package notice

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/notice/list" method:"get" tags:"公告" summary:"获取公告列表"`
	adminin.NoticeListInp
}

type ListRes struct {
	List []*adminin.NoticeListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取指定信息
type ViewReq struct {
	g.Meta `path:"/notice/view" method:"get" tags:"公告" summary:"获取指定信息"`
	adminin.NoticeViewInp
}

type ViewRes struct {
	*adminin.NoticeViewModel
}

// EditReq 修改/新增字典数据
type EditReq struct {
	g.Meta `path:"/notice/edit" method:"post" tags:"公告" summary:"修改/新增公告"`
	adminin.NoticeEditInp
}

type EditRes struct{}

// DeleteReq 删除字典类型
type DeleteReq struct {
	g.Meta `path:"/notice/delete" method:"post" tags:"公告" summary:"删除公告"`
	adminin.NoticeDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 最大排序
type MaxSortReq struct {
	g.Meta `path:"/notice/maxSort" method:"get" tags:"公告" summary:"公告最大排序"`
	adminin.NoticeMaxSortInp
}

type MaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

// StatusReq 更新公告状态
type StatusReq struct {
	g.Meta `path:"/notice/status" method:"post" tags:"公告" summary:"更新公告状态"`
	adminin.NoticeStatusInp
}

type StatusRes struct{}

// EditNotifyReq 修改/新增通知
type EditNotifyReq struct {
	g.Meta `path:"/notice/editNotify" method:"post" tags:"公告" summary:"修改/新增通知"`
	adminin.NoticeEditInp
}

type EditNotifyRes struct{}

// EditNoticeReq 修改/新增公告
type EditNoticeReq struct {
	g.Meta `path:"/notice/editNotice" method:"post" tags:"公告" summary:"修改/新增公告"`
	adminin.NoticeEditInp
}

type EditNoticeRes struct{}

// EditLetterReq 修改/新增公告
type EditLetterReq struct {
	g.Meta `path:"/notice/editLetter" method:"post" tags:"公告" summary:"修改/新增私信"`
	adminin.NoticeEditInp
}

type EditLetterRes struct{}

// PullMessagesReq 拉取我的消息
type PullMessagesReq struct {
	g.Meta `path:"/notice/pullMessages" method:"get" tags:"公告" summary:"拉取我的消息"`
	adminin.PullMessagesInp
}

type PullMessagesRes struct {
	*adminin.PullMessagesModel
}

// ReadAllReq 全部已读
type ReadAllReq struct {
	g.Meta `path:"/notice/readAll" method:"post" tags:"公告" summary:"全部已读"`
	adminin.NoticeReadAllInp
}

type ReadAllRes struct {
}

// UpReadReq 更新已读
type UpReadReq struct {
	g.Meta `path:"/notice/upRead" method:"post" tags:"公告" summary:"更新已读"`
	adminin.NoticeUpReadInp
}
type UpReadRes struct{}

// MessageListReq 我的消息列表
type MessageListReq struct {
	g.Meta `path:"/notice/messageList" method:"get" tags:"公告" summary:"我的消息列表"`
	adminin.NoticeMessageListInp
}

type MessageListRes struct {
	List []*adminin.NoticeMessageListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}
