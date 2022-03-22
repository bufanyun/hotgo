//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminController

import (
	"context"
	"github.com/bufanyun/hotgo/app/com"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/adminService"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

// 会员
var Member = member{}

type member struct{}

//
//  @Title  修改登录密码
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *member) UpdateProfile(ctx context.Context, req *adminForm.MemberUpdateProfileReq) (res *adminForm.MemberUpdateProfileRes, err error) {

	var in input.AdminMemberUpdateProfileInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = adminService.Member.UpdateProfile(ctx, in); err != nil {
		return nil, err
	}

	return
}

//
//  @Title  修改登录密码
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *member) UpdatePwd(ctx context.Context, req *adminForm.MemberUpdatePwdReq) (res *adminForm.MemberUpdatePwdRes, err error) {

	memberId := com.Context.Get(ctx).User.Id
	if memberId <= 0 {
		err := gerror.New("获取用户信息失败！")
		return nil, err
	}

	if err = adminService.Member.
		UpdatePwd(ctx, input.AdminMemberUpdatePwdInp{Id: memberId, OldPassword: req.OldPassword, NewPassword: req.NewPassword}); err != nil {
		return nil, err
	}

	return
}

//
//  @Title  获取登录用户的基本信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *member) Profile(ctx context.Context, req *adminForm.MemberProfileReq) (*adminForm.MemberProfileRes, error) {

	var res adminForm.MemberProfileRes

	memberId := com.Context.Get(ctx).User.Id
	if memberId <= 0 {
		err := gerror.New("获取用户信息失败！")
		return nil, err
	}

	// TODO  用户基本信息
	memberInfo, err := adminService.Member.View(ctx, input.AdminMemberViewInp{Id: memberId})
	if err != nil {
		return nil, err
	}
	res.User = memberInfo

	// TODO  所在部门
	sysDept, err := adminService.Dept.View(ctx, input.AdminDeptViewInp{Id: memberInfo.DeptId})
	if err != nil {
		return nil, err
	}
	res.SysDept = sysDept

	// TODO  角色列表
	sysRoles, err := adminService.Role.GetMemberList(ctx, memberInfo.Role)
	if err != nil {
		return nil, err
	}
	res.SysRoles = sysRoles

	// TODO  获取角色名称
	roleGroup, err := adminService.Role.GetName(ctx, memberInfo.Role)
	if err != nil {
		return nil, err
	}
	res.RoleGroup = roleGroup

	// TODO  获取第一岗位名称
	postGroup, err := adminService.Post.GetMemberByStartName(ctx, memberInfo.Id)
	if err != nil {
		return nil, err
	}
	res.PostGroup = postGroup

	return &res, nil
}

//
//  @Title  重置密码
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *member) ResetPwd(ctx context.Context, req *adminForm.MemberResetPwdReq) (res *adminForm.MemberResetPwdRes, err error) {

	if err = adminService.Member.
		ResetPwd(ctx, input.AdminMemberResetPwdInp{Id: req.Id, Password: req.Password}); err != nil {
		return nil, err
	}

	return
}

//
//  @Title  邮箱是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *member) EmailUnique(ctx context.Context, req *adminForm.MemberEmailUniqueReq) (*adminForm.MemberEmailUniqueRes, error) {

	data, err := adminService.Member.EmailUnique(ctx, input.AdminMemberEmailUniqueInp{Id: req.Id, Email: req.Email})
	if err != nil {
		return nil, err
	}

	var res adminForm.MemberEmailUniqueRes
	res.IsUnique = data.IsUnique
	return &res, nil
}

//
//  @Title  手机号是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *member) MobileUnique(ctx context.Context, req *adminForm.MemberMobileUniqueReq) (*adminForm.MemberMobileUniqueRes, error) {

	data, err := adminService.Member.MobileUnique(ctx, input.AdminMemberMobileUniqueInp{Id: req.Id, Mobile: req.Mobile})
	if err != nil {
		return nil, err
	}

	var res adminForm.MemberMobileUniqueRes
	res.IsUnique = data.IsUnique
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
func (controller *member) NameUnique(ctx context.Context, req *adminForm.MemberNameUniqueReq) (*adminForm.MemberNameUniqueRes, error) {

	data, err := adminService.Member.NameUnique(ctx, input.AdminMemberNameUniqueInp{Id: req.Id, Username: req.Username})
	if err != nil {
		return nil, err
	}

	var res adminForm.MemberNameUniqueRes
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
func (controller *member) Delete(ctx context.Context, req *adminForm.MemberDeleteReq) (res *adminForm.MemberDeleteRes, err error) {

	err = gerror.New("考虑安全暂时不允许删除用户，请选择禁用！")
	return nil, err

	var in input.AdminMemberDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = adminService.Member.Delete(ctx, in); err != nil {
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
func (controller *member) Edit(ctx context.Context, req *adminForm.MemberEditReq) (res *adminForm.MemberEditRes, err error) {

	var in input.AdminMemberEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = adminService.Member.Edit(ctx, in); err != nil {
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
func (controller *member) MaxSort(ctx context.Context, req *adminForm.MemberMaxSortReq) (*adminForm.MemberMaxSortRes, error) {

	data, err := adminService.Member.MaxSort(ctx, input.AdminMemberMaxSortInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res adminForm.MemberMaxSortRes
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
func (controller *member) View(ctx context.Context, req *adminForm.MemberViewReq) (*adminForm.MemberViewRes, error) {

	postsList, _, err := adminService.Post.List(ctx, input.AdminPostListInp{})
	if err != nil {
		return nil, err
	}

	roleList, _, err := adminService.Role.List(ctx, input.AdminRoleListInp{})
	if err != nil {
		return nil, err
	}

	var res adminForm.MemberViewRes
	res.Posts = postsList
	res.Roles = roleList

	if req.Id <= 0 {
		return &res, err
	}

	memberInfo, err := adminService.Member.View(ctx, input.AdminMemberViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	res.AdminMemberViewModel = memberInfo

	res.PostIds, err = adminService.MemberPost.GetMemberByIds(ctx, memberInfo.Id)
	if err != nil {
		return nil, err
	}

	res.RoleIds = []int64{memberInfo.Role}
	res.DeptName, err = adminService.Dept.GetName(ctx, memberInfo.DeptId)
	if err != nil {
		return nil, err
	}
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
func (controller *member) List(ctx context.Context, req *adminForm.MemberListReq) (*adminForm.MemberListRes, error) {

	var (
		in  input.AdminMemberListInp
		res adminForm.MemberListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := adminService.Member.List(ctx, in)
	if err != nil {
		return nil, err
	}

	res.List = list
	res.TotalCount = totalCount
	res.Limit = req.Page
	res.Limit = req.Limit

	return &res, nil
}

//
//  @Title  登录用户信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *member) Info(ctx context.Context, req *adminForm.MemberInfoReq) (res *adminForm.MemberInfoRes, err error) {

	return adminService.Member.LoginMemberInfo(ctx, req)
}
