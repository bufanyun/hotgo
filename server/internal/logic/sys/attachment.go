// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/format"
)

type sSysAttachment struct{}

func NewSysAttachment() *sSysAttachment {
	return &sSysAttachment{}
}

func init() {
	service.RegisterSysAttachment(NewSysAttachment())
}

// Model ORM模型
func (s *sSysAttachment) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysAttachment.Ctx(ctx), option...)
}

// Delete 删除附件
func (s *sSysAttachment) Delete(ctx context.Context, in *sysin.AttachmentDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除附件失败，请稍后重试！")
	}
	return
}

// View 获取附件信息
func (s *sSysAttachment) View(ctx context.Context, in *sysin.AttachmentViewInp) (res *sysin.AttachmentViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取附件信息失败，请稍后重试！")
	}
	return
}

// List 获取附件列表
func (s *sSysAttachment) List(ctx context.Context, in *sysin.AttachmentListInp) (list []*sysin.AttachmentListModel, totalCount int, err error) {
	mod := s.Model(ctx)
	memberId := contexts.GetUserId(ctx)

	// 超管允许查看指定用户的附件
	if service.AdminMember().VerifySuperId(ctx, memberId) && in.MemberId > 0 {
		mod = mod.Where(dao.SysAttachment.Columns().MemberId, in.MemberId)
	} else {
		mod = mod.Where(dao.SysAttachment.Columns().MemberId, memberId)
	}

	if in.Drive != "" {
		mod = mod.Where(dao.SysAttachment.Columns().Drive, in.Drive)
	}

	if in.Name != "" {
		mod = mod.WhereLike(dao.SysAttachment.Columns().Name, "%"+in.Name+"%")
	}

	if in.Status > 0 {
		mod = mod.Where(dao.SysAttachment.Columns().Status, in.Status)
	}

	if len(in.UpdatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysAttachment.Columns().UpdatedAt, in.UpdatedAt[0], in.UpdatedAt[1])
	}

	if in.Kind != "" {
		mod = mod.Where(dao.SysAttachment.Columns().Kind, in.Kind)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取附件数据行失败！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).OrderDesc(dao.SysAttachment.Columns().UpdatedAt).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取附件列表失败！")
		return
	}

	for _, v := range list {
		v.SizeFormat = format.FileSize(v.Size)
		v.FileUrl = storager.LastUrl(ctx, v.FileUrl, v.Drive)
	}
	return
}

// ClearKind 清空上传类型
func (s *sSysAttachment) ClearKind(ctx context.Context, in *sysin.AttachmentClearKindInp) (err error) {
	memberId := contexts.GetUserId(ctx)
	if _, err = s.Model(ctx).Where(dao.SysAttachment.Columns().MemberId, memberId).Where(dao.SysAttachment.Columns().Kind, in.Kind).Delete(); err != nil {
		err = gerror.Wrap(err, "删除附件上传类型失败，请稍后重试！")
	}
	return
}
