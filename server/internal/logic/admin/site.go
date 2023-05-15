package admin

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/token"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

type sAdminSite struct{}

func NewAdminSite() *sAdminSite {
	return &sAdminSite{}
}

func init() {
	service.RegisterAdminSite(NewAdminSite())
}

// Register 账号注册
func (s *sAdminSite) Register(ctx context.Context, in adminin.RegisterInp) (err error) {
	config, err := service.SysConfig().GetLogin(ctx)
	if err != nil {
		return
	}

	if config.ForceInvite == 1 && in.InviteCode == "" {
		err = gerror.New("请填写邀请码")
		return
	}

	var data adminin.MemberAddInp
	// 默认上级
	data.Pid = 1

	// 存在邀请人
	if in.InviteCode != "" {
		pmb, err := service.AdminMember().GetIdByCode(ctx, adminin.GetIdByCodeInp{Code: in.InviteCode})
		if err != nil {
			return err
		}

		if pmb == nil {
			err = gerror.New("邀请人信息不存在")
			return err
		}

		data.Pid = pmb.Id
	}

	if config.RegisterSwitch != 1 {
		err = gerror.New("管理员未开放注册")
		return
	}

	if config.RoleId < 1 {
		err = gerror.New("管理员未配置默认角色")
		return
	}

	if config.DeptId < 1 {
		err = gerror.New("管理员未配置默认部门")
		return
	}

	if len(config.PostIds) == 0 {
		err = gerror.New("管理员未配置默认岗位")
		return
	}

	// 验证唯一性
	err = service.AdminMember().VerifyUnique(ctx, adminin.VerifyUniqueInp{
		Where: g.Map{
			dao.AdminMember.Columns().Username: in.Username,
			dao.AdminMember.Columns().Mobile:   in.Mobile,
		},
	})
	if err != nil {
		return
	}

	// 验证短信验证码
	err = service.SysSmsLog().VerifyCode(ctx, sysin.VerifyCodeInp{
		Event:  consts.SmsTemplateRegister,
		Mobile: in.Mobile,
		Code:   in.Code,
	})
	if err != nil {
		return
	}

	data.MemberEditInp = adminin.MemberEditInp{
		Id:       0,
		RoleId:   config.RoleId,
		PostIds:  config.PostIds,
		DeptId:   config.DeptId,
		Username: in.Username,
		Password: in.Password,
		RealName: "",
		Avatar:   config.Avatar,
		Sex:      3, // 保密
		Mobile:   in.Mobile,
		Status:   consts.StatusEnabled,
	}
	data.Salt = grand.S(6)
	data.InviteCode = grand.S(12)
	data.PasswordHash = gmd5.MustEncryptString(data.Password + data.Salt)
	data.Level, data.Tree, err = service.AdminMember().GenTree(ctx, data.Pid)
	if err != nil {
		return
	}

	// 提交注册信息
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		id, err := dao.AdminMember.Ctx(ctx).Data(data).InsertAndGetId()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}

		// 更新岗位
		if err = dao.AdminMemberPost.UpdatePostIds(ctx, id, config.PostIds); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
		}
		return
	})
}

// AccountLogin 账号登录
func (s *sAdminSite) AccountLogin(ctx context.Context, in adminin.AccountLoginInp) (res *adminin.LoginModel, err error) {
	defer func() {
		service.SysLoginLog().Push(ctx, sysin.LoginLogPushInp{Response: res, Err: err})
	}()

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("username", in.Username).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if mb == nil {
		err = gerror.New("账号不存在")
		return
	}

	res = new(adminin.LoginModel)
	res.Id = mb.Id
	res.Username = mb.Username
	if mb.Salt == "" {
		err = gerror.New("用户信息错误")
		return
	}

	if err = simple.CheckPassword(in.Password, mb.Salt, mb.PasswordHash); err != nil {
		return
	}

	if mb.Status != consts.StatusEnabled {
		err = gerror.New("账号已被禁用")
		return
	}

	res, err = s.handleLogin(ctx, mb)
	return
}

// MobileLogin 手机号登录
func (s *sAdminSite) MobileLogin(ctx context.Context, in adminin.MobileLoginInp) (res *adminin.LoginModel, err error) {
	defer func() {
		service.SysLoginLog().Push(ctx, sysin.LoginLogPushInp{Response: res, Err: err})
	}()

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("mobile ", in.Mobile).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if mb == nil {
		err = gerror.New("账号不存在")
		return
	}

	res = new(adminin.LoginModel)
	res.Id = mb.Id
	res.Username = mb.Username

	err = service.SysSmsLog().VerifyCode(ctx, sysin.VerifyCodeInp{
		Event:  consts.SmsTemplateLogin,
		Mobile: in.Mobile,
		Code:   in.Code,
	})

	if err != nil {
		return
	}

	if mb.Status != consts.StatusEnabled {
		err = gerror.New("账号已被禁用")
		return
	}

	res, err = s.handleLogin(ctx, mb)
	return
}

// handleLogin .
func (s *sAdminSite) handleLogin(ctx context.Context, mb *entity.AdminMember) (res *adminin.LoginModel, err error) {
	var ro *entity.AdminRole
	if err = dao.AdminRole.Ctx(ctx).Fields("id,key,status").Where("id", mb.RoleId).Scan(&ro); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if ro == nil {
		err = gerror.New("角色不存在")
		return
	}

	if ro.Status != consts.StatusEnabled {
		err = gerror.New("角色已被禁用")
		return
	}

	user := &model.Identity{
		Id:       mb.Id,
		Pid:      mb.Pid,
		DeptId:   mb.DeptId,
		RoleId:   ro.Id,
		RoleKey:  ro.Key,
		Username: mb.Username,
		RealName: mb.RealName,
		Avatar:   mb.Avatar,
		Email:    mb.Email,
		Mobile:   mb.Mobile,
		App:      consts.AppAdmin,
		LoginAt:  gtime.Now(),
	}

	loginToken, expires, err := token.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	res = &adminin.LoginModel{
		Username: user.Username,
		Id:       user.Id,
		Token:    loginToken,
		Expires:  expires,
	}

	return
}
