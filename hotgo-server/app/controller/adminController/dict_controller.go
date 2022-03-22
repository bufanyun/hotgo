package adminController

import (
	"context"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/service/sysService"
)

// 字典
var Dict = dict{}

type dict struct{}

//
//  @Title  数据键值是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) DataUnique(ctx context.Context, req *adminForm.DictDataUniqueReq) (res *adminForm.DictDataUniqueRes, err error) {

	res, err = sysService.Dict.DataUnique(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  查询字典数据最大排序
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) DataMaxSort(ctx context.Context, req *adminForm.DictDataMaxSortReq) (res *adminForm.DictDataMaxSortRes, err error) {

	res, err = sysService.Dict.DataMaxSort(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  删除字典数据
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) DataDelete(ctx context.Context, req *adminForm.DictDataDeleteReq) (res *adminForm.DictDataDeleteRes, err error) {

	if err = sysService.Dict.DataDelete(ctx, req); err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  修改/新增字典数据
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) DataEdit(ctx context.Context, req *adminForm.DictDataEditReq) (res *adminForm.DictDataEditRes, err error) {

	if err = sysService.Dict.DataEdit(ctx, req); err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  获取指定字典类型信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) DataView(ctx context.Context, req *adminForm.DictDataViewReq) (res *adminForm.DictDataViewRes, err error) {

	res, err = sysService.Dict.DataView(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  获取字典数据列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) DataList(ctx context.Context, req *adminForm.DictDataListReq) (res *adminForm.DictDataListRes, err error) {

	res, err = sysService.Dict.DataList(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  获取指定字典类型的属性数据
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) Attribute(ctx context.Context, req *adminForm.DictAttributeReq) (res *adminForm.DictAttributeRes, err error) {

	res, err = sysService.Dict.Attribute(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  导出字典类型
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) TypeExport(ctx context.Context, req *adminForm.DictTypeExportReq) (res *adminForm.DictTypeExportRes, err error) {
	if err = sysService.Dict.TypeExport(ctx, req); err != nil {
		return nil, err
	}
	return
}

//
//  @Title  刷新字典缓存
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) TypeRefreshCache(ctx context.Context, req *adminForm.DictTypeRefreshCacheReq) (res *adminForm.DictTypeRefreshCacheRes, err error) {
	return nil, nil
}

//
//  @Title  删除字典类型
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) TypeDelete(ctx context.Context, req *adminForm.DictTypeDeleteReq) (res *adminForm.DictTypeDeleteRes, err error) {

	if err = sysService.Dict.TypeDelete(ctx, req); err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  修改/新增字典类型
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) TypeEdit(ctx context.Context, req *adminForm.DictTypeEditReq) (res *adminForm.DictTypeEditRes, err error) {

	if err = sysService.Dict.TypeEdit(ctx, req); err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  类型是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) TypeUnique(ctx context.Context, req *adminForm.DictTypeUniqueReq) (res *adminForm.DictTypeUniqueRes, err error) {

	res, err = sysService.Dict.TypeUnique(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  获取指定字典类型信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) TypeView(ctx context.Context, req *adminForm.DictTypeViewReq) (res *adminForm.DictTypeViewRes, err error) {

	res, err = sysService.Dict.TypeView(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//  @Title  获取字典类型列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dict) TypeList(ctx context.Context, req *adminForm.DictTypeListReq) (res *adminForm.DictTypeListRes, err error) {

	res, err = sysService.Dict.TypeList(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
