package sysService

import (
	"context"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/internal/dao"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

var Config = new(config)

type config struct{}

//
//  @Title  最大排序
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictDataMaxSortRes
//  @Return  error
//
func (service *config) GetValue(ctx context.Context, in input.SysConfigGetValueInp) (*input.SysConfigGetValueModel, error) {
	var res input.SysConfigGetValueModel

	if err := dao.SysConfig.Ctx(ctx).
		Fields("value").
		Where("key", in.Key).
		Order("id desc").
		Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

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
func (service *config) NameUnique(ctx context.Context, in input.SysConfigNameUniqueInp) (*input.SysConfigNameUniqueModel, error) {

	var res input.SysConfigNameUniqueModel
	isUnique, err := dao.SysConfig.IsUniqueName(ctx, in.Id, in.Name)
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
func (service *config) Delete(ctx context.Context, in input.SysConfigDeleteInp) error {

	exist, err := dao.SysConfig.Ctx(ctx).Where("Member_id", in.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !exist.IsEmpty() {
		return gerror.New("请先解除该部门下所有已关联用户关联关系！")
	}
	_, err = dao.SysConfig.Ctx(ctx).Where("id", in.Id).Delete()
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
func (service *config) Edit(ctx context.Context, in input.SysConfigEditInp) (err error) {

	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return err
	}

	uniqueName, err := dao.SysConfig.IsUniqueName(ctx, in.Id, in.Name)
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
		_, err = dao.SysConfig.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()
	_, err = dao.SysConfig.Ctx(ctx).Data(in).Insert()
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
func (service *config) MaxSort(ctx context.Context, in input.SysConfigMaxSortInp) (*input.SysConfigMaxSortModel, error) {
	var res input.SysConfigMaxSortModel

	if in.Id > 0 {
		if err := dao.SysConfig.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
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
func (service *config) View(ctx context.Context, in input.SysConfigViewInp) (res *input.SysConfigViewModel, err error) {

	if err = dao.SysConfig.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
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
func (service *config) List(ctx context.Context, in input.SysConfigListInp) (list []*input.SysConfigListModel, totalCount int, err error) {

	mod := dao.SysConfig.Ctx(ctx)

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
