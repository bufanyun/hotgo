// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"hotgo/api/admin/notice"
	"hotgo/internal/consts"
	"hotgo/internal/service"
)

var (
	Notice = cNotice{}
)

type cNotice struct{}

// Delete 删除
func (c *cNotice) Delete(ctx context.Context, req *notice.DeleteReq) (res *notice.DeleteRes, err error) {
	err = service.AdminNotice().Delete(ctx, &req.NoticeDeleteInp)
	return
}

// Edit 更新
func (c *cNotice) Edit(ctx context.Context, req *notice.EditReq) (res *notice.EditRes, err error) {
	err = service.AdminNotice().Edit(ctx, &req.NoticeEditInp)
	return
}

// MaxSort 最大排序
func (c *cNotice) MaxSort(ctx context.Context, req *notice.MaxSortReq) (res *notice.MaxSortRes, err error) {
	res = new(notice.MaxSortRes)
	res.NoticeMaxSortModel, err = service.AdminNotice().MaxSort(ctx, &req.NoticeMaxSortInp)
	return
}

// View 获取指定信息
func (c *cNotice) View(ctx context.Context, req *notice.ViewReq) (res *notice.ViewRes, err error) {
	data, err := service.AdminNotice().View(ctx, &req.NoticeViewInp)
	if err != nil {
		return
	}

	res = new(notice.ViewRes)
	res.NoticeViewModel = data
	return
}

// List 查看列表
func (c *cNotice) List(ctx context.Context, req *notice.ListReq) (res *notice.ListRes, err error) {
	list, totalCount, err := service.AdminNotice().List(ctx, &req.NoticeListInp)
	if err != nil {
		return nil, err
	}

	res = new(notice.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Status 更新状态
func (c *cNotice) Status(ctx context.Context, req *notice.StatusReq) (res *notice.StatusRes, err error) {
	err = service.AdminNotice().Status(ctx, &req.NoticeStatusInp)
	return
}

// Notify 更新通知
func (c *cNotice) Notify(ctx context.Context, req *notice.EditNotifyReq) (res *notice.EditNotifyRes, err error) {
	req.Type = consts.NoticeTypeNotify
	err = service.AdminNotice().Edit(ctx, &req.NoticeEditInp)
	return
}

// Notice 更新公告
func (c *cNotice) Notice(ctx context.Context, req *notice.EditNoticeReq) (res *notice.EditNoticeRes, err error) {
	req.Type = consts.NoticeTypeNotice
	err = service.AdminNotice().Edit(ctx, &req.NoticeEditInp)
	return
}

// Letter 更新私信
func (c *cNotice) Letter(ctx context.Context, req *notice.EditLetterReq) (res *notice.EditLetterRes, err error) {
	req.Type = consts.NoticeTypeLetter
	err = service.AdminNotice().Edit(ctx, &req.NoticeEditInp)
	return
}

// UpRead 更新已读
func (c *cNotice) UpRead(ctx context.Context, req *notice.UpReadReq) (res *notice.UpReadRes, err error) {
	err = service.AdminNotice().UpRead(ctx, &req.NoticeUpReadInp)
	return
}

// PullMessages 拉取未读消息列表
func (c *cNotice) PullMessages(ctx context.Context, req *notice.PullMessagesReq) (res *notice.PullMessagesRes, err error) {
	data, err := service.AdminNotice().PullMessages(ctx, &req.PullMessagesInp)
	if err != nil {
		return
	}

	res = new(notice.PullMessagesRes)
	res.PullMessagesModel = data
	return
}

// ReadAll 全部已读
func (c *cNotice) ReadAll(ctx context.Context, req *notice.ReadAllReq) (res *notice.ReadAllRes, err error) {
	err = service.AdminNotice().ReadAll(ctx, &req.NoticeReadAllInp)
	return
}

// MessageList 我的消息列表
func (c *cNotice) MessageList(ctx context.Context, req *notice.MessageListReq) (res *notice.MessageListRes, err error) {
	list, totalCount, err := service.AdminNotice().MessageList(ctx, &req.NoticeMessageListInp)
	if err != nil {
		return
	}

	res = new(notice.MessageListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
