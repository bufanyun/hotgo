// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/tree"
	"hotgo/utility/validate"
)

type sAdminMember struct{}

func NewAdminMember() *sAdminMember {
	return &sAdminMember{}
}

func init() {
	service.RegisterAdminMember(NewAdminMember())
}

// AddBalance 增加余额
func (s *sAdminMember) AddBalance(ctx context.Context, in adminin.MemberAddBalanceInp) (err error) {
	var (
		mb       *entity.AdminMember
		memberId = contexts.GetUserId(ctx)
	)

	if err = s.FilterAuthModel(ctx, memberId).Where("id", in.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if mb == nil {
		err = gerror.New("用户信息不存在")
		return err
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 更新我的余额
		_, err = service.AdminCreditsLog().SaveBalance(ctx, adminin.CreditsLogSaveBalanceInp{
			MemberId:    memberId,
			AppId:       in.AppId,
			AddonsName:  in.AddonsName,
			CreditGroup: in.SelfCreditGroup,
			Num:         in.SelfNum,
			Remark:      fmt.Sprintf("为后台用户:%v 操作%v", mb.Id, in.Remark),
		})
		if err != nil {
			return err
		}

		// 更新对方余额
		_, err = service.AdminCreditsLog().SaveBalance(ctx, adminin.CreditsLogSaveBalanceInp{
			MemberId:    mb.Id,
			AppId:       in.AppId,
			AddonsName:  in.AddonsName,
			CreditGroup: in.OtherCreditGroup,
			Num:         in.OtherNum,
			Remark:      fmt.Sprintf("后台用户:%v 为你操作%v", memberId, in.Remark),
		})
		return
	})

	return
}

// AddIntegral 增加积分
func (s *sAdminMember) AddIntegral(ctx context.Context, in adminin.MemberAddIntegralInp) (err error) {
	var (
		mb       *entity.AdminMember
		memberId = contexts.GetUserId(ctx)
	)

	if err = s.FilterAuthModel(ctx, memberId).Where("id", in.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if mb == nil {
		err = gerror.New("用户信息不存在")
		return err
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 更新我的余额
		_, err = service.AdminCreditsLog().SaveIntegral(ctx, adminin.CreditsLogSaveIntegralInp{
			MemberId:    memberId,
			AppId:       in.AppId,
			AddonsName:  in.AddonsName,
			CreditGroup: in.SelfCreditGroup,
			Num:         in.SelfNum,
			Remark:      fmt.Sprintf("为后台用户:%v 操作%v", mb.Id, in.Remark),
		})
		if err != nil {
			return err
		}

		// 更新对方余额
		_, err = service.AdminCreditsLog().SaveIntegral(ctx, adminin.CreditsLogSaveIntegralInp{
			MemberId:    mb.Id,
			AppId:       in.AppId,
			AddonsName:  in.AddonsName,
			CreditGroup: in.OtherCreditGroup,
			Num:         in.OtherNum,
			Remark:      fmt.Sprintf("后台用户:%v 为你操作%v", memberId, in.Remark),
		})
		return
	})

	return
}

// UpdateCash 修改提现信息
func (s *sAdminMember) UpdateCash(ctx context.Context, in adminin.MemberUpdateCashInp) (err error) {
	memberId := contexts.Get(ctx).User.Id
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	var memberInfo entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("id", memberId).Scan(&memberInfo); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if gmd5.MustEncryptString(in.Password+memberInfo.Salt) != memberInfo.PasswordHash {
		err = gerror.New("登录密码不正确")
		return
	}

	_, err = dao.AdminMember.Ctx(ctx).
		Where("id", memberId).
		Data(g.Map{
			dao.AdminMember.Columns().Cash: adminin.MemberCash{
				Name:      in.Name,
				Account:   in.Account,
				PayeeCode: in.PayeeCode,
			},
		}).
		Update()

	return
}

// UpdateEmail 换绑邮箱
func (s *sAdminMember) UpdateEmail(ctx context.Context, in adminin.MemberUpdateEmailInp) (err error) {
	memberId := contexts.Get(ctx).User.Id
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return err
	}

	var memberInfo *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("id", memberId).Scan(&memberInfo); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if memberInfo == nil {
		err = gerror.New("用户信息不存在")
		return err
	}

	if memberInfo.Email == in.Email {
		err = gerror.New("新旧邮箱不能一样")
		return
	}

	if !validate.IsEmail(in.Email) {
		err = gerror.New("邮箱地址不正确")
		return
	}

	// 存在原绑定号码，需要进行验证
	if memberInfo.Email != "" {
		err = service.SysEmsLog().VerifyCode(ctx, sysin.VerifyEmsCodeInp{
			Event: consts.EmsTemplateBind,
			Email: memberInfo.Email,
			Code:  in.Code,
		})
		if err != nil {
			return err
		}
	}

	update := g.Map{
		dao.AdminMember.Columns().Email: in.Email,
	}

	_, err = dao.AdminMember.Ctx(ctx).Where("id", memberId).Data(update).Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return
}

