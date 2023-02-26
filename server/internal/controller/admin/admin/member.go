// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/member"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var (
	Member = cMember{}
)

type cMember struct{}

// UpdateCash 修改代理商提现信息
func (c *cMember) UpdateCash(ctx context.Context, req *member.UpdateCashReq) (res *member.UpdateCashRes, err error) {
	var in adminin.MemberUpdateCashInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminMember().UpdateCash(ctx, in)
	return
}

// UpdateEmail 换绑邮箱
func (c *cMember) UpdateEmail(ctx context.Context, req *member.UpdateEmailReq) (res *member.UpdateEmailRes, err error) {
	var in adminin.MemberUpdateEmailInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminMember().UpdateEmail(ctx, in)
	return
}

// UpdateMobile 换绑手机号
func (c *cMember) UpdateMobile(ctx context.Context, req *member.UpdateMobileReq) (res *member.UpdateMobileRes, err error) {
	var in adminin.MemberUpdateMobileInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminMember().UpdateMobile(ctx, in)
	return
}

// UpdateProfile 更新用户资料
func (c *cMember) UpdateProfile(ctx context.Context, req *member.UpdateProfileReq) (res *member.UpdateProfileRes, err error) {
	var in adminin.MemberUpdateProfileInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminMember().UpdateProfile(ctx, in)
	return
}

// UpdatePwd 修改登录密码
func (c *cMember) UpdatePwd(ctx context.Context, req *member.UpdatePwdReq) (res *member.UpdatePwdRes, err error) {
	var memberId = contexts.Get(ctx).User.Id
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return nil, err
	}

	var in = adminin.MemberUpdatePwdInp{
		Id:          memberId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}

	err = service.AdminMember().UpdatePwd(ctx, in)
	return
}

// ResetPwd 重置密码
func (c *cMember) ResetPwd(ctx context.Context, req *member.ResetPwdReq) (res *member.ResetPwdRes, err error) {
	var in = adminin.MemberResetPwdInp{
		Id:       req.Id,
		Password: req.Password,
	}

	err = service.AdminMember().ResetPwd(ctx, in)
	return
}

// EmailUnique 邮箱是否唯一
func (c *cMember) EmailUnique(ctx context.Context, req *member.EmailUniqueReq) (res *member.EmailUniqueRes, err error) {
	var in = adminin.MemberEmailUniqueInp{
		Id:    req.Id,
		Email: req.Email,
	}

	data, err := service.AdminMember().EmailUnique(ctx, in)
	if err != nil {
		return
	}

	res = new(member.EmailUniqueRes)
	res.IsUnique = data.IsUnique
	return
}

// MobileUnique 手机号是否唯一
func (c *cMember) MobileUnique(ctx context.Context, req *member.MobileUniqueReq) (res *member.MobileUniqueRes, err error) {
	var in = adminin.MemberMobileUniqueInp{
		Id:     req.Id,
		Mobile: req.Mobile,
	}

	data, err := service.AdminMember().MobileUnique(ctx, in)
	if err != nil {
		return
	}

	res = new(member.MobileUniqueRes)
	res.IsUnique = data.IsUnique
	return
}

// NameUnique 名称是否唯一
func (c *cMember) NameUnique(ctx context.Context, req *member.NameUniqueReq) (res *member.NameUniqueRes, err error) {
	var in = adminin.MemberNameUniqueInp{
		Id:       req.Id,
		Username: req.Username,
	}

	data, err := service.AdminMember().NameUnique(ctx, in)
	if err != nil {
		return
	}

	res = new(member.NameUniqueRes)
	res.IsUnique = data.IsUnique
	return
}

// Delete 删除
func (c *cMember) Delete(ctx context.Context, req *member.DeleteReq) (res *member.DeleteRes, err error) {
	var in adminin.MemberDeleteInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminMember().Delete(ctx, in)
	return
}

// Edit 修改/新增
func (c *cMember) Edit(ctx context.Context, req *member.EditReq) (res *member.EditRes, err error) {
	var in adminin.MemberEditInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	if err = validate.PreFilter(ctx, &in); err != nil {
		return
	}

	in.PostIds = req.PostIds
	err = service.AdminMember().Edit(ctx, in)
	return
}

// MaxSort 最大排序
func (c *cMember) MaxSort(ctx context.Context, req *member.MaxSortReq) (res *member.MaxSortRes, err error) {
	var in = adminin.MemberMaxSortInp{Id: req.Id}
	data, err := service.AdminMember().MaxSort(ctx, in)
	if err != nil {
		return
	}

	res = new(member.MaxSortRes)
	res.Sort = data.Sort
	return
}

// View 获取指定信息
func (c *cMember) View(ctx context.Context, req *member.ViewReq) (res *member.ViewRes, err error) {
	data, err := service.AdminMember().View(ctx, adminin.MemberViewInp{Id: req.Id})
	if err != nil {
		return
	}

	res = new(member.ViewRes)
	res.MemberViewModel = data
	return
}

// List 查看列表
func (c *cMember) List(ctx context.Context, req *member.ListReq) (res *member.ListRes, err error) {
	var in adminin.MemberListInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	list, totalCount, err := service.AdminMember().List(ctx, in)
	if err != nil {
		return
	}

	res = new(member.ListRes)
	res.List = list
	res.PageCount = form.CalPageCount(totalCount, req.PerPage)
	res.Page = req.Page
	res.PerPage = req.PerPage
	return
}

// LoginInfo 登录用户信息
func (c *cMember) LoginInfo(ctx context.Context, req *member.LoginInfoReq) (res *member.LoginInfoRes, err error) {
	data, err := service.AdminMember().LoginMemberInfo(ctx)
	if err != nil {
		return
	}

	res = new(member.LoginInfoRes)
	res.LoginMemberInfoModel = data
	return
}

// Status 更新状态
func (c *cMember) Status(ctx context.Context, req *member.StatusReq) (res *member.StatusRes, err error) {
	var in adminin.MemberStatusInp
	if err = gconv.Scan(req, &in); err != nil {
		return
	}

	err = service.AdminMember().Status(ctx, in)
	return
}

// Select 获取可选的后台用户选项
func (c *cMember) Select(ctx context.Context, req *member.SelectReq) (res *member.SelectRes, err error) {
	data, err := service.AdminMember().Select(ctx, adminin.MemberSelectInp{})
	if err != nil {
		return
	}

	res = (*member.SelectRes)(&data)
	return
}
