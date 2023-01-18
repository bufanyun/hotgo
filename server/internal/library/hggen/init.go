// Package hggen
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package hggen

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"gopkg.in/yaml.v3"
	"hotgo/internal/library/hggen/internal/cmd/gendao"
	"hotgo/internal/library/hggen/internal/cmd/genservice"
)

const (
	cliFolderName    = `hack/config.yaml`
	RequiredErrorTag = `the cli configuration file must be configured %s`
)

var (
	config        g.Map
	daoConfig     []interface{}
	serviceConfig g.Map

	// 生成service默认参数，请不要直接修改以下配置，如需调整请到/hack/config.yaml，可参考：https://goframe.org/pages/viewpage.action?pageId=49770772
	defaultGenServiceInput = genservice.CGenServiceInput{
		SrcFolder:       "internal/logic",
		DstFolder:       "internal/service",
		DstFileNameCase: "Snake",
		StPattern:       `s([A-Z]\w+)`,
		Clear:           false,
	}

	// 生成dao默认参数，请不要直接修改以下配置，如需调整请到/hack/config.yaml，可参考：https://goframe.org/pages/viewpage.action?pageId=3673173
	defaultGenDaoInput = gendao.CGenDaoInput{
		Path:           "internal",
		Group:          "default",
		JsonCase:       "CamelLower",
		DaoPath:        "dao",
		DoPath:         "model/do",
		EntityPath:     "model/entity",
		StdTime:        false,
		WithTime:       false,
		GJsonSupport:   false,
		OverwriteDao:   false,
		DescriptionTag: true,
		NoJsonTag:      false,
		NoModelComment: false,
		Clear:          false,
	}
)

func GetServiceConfig() genservice.CGenServiceInput {
	inp := defaultGenServiceInput
	_ = gconv.Scan(serviceConfig, &inp)
	return inp
}

func GetDaoConfig(group string) gendao.CGenDaoInput {
	inp := defaultGenDaoInput
	find := func(group string) g.Map {
		for _, v := range daoConfig {
			if v.(g.Map)["group"].(string) == group {
				return v.(g.Map)
			}
		}
		return nil
	}

	v := find(group)
	if v != nil {
		err := gconv.Scan(v, &inp)
		if err != nil {
			panic(err)
		}
	}

	return inp
}

func InIt(ctx context.Context) {
	path, err := gfile.Search(cliFolderName)
	if err != nil {
		g.Log().Fatalf(ctx, "get cli configuration file:%v, err:%+v", cliFolderName, err)
	}
	if path == "" {
		g.Log().Fatalf(ctx, "get cli configuration file:%v fail", cliFolderName)
	}
	if config == nil {
		config = make(g.Map)
	}
	err = yaml.Unmarshal(gfile.GetBytes(path), &config)
	if err != nil {
		g.Log().Fatalf(ctx, "load cli configuration file:%v, yaml err:%+v", cliFolderName, err)
	}
	loadConfig(ctx)
}

func loadConfig(ctx context.Context) {
	if _, ok := config["gfcli"]; !ok {
		g.Log().Fatalf(ctx, RequiredErrorTag, "gfcli")
	}

	if _, ok := config["gfcli"].(g.Map)["gen"]; !ok {
		g.Log().Fatalf(ctx, RequiredErrorTag, "gfcli.gen")
	}

	dao, ok := config["gfcli"].(g.Map)["gen"].(map[string]interface{})["dao"]
	if !ok {
		g.Log().Fatalf(ctx, RequiredErrorTag, "gfcli.gen.dao")
	}
	daoConf, ok := dao.([]interface{})
	if !ok {
		g.Log().Fatalf(ctx, RequiredErrorTag, "gfcli.gen.dao format error")
	}
	daoConfig = daoConf
	for _, v := range daoConfig {
		if _, ok := v.(g.Map)["group"].(string); !ok {
			g.Log().Fatalf(ctx, "group must be configured in %s: `gfcli.gen.dao` and must be the same as the database group", cliFolderName)
		}
	}

	if serviceConf, ok := config["gfcli"].(g.Map)["gen"].(map[string]interface{})["service"]; ok {
		if serviceConfig == nil {
			serviceConfig = make(g.Map)
		}
		serviceConfig = serviceConf.(g.Map)
	}
}