// UpdateMobile 换绑手机号
func (s *sAdminMember) UpdateMobile(ctx context.Context, in adminin.MemberUpdateMobileInp) (err error) {
	memberId := contexts.Get(ctx).User.Id
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return err
	}

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("id", memberId).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if mb == nil {
		err = gerror.New("用户信息不存在")
		return err
	}

	if mb.Mobile == in.Mobile {
		err = gerror.New("新旧手机号不能一样")
		return
	}

	if !validate.IsMobile(in.Mobile) {
		err = gerror.New("手机号码不正确")
		return
	}

	// 存在原绑定号码，需要进行验证
	if mb.Mobile != "" {
		err = service.SysSmsLog().VerifyCode(ctx, sysin.VerifyCodeInp{
			Event:  consts.SmsTemplateBind,
			Mobile: mb.Mobile,
			Code:   in.Code,
		})
		if err != nil {
			return err
		}
	}

	update := g.Map{
		dao.AdminMember.Columns().Mobile: in.Mobile,
	}

	if _, err = dao.AdminMember.Ctx(ctx).Where("id", memberId).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return
}

// UpdateProfile 更新用户资料
func (s *sAdminMember) UpdateProfile(ctx context.Context, in adminin.MemberUpdateProfileInp) (err error) {
	memberId := contexts.Get(ctx).User.Id
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return err
	}

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("id", memberId).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if mb == nil {
		err = gerror.New("用户信息不存在")
		return err
	}

	update := g.Map{
		dao.AdminMember.Columns().Avatar:   in.Avatar,
		dao.AdminMember.Columns().RealName: in.RealName,
		dao.AdminMember.Columns().Qq:       in.Qq,
		dao.AdminMember.Columns().Birthday: in.Birthday,
		dao.AdminMember.Columns().Sex:      in.Sex,
		dao.AdminMember.Columns().CityId:   in.CityId,
		dao.AdminMember.Columns().Address:  in.Address,
	}

	if _, err = dao.AdminMember.Ctx(ctx).Where("id", memberId).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return
}

