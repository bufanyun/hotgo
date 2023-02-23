package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/addons/hgexample/global"
	"hotgo/addons/hgexample/model"
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/addons/hgexample/service"
	isc "hotgo/internal/service"
)

type sSysConfig struct{}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

// GetBasic 获取基础配置
func (s *sSysConfig) GetBasic(ctx context.Context) (conf *model.BasicConfig, err error) {
	var in sysin.GetConfigInp
	in.GetAddonsConfigInp.AddonName = global.GetSkeleton().Name
	in.GetAddonsConfigInp.Group = "basic"
	models, err := isc.SysAddonsConfig().GetConfigByGroup(ctx, in.GetAddonsConfigInp)
	if err != nil {
		return
	}

	err = gconv.Struct(models.List, &conf)
	return
}

// GetConfigByGroup 获取指定分组配置
func (s *sSysConfig) GetConfigByGroup(ctx context.Context, in sysin.GetConfigInp) (res *sysin.GetConfigModel, err error) {
	in.GetAddonsConfigInp.AddonName = global.GetSkeleton().Name
	models, err := isc.SysAddonsConfig().GetConfigByGroup(ctx, in.GetAddonsConfigInp)
	if err != nil {
		return
	}

	res = new(sysin.GetConfigModel)
	res.List = models.List
	return
}

// UpdateConfigByGroup 更新指定分组的配置
func (s *sSysConfig) UpdateConfigByGroup(ctx context.Context, in sysin.UpdateConfigInp) error {
	in.UpdateAddonsConfigInp.AddonName = global.GetSkeleton().Name
	return isc.SysAddonsConfig().UpdateConfigByGroup(ctx, in.UpdateAddonsConfigInp)
}
