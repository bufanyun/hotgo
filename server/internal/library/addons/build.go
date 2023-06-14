package addons

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/model"
	"strconv"
	"strings"
)

// Build 构建新插件
func Build(ctx context.Context, sk Skeleton, conf *model.BuildAddonConfig) (err error) {
	var (
		buildPath    = "./" + consts.AddonsDir + "/" + sk.Name
		modulesPath  = "./" + consts.AddonsDir + "/modules/" + sk.Name + ".go"
		webApiPath   = gstr.Replace(conf.WebApiPath, "{$name}", sk.Name)
		webViewsPath = gstr.Replace(conf.WebViewsPath, "{$name}", sk.Name)
		replaces     = map[string]string{
			"@{.label}":       sk.Label,
			"@{.name}":        sk.Name,
			"@{.group}":       strconv.Itoa(sk.Group),
			"@{.brief}":       sk.Brief,
			"@{.description}": sk.Description,
			"@{.author}":      sk.Author,
			"@{.version}":     sk.Version,
			"@{.hgVersion}":   consts.VersionApp, // HG 版本
		}
	)

	if err = checkBuildDir(buildPath, modulesPath, webApiPath, webViewsPath); err != nil {
		return
	}

	// scans directory recursively
	list, err := gfile.ScanDirFunc(conf.SrcPath, "*", true, func(path string) string {
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
			gfile.RealPath(conf.SrcPath): "",
			".template":                  "",
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
