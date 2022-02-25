//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminService

import (
	"context"
	"github.com/bufanyun/hotgo/app/com"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/bufanyun/hotgo/app/service/internal/dao"
	"github.com/bufanyun/hotgo/app/service/internal/dto"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

var Member = new(member)

type member struct{}

//
//  @Title  修改登录密码
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *member) UpdateProfile(ctx context.Context, in input.AdminMemberUpdateProfileInp) (err error) {

	memberId := com.Context.Get(ctx).User.Id
	if memberId <= 0 {
		err := gerror.New("获取用户信息失败！")
		return err
	}

	var memberInfo entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("id", memberId).Scan(&memberInfo); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	_, err = dao.AdminMember.Ctx(ctx).
		Where("id", memberId).
		Data(g.Map{
			"mobile":   in.Mobile,
			"email":    in.Email,
			"realname": in.Realname,
		}).
		Update()

	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return
}

//
//  @Title  修改登录密码
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *member) UpdatePwd(ctx context.Context, in input.AdminMemberUpdatePwdInp) (err error) {

	var memberInfo entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("id", in.Id).Scan(&memberInfo); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if gmd5.MustEncryptString(in.OldPassword+memberInfo.Salt) != memberInfo.PasswordHash {
		err = gerror.New("原密码不正确")
		return err
	}

	_, err = dao.AdminMember.Ctx(ctx).
		Where("id", in.Id).
		Data(g.Map{
			"password_hash": gmd5.MustEncryptString(in.NewPassword + memberInfo.Salt),
			"updated_at":    gtime.Now(),
		}).
		Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return
}

//
//  @Title  重置密码
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *member) ResetPwd(ctx context.Context, in input.AdminMemberResetPwdInp) (err error) {

	var memberInfo entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("id", in.Id).Scan(&memberInfo); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	_, err = dao.AdminMember.Ctx(ctx).
		Where("id", in.Id).
		Data(g.Map{
			"password_hash": gmd5.MustEncryptString(in.Password + memberInfo.Salt),
			"updated_at":    gtime.Now(),
		}).
		Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return
}

//
//  @Title  菜单名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *member) EmailUnique(ctx context.Context, in input.AdminMemberEmailUniqueInp) (*input.AdminMemberEmailUniqueModel, error) {

	var res input.AdminMemberEmailUniqueModel
	isUnique, err := dao.AdminMember.IsUniqueEmail(ctx, in.Id, in.Email)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.IsUnique = isUnique
	return &res, nil
}

//
//  @Title  手机号是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *member) MobileUnique(ctx context.Context, in input.AdminMemberMobileUniqueInp) (*input.AdminMemberMobileUniqueModel, error) {

	var res input.AdminMemberMobileUniqueModel
	isUnique, err := dao.AdminMember.IsUniqueMobile(ctx, in.Id, in.Mobile)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.IsUnique = isUnique
	return &res, nil
}

//
//  @Title  菜单名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *member) NameUnique(ctx context.Context, in input.AdminMemberNameUniqueInp) (*input.AdminMemberNameUniqueModel, error) {

	var res input.AdminMemberNameUniqueModel
	isUnique, err := dao.AdminMember.IsUniqueName(ctx, in.Id, in.Username)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.IsUnique = isUnique
	return &res, nil
}