// UpdatePwd 修改登录密码
func (s *sAdminMember) UpdatePwd(ctx context.Context, in adminin.MemberUpdatePwdInp) (err error) {
	var mb entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("id", in.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if gmd5.MustEncryptString(in.OldPassword+mb.Salt) != mb.PasswordHash {
		err = gerror.New("原密码不正确")
		return err
	}

	_, err = dao.AdminMember.Ctx(ctx).
		Where("id", in.Id).
		Data(g.Map{
			dao.AdminMember.Columns().PasswordHash: gmd5.MustEncryptString(in.NewPassword + mb.Salt),
		}).
		Update()

	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return
}

// ResetPwd 重置密码
func (s *sAdminMember) ResetPwd(ctx context.Context, in adminin.MemberResetPwdInp) (err error) {
	var (
		mb       *entity.AdminMember
		memberId = contexts.GetUserId(ctx)
	)

	if err = s.FilterAuthModel(ctx, memberId).Where("id", in.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if mb == nil {
		err = gerror.New("用户信息不存在")
		return
	}

	_, err = s.FilterAuthModel(ctx, memberId).
		Where("id", in.Id).
		Data(g.Map{
			dao.AdminMember.Columns().PasswordHash: gmd5.MustEncryptString(in.Password + mb.Salt),
		}).
		Update()
	return
}

// VerifyUnique 验证管理员唯一属性
func (s *sAdminMember) VerifyUnique(ctx context.Context, in adminin.VerifyUniqueInp) (err error) {
	if in.Where == nil {
		return
	}

	msgMap := g.MapStrStr{
		"username":    "用户名已存在，请换一个",
		"email":       "邮箱已存在，请换一个",
		"mobile":      "手机号已存在，请换一个",
		"invite_code": "邀请码已存在，请换一个",
	}

	for k, v := range in.Where {
		if v == "" {
			continue
		}
		message, ok := msgMap[k]
		if !ok {
			err = gerror.Newf("字段 [ %v ] 未配置唯一属性验证", k)
			return
		}
		if err = hgorm.IsUnique(ctx, dao.AdminMember, g.Map{k: v}, message, in.Id); err != nil {
			return
		}
	}

	return
}

// Delete 删除
func (s *sAdminMember) Delete(ctx context.Context, in adminin.MemberDeleteInp) (err error) {
	if s.VerifySuperId(ctx, gconv.Int64(in.Id)) {
		err = gerror.New("超管账号禁止删除！")
		return
	}

	memberId := contexts.GetUserId(ctx)
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	_, err = s.FilterAuthModel(ctx, memberId).Where("id", in.Id).Delete()
	return
}

// Edit 修改/新增
func (s *sAdminMember) Edit(ctx context.Context, in adminin.MemberEditInp) (err error) {
	opMemberId := contexts.GetUserId(ctx)
	if opMemberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	if in.Username == "" {
		err = gerror.New("帐号不能为空")
		return
	}

	err = s.VerifyUnique(ctx, adminin.VerifyUniqueInp{
		Id: in.Id,
		Where: g.Map{
			dao.AdminMember.Columns().Username: in.Username,
			dao.AdminMember.Columns().Mobile:   in.Mobile,
			dao.AdminMember.Columns().Email:    in.Email,
		},
	})
	if err != nil {
		return
	}

	config, err := service.SysConfig().GetLogin(ctx)
	if err != nil {
		return
	}

	// 修改
	if in.Id > 0 {
		if s.VerifySuperId(ctx, in.Id) {
			err = gerror.New("超管账号禁止编辑！")
			return
		}

		mod := s.FilterAuthModel(ctx, opMemberId)

		if in.Password != "" {
			// 修改密码
			salt, err := s.FilterAuthModel(ctx, contexts.GetUserId(ctx)).
				Fields(dao.AdminMember.Columns().Salt).
				Where("id", in.Id).
				Value()
			if err != nil {
				return err
			}
			in.PasswordHash = gmd5.MustEncryptString(in.Password + salt.String())
		} else {
			mod = mod.FieldsEx(dao.AdminMember.Columns().PasswordHash)
		}

		return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			if _, err = mod.Where("id", in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return
			}

			// 更新岗位
			if err = dao.AdminMemberPost.UpdatePostIds(ctx, in.Id, in.PostIds); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
			}
			return
		})
	}

	// 新增用户时的额外属性
	var data adminin.MemberAddInp
	data.Salt = grand.S(6)
	data.InviteCode = grand.S(12)
	data.PasswordHash = gmd5.MustEncryptString(data.Password + data.Salt)

	// 关系树
	data.Pid = opMemberId
	data.Level, data.Tree, err = s.GenTree(ctx, opMemberId)
	if err != nil {
		return
	}

	// 默认头像
	if data.Avatar == "" {
		data.Avatar = config.Avatar
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		id, err := dao.AdminMember.Ctx(ctx).Data(data).InsertAndGetId()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}

		// 更新岗位
		if err = dao.AdminMemberPost.UpdatePostIds(ctx, id, in.PostIds); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
		}
		return
	})
}

// View 获取信息
func (s *sAdminMember) View(ctx context.Context, in adminin.MemberViewInp) (res *adminin.MemberViewModel, err error) {
	err = s.FilterAuthModel(ctx, contexts.GetUserId(ctx)).
		Hook(hook.MemberInfo).
		Where("id", in.Id).
		Scan(&res)
	return
}

// List 获取列表
func (s *sAdminMember) List(ctx context.Context, in adminin.MemberListInp) (list []*adminin.MemberListModel, totalCount int, err error) {
	mod := s.FilterAuthModel(ctx, contexts.GetUserId(ctx))

	if in.RealName != "" {
		mod = mod.WhereLike("real_name", "%"+in.RealName+"%")
	}

	if in.Username != "" {
		mod = mod.WhereLike("username", "%"+in.Username+"%")
	}

	if in.Mobile > 0 {
		mod = mod.Where("mobile", in.Mobile)
	}

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	if in.DeptId > 0 {
		mod = mod.Where("dept_id", in.DeptId)
	}

	if in.RoleId > 0 {
		mod = mod.Where("role_id", in.RoleId)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween("created_at", gtime.New(in.CreatedAt[0]), gtime.New(in.CreatedAt[1]))
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Hook(hook.MemberInfo).Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		return
	}

	for i := 0; i < len(list); i++ {
		posts, err := dao.AdminMemberPost.Ctx(ctx).
			Fields("post_id").
			Where("member_id", list[i].Id).
			Array()

		if err != nil {
			return nil, 0, err
		}

		for _, v := range posts {
			list[i].PostIds = append(list[i].PostIds, v.Int64())
		}
	}

	return
}

