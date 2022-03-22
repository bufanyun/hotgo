package adminController

import (
	"context"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/sysService"
	"github.com/gogf/gf/v2/util/gconv"
)

// 配置
var Config = config{}

type config struct{}

//
//  @Title  名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *config) GetValue(ctx context.Context, req *adminForm.ConfigGetValueReq) (*adminForm.ConfigGetValueRes, error) {

	data, err := sysService.Config.GetValue(ctx, input.SysConfigGetValueInp{Key: req.Key})
	if err != nil {
		return nil, err
	}

	var res adminForm.ConfigGetValueRes
	res.Value = data.Value
	return &res, nil
}

//
//  @Title  名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *config) NameUnique(ctx context.Context, req *adminForm.ConfigNameUniqueReq) (*adminForm.ConfigNameUniqueRes, error) {

	data, err := sysService.Config.NameUnique(ctx, input.SysConfigNameUniqueInp{Id: req.Id, Name: req.Name})
	if err != nil {
		return nil, err
	}

	var res adminForm.ConfigNameUniqueRes
	res.IsUnique = data.IsUnique
	return &res, nil
}

//
//  @Title  删除
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *config) Delete(ctx context.Context, req *adminForm.ConfigDeleteReq) (res *adminForm.ConfigDeleteRes, err error) {
	var in input.SysConfigDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = sysService.Config.Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  修改/新增
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *config) Edit(ctx context.Context, req *adminForm.ConfigEditReq) (res *adminForm.ConfigEditRes, err error) {

	var in input.SysConfigEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = sysService.Config.Edit(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

//
//  @Title  最大排序
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *config) MaxSort(ctx context.Context, req *adminForm.ConfigMaxSortReq) (*adminForm.ConfigMaxSortRes, error) {

	data, err := sysService.Config.MaxSort(ctx, input.SysConfigMaxSortInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res adminForm.ConfigMaxSortRes
	res.Sort = data.Sort
	return &res, nil
}

//
//  @Title  获取指定信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *config) View(ctx context.Context, req *adminForm.ConfigViewReq) (*adminForm.ConfigViewRes, error) {

	data, err := sysService.Config.View(ctx, input.SysConfigViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res adminForm.ConfigViewRes
	res.SysConfigViewModel = data
	return &res, nil
}

//
//  @Title  查看列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *config) List(ctx context.Context, req *adminForm.ConfigListReq) (*adminForm.ConfigListRes, error) {

	var (
		in  input.SysConfigListInp
		res adminForm.ConfigListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := sysService.Config.List(ctx, in)
	if err != nil {
		return nil, err
	}

	res.List = list
	res.TotalCount = totalCount
	res.Limit = req.Page
	res.Limit = req.Limit

	return &res, nil
}
