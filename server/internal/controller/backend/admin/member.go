// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/member"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
)

var (
	Member = cMember{}
)

type cMember struct{}

// UpdateProfile 修改登录密码
func (c *cMember) UpdateProfile(ctx context.Context, req *member.UpdateProfileReq) (res *member.UpdateProfileRes, err error) {

	var in adminin.MemberUpdateProfileInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	if err = service.AdminMember().UpdateProfile(ctx, in); err != nil {
		return nil, err
	}

	return
}

// UpdatePwd 修改登录密码
func (c *cMember) UpdatePwd(ctx context.Context, req *member.UpdatePwdReq) (res *member.UpdatePwdRes, err error) {

	memberId := contexts.Get(ctx).User.Id
	if memberId <= 0 {
		err := gerror.New("获取用户信息失败！")
		return nil, err
	}

	if err = service.AdminMember().
		UpdatePwd(ctx, adminin.MemberUpdatePwdInp{Id: memberId, OldPassword: req.OldPassword, NewPassword: req.NewPassword}); err != nil {
		return nil, err
	}

	return
}

// Profile 获取登录用户的基本信息
func (c *cMember) Profile(ctx context.Context, req *member.ProfileReq) (*member.ProfileRes, error) {

	var res member.ProfileRes

	memberId := contexts.Get(ctx).User.Id
	if memberId <= 0 {
		err := gerror.New("获取用户信息失败！")
		return nil, err
	}

	// 用户基本信息
	memberInfo, err := service.AdminMember().View(ctx, adminin.MemberViewInp{Id: memberId})
	if err != nil {
		return nil, err
	}
	res.User = memberInfo

	// 所在部门
	sysDept, err := service.AdminDept().View(ctx, adminin.DeptViewInp{Id: memberInfo.DeptId})
	if err != nil {
		return nil, err
	}
	res.SysDept = sysDept

	// 角色列表
	sysRoles, err := service.AdminRole().GetMemberList(ctx, memberInfo.RoleId)
	if err != nil {
		return nil, err
	}
	res.SysRoles = sysRoles

	// 获取角色名称
	roleGroup, err := service.AdminRole().GetName(ctx, memberInfo.RoleId)
	if err != nil {
		return nil, err
	}
	res.RoleGroup = roleGroup

	// 获取第一岗位名称
	postGroup, err := service.AdminPost().GetMemberByStartName(ctx, memberInfo.Id)
	if err != nil {
		return nil, err
	}
	res.PostGroup = postGroup

	return &res, nil
}

// ResetPwd 重置密码
func (c *cMember) ResetPwd(ctx context.Context, req *member.ResetPwdReq) (res *member.ResetPwdRes, err error) {

	if err = service.AdminMember().
		ResetPwd(ctx, adminin.MemberResetPwdInp{Id: req.Id, Password: req.Password}); err != nil {
		return nil, err
	}

	return
}

// EmailUnique 邮箱是否唯一
func (c *cMember) EmailUnique(ctx context.Context, req *member.EmailUniqueReq) (*member.EmailUniqueRes, error) {

	data, err := service.AdminMember().EmailUnique(ctx, adminin.MemberEmailUniqueInp{Id: req.Id, Email: req.Email})
	if err != nil {
		return nil, err
	}

	var res member.EmailUniqueRes
	res.IsUnique = data.IsUnique
	return &res, nil
}

// MobileUnique 手机号是否唯一
func (c *cMember) MobileUnique(ctx context.Context, req *member.MobileUniqueReq) (*member.MobileUniqueRes, error) {

	data, err := service.AdminMember().MobileUnique(ctx, adminin.MemberMobileUniqueInp{Id: req.Id, Mobile: req.Mobile})
	if err != nil {
		return nil, err
	}

	var res member.MobileUniqueRes
	res.IsUnique = data.IsUnique
	return &res, nil
}

// NameUnique 名称是否唯一
func (c *cMember) NameUnique(ctx context.Context, req *member.NameUniqueReq) (*member.NameUniqueRes, error) {

	data, err := service.AdminMember().NameUnique(ctx, adminin.MemberNameUniqueInp{Id: req.Id, Username: req.Username})
	if err != nil {
		return nil, err
	}

	var res member.NameUniqueRes
	res.IsUnique = data.IsUnique
	return &res, nil
}

// Delete 删除
func (c *cMember) Delete(ctx context.Context, req *member.DeleteReq) (res *member.DeleteRes, err error) {

	var in adminin.MemberDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.AdminMember().Delete(ctx, in); err != nil {
		return nil, err
	}
	return res, nil
}

// Edit 修改/新增
func (c *cMember) Edit(ctx context.Context, req *member.EditReq) (res *member.EditRes, err error) {

	var in adminin.MemberEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	in.PostIds = req.PostIds
	if err = service.AdminMember().Edit(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}

// MaxSort 最大排序
func (c *cMember) MaxSort(ctx context.Context, req *member.MaxSortReq) (*member.MaxSortRes, error) {

	data, err := service.AdminMember().MaxSort(ctx, adminin.MemberMaxSortInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var res member.MaxSortRes
	res.Sort = data.Sort
	return &res, nil
}

// View 获取指定信息
func (c *cMember) View(ctx context.Context, req *member.ViewReq) (*member.ViewRes, error) {

	postsList, _, err := service.AdminPost().List(ctx, adminin.PostListInp{})
	if err != nil {
		return nil, err
	}

	roleList, _, err := service.AdminRole().List(ctx, adminin.RoleListInp{})
	if err != nil {
		return nil, err
	}

	var res member.ViewRes
	res.Posts = postsList
	res.Roles = roleList

	if req.Id <= 0 {
		return &res, err
	}

	memberInfo, err := service.AdminMember().View(ctx, adminin.MemberViewInp{Id: req.Id})
	if err != nil {
		return nil, err
	}

	res.MemberViewModel = memberInfo

	res.PostIds, err = service.AdminMemberPost().GetMemberByIds(ctx, memberInfo.Id)
	if err != nil {
		return nil, err
	}

	res.RoleIds = []int64{memberInfo.RoleId}
	res.DeptName, err = service.AdminDept().GetName(ctx, memberInfo.DeptId)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// List 查看列表
func (c *cMember) List(ctx context.Context, req *member.ListReq) (*member.ListRes, error) {

	var (
		in  adminin.MemberListInp
		res member.ListRes
	)

	if err := gconv.Scan(req, &in); err != nil {
		return nil, err
	}

	list, totalCount, err := service.AdminMember().List(ctx, in)
	if err != nil {
		return nil, err
	}

	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage

	return &res, nil
}

// Info 登录用户信息
func (c *cMember) Info(ctx context.Context, req *member.InfoReq) (res *member.InfoRes, err error) {

	model, err := service.AdminMember().LoginMemberInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	if err = gconv.Scan(model, &res); err != nil {
		return nil, err
	}
	return
}

// Status 更新状态
func (c *cMember) Status(ctx context.Context, req *member.StatusReq) (res *member.StatusRes, err error) {

	var in adminin.MemberStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return nil, err
	}
	if err = service.AdminMember().Status(ctx, in); err != nil {
		return nil, err
	}

	return res, nil
}
