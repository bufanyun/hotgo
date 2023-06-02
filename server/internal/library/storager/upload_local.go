package storager

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"strings"
)

// LocalDrive 本地驱动
type LocalDrive struct {
}

// Upload 上传到本地
func (d *LocalDrive) Upload(ctx context.Context, file *ghttp.UploadFile) (fullPath string, err error) {
	var (
		sp      = g.Cfg().MustGet(ctx, "server.serverRoot")
		nowDate = gtime.Date()
	)

	if sp.IsEmpty() {
		err = gerror.New("本地上传驱动必须配置静态路径!")
		return
	}

	if config.LocalPath == "" {
		err = gerror.New("本地上传驱动必须配置本地存储路径!")
		return
	}

	// 包含静态文件夹的路径
	fullDirPath := strings.Trim(sp.String(), "/") + "/" + config.LocalPath + nowDate
	fileName, err := file.Save(fullDirPath, true)
	if err != nil {
		return
	}
	// 不含静态文件夹的路径
	fullPath = config.LocalPath + nowDate + "/" + fileName
	return
}
