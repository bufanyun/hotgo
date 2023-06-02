package storager

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
)

// OssDrive 阿里云oss驱动
type OssDrive struct {
}

// Upload 上传到阿里云oss
func (d *OssDrive) Upload(ctx context.Context, file *ghttp.UploadFile) (fullPath string, err error) {
	if config.OssPath == "" {
		err = gerror.New("OSS存储驱动必须配置存储路径!")
		return
	}

	// 流式上传本地小文件
	f2, err := file.Open()
	defer func() { _ = f2.Close() }()
	if err != nil {
		return
	}

	client, err := oss.New(config.OssEndpoint, config.OssSecretId, config.OssSecretKey)
	if err != nil {
		return
	}

	bucket, err := client.Bucket(config.OssBucket)
	if err != nil {
		return
	}

	fullPath = GenFullPath(config.UCloudPath, gfile.Ext(file.Filename))
	err = bucket.PutObject(fullPath, f2)
	return
}
