// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
)

type sAdminPost struct{}

func NewAdminPost() *sAdminPost {
	return &sAdminPost{}
}

func init() {
	service.RegisterAdminPost(NewAdminPost())
}

// Delete 删除
func (s *sAdminPost) Delete(ctx context.Context, in adminin.PostDeleteInp) error {
	exist, err := dao.AdminMemberPost.Ctx(ctx).Where("post_id", in.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !exist.IsEmpty() {
		return gerror.New("请先解除该岗位下所有已关联用户关联关系！")
	}
	_, err = dao.AdminPost.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sAdminPost) Edit(ctx context.Context, in adminin.PostEditInp) (err error) {
	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return err
	}
	if in.Code == "" {
		err = gerror.New("编码不能为空")
		return err
	}

	uniqueName, err := dao.AdminPost.IsUniqueName(ctx, in.Id, in.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !uniqueName {
		err = gerror.New("名称已存在")
		return err
	}

	uniqueCode, err := dao.AdminPost.IsUniqueCode(ctx, in.Id, in.Code)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !uniqueCode {
		err = gerror.New("编码已存在")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.AdminPost.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()
	_, err = dao.AdminPost.Ctx(ctx).Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

// MaxSort 最大排序
func (s *sAdminPost) MaxSort(ctx context.Context, in adminin.PostMaxSortInp) (*adminin.PostMaxSortModel, error) {
	var res adminin.PostMaxSortModel

	if in.Id > 0 {
		if err := dao.AdminMenu.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	res.Sort = res.Sort + 10

	return &res, nil
}

// NameUnique 菜单名称是否唯一
func (s *sAdminPost) NameUnique(ctx context.Context, in adminin.PostNameUniqueInp) (*adminin.PostNameUniqueModel, error) {
	var res adminin.PostNameUniqueModel
	isUnique, err := dao.AdminPost.IsUniqueName(ctx, in.Id, in.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.IsUnique = isUnique
	return &res, nil
}

// CodeUnique 编码是否唯一
func (s *sAdminPost) CodeUnique(ctx context.Context, in adminin.PostCodeUniqueInp) (*adminin.PostCodeUniqueModel, error) {
	var res adminin.PostCodeUniqueModel
	isUnique, err := dao.AdminPost.IsUniqueCode(ctx, in.Id, in.Code)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.IsUnique = isUnique

	return &res, nil
}

// View 获取指定字典类型信息
func (s *sAdminPost) View(ctx context.Context, in adminin.PostViewInp) (res *adminin.PostViewModel, err error) {
	if err = dao.AdminPost.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

// List 获取列表
func (s *sAdminPost) List(ctx context.Context, in adminin.PostListInp) (list []*adminin.PostListModel, totalCount int, err error) {
	mod := dao.AdminPost.Ctx(ctx)

	// 访问路径
	if in.Name != "" {
		mod = mod.WhereLike("name", "%"+in.Name+"%")
	}

	// 模块
	if in.Code != "" {
		mod = mod.Where("code", in.Code)
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

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}
	return list, totalCount, err
}

// GetMemberByStartName 获取指定用户的第一岗位
func (s *sAdminPost) GetMemberByStartName(ctx context.Context, memberId int64) (name string, err error) {
	// 默认取第一岗位
	postId, err := dao.AdminMemberPost.Ctx(ctx).
		Fields("post_id").
		Where("member_id", memberId).
		Limit(1).
		Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return name, err
	}

	val, err := dao.AdminPost.Ctx(ctx).
		Fields("name").
		Where("id", postId.Int()).
		Order("id desc").
		Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return name, err
	}

	return val.String(), nil
}

// Status 更新状态
func (s *sAdminPost) Status(ctx context.Context, in adminin.PostStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return err
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return err
	}

	if !convert.InSliceInt(consts.StatusMap, in.Status) {
		err = gerror.New("状态不正确")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	_, err = dao.AdminPost.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}
