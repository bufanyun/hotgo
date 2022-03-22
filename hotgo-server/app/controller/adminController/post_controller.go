package adminController

import (
	"context"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/adminService"
	"github.com/gogf/gf/v2/util/gconv"
)

// 岗位
var Post = post{}

type post struct{}

//
//  @Title  删除
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *post) Delete(ctx context.Context, req *adminForm.PostDeleteReq) (res *adminForm.PostDeleteRes, err error) {
	var in input.AdminPostDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = adminService.Post.Delete(ctx, in); err != nil {
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
func (controller *post) Edit(ctx context.Context, req *adminForm.PostEditReq) (res *adminForm.PostEditRes, err error) {

	var in input.AdminPostEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = adminService.Post.Edit(ctx, in); err != nil {
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
func (controller *post) MaxSort(ctx context.Context, req *adminForm.PostMaxSortReq) (*adminForm.PostMaxSortRes, error) {

	data, err := adminService.Post.MaxSort(ctx, input.AdminPostMaxSortInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res adminForm.PostMaxSortRes
	res.Sort = data.Sort
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
func (controller *post) NameUnique(ctx context.Context, req *adminForm.PostNameUniqueReq) (*adminForm.PostNameUniqueRes, error) {

	data, err := adminService.Post.NameUnique(ctx, input.AdminPostNameUniqueInp{Id: req.Id, Name: req.Name})
	if err != nil {
		return nil, err
	}

	var res adminForm.PostNameUniqueRes
	res.IsUnique = data.IsUnique
	return &res, nil
}

//
//  @Title  编码是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *post) CodeUnique(ctx context.Context, req *adminForm.PostCodeUniqueReq) (*adminForm.PostCodeUniqueRes, error) {

	data, err := adminService.Post.CodeUnique(ctx, input.AdminPostCodeUniqueInp{Id: req.Id, Code: req.Code})
	if err != nil {
		return nil, err
	}

	var res adminForm.PostCodeUniqueRes
	res.IsUnique = data.IsUnique
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
func (controller *post) View(ctx context.Context, req *adminForm.PostViewReq) (*adminForm.PostViewRes, error) {

	data, err := adminService.Post.View(ctx, input.AdminPostViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res adminForm.PostViewRes
	res.AdminPostViewModel = data
	return &res, nil
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
func (controller *post) List(ctx context.Context, req *adminForm.PostListReq) (*adminForm.PostListRes, error) {

	list, totalCount, err := adminService.Post.List(ctx, input.AdminPostListInp{
		Page:   req.Page,
		Limit:  req.Limit,
		Name:   req.Name,
		Code:   req.Code,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	var res adminForm.PostListRes
	res.List = list
	res.TotalCount = totalCount
	res.Limit = req.Page
	res.Limit = req.Limit

	return &res, nil
}
