// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/payment"
	"hotgo/internal/library/storager"
	"hotgo/internal/library/token"
	"hotgo/internal/library/wechat"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

type sSysConfig struct{}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

// InitConfig 初始化一些系统启动就需要用到的配置
func (s *sSysConfig) InitConfig(ctx context.Context) {
	wx, err := s.GetWechat(ctx)
	if err != nil {
		g.Log().Fatalf(ctx, "init wechat conifg fail：%+v", err)
	}
	wechat.SetConfig(wx)

	pay, err := s.GetPay(ctx)
	if err != nil {
		g.Log().Fatalf(ctx, "init pay conifg fail：%+v", err)
	}
	payment.SetConfig(pay)

	upload, err := s.GetUpload(ctx)
	if err != nil {
		g.Log().Fatalf(ctx, "init upload conifg fail：%+v", err)
	}
	storager.SetConfig(upload)

	tk, err := s.GetLoadToken(ctx)
	if err != nil {
		g.Log().Fatalf(ctx, "init token conifg fail：%+v", err)
	}
	token.SetConfig(tk)
}

// GetLogin 获取登录配置
func (s *sSysConfig) GetLogin(ctx context.Context) (conf *model.LoginConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "login"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetWechat 获取微信配置
func (s *sSysConfig) GetWechat(ctx context.Context) (conf *model.WechatConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "wechat"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetPay 获取支付配置
func (s *sSysConfig) GetPay(ctx context.Context) (conf *model.PayConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "pay"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetSms 获取短信配置
func (s *sSysConfig) GetSms(ctx context.Context) (conf *model.SmsConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "sms"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetGeo 获取地理配置
func (s *sSysConfig) GetGeo(ctx context.Context) (conf *model.GeoConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "geo"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetUpload 获取上传配置
func (s *sSysConfig) GetUpload(ctx context.Context) (conf *model.UploadConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "upload"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetSmtp 获取邮件配置
func (s *sSysConfig) GetSmtp(ctx context.Context) (conf *model.EmailConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "smtp"})
	if err != nil {
		return
	}
	if err = gconv.Scan(models.List, &conf); err != nil {
		return
	}

	conf.Addr = fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	return
}

// GetBasic 获取基础配置
func (s *sSysConfig) GetBasic(ctx context.Context) (conf *model.BasicConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "basic"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetLoadTCP 获取本地tcp配置
func (s *sSysConfig) GetLoadTCP(ctx context.Context) (conf *model.TCPConfig, err error) {
	err = g.Cfg().MustGet(ctx, "tcp").Scan(&conf)
	return
}

// GetLoadCache 获取本地缓存配置
func (s *sSysConfig) GetLoadCache(ctx context.Context) (conf *model.CacheConfig, err error) {
	err = g.Cfg().MustGet(ctx, "cache").Scan(&conf)
	return
}

// GetLoadGenerate 获取本地生成配置
func (s *sSysConfig) GetLoadGenerate(ctx context.Context) (conf *model.GenerateConfig, err error) {
	err = g.Cfg().MustGet(ctx, "hggen").Scan(&conf)
	return
}

// GetLoadToken 获取本地token配置
func (s *sSysConfig) GetLoadToken(ctx context.Context) (conf *model.TokenConfig, err error) {
	err = g.Cfg().MustGet(ctx, "token").Scan(&conf)
	return
}

// GetLoadSSL 获取本地日志配置
func (s *sSysConfig) GetLoadSSL(ctx context.Context) (conf *model.SSLConfig, err error) {
	err = g.Cfg().MustGet(ctx, "hotgo.ssl").Scan(&conf)
	return
}

// GetLoadLog 获取本地日志配置
func (s *sSysConfig) GetLoadLog(ctx context.Context) (conf *model.LogConfig, err error) {
	err = g.Cfg().MustGet(ctx, "hotgo.log").Scan(&conf)
	return
}

// GetLoadServeLog 获取本地服务日志配置
func (s *sSysConfig) GetLoadServeLog(ctx context.Context) (conf *model.ServeLogConfig, err error) {
	err = g.Cfg().MustGet(ctx, "hotgo.serveLog").Scan(&conf)
	return
}

// GetConfigByGroup 获取指定分组的配置
func (s *sSysConfig) GetConfigByGroup(ctx context.Context, in sysin.GetConfigInp) (res *sysin.GetConfigModel, err error) {
	if in.Group == "" {
		err = gerror.New("分组不能为空")
		return
	}

	var models []*entity.SysConfig
	if err = dao.SysConfig.Ctx(ctx).Fields("key", "value", "type").Where("group", in.Group).Scan(&models); err != nil {
		err = gerror.Wrapf(err, "获取配置分组[ %v ]失败，请稍后重试！", in.Group)
		return
	}

	res = new(sysin.GetConfigModel)
	if len(models) > 0 {
		res.List = make(g.Map, len(models))
		for _, v := range models {
			val, err := s.ConversionType(ctx, v)
			if err != nil {
				return nil, err
			}
			res.List[v.Key] = val
		}
	}

	res.List = simple.FilterMaskDemo(ctx, res.List)
	return
}

// ConversionType 转换类型
func (s *sSysConfig) ConversionType(ctx context.Context, models *entity.SysConfig) (value interface{}, err error) {
	if models == nil {
		err = gerror.New("数据不存在")
		return
	}
	return consts.ConvType(models.Value, models.Type), nil
}

// UpdateConfigByGroup 更新指定分组的配置
func (s *sSysConfig) UpdateConfigByGroup(ctx context.Context, in sysin.UpdateConfigInp) (err error) {
	if in.Group == "" {
		err = gerror.New("分组不能为空")
		return
	}
	var (
		mod    = dao.SysConfig.Ctx(ctx)
		models []*entity.SysConfig
	)

	if err = mod.Where("group", in.Group).Scan(&models); err != nil {
		return
	}

	err = dao.SysConfig.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		for k, v := range in.List {
			row := s.getConfigByKey(k, models)
			// 新增
			if row == nil {
				//row.Id = 0
				//row.Key = k
				//row.Value = gconv.String(v)
				//row.Group = in.Group
				//row.Status = consts.StatusEnabled
				//row.CreatedAt = gtime.Now()
				//row.UpdatedAt = gtime.Now()
				//_, err := dao.SysConfig.Ctx(ctx).Data(row).Insert()
				//if err != nil {
				//	err = gerror.Wrap(err, consts.ErrorORM)
				//	return err
				//}
				//continue
				err = gerror.Newf("暂不支持从前台添加变量，请先在数据库表[%v]中配置变量：%v", dao.SysConfig.Table(), k)
				return
			}

			// 更新
			_, err = dao.SysConfig.Ctx(ctx).Where("id", row.Id).Data(g.Map{"value": v, "updated_at": gtime.Now()}).Update()
			if err != nil {
				return
			}
		}

		return s.syncUpdate(ctx, in)
	})
	return
}

func (s *sSysConfig) getConfigByKey(key string, models []*entity.SysConfig) *entity.SysConfig {
	if len(models) == 0 {
		return nil
	}

	for _, v := range models {
		if key == v.Key {
			return v
		}
	}
	return nil
}

// syncUpdate 同步更新一些加载配置
func (s *sSysConfig) syncUpdate(ctx context.Context, in sysin.UpdateConfigInp) (err error) {
	switch in.Group {
	case "wechat":
		wx, err := s.GetWechat(ctx)
		if err == nil {
			wechat.SetConfig(wx)
		}
	case "pay":
		pay, err := s.GetPay(ctx)
		if err == nil {
			payment.SetConfig(pay)
		}
	case "upload":
		upload, err := s.GetUpload(ctx)
		if err == nil {
			storager.SetConfig(upload)
		}
	}

	if err != nil {
		err = gerror.Newf("syncUpdate %v conifg fail：%+v", in.Group, err.Error())
	}
	return
}
