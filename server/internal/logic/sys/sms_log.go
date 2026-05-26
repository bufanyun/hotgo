// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/location"
	"hotgo/internal/library/sms"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"time"
)

type sSysSmsLog struct{}

func NewSysSmsLog() *sSysSmsLog {
	return &sSysSmsLog{}
}

func init() {
	service.RegisterSysSmsLog(NewSysSmsLog())
}

// Delete 删除
func (s *sSysSmsLog) Delete(ctx context.Context, in *sysin.SmsLogDeleteInp) (err error) {
	_, err = dao.SysSmsLog.Ctx(ctx).WherePri(in.Id).Delete()
	return
}

// View 获取指定字典类型信息
func (s *sSysSmsLog) View(ctx context.Context, in *sysin.SmsLogViewInp) (res *sysin.SmsLogViewModel, err error) {
	if err = dao.SysSmsLog.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return
}

// List 获取列表
func (s *sSysSmsLog) List(ctx context.Context, in *sysin.SmsLogListInp) (list []*sysin.SmsLogListModel, totalCount int, err error) {
	mod := dao.SysSmsLog.Ctx(ctx)
	cols := dao.SysSmsLog.Columns()

	if in.Mobile != "" {
		mod = mod.WhereLike(cols.Mobile, "%"+in.Mobile+"%")
	}

	if in.Ip != "" {
		mod = mod.Where(cols.Ip, in.Ip)
	}

	if in.Event != "" {
		mod = mod.Where(cols.Event, in.Event)
	}

	if in.Status > 0 {
		mod = mod.Where(cols.Status, in.Status)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(cols.CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).OrderDesc(cols.Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return
}

// SendCode 发送验证码
func (s *sSysSmsLog) SendCode(ctx context.Context, in *sysin.SendCodeInp) (err error) {
	if in.Event == "" {
		err = gerror.New("事件不能为空")
		return
	}

	if in.Mobile == "" {
		err = gerror.New("手机号不能为空")
		return
	}

	var models *entity.SysSmsLog
	if err = dao.SysSmsLog.Ctx(ctx).
		Where(dao.SysSmsLog.Columns().Event, in.Event).
		Where(dao.SysSmsLog.Columns().Mobile, in.Mobile).
		OrderDesc(dao.SysSmsLog.Columns().Id).
		Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	config, err := service.SysConfig().GetSms(ctx)
	if err != nil {
		return
	}

	if in.Template, err = s.GetTemplate(ctx, in.Event, config); err != nil {
		return
	}

	if err = s.AllowSend(ctx, models, config); err != nil {
		return
	}

	if in.Code == "" {
		in.Code = grand.Digits(4)
	}

	if err = sms.New(config.SmsDrive).SendCode(ctx, in); err != nil {
		return
	}

	var data = new(entity.SysSmsLog)
	data.Event = in.Event
	data.Mobile = in.Mobile
	data.Code = in.Code
	data.Ip = location.GetClientIp(ghttp.RequestFromCtx(ctx))
	data.Status = consts.CodeStatusNotUsed
	data.CreatedAt = gtime.Now()
	data.UpdatedAt = gtime.Now()

	_, err = dao.SysSmsLog.Ctx(ctx).Data(data).OmitEmptyData().Insert()
	return
}

// GetTemplate 获取指定短信模板
func (s *sSysSmsLog) GetTemplate(ctx context.Context, template string, config *model.SmsConfig) (val string, err error) {
	if template == "" {
		err = gerror.New("模板不能为空")
		return
	}
	if config == nil {
		config, err = service.SysConfig().GetSms(ctx)
		if err != nil {
			return
		}
	}

	switch config.SmsDrive {
	case consts.SmsDriveAliYun:
		if len(config.AliYunTemplate) == 0 {
			err = gerror.New("管理员还没有配置任何阿里云短信模板！")
			return
		}

		for _, v := range config.AliYunTemplate {
			if v.Key == template {
				return v.Value, nil
			}
		}

	case consts.SmsDriveTencent:
		if len(config.TencentTemplate) == 0 {
			err = gerror.New("管理员还没有配置任何腾讯云短信模板！")
			return
		}

		for _, v := range config.TencentTemplate {
			if v.Key == template {
				return v.Value, nil
			}
		}
	default:
		err = gerror.Newf("暂不支持短信驱动:%v", config.SmsDrive)
		return
	}
	return
}

// AllowSend 是否允许发送
func (s *sSysSmsLog) AllowSend(ctx context.Context, models *entity.SysSmsLog, config *model.SmsConfig) (err error) {
	if models == nil {
		return
	}

	if config == nil {
		if config, err = service.SysConfig().GetSms(ctx); err != nil {
			return
		}
	}

	if gtime.Now().Before(models.CreatedAt.Add(time.Second * time.Duration(config.SmsMinInterval))) {
		err = gerror.New("发送频繁，请稍后再试！")
		return
	}

	if config.SmsMaxIpLimit > 0 {
		count, err := s.NowDayIpSendCount(ctx, models.Event)
		if err != nil {
			return err
		}

		if count >= config.SmsMaxIpLimit {
			err = gerror.New("今天发送短信过多，请次日后再试！")
			return err
		}
	}
	return
}

// NowDayIpSendCount 当天IP累计发送次数
func (s *sSysSmsLog) NowDayIpSendCount(ctx context.Context, event string) (count int, err error) {
	return dao.SysSmsLog.Ctx(ctx).
		Where("ip", location.GetClientIp(ghttp.RequestFromCtx(ctx))).
		Where("event", event).
		WhereGTE("created_at", gtime.Now().Format("Y-m-d")).
		Count()
}

// VerifyCode 效验验证码
func (s *sSysSmsLog) VerifyCode(ctx context.Context, in *sysin.VerifyCodeInp) (err error) {
	if in.Event == "" {
		err = gerror.New("事件不能为空")
		return
	}

	if in.Mobile == "" {
		err = gerror.New("手机号不能为空")
		return
	}

	config, err := service.SysConfig().GetSms(ctx)
	if err != nil {
		return
	}

	var models *entity.SysSmsLog
	cols := dao.SysSmsLog.Columns()
	if err = dao.SysSmsLog.Ctx(ctx).Where(cols.Event, in.Event).Where(cols.Mobile, in.Mobile).OrderDesc(cols.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("验证码错误")
		return
	}

	if models.Times >= 10 {
		err = gerror.New("验证码错误次数过多，请重新发送！")
		return
	}

	if in.Event != consts.SmsTemplateCode {
		if models.Status == consts.CodeStatusUsed {
			err = gerror.New("验证码已使用，请重新发送！")
			return
		}
	}

	if gtime.Now().After(models.CreatedAt.Add(time.Second * time.Duration(config.SmsCodeExpire))) {
		err = gerror.New("验证码已过期，请重新发送")
		return
	}

	if models.Code != in.Code {
		_, _ = dao.SysSmsLog.Ctx(ctx).WherePri(models.Id).Increment(cols.Times, 1)
		err = gerror.New("验证码错误！")
		return
	}

	_, err = dao.SysSmsLog.Ctx(ctx).WherePri(models.Id).Data(g.Map{
		cols.Times:     models.Times + 1,
		cols.Status:    consts.CodeStatusUsed,
		cols.UpdatedAt: gtime.Now(),
	}).Update()
	return
}
