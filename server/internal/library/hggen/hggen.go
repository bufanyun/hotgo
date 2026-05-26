// Package hggen
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package hggen

import (
	_ "hotgo/internal/library/hggen/internal/cmd/gendao"
	"hotgo/internal/library/hggen/internal/utility/utils"
	_ "unsafe"

	"github.com/gogf/gf/v2/os/gfile"

	"context"
	"hotgo/internal/consts"
	"hotgo/internal/library/addons"
	"hotgo/internal/library/hggen/internal/cmd"
	"hotgo/internal/library/hggen/internal/cmd/gendao"
	"hotgo/internal/library/hggen/internal/cmd/genservice"
	"hotgo/internal/library/hggen/views"
	"hotgo/internal/model"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"sort"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

//go:linkname doGenDaoForArray hotgo/internal/library/hggen/internal/cmd/gendao.doGenDaoForArray
func doGenDaoForArray(ctx context.Context, index int, in gendao.CGenDaoInput)

// Dao 生成数据库实体
func Dao(ctx context.Context) (err error) {
	// 在执行gf gen dao时，先将生成文件放在临时路径，生成完成后再拷贝到项目
	// 目的是希望减少触发gf热编译的几率，防止热编译运行时代码生成流程未结束被自动重启打断
	// gf gen dao 的执行时长主要取决于需要生成数据库表的数量，表越多速度越慢
	tempPathPrefix := views.GetTempGeneratePath(ctx) + "/dao"

	for _, v := range daoConfig {
		inp := defaultGenDaoInput
		if err = gconv.Scan(v, &inp); err != nil {
			return
		}
		oldPath := inp.Path
		inp.ImportPrefix = utils.GetImportPath(inp.Path)
		inp.Path = tempPathPrefix + "/" + inp.Path

		if err = gfile.RemoveAll(inp.Path); err != nil {
			err = gerror.Newf("清理临时生成目录失败:%v", err)
			return err
		}

		if err = gfile.Mkdir(inp.Path); err != nil {
			err = gerror.Newf("创建临时生成目录失败:%v", err)
			return err
		}

		doGenDaoForArray(ctx, -1, inp)

		if err = gfile.CopyDir(inp.Path, gfile.Pwd()+"/"+oldPath); err != nil {
			err = gerror.Newf("拷贝生成文件失败:%v", err)
			return err
		}
	}
	return
}

// Service 生成业务接口
func Service(ctx context.Context) (err error) {
	return ServiceWithCfg(ctx, GetServiceConfig())
}

// ServiceWithCfg 生成业务接口
func ServiceWithCfg(ctx context.Context, cfg ...genservice.CGenServiceInput) (err error) {
	c := GetServiceConfig()
	if len(cfg) > 0 {
		c = cfg[0]
	}
	_, err = cmd.Gen.Service(ctx, c)
	return
}

// TableColumns 获取指定表生成字段列表
func TableColumns(ctx context.Context, in *sysin.GenCodesColumnListInp) (fields []*sysin.GenCodesColumnListModel, err error) {
	return views.DoTableColumns(ctx, in, GetDaoConfig(in.Name))
}

func TableSelects(ctx context.Context, in *sysin.GenCodesSelectsInp) (res *sysin.GenCodesSelectsModel, err error) {
	res = new(sysin.GenCodesSelectsModel)
	res.GenType, err = GenTypeSelect(ctx)
	if err != nil {
		return
	}

	res.Db = DbSelect(ctx)

	for k, v := range consts.GenCodesStatusNameMap {
		res.Status = append(res.Status, &form.Select{
			Value: k,
			Name:  v,
			Label: v,
		})
	}
	sort.Sort(res.Status)

	for k, v := range consts.GenCodesJoinNameMap {
		res.LinkMode = append(res.LinkMode, &form.Select{
			Value: k,
			Name:  v,
			Label: v,
		})
	}
	sort.Sort(res.LinkMode)

	for k, v := range consts.GenCodesBuildMethNameMap {
		res.BuildMeth = append(res.BuildMeth, &form.Select{
			Value: k,
			Name:  v,
			Label: v,
		})
	}
	sort.Sort(res.BuildMeth)

	for _, v := range views.FormModes {
		res.FormMode = append(res.FormMode, &form.Select{
			Value: v,
			Name:  views.FormModeMap[v],
			Label: views.FormModeMap[v],
		})
	}
	sort.Sort(res.FormMode)

	for k, v := range views.FormRoleMap {
		res.FormRole = append(res.FormRole, &form.Select{
			Value: k,
			Name:  v,
			Label: v,
		})
	}
	sort.Sort(res.FormRole)

	dictMode, err := service.SysDictType().TreeSelect(ctx, &sysin.DictTreeSelectInp{})
	if err != nil {
		return
	}
	res.DictMode = dictMode

	for _, v := range views.WhereModes {
		res.WhereMode = append(res.WhereMode, &form.Select{
			Value: v,
			Name:  v,
			Label: v,
		})
	}
	for _, v := range views.TableAligns {
		res.TableAlign = append(res.TableAlign, &form.Select{
			Value: v,
			Name:  views.TableAlignMap[v],
			Label: views.TableAlignMap[v],
		})
	}

	res.Addons = addons.ModuleSelect()
	res.TreeStyleType = consts.GenCodesTreeStyleTypeOptions
	return
}

// GenTypeSelect 获取生成类型选项
func GenTypeSelect(ctx context.Context) (res sysin.GenTypeSelects, err error) {
	for k, v := range consts.GenCodesTypeNameMap {
		row := &sysin.GenTypeSelect{
			Value:     k,
			Name:      v,
			Label:     v,
			Templates: make(sysin.GenTemplateSelects, 0),
		}

		km := k
		if km == consts.GenCodesTypeTree {
			km = consts.GenCodesTypeCurd
		}
		confName, ok := consts.GenCodesTypeConfMap[km]
		if ok {
			var temps []*model.GenerateAppCrudTemplate
			err = g.Cfg().MustGet(ctx, "hggen.application."+confName+".templates").Scan(&temps)
			if err != nil {
				return
			}
			if len(temps) > 0 {
				for index, temp := range temps {
					row.Templates = append(row.Templates, &sysin.GenTemplateSelect{
						Value:   index,
						Label:   temp.Group,
						Name:    temp.Group,
						IsAddon: temp.IsAddon,
					})
				}
				sort.Sort(row.Templates)
			}
		}

		res = append(res, row)
	}
	sort.Sort(res)
	return
}

// DbSelect db选项
func DbSelect(ctx context.Context) (res form.Selects) {
	dbs := g.Cfg().MustGet(ctx, "hggen.selectDbs")
	if len(dbs.Strings()) == 0 {
		res = make(form.Selects, 0)
		return res
	}

	for _, v := range dbs.Strings() {
		res = append(res, &form.Select{
			Value: v,
			Label: v,
			Name:  v,
		})
	}
	return res
}

// Preview 生成预览
func Preview(ctx context.Context, in *sysin.GenCodesPreviewInp) (res *sysin.GenCodesPreviewModel, err error) {
	genConfig, err := service.SysConfig().GetLoadGenerate(ctx)
	if err != nil {
		return nil, err
	}

	switch in.GenType {
	case consts.GenCodesTypeCurd, consts.GenCodesTypeTree:
		return views.Curd.DoPreview(ctx, &views.CurdPreviewInput{
			In:        in,
			DaoConfig: GetDaoConfig(in.DbName),
			Config:    genConfig,
		})
	case consts.GenCodesTypeQueue:
		err = gerror.Newf("生成类型开发中！")
		return
	default:
		err = gerror.Newf("生成类型暂不支持！")
		return
	}
}

// Build 提交生成
func Build(ctx context.Context, in *sysin.GenCodesBuildInp) (err error) {
	genConfig, err := service.SysConfig().GetLoadGenerate(ctx)
	if err != nil {
		return err
	}

	switch in.GenType {
	case consts.GenCodesTypeCurd, consts.GenCodesTypeTree:
		pin := &sysin.GenCodesPreviewInp{SysGenCodes: in.SysGenCodes}
		return views.Curd.DoBuild(ctx, &views.CurdBuildInput{
			PreviewIn: &views.CurdPreviewInput{
				In:        pin,
				DaoConfig: GetDaoConfig(in.DbName),
				Config:    genConfig,
			},
			BeforeEvent: views.CurdBuildEvent{"runDao": Dao},
			AfterEvent: views.CurdBuildEvent{"runService": func(ctx context.Context) (err error) {
				cfg := GetServiceConfig()

				// 插件模块，切换到插件下运行gen service
				if genConfig.Application.Crud.Templates[pin.GenTemplate].IsAddon {
					// 依然使用配置中的参数，只是将生成路径指向插件模块路径
					cfg.SrcFolder = "addons/" + pin.AddonName + "/logic"
					cfg.DstFolder = "addons/" + pin.AddonName + "/service"
				}
				err = ServiceWithCfg(ctx, cfg)
				return
			}},
		})
	case consts.GenCodesTypeQueue:
		err = gerror.Newf("生成类型开发中！")
		return
	default:
		err = gerror.Newf("生成类型暂不支持！")
		return
	}
}
