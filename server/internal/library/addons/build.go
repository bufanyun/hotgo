package addons

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/model"
	"strconv"
	"strings"
)

// Build 构建新插件
func Build(ctx context.Context, sk Skeleton, conf *model.BuildAddonConfig) (err error) {
	buildPath := "./" + consts.AddonsDir + "/" + sk.Name
	modulesPath := "./" + consts.AddonsDir + "/modules/" + sk.Name + ".go"
	templatePath := gstr.Replace(conf.TemplatePath, "{$name}", sk.Name)
	replaces := map[string]string{
		"@{.label}":       sk.Label,
		"@{.name}":        sk.Name,
		"@{.group}":       strconv.Itoa(sk.Group),
		"@{.brief}":       sk.Brief,
		"@{.description}": sk.Description,
		"@{.author}":      sk.Author,
		"@{.version}":     sk.Version,
	}

	if err = checkBuildDir(buildPath, modulesPath, templatePath); err != nil {
		return
	}

	// scans directory recursively
	list, err := gfile.ScanDirFunc(conf.SrcPath, "*", true, func(path string) string {
		return path
	})

	for _, path := range list {
		if !gfile.IsReadable(path) {
			err = fmt.Errorf("file：%v is unreadable, please check permissions", path)
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

	if err = gfile.PutContents(templatePath+"/home/index.html", homeLayout); err != nil {
		return
	}
	
	err = gfile.PutContents(modulesPath, gstr.ReplaceByMap(importModules, replaces))
	return
}

func checkBuildDir(paths ...string) error {
	if len(paths) == 0 {
		return nil
	}

	for _, path := range paths {
		if gfile.Exists(path) {
			return fmt.Errorf("插件已存在，请换一个插件名称或者经确认无误后依次删除文件夹： [%v] 后重新生成", strings.Join(paths, "、\t"))
		}
	}
	return nil
}

const (
	importModules = `// Package modules
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package modules

import _ "hotgo/addons/@{.name}"
`

	homeLayout = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0,user-scalable=no">
    <meta name="keywords" content="@{.Keywords}"/>
    <meta name="description" content="@{.Description}"/>
    <title>@{.Title}</title>
    <script type="text/javascript" src="/resource/home/js/jquery-3.6.0.min.js"></script>
    <style>
        html, body {
            width: 100%;
            height: 100%;
            margin: 0;
            padding: 0;
            background-color: #f6f6f6;
        }
    </style>
</head>
<body>
<div style="padding-top: 100px;text-align:center;">
    <h1><p>Hello，@{.Data.name}!!</p></h1>
    <h2><p>@{.Data.module}</p></h2>
    <h2><p>服务器时间：@{.Data.time}</p></h2>
</div>

</body>
<script>

</script>
</html>`
)
