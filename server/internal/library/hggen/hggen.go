// Package hggen
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package hggen

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/library/hggen/internal/cmd"
	"hotgo/internal/library/hggen/internal/cmd/gendao"
	"hotgo/internal/library/hggen/views"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"sort"
)

// Dao 生成数据库实体
func Dao(ctx context.Context) (err error) {
	for _, v := range daoConfig {
		inp := defaultGenDaoInput
		err = gconv.Scan(v, &inp)
		if err != nil {
			return
		}
		gendao.DoGenDaoForArray(ctx, inp)
	}

	return
}

// Service 生成业务接口
func Service(ctx context.Context) (err error) {
	_, err = cmd.Gen.Service(ctx, GetServiceConfig())
	return
}

// TableColumns 获取指定表生成字段列表
func TableColumns(ctx context.Context, in sysin.GenCodesColumnListInp) (fields []*sysin.GenCodesColumnListModel, err error) {
	return views.DoTableColumns(ctx, in, GetDaoConfig(in.Name))
}

func TableSelects(ctx context.Context, in sysin.GenCodesSelectsInp) (res *sysin.GenCodesSelectsModel, err error) {
	res = new(sysin.GenCodesSelectsModel)
	for k, v := range consts.GenCodesTypeNameMap {
		res.GenType = append(res.GenType, &form.Select{
			Value: k,
			Name:  v,
			Label: v,
		})
	}
	sort.Sort(res.GenType)
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

	dictMode, _ := service.SysDictType().TreeSelect(ctx, sysin.DictTreeSelectInp{})

	res.DictMode = dictMode
	for _, v := range views.WhereModes {
		res.WhereMode = append(res.WhereMode, &form.Select{
			Value: v,
			Name:  v,
			Label: v,
		})
	}

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
func Preview(ctx context.Context, in sysin.GenCodesPreviewInp) (res *sysin.GenCodesPreviewModel, err error) {
	genConfig, err := service.SysConfig().GetLoadGenerate(ctx)
	if err != nil {
		return nil, err
	}

	switch in.GenType {
	case consts.GenCodesTypeCurd:
		return views.Curd.DoPreview(ctx, &views.CurdPreviewInput{
			In:        in,
			DaoConfig: GetDaoConfig(in.DbName),
			Config:    genConfig,
		})
	case consts.GenCodesTypeTree:
		err = gerror.Newf("生成类型开发中！")
		return
	case consts.GenCodesTypeQueue:
		err = gerror.Newf("生成类型开发中！")
		return
	default:
		err = gerror.Newf("生成类型暂不支持！")
		return
	}
}

// Build 提交生成
func Build(ctx context.Context, in sysin.GenCodesBuildInp) (err error) {
	genConfig, err := service.SysConfig().GetLoadGenerate(ctx)
	if err != nil {
		return err
	}

	switch in.GenType {
	case consts.GenCodesTypeCurd:
		return views.Curd.DoBuild(ctx, &views.CurdBuildInput{
			PreviewIn: &views.CurdPreviewInput{
				In:        sysin.GenCodesPreviewInp(in),
				DaoConfig: GetDaoConfig(in.DbName),
				Config:    genConfig,
			},
			BeforeEvent: views.CurdBuildEvent{"runDao": Dao},
			AfterEvent:  views.CurdBuildEvent{"runService": Service},
		})
	case consts.GenCodesTypeTree:
		err = gerror.Newf("生成类型开发中！")
		return
	case consts.GenCodesTypeQueue:
		err = gerror.Newf("生成类型开发中！")
		return
	default:
		err = gerror.Newf("生成类型暂不支持！")
		return
	}
}