//
//  @Title  删除
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  error
//
func (service *member) Delete(ctx context.Context, in input.AdminMemberDeleteInp) error {

	exist, err := dao.AdminMember.Ctx(ctx).Where("member_id", in.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !exist.IsEmpty() {
		return gerror.New("请先解除该部门下所有已关联用户关联关系！")
	}
	_, err = dao.AdminMember.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

//
//  @Title  修改/新增
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  error
//
func (service *member) Edit(ctx context.Context, in input.AdminMemberEditInp) (err error) {

	if in.Username == "" {
		err = gerror.New("帐号不能为空")
		return err
	}

	uniqueName, err := dao.AdminMember.IsUniqueName(ctx, in.Id, in.Username)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !uniqueName {
		err = gerror.New("帐号已存在")
		return err
	}

	if in.Mobile != "" {
		uniqueMobile, err := dao.AdminMember.IsUniqueMobile(ctx, in.Id, in.Mobile)
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}
		if !uniqueMobile {
			err = gerror.New("手机号已存在")
			return err
		}
	}

	if in.Email != "" {
		uniqueEmail, err := dao.AdminMember.IsUniqueMobile(ctx, in.Id, in.Email)
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}
		if !uniqueEmail {
			err = gerror.New("邮箱已存在")
			return err
		}
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.AdminMember.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		// 更新岗位
		err = MemberPost.UpdatePostIds(ctx, in.Id, in.PostIds)
		if err != nil {
			return err
		}
		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()

	// 新增用户时的额外属性
	var data input.AdminMemberAddInp
	data.AdminMemberEditInp = in
	data.Salt = grand.S(6)
	data.PasswordHash = gmd5.MustEncryptString(data.Password + data.Salt)

	g.Log().Print(ctx, "data.Salt:", data)
	insert, err := dao.AdminMember.Ctx(ctx).Data(data).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	// 更新岗位
	id, err := insert.LastInsertId()
	if err != nil {
		return err
	}
	err = MemberPost.UpdatePostIds(ctx, id, in.PostIds)
	if err != nil {
		return err
	}
	return nil
}

//
//  @Title  最大排序
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictDataMaxSortRes
//  @Return  error
//
func (service *member) MaxSort(ctx context.Context, in input.AdminMemberMaxSortInp) (*input.AdminMemberMaxSortModel, error) {
	var res input.AdminMemberMaxSortModel

	if in.Id > 0 {
		if err := dao.AdminMember.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	res.Sort = res.Sort + 10

	return &res, nil
}

//
//  @Title  获取指定字典类型信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeViewRes
//  @Return  error
//
func (service *member) View(ctx context.Context, in input.AdminMemberViewInp) (res *input.AdminMemberViewModel, err error) {

	if err = dao.AdminMember.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
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
func (service *member) List(ctx context.Context, in input.AdminMemberListInp) (list []*input.AdminMemberListModel, totalCount int, err error) {

	var authorization = com.Jwt.GetAuthorization(com.Context.Get(ctx).Request)
	// TODO  获取jwtToken
	jwtToken := consts.RedisJwtToken + gmd5.MustEncryptString(authorization)
	g.Log().Print(ctx, "jwtToken:", jwtToken)

	mod := dao.AdminMember.Ctx(ctx)

	if in.Realname != "" {
		mod = mod.WhereLike("realname", "%"+in.Realname+"%")
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

	// 日期范围
	if in.StartTime != "" {
		mod = mod.WhereGTE("created_at", in.StartTime)
	}
	if in.EndTime != "" {
		mod = mod.WhereLTE("created_at", in.EndTime)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	err = mod.Page(in.Page, in.Limit).Order("id desc").Scan(&list)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	// TODO  重写树入参
	for i := 0; i < len(list); i++ {
		// TODO  部门
		deptName, err := dao.AdminDept.Ctx(ctx).
			Fields("name").
			Where("id", list[i].DeptId).
			Value()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return list, totalCount, err
		}
		list[i].DeptName = deptName.String()

		// TODO  角色
		roleName, err := dao.AdminRole.Ctx(ctx).
			Fields("name").
			Where("id", list[i].Role).
			Value()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return list, totalCount, err
		}
		list[i].RoleName = roleName.String()
	}

	return list, totalCount, err
}

// //
//  @Title  获取登录用户信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (service *member) LoginMemberInfo(ctx context.Context, req *adminForm.MemberInfoReq) (res *adminForm.MemberInfoRes, err error) {

	var (
		defaultPortalConfig  adminForm.PortalConfig
		defaultPortalConfigs []*adminForm.PortalConfig
		configContent        adminForm.PortalConfigContent
		configContents       []*adminForm.PortalConfigContent
		configContentOptions []*adminForm.PortalConfigContentOptions
		Options              adminForm.PortalConfigContentOptions
	)

	g.Log().Print(ctx, "测试")

	// TODO  配置内容选项
	Options.TitleRequired = true
	Options.Refresh = 1
	configContentOptions = append(configContentOptions, &Options)

	// TODO  配置内容
	configContent.Options = configContentOptions
	configContent.Id = 1
	configContent.X = 0
	configContent.Y = 0
	configContent.W = 3
	configContent.H = 262
	configContent.I = 1
	configContent.Key = "kuaijierukou"
	configContent.IsShowTitle = "N"
	configContent.IsAllowDrag = false
	configContent.Name = "快捷入口"
	configContent.Type = "smallPage"
	configContent.Url = "dashboard/portal/CommonUse"
	configContent.Moved = true

	configContents = append(configContents, &configContent)

	// TODO  默认配置
	defaultPortalConfig.Id = "4ae60dd1debe462096698e1da993317a"
	defaultPortalConfig.Name = "首页"
	defaultPortalConfig.Code = "6c297eb4651940edbb45c87c75be00d7"
	defaultPortalConfig.ApplicationRange = "U"
	defaultPortalConfig.IsDefault = "Y"
	defaultPortalConfig.ResourceId = "1"
	defaultPortalConfig.SystemDefinedId = "app1"
	defaultPortalConfig.PortalConfigContent = gconv.String(configContents)

	defaultPortalConfigs = append(defaultPortalConfigs, &defaultPortalConfig)

	member := com.Context.Get(ctx).User

	noticeList, err := Notice.WhereAll(ctx, dto.AdminNotice{
		Status: consts.StatusEnabled,
	})
	if err != nil {
		noticeList = nil
	}

	res = &adminForm.MemberInfoRes{
		LincenseInfo:        consts.VersionApp,
		Permissions:         []string{"*:*:*"},
		Roles:               []string{"admin"},
		User:                *member,
		DefaultPortalConfig: defaultPortalConfigs,
		UserPortalConfig:    defaultPortalConfigs,
		SysNoticeList:       noticeList,
	}

	return
}

//
//  @Title  提交登录
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (service *member) Login(ctx context.Context, in input.AdminMemberLoginSignInp) (res *input.AdminMemberLoginSignModel, err error) {

	var member *entity.AdminMember
	err = dao.AdminMember.Ctx(ctx).Where("username", in.Username).Scan(&member)

	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	if member == nil {
		err = gerror.New(consts.ErrorNotData)
		return
	}

	if member.Salt == "" {
		err = gerror.New("用户信息错误")
		return
	}

	if member.PasswordHash != gmd5.MustEncryptString(in.Password+member.Salt) {
		err = gerror.New("用户密码不正确")
		return
	}

	// 默认设备
	if in.Device != consts.AppAdmin && in.Device != consts.AppApi {
		in.Device = consts.AppAdmin
	}

	// TODO  生成token
	jwtExpires, err := g.Cfg().Get(ctx, "jwt.expires", 1)
	if err != nil {
		err := gerror.New(err.Error())
		return nil, err
	}
	// TODO  有效期
	expires := jwtExpires.Int64()

	// TODO  过期时间戳
	exp := gconv.Int64(gtime.Timestamp()) + expires

	var identity *model.Identity
	identity = &model.Identity{
		Id:         member.Id,
		Username:   member.Username,
		Realname:   member.Realname,
		Avatar:     member.Avatar,
		Email:      member.Email,
		Mobile:     member.Mobile,
		VisitCount: member.VisitCount,
		LastTime:   member.LastTime,
		LastIp:     member.LastIp,
		Role:       member.Role,
		Exp:        exp,
		Expires:    expires,
		App:        consts.AppAdmin,
	}
	token, err := com.Jwt.GenerateLoginToken(ctx, identity, false)
	if err != nil {
		err = gerror.New(err.Error())
		return
	}

	// TODO  更新登录信息
	authKey := gmd5.MustEncryptString(gconv.String(token))

	_, err = dao.AdminMember.Ctx(ctx).Data(dto.AdminMember{
		AuthKey:    gmd5.MustEncryptString(authKey),
		VisitCount: member.VisitCount + 1,
		LastTime:   gtime.Timestamp(),
		LastIp:     com.Context.Get(ctx).Request.GetClientIp(),
	}).Where(dto.AdminMember{
		Id: member.Id,
	}).Update()

	if err != nil {
		err = gerror.New(err.Error())
		return
	}

	res = &input.AdminMemberLoginSignModel{
		Identity: *identity,
		Token:    gconv.String(token),
	}

	return res, nil
}

//
//  @Title  获取角色下的会员列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (service *member) RoleMemberList(ctx context.Context, in input.AdminRoleMemberListInp) (list []*input.AdminMemberListModel, totalCount int, err error) {

	mod := dao.AdminMember.Ctx(ctx)
	if in.Role > 0 {
		mod = mod.Where("role", in.Role)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	err = mod.Page(in.Page, in.Limit).Order("id desc").Scan(&list)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	return list, totalCount, err
}
