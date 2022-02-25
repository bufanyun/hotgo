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
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/bufanyun/hotgo/app/service/internal/dao"
	"github.com/bufanyun/hotgo/app/service/internal/dto"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var Notice = new(notice)

type notice struct{}

//
//  @Title  菜单名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *notice) NameUnique(ctx context.Context, in input.AdminNoticeNameUniqueInp) (*input.AdminNoticeNameUniqueModel, error) {

	var res input.AdminNoticeNameUniqueModel
	isUnique, err := dao.AdminNotice.IsUniqueTitle(ctx, in.Id, in.Title)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.IsUnique = isUnique
	return &res, nil
}

//
//  @Title  删除
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  error
//
func (service *notice) Delete(ctx context.Context, in input.AdminNoticeDeleteInp) error {

	_, err := dao.AdminNotice.Ctx(ctx).Where("id", in.Id).Delete()
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
func (service *notice) Edit(ctx context.Context, in input.AdminNoticeEditInp) (err error) {

	if in.Title == "" {
		err = gerror.New("名称不能为空")
		return err
	}

	uniqueName, err := dao.AdminNotice.IsUniqueTitle(ctx, in.Id, in.Title)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !uniqueName {
		err = gerror.New("名称已存在")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.AdminNotice.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()
	_, err = dao.AdminNotice.Ctx(ctx).Data(in).Insert()
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
func (service *notice) MaxSort(ctx context.Context, in input.AdminNoticeMaxSortInp) (*input.AdminNoticeMaxSortModel, error) {
	var res input.AdminNoticeMaxSortModel

	if in.Id > 0 {
		if err := dao.AdminNotice.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	res.Sort = res.Sort + 10

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
func (service *notice) View(ctx context.Context, in input.AdminNoticeViewInp) (res *input.AdminNoticeViewModel, err error) {

	if err = dao.AdminNotice.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
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
func (service *notice) List(ctx context.Context, in input.AdminNoticeListInp) (list []*input.AdminNoticeListModel, totalCount int, err error) {

	mod := dao.AdminNotice.Ctx(ctx)

	if in.Realname != "" {
		mod = mod.WhereLike("realname", "%"+in.Realname+"%")
	}
	if in.Username != "" {
		mod = mod.WhereLike("username", "%"+in.Username+"%")
	}
	if in.Mobile > 0 {
		mod = mod.Where("mobile", in.Mobile)
	}
	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}
	if in.DeptId > 0 {
		mod = mod.Where("dept_id", in.DeptId)
	}

	// 日期范围
	if in.StartTime != "" {
		mod = mod.WhereGTE("created_at", in.StartTime)
	}
	if in.EndTime != "" {
		mod = mod.WhereLTE("created_at", in.EndTime)
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

	//// TODO  重写树入参
	//for i := 0; i < len(list); i++ {
	//}

	return list, totalCount, err
}

//
//  @Title  根据条件查询所有数据
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   where
//  @Return  []*entity.AdminNotice
//  @Return  error
//
func (service *notice) WhereAll(ctx context.Context, where dto.AdminNotice) ([]*entity.AdminNotice, error) {
	var (
		model  []*entity.AdminNotice
		err    error
		result gdb.Result
	)
	result, err = dao.AdminNotice.Ctx(ctx).Where(where).All()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	err = gconv.Scan(result, &model)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorRotaPointer)
		return nil, err
	}

	return model, nil
}

//
//  @Title  根据条件查询一行的数据
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   where
//  @Return  *entity.AdminMenu
//
func (service *notice) WhereScan(ctx context.Context, where dto.AdminNotice) *entity.AdminNotice {
	var (
		model *entity.AdminNotice
		err   error
	)

	if err = dao.AdminMenu.Ctx(ctx).Where(where).Scan(&model); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil
	}

	return model
}
