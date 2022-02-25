//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminService

import (
	"context"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/internal/dao"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

var Post = new(post)

type post struct{}

//
//  @Title  删除
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  error
//
func (service *post) Delete(ctx context.Context, in input.AdminPostDeleteInp) error {

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

//
//  @Title  修改/新增
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  error
//
func (service *post) Edit(ctx context.Context, in input.AdminPostEditInp) (err error) {

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

//
//  @Title  最大排序
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictDataMaxSortRes
//  @Return  error
//
func (service *post) MaxSort(ctx context.Context, in input.AdminPostMaxSortInp) (*input.AdminPostMaxSortModel, error) {
	var res input.AdminPostMaxSortModel

	if in.Id > 0 {
		if err := dao.AdminMenu.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	res.Sort = res.Sort + 10

	return &res, nil
}

//
//  @Title  菜单名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *post) NameUnique(ctx context.Context, in input.AdminPostNameUniqueInp) (*input.AdminPostNameUniqueModel, error) {

	var res input.AdminPostNameUniqueModel
	isUnique, err := dao.AdminPost.IsUniqueName(ctx, in.Id, in.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.IsUnique = isUnique
	return &res, nil
}

//
//  @Title  编码是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *post) CodeUnique(ctx context.Context, in input.AdminPostCodeUniqueInp) (*input.AdminPostCodeUniqueModel, error) {

	var res input.AdminPostCodeUniqueModel
	isUnique, err := dao.AdminPost.IsUniqueCode(ctx, in.Id, in.Code)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.IsUnique = isUnique

	return &res, nil
}

//
//  @Title  获取指定字典类型信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeViewRes
//  @Return  error
//
func (service *post) View(ctx context.Context, in input.AdminPostViewInp) (res *input.AdminPostViewModel, err error) {

	if err = dao.AdminPost.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

//
//  @Title  获取列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (service *post) List(ctx context.Context, in input.AdminPostListInp) (list []*input.AdminPostListModel, totalCount int, err error) {

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

	err = mod.Page(in.Page, in.Limit).Order("id desc").Scan(&list)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	return list, totalCount, err
}

//
//  @Title  获取指定用户的第一岗位
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   RoleId
//  @Return  name
//  @Return  err
//
func (service *post) GetMemberByStartName(ctx context.Context, memberId int64) (name string, err error) {

	// TODO  默认取第一岗位
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
