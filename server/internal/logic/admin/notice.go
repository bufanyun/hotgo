// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
	"hotgo/utility/simple"
	"hotgo/utility/validate"
)

type sAdminNotice struct{}

func NewAdminNotice() *sAdminNotice {
	return &sAdminNotice{}
}

func init() {
	service.RegisterAdminNotice(NewAdminNotice())
}

// Model Orm模型
func (s *sAdminNotice) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AdminNotice.Ctx(ctx), option...)
}

// Delete 删除
func (s *sAdminNotice) Delete(ctx context.Context, in adminin.NoticeDeleteInp) error {
	_, err := s.Model(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sAdminNotice) Edit(ctx context.Context, in adminin.NoticeEditInp) (err error) {
	var member = contexts.Get(ctx).User
	if member == nil {
		err = gerror.New("获取用户信息失败！")
		return
	}

	if in.Title == "" {
		err = gerror.New("标题不能为空")
		return
	}

	if in.Type == consts.NoticeTypeLetter && len(in.Receiver) == 0 {
		err = gerror.New("私信类型必须选择接收人")
		return
	}

	// 检查选项接收人是否合法
	if in.Type == consts.NoticeTypeLetter {
		count, _ := dao.AdminMember.Ctx(ctx).Handler(handler.FilterAuthWithField("id")).WhereIn("id", in.Receiver).Count()
		if count != len(in.Receiver) {
			err = gerror.New("接收人不合法")
			return
		}
		in.SenderAvatar = member.Avatar
	}

	// 修改
	if in.Id > 0 {
		in.UpdatedBy = member.Id
		_, err = s.Model(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}
		return nil
	}

	// 新增
	in.CreatedBy = member.Id
	in.CreatedAt = gtime.Now()
	in.Id, err = s.Model(ctx, &handler.Option{FilterAuth: false}).Data(in).InsertAndGetId()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	// 推送通知
	response := &websocket.WResponse{
		Event: "notice",
		Data:  in,
	}
	simple.SafeGo(ctx, func(ctx context.Context) {
		if in.Type == consts.NoticeTypeLetter {
			for _, receiverId := range in.Receiver {
				websocket.SendToUser(receiverId, response)
			}
		} else {
			websocket.SendToAll(response)
		}
	})

	return nil
}

// Status 更新部门状态
func (s *sAdminNotice) Status(ctx context.Context, in adminin.NoticeStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return err
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return err
	}

	if !validate.InSliceInt(consts.StatusMap, in.Status) {
		err = gerror.New("状态不正确")
		return err
	}

	// 修改
	_, err = s.Model(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// MaxSort 最大排序
func (s *sAdminNotice) MaxSort(ctx context.Context, in adminin.NoticeMaxSortInp) (res *adminin.NoticeMaxSortModel, err error) {
	if err = dao.AdminNotice.Ctx(ctx).Order("sort desc").Scan(&res); err != nil {
		return
	}

	if res == nil {
		res = new(adminin.NoticeMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(ctx, res.Sort)
	return
}

// View 获取指定字典类型信息
func (s *sAdminNotice) View(ctx context.Context, in adminin.NoticeViewInp) (res *adminin.NoticeViewModel, err error) {
	if err = s.Model(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

// List 获取列表
func (s *sAdminNotice) List(ctx context.Context, in adminin.NoticeListInp) (list []*adminin.NoticeListModel, totalCount int, err error) {
	var memberId = contexts.GetUserId(ctx)
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	mod := s.Model(ctx)

	if in.Title != "" {
		mod = mod.WhereLike("title", "%"+in.Title+"%")
	}

	if in.Content != "" {
		mod = mod.WhereLike("content", "%"+in.Content+"%")
	}

	if in.Type > 0 {
		mod = mod.Where("type", in.Type)
	}

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	if totalCount == 0 {
		return list, totalCount, nil
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	for _, v := range list {
		// 接收人头像组
		if v.Type == consts.NoticeTypeLetter {
			err = dao.AdminMember.Ctx(ctx).
				Fields("real_name as name,avatar as src").
				WhereIn("id", v.Receiver.Var().Int64s()).
				Scan(&v.ReceiverGroup)
			if err != nil {
				return
			}
		}
		if v.ReceiverGroup == nil || len(v.ReceiverGroup) == 0 {
			v.ReceiverGroup = make([]form.AvatarGroup, 0)
		}

		// 阅读次数
		v.ReadCount, err = dao.AdminNoticeRead.Ctx(ctx).Where("notice_id", v.Id).Sum("clicks")
		if err != nil {
			return
		}
	}
	return list, totalCount, err
}

// PullMessages 拉取未读消息列表
func (s *sAdminNotice) PullMessages(ctx context.Context, in adminin.PullMessagesInp) (res *adminin.PullMessagesModel, err error) {
	var memberId = contexts.GetUserId(ctx)
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	messageIds, err := s.messageIds(ctx, memberId)
	if err != nil {
		return
	}

	res = new(adminin.PullMessagesModel)
	unread, err := s.UnreadCount(ctx, adminin.NoticeUnreadCountInp{MemberId: memberId, MessageIds: messageIds})
	if err != nil {
		return
	}

	if unread != nil {
		res.NoticeUnreadCountModel = unread
	}

	if err = s.Model(ctx).WhereIn("id", messageIds).Limit(in.Limit).Order("id desc").Scan(&res.List); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	for _, v := range res.List {
		count, _ := dao.AdminNoticeRead.Ctx(ctx).Where("notice_id", v.Id).Where("member_id", memberId).Count()
		if count > 0 {
			v.IsRead = true
		}

		if v.Type == consts.NoticeTypeLetter {
			val, err := dao.AdminMember.Ctx(ctx).Fields("avatar").Where("id", v.CreatedBy).Value()
			if err == nil {
				v.SenderAvatar = val.String()
			}
		}
	}
	return
}

// UnreadCount 获取所有类型消息的未读数量
func (s *sAdminNotice) UnreadCount(ctx context.Context, in adminin.NoticeUnreadCountInp) (res *adminin.NoticeUnreadCountModel, err error) {
	if in.MemberId <= 0 {
		if in.MemberId = contexts.GetUserId(ctx); in.MemberId <= 0 {
			err = gerror.New("获取用户信息失败！")
			return
		}
	}

	if len(in.MessageIds) == 0 {
		in.MessageIds, err = s.messageIds(ctx, in.MemberId)
		if err != nil {
			return
		}

		if len(in.MessageIds) == 0 {
			return
		}
	}

	stat := func(t int) (count int) {
		all, err := dao.AdminNotice.Ctx(ctx).As("nr").
			Where("type =? and id IN(?)", t, in.MessageIds).
			Count()
		if err != nil {
			g.Log().Infof(ctx, "UnreadCount stat err:%+v", err)
			return
		}

		if all == 0 {
			return
		}

		read, err := dao.AdminNoticeRead.Ctx(ctx).As("nr").
			LeftJoin("admin_notice n", "nr.notice_id=n.id").
			Where("n.type = ? and n.id IN(?)", t, in.MessageIds).
			Where("nr.member_id", in.MemberId).
			Count()
		if err != nil {
			g.Log().Infof(ctx, "UnreadCount stat2 err:%+v", err)
			return
		}
		count = all - read
		return
	}

	res = new(adminin.NoticeUnreadCountModel)
	res.NotifyCount = stat(consts.NoticeTypeNotify)
	res.NoticeCount = stat(consts.NoticeTypeNotice)
	res.LetterCount = stat(consts.NoticeTypeLetter)
	return
}

// messageIds 获取我的消息所有的消息ID
func (s *sAdminNotice) messageIds(ctx context.Context, memberId int64) (ids []int64, err error) {
	mod := s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields("id").
		Where("status", consts.StatusEnabled).
		Where("(`type` IN(?) OR (`type` = ? and JSON_CONTAINS(`receiver`,'"+gconv.String(memberId)+"')))",
			[]int{consts.NoticeTypeNotify, consts.NoticeTypeNotice}, consts.NoticeTypeLetter,
		)

	array, err := mod.Array()
	if err != nil {
		return nil, err
	}

	for _, v := range array {
		ids = append(ids, v.Int64())
	}
	return
}

// UpRead 更新已读
func (s *sAdminNotice) UpRead(ctx context.Context, in adminin.NoticeUpReadInp) (err error) {
	var (
		data     *entity.AdminNotice
		memberId = contexts.GetUserId(ctx)
	)

	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}
	if err = dao.AdminNotice.Ctx(ctx).Where("id", in.Id).Scan(&data); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if data == nil {
		return gerror.New("公告不存在")
	}

	return s.updatedReadClicks(ctx, in.Id, memberId)
}

// ReadAll 已读全部
func (s *sAdminNotice) ReadAll(ctx context.Context, in adminin.NoticeReadAllInp) (err error) {
	var memberId = contexts.GetUserId(ctx)
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	allMessageIds, err := s.messageIds(ctx, memberId)
	if err != nil {
		return
	}

	if len(allMessageIds) == 0 {
		return
	}

	array, err := dao.AdminNotice.Ctx(ctx).
		Fields("id").
		Where("type = ? and id IN(?)", in.Type, allMessageIds).
		Array()
	if err != nil {
		return
	}

	var messageIds []int64
	for _, v := range array {
		messageIds = append(messageIds, v.Int64())
	}

	array, err = dao.AdminNoticeRead.Ctx(ctx).As("nr").
		Fields("nr.notice_id").
		LeftJoin("admin_notice n", "nr.notice_id=n.id").
		Where("n.type = ? and n.id IN(?)", in.Type, messageIds).
		Where("nr.member_id", memberId).
		Array()
	if err != nil {
		return
	}

	var readIds []int64
	for _, v := range array {
		readIds = append(readIds, v.Int64())
	}

	for _, messageId := range messageIds {
		if !validate.InSliceInt64(readIds, messageId) {
			if err = s.updatedReadClicks(ctx, messageId, memberId); err != nil {
				return
			}
		}
	}

	return
}

// updatedReadClicks 更新公告已读次数
func (s *sAdminNotice) updatedReadClicks(ctx context.Context, noticeId, memberId int64) (err error) {
	var (
		models *entity.AdminNoticeRead
	)
	err = dao.AdminNoticeRead.Ctx(ctx).
		Where(dao.AdminNoticeRead.Columns().NoticeId, noticeId).
		Where(dao.AdminNoticeRead.Columns().MemberId, memberId).
		Scan(&models)
	if err != nil {
		return
	}

	if models == nil {
		_, err = dao.AdminNoticeRead.Ctx(ctx).Data(entity.AdminNoticeRead{NoticeId: noticeId, MemberId: memberId}).Insert()
		return
	}
	_, err = dao.AdminNoticeRead.Ctx(ctx).Where(dao.AdminNoticeRead.Columns().Id, models.Id).Increment("clicks", 1)
	return
}

// MessageList 我的消息列表
func (s *sAdminNotice) MessageList(ctx context.Context, in adminin.NoticeMessageListInp) (list []*adminin.NoticeMessageListModel, totalCount int, err error) {
	var memberId = contexts.GetUserId(ctx)
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	allMessageIds, err := s.messageIds(ctx, memberId)
	if err != nil {
		return
	}

	if len(allMessageIds) == 0 {
		return
	}

	mod := s.Model(ctx, &handler.Option{FilterAuth: false}).WhereIn("id", allMessageIds).Where("type", in.Type)
	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	if totalCount == 0 {
		return list, totalCount, nil
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	for _, v := range list {
		count, _ := dao.AdminNoticeRead.Ctx(ctx).Where("notice_id", v.Id).Where("member_id", memberId).Count()
		if count > 0 {
			v.IsRead = true
		}

		if v.Type == consts.NoticeTypeLetter {
			val, err := dao.AdminMember.Ctx(ctx).Fields("avatar").Where("id", v.CreatedBy).Value()
			if err == nil {
				v.SenderAvatar = val.String()
			}
		}
	}
	return list, totalCount, err
}
