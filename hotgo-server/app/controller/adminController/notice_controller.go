package adminController

import (
	"context"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/adminService"
	"github.com/gogf/gf/v2/util/gconv"
)

// 公告
var Notice = notice{}

type notice struct{}

//
//  @Title  名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *notice) NameUnique(ctx context.Context, req *adminForm.NoticeNameUniqueReq) (*adminForm.NoticeNameUniqueRes, error) {

	data, err := adminService.Notice.NameUnique(ctx, input.AdminNoticeNameUniqueInp{Id: req.Id, Title: req.Title})
	if err != nil {
		return nil, err
	}

	var res adminForm.NoticeNameUniqueRes
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
func (controller *notice) Delete(ctx context.Context, req *adminForm.NoticeDeleteReq) (res *adminForm.NoticeDeleteRes, err error) {
	var in input.AdminNoticeDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = adminService.Notice.Delete(ctx, in); err != nil {
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
func (controller *notice) Edit(ctx context.Context, req *adminForm.NoticeEditReq) (res *adminForm.NoticeEditRes, err error) {

	var in input.AdminNoticeEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = adminService.Notice.Edit(ctx, in); err != nil {
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
func (controller *notice) MaxSort(ctx context.Context, req *adminForm.NoticeMaxSortReq) (*adminForm.NoticeMaxSortRes, error) {

	data, err := adminService.Notice.MaxSort(ctx, input.AdminNoticeMaxSortInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res adminForm.NoticeMaxSortRes
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
func (controller *notice) View(ctx context.Context, req *adminForm.NoticeViewReq) (*adminForm.NoticeViewRes, error) {

	data, err := adminService.Notice.View(ctx, input.AdminNoticeViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res adminForm.NoticeViewRes
	res.AdminNoticeViewModel = data
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
func (controller *notice) List(ctx context.Context, req *adminForm.NoticeListReq) (*adminForm.NoticeListRes, error) {

	var (
		in  input.AdminNoticeListInp
		res adminForm.NoticeListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := adminService.Notice.List(ctx, in)
	if err != nil {
		return nil, err
	}

	res.List = list
	res.TotalCount = totalCount
	res.Limit = req.Page
	res.Limit = req.Limit

	return &res, nil
}
