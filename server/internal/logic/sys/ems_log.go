// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/ems"
	"hotgo/internal/library/location"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/charset"
	"hotgo/utility/simple"
	"hotgo/utility/useragent"
	"hotgo/utility/validate"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
)

type sSysEmsLog struct{}

func NewSysEmsLog() *sSysEmsLog {
	return &sSysEmsLog{}
}

func init() {
	service.RegisterSysEmsLog(NewSysEmsLog())
}

// Delete 删除
func (s *sSysEmsLog) Delete(ctx context.Context, in *sysin.EmsLogDeleteInp) (err error) {
	_, err = dao.SysEmsLog.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

// Edit 修改/新增
func (s *sSysEmsLog) Edit(ctx context.Context, in *sysin.EmsLogEditInp) (err error) {
	if in.Ip == "" {
		err = gerror.New("ip不能为空")
		return
	}

	// 修改
	if in.Id > 0 {
		_, err = dao.SysEmsLog.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		return
	}

	// 新增
	_, err = dao.SysEmsLog.Ctx(ctx).Data(in).OmitEmptyData().Insert()
	return
}

// Status 更新部门状态
func (s *sSysEmsLog) Status(ctx context.Context, in *sysin.EmsLogStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}

	// 修改
	_, err = dao.SysEmsLog.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	return
}

// View 获取指定字典类型信息
func (s *sSysEmsLog) View(ctx context.Context, in *sysin.EmsLogViewInp) (res *sysin.EmsLogViewModel, err error) {
	err = dao.SysEmsLog.Ctx(ctx).Where("id", in.Id).Scan(&res)
	return
}

// List 获取列表
func (s *sSysEmsLog) List(ctx context.Context, in *sysin.EmsLogListInp) (list []*sysin.EmsLogListModel, totalCount int, err error) {
	mod := dao.SysEmsLog.Ctx(ctx)

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if totalCount == 0 {
		return
	}

	err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list)
	return
}

// Send 发送邮件
func (s *sSysEmsLog) Send(ctx context.Context, in *sysin.SendEmsInp) (err error) {
	var models *entity.SysEmsLog
	if err = dao.SysEmsLog.Ctx(ctx).
		Where("event", in.Event).
		Where("email", in.Email).
		OrderDesc(dao.SysEmsLog.Columns().Id).
		Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	config, err := service.SysConfig().GetSmtp(ctx)
	if err != nil {
		return
	}

	in.Template, err = s.GetTemplate(ctx, in.Event, config)
	if err != nil {
		return
	}

	if err = s.AllowSend(ctx, models, config); err != nil {
		return
	}

	if consts.IsCodeEmsTemplate(in.Event) && in.Code == "" {
		in.Code = grand.Digits(4)
	}

	view, err := s.newView(ctx, in, config)
	if err != nil {
		return
	}

	if in.TplData == nil {
		in.TplData = make(g.Map)
	}

	switch in.Event {
	// 富文本
	case consts.EmsTemplateText:
		if in.Content == "" {
			err = gerror.New("富文本类型邮件内容不能为空")
			return
		}
		in.TplData["content"] = in.Content
		in.Content, err = view.Parse(ctx, in.Template, in.TplData)
		if err != nil {
			return err
		}
	// 验证码、重置密码类
	default:
		in.Content, err = view.Parse(ctx, in.Template, in.TplData)
		if err != nil {
			return err
		}
	}

	subject, ok := consts.EmsSubjectMap[in.Event]
	if !ok {
		subject = simple.AppName(ctx)
	}

	err = ems.Send(config, in.Email, subject, in.Content)
	if err != nil {
		return
	}

	var data = new(entity.SysEmsLog)
	data.Event = in.Event
	data.Email = in.Email
	data.Content = in.Content
	data.Code = in.Code
	data.Ip = location.GetClientIp(ghttp.RequestFromCtx(ctx))
	data.Status = consts.CodeStatusNotUsed
	data.CreatedAt = gtime.Now()
	data.UpdatedAt = gtime.Now()

	_, err = dao.SysEmsLog.Ctx(ctx).Data(data).OmitEmptyData().Insert()
	return
}

func (s *sSysEmsLog) newView(ctx context.Context, in *sysin.SendEmsInp, config *model.EmailConfig) (view *gview.View, err error) {
	view = gview.New()
	err = view.SetConfig(gview.Config{
		Delimiters: g.Cfg().MustGet(ctx, "viewer.delimiters").Strings(),
	})
	if err != nil {
		return
	}

	// 富文本为自定义内容，可能不需要变量
	if in.Event == consts.EmsTemplateText {
		return
	}

	var (
		username string
		user     = contexts.GetUser(ctx)
		request  = ghttp.RequestFromCtx(ctx)
		ip       = location.GetClientIp(request)
	)

	loc, err := location.GetLocation(ctx, ip)
	if err != nil {
		return
	}

	if loc == nil {
		loc = new(location.IpLocationData)
	}

	cityLabel, err := location.ParseRegion(ctx, loc.ProvinceCode, loc.CityCode, 0)
	if err != nil {
		return
	}

	basic, err := service.SysConfig().GetBasic(ctx)
	if err != nil {
		return
	}

	if basic == nil {
		basic = new(model.BasicConfig)
		basic.Name = simple.AppName(ctx)
		basic.Domain = "https://hotgo.facms.cn"
		basic.Logo = "https://gmycos.facms.cn/haoka/attachment/images/2023-02-04/cq9kf7s66jt7hkpvbh.png"
		basic.SystemOpen = true
	}

	if user != nil {
		username = user.Username
	}

	// 公共变量
	view.Assigns(gview.Params{
		"code":      in.Code,                                           // 验证码
		"expires":   config.CodeExpire / 60,                            // 验证码有效期(分钟)
		"username":  username,                                          // 发送者用户名
		"name":      basic.Name,                                        // 网站名称
		"logo":      basic.Logo,                                        // 网站logo
		"domain":    basic.Domain,                                      // 网站域名
		"github":    "https://github.com/bufanyun/hotgo",               // github
		"os":        useragent.GetOs(request.Header.Get("User-Agent")), // 发送者操作系统
		"ip":        gstr.HideStr(ip, 30, `*`),                         // 发送者IP
		"cityLabel": cityLabel,                                         // IP归属地
	})

	// 重置密码
	if in.Event == consts.EmsTemplateResetPwd {
		var (
			passwordResetLink string
			resetToken        = charset.RandomCreateBytes(32)
		)
		if user != nil {
			switch user.App {
			// 后台用户
			case consts.AppAdmin:
				_, err = dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Id, user.Id).Data(g.Map{dao.AdminMember.Columns().PasswordResetToken: resetToken}).Update()
				if err != nil {
					return
				}
				passwordResetLink = fmt.Sprintf("%s/admin/passwordReset?token=%s", basic.Domain, resetToken)
			// 前台用户
			case consts.AppApi:
				// ...
			}
		}
		view.Assign("passwordResetLink", passwordResetLink)
	}
	return
}