// Status 更新状态
func (s *sAdminMember) Status(ctx context.Context, in adminin.MemberStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if s.VerifySuperId(ctx, in.Id) {
		err = gerror.New("超管账号不能更改状态")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSliceInt(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	_, err = s.FilterAuthModel(ctx, contexts.GetUserId(ctx)).Where("id", in.Id).Data("status", in.Status).Update()
	return
}

// GenTree 生成关系树
func (s *sAdminMember) GenTree(ctx context.Context, pid int64) (level int, newTree string, err error) {
	var pmb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("id", pid).Scan(&pmb); err != nil {
		return
	}

	if pmb == nil {
		err = gerror.New("上级信息不存在")
		return
	}

	level = pmb.Level + 1
	newTree = tree.GenLabel(pmb.Tree, pmb.Id)
	return
}

// RoleMemberList 获取角色下的用户列表
func (s *sAdminMember) RoleMemberList(ctx context.Context, in adminin.RoleMemberListInp) (list []*adminin.MemberListModel, totalCount int, err error) {
	mod := dao.AdminMember.Ctx(ctx)
	if in.Role > 0 {
		mod = mod.Where("role", in.Role)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return
}

// LoginMemberInfo 获取登录用户信息
func (s *sAdminMember) LoginMemberInfo(ctx context.Context) (res *adminin.LoginMemberInfoModel, err error) {
	var memberId = contexts.GetUserId(ctx)
	if memberId <= 0 {
		err = gerror.New("用户身份异常，请重新登录！")
		return
	}

	if err = dao.AdminMember.Ctx(ctx).Hook(hook.MemberInfo).Where("id", memberId).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if res == nil {
		err = gerror.New("用户不存在！")
		return
	}

	// 细粒度权限
	permissions, err := service.AdminMenu().LoginPermissions(ctx, memberId)
	if err != nil {
		return
	}
	res.Permissions = permissions

	// 登录统计
	stat, err := s.MemberLoginStat(ctx, adminin.MemberLoginStatInp{MemberId: memberId})
	if err != nil {
		return nil, err
	}

	res.MemberLoginStatModel = stat
	res.Mobile = gstr.HideStr(res.Mobile, 40, `*`)
	res.Email = gstr.HideStr(res.Email, 40, `*`)
	res.OpenId, _ = service.CommonWechat().GetOpenId(ctx)
	return
}

// MemberLoginStat 用户登录统计
func (s *sAdminMember) MemberLoginStat(ctx context.Context, in adminin.MemberLoginStatInp) (res *adminin.MemberLoginStatModel, err error) {
	var models *entity.SysLoginLog
	err = dao.SysLoginLog.Ctx(ctx).
		Fields("login_at,login_ip").
		Where("member_id", in.MemberId).
		Where("status", consts.StatusEnabled).
		Scan(&models)

	if err != nil {
		return
	}

	res = new(adminin.MemberLoginStatModel)
	if models == nil {
		return
	}

	res.LastLoginAt = models.LoginAt
	res.LastLoginIp = models.LoginIp
	res.LoginCount, err = dao.SysLoginLog.Ctx(ctx).
		Where("member_id", in.MemberId).
		Where("status", consts.StatusEnabled).
		Count()
	return
}

// GetIdByCode 通过邀请码获取用户ID
func (s *sAdminMember) GetIdByCode(ctx context.Context, in adminin.GetIdByCodeInp) (res *adminin.GetIdByCodeModel, err error) {
	if err = dao.AdminMember.Ctx(ctx).Fields(adminin.GetIdByCodeModel{}).Where("invite_code", in.Code).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}
	return
}

// Select 获取可选的用户选项
func (s *sAdminMember) Select(ctx context.Context, in adminin.MemberSelectInp) (res []*adminin.MemberSelectModel, err error) {
	err = dao.AdminMember.Ctx(ctx).
		Fields("id as value,real_name as label,username,avatar").
		Handler(handler.FilterAuthWithField("id")).
		Scan(&res)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}
	return
}

// VerifySuperId 验证是否为超管
func (s *sAdminMember) VerifySuperId(ctx context.Context, verifyId int64) bool {
	superIds := g.Cfg().MustGet(ctx, "hotgo.admin.superIds")
	for _, id := range superIds.Int64s() {
		if id == verifyId {
			return true
		}
	}
	return false
}

func (s *sAdminMember) FilterAuthModel(ctx context.Context, memberId int64) *gdb.Model {
	m := dao.AdminMember.Ctx(ctx)
	if !s.VerifySuperId(ctx, memberId) {
		m = m.Where("id <> ?", memberId)
	}
	return m.Handler(handler.FilterAuthWithField("id"))
}
