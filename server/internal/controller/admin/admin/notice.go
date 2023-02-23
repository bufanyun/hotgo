// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/notice"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
)

var (
	Notice = cNotice{}
)

type cNotice struct{}

// Delete 删除
func (c *cNotice) Delete(ctx context.Context, req *notice.DeleteReq) (res *notice.DeleteRes, err error) {
	var in adminin.NoticeDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminNotice().Delete(ctx, in)
	return
}

// Edit 更新
func (c *cNotice) Edit(ctx context.Context, req *notice.EditReq) (res *notice.EditRes, err error) {
	var in adminin.NoticeEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	in.Receiver = req.Receiver
	err = service.AdminNotice().Edit(ctx, in)
	return
}

// MaxSort 最大排序
func (c *cNotice) MaxSort(ctx context.Context, req *notice.MaxSortReq) (res *notice.MaxSortRes, err error) {
	data, err := service.AdminNotice().MaxSort(ctx, adminin.NoticeMaxSortInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(notice.MaxSortRes)
	res.Sort = data.Sort
	return
}

// View 获取指定信息
func (c *cNotice) View(ctx context.Context, req *notice.ViewReq) (res *notice.ViewRes, err error) {
	data, err := service.AdminNotice().View(ctx, adminin.NoticeViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(notice.ViewRes)
	res.NoticeViewModel = data
	return
}

// List 查看列表
func (c *cNotice) List(ctx context.Context, req *notice.ListReq) (res *notice.ListRes, err error) {
	var in adminin.NoticeListInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := service.AdminNotice().List(ctx, in)
	if err != nil {
		return nil, err
	}

	res = new(notice.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// Status 更新部门状态
func (c *cNotice) Status(ctx context.Context, req *notice.StatusReq) (res *notice.StatusRes, err error) {
	var in adminin.NoticeStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	err = service.AdminNotice().Status(ctx, in)
	return
}

// Notify 更新通知
func (c *cNotice) Notify(ctx context.Context, req *notice.EditNotifyReq) (res *notice.EditNotifyRes, err error) {
	var in adminin.NoticeEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	in.Type = consts.NoticeTypeNotify
	in.Receiver = req.Receiver
	err = service.AdminNotice().Edit(ctx, in)
	return
}

// Notice 更新公告
func (c *cNotice) Notice(ctx context.Context, req *notice.EditNoticeReq) (res *notice.EditNoticeRes, err error) {
	var in adminin.NoticeEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	in.Type = consts.NoticeTypeNotice
	in.Receiver = req.Receiver
	err = service.AdminNotice().Edit(ctx, in)
	return
}

// Letter 更新私信
func (c *cNotice) Letter(ctx context.Context, req *notice.EditLetterReq) (res *notice.EditLetterRes, err error) {
	var in adminin.NoticeEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	in.Type = consts.NoticeTypeLetter
	in.Receiver = req.Receiver
	err = service.AdminNotice().Edit(ctx, in)
	return
}

// UpRead 更新已读
func (c *cNotice) UpRead(ctx context.Context, req *notice.UpReadReq) (res *notice.UpReadRes, err error) {
	var in adminin.NoticeUpReadInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	err = service.AdminNotice().UpRead(ctx, in)
	return res, nil
}

// PullMessages 拉取未读消息列表
func (c *cNotice) PullMessages(ctx context.Context, req *notice.PullMessagesReq) (res *notice.PullMessagesRes, err error) {
	var in adminin.PullMessagesInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if in.Limit == 0 {
		in.Limit = 100
	}

	data, err := service.AdminNotice().PullMessages(ctx, in)
	if err != nil {
		return
	}

	res = new(notice.PullMessagesRes)
	res.PullMessagesModel = data
	return
}

// ReadAll 全部已读
func (c *cNotice) ReadAll(ctx context.Context, req *notice.ReadAllReq) (res *notice.ReadAllRes, err error) {
	var in adminin.NoticeReadAllInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminNotice().ReadAll(ctx, in)
	return
}

// MessageList 我的消息列表
func (c *cNotice) MessageList(ctx context.Context, req *notice.MessageListReq) (res *notice.MessageListRes, err error) {
	var in adminin.NoticeMessageListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.AdminNotice().MessageList(ctx, in)
	if err != nil {
		return
	}

	res = new(notice.MessageListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}
