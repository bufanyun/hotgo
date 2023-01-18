// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
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

// Delete 删除
func (s *sSysAttachment) Delete(ctx context.Context, in sysin.AttachmentDeleteInp) error {
	_, err := dao.SysAttachment.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sSysAttachment) Edit(ctx context.Context, in sysin.AttachmentEditInp) (err error) {
	if in.Name == "" {
		err = gerror.New("标题不能为空")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.SysAttachment.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()
	_, err = dao.SysAttachment.Ctx(ctx).Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

// Status 更新部门状态
func (s *sSysAttachment) Status(ctx context.Context, in sysin.AttachmentStatusInp) (err error) {
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
	in.UpdatedAt = gtime.Now()
	_, err = dao.SysAttachment.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// MaxSort 最大排序
func (s *sSysAttachment) MaxSort(ctx context.Context, in sysin.AttachmentMaxSortInp) (*sysin.AttachmentMaxSortModel, error) {
	var res sysin.AttachmentMaxSortModel
	if in.Id > 0 {
		if err := dao.SysAttachment.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	res.Sort = res.Sort + 10

	return &res, nil
}

// View 获取指定字典类型信息
func (s *sSysAttachment) View(ctx context.Context, in sysin.AttachmentViewInp) (res *sysin.AttachmentViewModel, err error) {
	if err = dao.SysAttachment.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

// List 获取列表
func (s *sSysAttachment) List(ctx context.Context, in sysin.AttachmentListInp) (list []*sysin.AttachmentListModel, totalCount int, err error) {
	mod := dao.SysAttachment.Ctx(ctx)

	// 访问路径
	if in.MemberId > 0 {
		mod = mod.Where("member_id", in.MemberId)
	}

	// 模块
	if in.Drive != "" {
		mod = mod.Where("drive", in.Drive)
	}

	// 请求方式
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

	if err = mod.Page(in.Page, in.PerPage).Order("updated_at desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	conf, err := service.SysConfig().GetUpload(ctx)
	if err != nil {
		return list, totalCount, err
	}
	for _, v := range list {
		v.SizeFormat = format.FileSize(v.Size)
		v.FileUrl = service.CommonUpload().LastUrl(ctx, conf, v.FileUrl, v.Drive)
	}

	return list, totalCount, err
}

// Add 新增附件
func (s *sSysAttachment) Add(ctx context.Context, meta *sysin.UploadFileMeta, fullPath, drive string) (data *entity.SysAttachment, err error) {
	var (
		c              = contexts.Get(ctx)
		user           = c.User
		memberId int64 = 0
	)
	if user != nil {
		memberId = user.Id
	}

	models := &entity.SysAttachment{
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
		return nil, gerror.Wrap(err, consts.ErrorORM)
	}

	models.Id = id
	return models, nil
}
