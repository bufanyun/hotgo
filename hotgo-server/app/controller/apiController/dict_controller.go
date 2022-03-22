package apiController

import (
	"context"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/service/sysService"
)

// 字典
var Dict = dict{}

type dict struct{}

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
