package apiController

import (
	"context"
	"github.com/bufanyun/hotgo/app/form/apiForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/sysService"
	"github.com/gogf/gf/v2/errors/gerror"
)

// 日志
var Log = log{}

type log struct{}

//
//  @Title  清空日志
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *log) Clear(ctx context.Context, req *apiForm.LogClearReq) (res *apiForm.LogClearRes, err error) {
	err = gerror.New("考虑安全，请到数据库清空")
	return
}

//
//  @Title  导出
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *log) Export(ctx context.Context, req *apiForm.LogExportReq) (res *apiForm.LogExportRes, err error) {

	err = sysService.Log.Export(ctx, input.LogListInp{
		Page:       req.Page,
		Limit:      req.Limit,
		Module:     req.Module,
		Method:     req.Method,
		Url:        req.Url,
		Ip:         req.Ip,
		ErrorCode:  req.ErrorCode,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		MemberId:   req.MemberId,
		TakeUpTime: req.TakeUpTime,
	})
	if err != nil {
		return nil, err
	}

	return
}

//
//  @Title  获取全局日志列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *log) List(ctx context.Context, req *apiForm.LogListReq) (*apiForm.LogListRes, error) {

	list, totalCount, err := sysService.Log.List(ctx, input.LogListInp{
		Page:       req.Page,
		Limit:      req.Limit,
		Module:     req.Module,
		Method:     req.Method,
		Url:        req.Url,
		Ip:         req.Ip,
		ErrorCode:  req.ErrorCode,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		MemberId:   req.MemberId,
		TakeUpTime: req.TakeUpTime,
	})
	if err != nil {
		return nil, err
	}

	var res apiForm.LogListRes
	res.List = list
	res.TotalCount = totalCount
	res.Limit = req.Page
	res.Limit = req.Limit

	return &res, nil
}
