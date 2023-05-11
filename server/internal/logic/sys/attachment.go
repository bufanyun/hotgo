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
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/format"
	"hotgo/utility/validate"
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

// Delete 删除
func (s *sSysAttachment) Delete(ctx context.Context, in sysin.AttachmentDeleteInp) (err error) {
	_, err = s.Model(ctx).Where("id", in.Id).Delete()
	return
}

// Edit 修改/新增
func (s *sSysAttachment) Edit(ctx context.Context, in sysin.AttachmentEditInp) (err error) {
	if in.Name == "" {
		err = gerror.New("标题不能为空")
		return
	}

	// 修改
	if in.Id > 0 {
		_, err = s.Model(ctx).Where("id", in.Id).Data(in).Update()
		return
	}

	// 新增
	_, err = dao.SysAttachment.Ctx(ctx).Data(in).Insert()
	return
}

// Status 更新部门状态
func (s *sSysAttachment) Status(ctx context.Context, in sysin.AttachmentStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSliceInt(consts.StatusMap, in.Status) {
		err = gerror.New("状态不正确")
		return
	}

	// 修改
	_, err = s.Model(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	return
}

// MaxSort 最大排序
func (s *sSysAttachment) MaxSort(ctx context.Context, in sysin.AttachmentMaxSortInp) (res *sysin.AttachmentMaxSortModel, err error) {
	if in.Id > 0 {
		if err = s.Model(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			return
		}
	}

	if res == nil {
		res = new(sysin.AttachmentMaxSortModel)
	}
	res.Sort = form.DefaultMaxSort(ctx, res.Sort)
	return
}

// View 获取指定字典类型信息
func (s *sSysAttachment) View(ctx context.Context, in sysin.AttachmentViewInp) (res *sysin.AttachmentViewModel, err error) {
	err = s.Model(ctx).Where("id", in.Id).Scan(&res)
	return
}

// List 获取列表
func (s *sSysAttachment) List(ctx context.Context, in sysin.AttachmentListInp) (list []*sysin.AttachmentListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	if in.MemberId > 0 {
		mod = mod.Where("member_id", in.MemberId)
	}

	if in.Drive != "" {
		mod = mod.Where("drive", in.Drive)
	}

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("updated_at desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	conf, err := service.SysConfig().GetUpload(ctx)
	if err != nil {
		return
	}

	for _, v := range list {
		v.SizeFormat = format.FileSize(v.Size)
		v.FileUrl = service.CommonUpload().LastUrl(ctx, conf, v.FileUrl, v.Drive)
	}

	return
}

// Add 新增附件
func (s *sSysAttachment) Add(ctx context.Context, meta *sysin.UploadFileMeta, fullPath, drive string) (models *entity.SysAttachment, err error) {
	var (
		c              = contexts.Get(ctx)
		user           = c.User
		memberId int64 = 0
	)

	if user != nil {
		memberId = user.Id
	}

	models = &entity.SysAttachment{
		Id:        0,
		AppId:     c.Module,
		MemberId:  memberId,
		Drive:     drive,
		Size:      meta.Size,
		Path:      fullPath,
		FileUrl:   fullPath,
		Name:      meta.Filename,
		Kind:      meta.Kind,
		MetaType:  meta.MetaType,
		NaiveType: meta.NaiveType,
		Ext:       meta.Ext,
		Md5:       meta.Md5,
		Status:    consts.StatusEnabled,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}
	id, err := dao.SysAttachment.Ctx(ctx).Data(models).InsertAndGetId()
	if err != nil {
		return
	}

	models.Id = id
	return
}
