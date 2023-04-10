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
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var MaskDemoField = []string{
	// 邮箱
	"smtpUser", "smtpPass",

	// 云存储
	"uploadUCloudPublicKey", "uploadUCloudPrivateKey",
	"uploadCosSecretId", "uploadCosSecretKey",
	"uploadOssSecretId", "uploadOssSecretKey",
	"uploadQiNiuAccessKey", "uploadQiNiuSecretKey",

	// 地图
	"geoAmapWebKey",
	
	// 短信
	"smsAliYunAccessKeyID", "smsAliYunAccessKeySecret",
	"smsTencentSecretId", "smsTencentSecretKey",
}

type sSysConfig struct{}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

// GetLoadCache 获取本地缓存配置
func (s *sSysConfig) GetLoadCache(ctx context.Context) (conf *model.CacheConfig, err error) {
	err = g.Cfg().MustGet(ctx, "cache").Scan(&conf)
	return
}

// GetLoadGenerate 获取本地生成配置
func (s *sSysConfig) GetLoadGenerate(ctx context.Context) (conf *model.GenerateConfig, err error) {
	generate := g.Cfg().MustGet(ctx, "hggen")
	if err = gconv.Struct(generate, &conf); err != nil {
		return nil, err
	}

	return conf, nil
}

// GetSms 获取短信配置
func (s *sSysConfig) GetSms(ctx context.Context) (conf *model.SmsConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "sms"})
	if err != nil {
		return nil, err
	}
	if err = gconv.Struct(models.List, &conf); err != nil {
		return nil, err
	}

	return conf, nil
}

// GetGeo 获取地理配置
func (s *sSysConfig) GetGeo(ctx context.Context) (conf *model.GeoConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "geo"})
	if err != nil {
		return nil, err
	}
	if err = gconv.Struct(models.List, &conf); err != nil {
		return nil, err
	}

	return conf, nil
}

// GetUpload 获取上传配置
func (s *sSysConfig) GetUpload(ctx context.Context) (conf *model.UploadConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "upload"})
	if err != nil {
		return nil, err
	}
	if err = gconv.Struct(models.List, &conf); err != nil {
		return nil, err
	}

	return conf, nil
}

// GetSmtp 获取邮件配置
func (s *sSysConfig) GetSmtp(ctx context.Context) (conf *model.EmailConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "smtp"})
	if err != nil {
		return nil, err
	}
	if err = gconv.Struct(models.List, &conf); err != nil {
		return nil, err
	}

	conf.Addr = fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	return conf, nil
}

// GetBasic 获取基础配置
func (s *sSysConfig) GetBasic(ctx context.Context) (conf *model.BasicConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, sysin.GetConfigInp{Group: "basic"})
	if err != nil {
		return nil, err
	}
	if err = gconv.Struct(models.List, &conf); err != nil {
		return nil, err
	}

	return conf, nil
}

// GetLoadSSL 获取本地日志配置
func (s *sSysConfig) GetLoadSSL(ctx context.Context) (conf *model.SSLConfig, err error) {
	if err = g.Cfg().MustGet(ctx, "hotgo.ssl").Struct(&conf); err != nil {
		return nil, err
	}
	return conf, nil
}

// GetLoadLog 获取本地日志配置
func (s *sSysConfig) GetLoadLog(ctx context.Context) (conf *model.LogConfig, err error) {
	if err = g.Cfg().MustGet(ctx, "hotgo.log").Struct(&conf); err != nil {
		return nil, err
	}
	return conf, nil
}

// GetLoadServeLog 获取本地服务日志配置
func (s *sSysConfig) GetLoadServeLog(ctx context.Context) (conf *model.ServeLogConfig, err error) {
	if err = g.Cfg().MustGet(ctx, "hotgo.serveLog").Struct(&conf); err != nil {
		return nil, err
	}
	return conf, nil
}

// GetConfigByGroup 获取指定分组的配置
func (s *sSysConfig) GetConfigByGroup(ctx context.Context, in sysin.GetConfigInp) (*sysin.GetConfigModel, error) {
	if in.Group == "" {
		return nil, gerror.New("分组不能为空")
	}
	var (
		mod    = dao.SysConfig.Ctx(ctx)
		models []*entity.SysConfig
		res    sysin.GetConfigModel
	)
	if err := mod.Fields("key", "value", "type").Where("group", in.Group).Scan(&models); err != nil {
		return nil, err
	}
	isDemo := g.Cfg().MustGet(ctx, "hotgo.isDemo", false)

	if len(models) > 0 {
		res.List = make(g.Map, len(models))
		for _, v := range models {
			val, err := s.ConversionType(ctx, v)
			if err != nil {
				return nil, err
			}
			res.List[v.Key] = val
			//if isDemo.Bool() && (v.Key == "smtpUser" || v.Key == "smtpPass") {
			//	res.List[v.Key] = consts.DemoTips
			//	res.List[v.Key] = consts.DemoTips
			//}

			if isDemo.Bool() && gstr.InArray(MaskDemoField, v.Key) {
				res.List[v.Key] = consts.DemoTips
				res.List[v.Key] = consts.DemoTips
			}
		}
	}
	return &res, nil
}

// ConversionType 转换类型
func (s *sSysConfig) ConversionType(ctx context.Context, models *entity.SysConfig) (value interface{}, err error) {
	if models == nil {
		return nil, gerror.New("数据不存在")
	}
	return consts.ConvType(models.Value, models.Type), nil
}

// UpdateConfigByGroup 更新指定分组的配置
func (s *sSysConfig) UpdateConfigByGroup(ctx context.Context, in sysin.UpdateConfigInp) error {
	if in.Group == "" {
		return gerror.New("分组不能为空")
	}
	var (
		mod    = dao.SysConfig.Ctx(ctx)
		models []*entity.SysConfig
	)
	if err := mod.Where("group", in.Group).Scan(&models); err != nil {
		return err
	}

	err := dao.SysConfig.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
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
				return gerror.Newf("暂不支持从前台添加变量，请先在数据库表[%v]中配置变量：%v", dao.SysConfig.Table(), k)
			}

			// 更新
			_, err := dao.SysConfig.Ctx(ctx).Where("id", row.Id).Data(g.Map{"value": v, "updated_at": gtime.Now()}).Update()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
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
