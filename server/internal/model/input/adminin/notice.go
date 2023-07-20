// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package adminin

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
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

// NoticeEditInp 修改/新增
type NoticeEditInp struct {
	entity.AdminNotice
	Receiver     []int64 `json:"receiver"     dc:"接收者"`
	SenderAvatar string  `json:"senderAvatar" dc:"发送者头像"`
}
type NoticeEditModel struct{}

// NoticeDeleteInp 删除字典类型
type NoticeDeleteInp struct {
	Id interface{} `json:"id" v:"required#公告ID不能为空" dc:"公告ID"`
}
type NoticeDeleteModel struct{}

// NoticeViewInp 获取信息
type NoticeViewInp struct {
	Id int64 `json:"id" v:"required#公告ID不能为空" dc:"公告ID"`
}

type NoticeViewModel struct {
	entity.AdminNotice
}

// NoticeListInp 获取列表
type NoticeListInp struct {
	form.PageReq
	form.StatusReq
	Title   string
	Content string
	Type    int64
}

type NoticeListModel struct {
	entity.AdminNotice
	ReadCount     float64            `json:"readCount"     dc:"阅读次数"`
	ReceiverGroup []form.AvatarGroup `json:"receiverGroup" dc:"接收人头像组"`
}

// NoticeStatusInp 更新状态
type NoticeStatusInp struct {
	entity.AdminNotice
}
type NoticeStatusModel struct{}

// NoticeUpReadInp 更新已读
type NoticeUpReadInp struct {
	Id int64 `json:"id"     v:"required#公告ID不能为空"   dc:"公告ID"`
}

// NoticeReadAllInp 全部已读
type NoticeReadAllInp struct {
	Type int64 `json:"type" v:"required#公告类型不能为空" dc:"公告类型"`
}

// PullMessagesInp 获取列表
type PullMessagesInp struct {
	Limit int `json:"limit" dc:"拉取最大数量限制"`
}

func (in *PullMessagesInp) Filter(ctx context.Context) (err error) {
	if in.Limit == 0 {
		in.Limit = 100
	}
	return
}

type PullMessagesRow struct {
	Id           int64       `json:"id"           dc:"消息ID"`
	Type         int64       `json:"type"         dc:"消息类型"`
	Title        string      `json:"title"        dc:"消息标题"`
	Content      string      `json:"content"      dc:"消息内容"`
	Tag          int64       `json:"tag"          dc:"标签"`
	Sort         int64       `json:"sort"         dc:"排序"`
	CreatedBy    int64       `json:"createdBy"    dc:"发送人"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	SenderAvatar string      `json:"senderAvatar" dc:"发送者头像"`
	IsRead       bool        `json:"isRead"       dc:"是否已读"`
}

type PullMessagesModel struct {
	List []*PullMessagesRow `json:"list" dc:"消息列表"`
	*NoticeUnreadCountModel
}

type NoticeUnreadCountInp struct {
	MemberId   int64
	MessageIds []int64
}

type NoticeUnreadCountModel struct {
	NotifyCount int `json:"notifyCount" dc:"未读通知数量"`
	NoticeCount int `json:"noticeCount" dc:"未读公告数量"`
	LetterCount int `json:"letterCount" dc:"未读私信数量"`
}

// NoticeMessageListInp 我的消息列表
type NoticeMessageListInp struct {
	form.PageReq
	Type int64 `json:"type" v:"required#公告类型不能为空" dc:"公告类型"`
}

type NoticeMessageListModel struct {
	Id           int64       `json:"id"           dc:"消息ID"`
	Type         int64       `json:"type"         dc:"消息类型"`
	Title        string      `json:"title"        dc:"消息标题"`
	Content      string      `json:"content"      dc:"消息内容"`
	Tag          int64       `json:"tag"          dc:"标签"`
	Sort         int64       `json:"sort"         dc:"排序"`
	CreatedBy    int64       `json:"createdBy"    dc:"发送人"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	SenderAvatar string      `json:"senderAvatar" dc:"发送者头像"`
	IsRead       bool        `json:"isRead"       dc:"是否已读"`
}
