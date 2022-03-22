package apiController

import (
	"context"
	"github.com/bufanyun/hotgo/app/com"
	"github.com/bufanyun/hotgo/app/form/apiForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/service/adminService"
	"github.com/gogf/gf/v2/errors/gerror"
)

// 会员
var Member = member{}

type member struct{}

//
//  @Title  获取登录用户的基本信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *member) Profile(ctx context.Context, req *apiForm.MemberProfileReq) (*apiForm.MemberProfileRes, error) {

	var res apiForm.MemberProfileRes

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
