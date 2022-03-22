package adminController

import (
	"context"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/adminService"
	"github.com/gogf/gf/v2/util/gconv"
)

// 部门
var Dept = dept{}

type dept struct{}

//
//  @Title  名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dept) NameUnique(ctx context.Context, req *adminForm.DeptNameUniqueReq) (*adminForm.DeptNameUniqueRes, error) {

	data, err := adminService.Dept.NameUnique(ctx, input.AdminDeptNameUniqueInp{Id: req.Id, Name: req.Name})
	if err != nil {
		return nil, err
	}

	var res adminForm.DeptNameUniqueRes
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
func (controller *dept) Delete(ctx context.Context, req *adminForm.DeptDeleteReq) (res *adminForm.DeptDeleteRes, err error) {
	var in input.AdminDeptDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = adminService.Dept.Delete(ctx, in); err != nil {
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
func (controller *dept) Edit(ctx context.Context, req *adminForm.DeptEditReq) (res *adminForm.DeptEditRes, err error) {

	var in input.AdminDeptEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = adminService.Dept.Edit(ctx, in); err != nil {
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
func (controller *dept) MaxSort(ctx context.Context, req *adminForm.DeptMaxSortReq) (*adminForm.DeptMaxSortRes, error) {

	data, err := adminService.Dept.MaxSort(ctx, input.AdminDeptMaxSortInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res adminForm.DeptMaxSortRes
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
func (controller *dept) View(ctx context.Context, req *adminForm.DeptViewReq) (*adminForm.DeptViewRes, error) {

	data, err := adminService.Dept.View(ctx, input.AdminDeptViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res adminForm.DeptViewRes
	res.AdminDeptViewModel = data
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
func (controller *dept) List(ctx context.Context, req *adminForm.DeptListReq) (*adminForm.DeptListRes, error) {

	var (
		in  input.AdminDeptListInp
		res adminForm.DeptListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	data, err := adminService.Dept.List(ctx, in)
	if err != nil {
		return nil, err
	}

	_ = gconv.Structs(data, &res)

	return &res, nil
}

//
//  @Title  查看列表树
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *dept) ListTree(ctx context.Context, req *adminForm.DeptListTreeReq) (*adminForm.DeptListTreeRes, error) {

	var (
		in  input.AdminDeptListTreeInp
		res adminForm.DeptListTreeRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	data, err := adminService.Dept.ListTree(ctx, in)
	if err != nil {
		return nil, err
	}

	_ = gconv.Structs(data, &res)

	return &res, nil
}
