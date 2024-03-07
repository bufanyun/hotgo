// Package addons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package addons

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/model"
	"hotgo/utility/validate"
	"strconv"
	"strings"
)

type BuildOption struct {
	Skeleton Skeleton
	Config   *model.BuildAddonConfig
	Extend   []string `json:"extend" dc:"扩展功能"`
}

// Build 构建新插件
func Build(ctx context.Context, option *BuildOption) (err error) {
	var (
		resourcePath = GetResourcePath(ctx)
		buildPath    = "./" + consts.AddonsDir + "/" + option.Skeleton.Name
		modulesPath  = "./" + consts.AddonsDir + "/modules/" + option.Skeleton.Name + ".go"
		webApiPath   = gstr.Replace(option.Config.WebApiPath, "{$name}", option.Skeleton.Name)
		webViewsPath = gstr.Replace(option.Config.WebViewsPath, "{$name}", option.Skeleton.Name)
		replaces     = map[string]string{
			"@{.label}":       option.Skeleton.Label,
			"@{.name}":        option.Skeleton.Name,
			"@{.group}":       strconv.Itoa(option.Skeleton.Group),
			"@{.brief}":       option.Skeleton.Brief,
			"@{.description}": option.Skeleton.Description,
			"@{.author}":      option.Skeleton.Author,
			"@{.version}":     option.Skeleton.Version,
			"@{.hgVersion}":   consts.VersionApp, // HG 版本
		}
	)

	if resourcePath == "" {
		err = gerror.New("请先设置一个有效的插件资源路径，配置名称:'hotgo.addonsResourcePath'")
		return
	}

	if err = checkBuildDir(buildPath, modulesPath, webApiPath, webViewsPath); err != nil {
		return
	}

	// scans directory recursively
	list, err := gfile.ScanDirFunc(option.Config.SrcPath, "*", true, func(path string) string {
		return path
	})

	if err != nil {
		return
	}

	for _, path := range list {
		if !gfile.IsReadable(path) {
			err = gerror.Newf("file：%v is unreadable, please check permissions", path)
			return
		}

		if gfile.IsDir(path) {
			continue
		}

		flowFile := gstr.ReplaceByMap(path, map[string]string{
			gfile.RealPath(option.Config.SrcPath): "",
			".template":                           "",
		})
		flowFile = buildPath + "/" + flowFile

		content := gstr.ReplaceByMap(gfile.GetContents(path), replaces)

		if err = gfile.PutContents(flowFile, content); err != nil {
			break
		}
	}

	// 隐式注入插件
	if err = gfile.PutContents(modulesPath, gstr.ReplaceByMap(importModules, replaces)); err != nil {
		return
	}

	// webApi
	if err = gfile.PutContents(webApiPath+"/config/index.ts", gstr.ReplaceByMap(webApiLayout, replaces)); err != nil {
		return
	}

	// web插件配置主页面
	if err = gfile.PutContents(webViewsPath+"/config/BasicSetting.vue", gstr.ReplaceByMap(webConfigBasicSetting, replaces)); err != nil {
		return
	}

	// web插件基础配置页面
	if err = gfile.PutContents(webViewsPath+"/config/system.vue", gstr.ReplaceByMap(webConfigSystem, replaces)); err != nil {
		return
	}

	// 创建静态目录
	if validate.InSlice(option.Extend, consts.AddonsExtendResourcePublic) {
		_, staticPath := StaticPath(option.Skeleton.Name, resourcePath)
		content := fmt.Sprintf(resourcePublicDefaultFile, option.Skeleton.Label)
		if err = gfile.PutContents(staticPath+"/default", content); err != nil {
			return
		}
	}

	// 创建模板目录
	if validate.InSlice(option.Extend, consts.AddonsExtendResourceTemplate) {
		viewPath := ViewPath(option.Skeleton.Name, resourcePath)
		if err = gfile.PutContents(viewPath+"/home/index.html", resourceTemplateHomeFile); err != nil {
			return
		}
	}
	return
}

func checkBuildDir(paths ...string) (err error) {
	if len(paths) == 0 {
		return
	}

	for _, path := range paths {
		if gfile.Exists(path) {
			return gerror.Newf("插件已存在，请换一个插件名称或者经确认无误后依次删除文件夹： [%v] 后重新生成", strings.Join(paths, "、\t"))
		}
	}
	return
}