// GetTemplate 获取指定邮件模板
func (s *sSysEmsLog) GetTemplate(ctx context.Context, template string, config *model.EmailConfig) (val string, err error) {
	if template == "" {
		err = gerror.New("模板不能为空")
		return
	}
	if config == nil {
		config, err = service.SysConfig().GetSmtp(ctx)
		if err != nil {
			return
		}
	}

	if len(config.Template) == 0 {
		err = gerror.New("管理员还没有配置任何模板！")
		return
	}

	for _, v := range config.Template {
		if v.Key == template {
			return v.Value, nil
		}
	}
	return
}

// AllowSend 是否允许发送
func (s *sSysEmsLog) AllowSend(ctx context.Context, models *entity.SysEmsLog, config *model.EmailConfig) (err error) {
	if models == nil {
		return
	}

	// 富文本事件不限制
	if models.Event == consts.EmsTemplateText {
		return
	}

	if config == nil {
		config, err = service.SysConfig().GetSmtp(ctx)
		if err != nil {
			return
		}
	}

	if gtime.Now().Before(models.CreatedAt.Add(time.Second * time.Duration(config.MinInterval))) {
		err = gerror.New("发送频繁，请稍后再试！")
		return
	}

	if config.MaxIpLimit > 0 {
		count, err := s.NowDayIpSendCount(ctx, models.Event)
		if err != nil {
			return err
		}

		if count >= config.MaxIpLimit {
			err = gerror.New("今天发送短信过多，请次日后再试！")
			return err
		}
	}
	return
}

// NowDayIpSendCount 当天IP累计发送次数
func (s *sSysEmsLog) NowDayIpSendCount(ctx context.Context, event string) (count int, err error) {
	return dao.SysEmsLog.Ctx(ctx).
		Where("ip", location.GetClientIp(ghttp.RequestFromCtx(ctx))).
		Where("event", event).
		WhereGTE("created_at", gtime.Now().Format("Y-m-d")).
		Count()
}

// VerifyCode 效验验证码
func (s *sSysEmsLog) VerifyCode(ctx context.Context, in *sysin.VerifyEmsCodeInp) (err error) {
	if in.Event == "" {
		err = gerror.New("事件不能为空")
		return
	}
	if in.Email == "" {
		err = gerror.New("邮箱不能为空")
		return
	}

	if in.Event == consts.EmsTemplateResetPwd || in.Event == consts.EmsTemplateText {
		err = gerror.Newf("事件类型无需验证:%v", in.Event)
		return
	}

	config, err := service.SysConfig().GetSmtp(ctx)
	if err != nil {
		return
	}

	var models *entity.SysEmsLog
	if err = dao.SysEmsLog.Ctx(ctx).
		Where(dao.SysEmsLog.Columns().Event, in.Event).
		Where(dao.SysEmsLog.Columns().Email, in.Email).
		OrderDesc(dao.SysEmsLog.Columns().Id).
		Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if models == nil {
		err = gerror.New("验证码错误")
		return
	}

	if models.Times >= 10 {
		err = gerror.New("验证码错误次数过多，请重新发送！")
		return
	}

	if in.Event != consts.EmsTemplateCode {
		if models.Status == consts.CodeStatusUsed {
			err = gerror.New("验证码已使用，请重新发送！")
			return
		}
	}

	if gtime.Now().After(models.CreatedAt.Add(time.Second * time.Duration(config.CodeExpire))) {
		err = gerror.New("验证码已过期，请重新发送")
		return
	}

	if models.Code != in.Code {
		_, _ = dao.SysEmsLog.Ctx(ctx).Where(dao.SysEmsLog.Columns().Id, models.Id).Increment(dao.SysEmsLog.Columns().Times, 1)
		err = gerror.New("验证码错误！")
		return
	}

	_, err = dao.SysEmsLog.Ctx(ctx).Where(dao.SysEmsLog.Columns().Id, models.Id).Data(g.Map{
		dao.SysEmsLog.Columns().Times:     models.Times + 1,
		dao.SysEmsLog.Columns().Status:    consts.CodeStatusUsed,
		dao.SysEmsLog.Columns().UpdatedAt: gtime.Now(),
	}).Update()
	return
}
