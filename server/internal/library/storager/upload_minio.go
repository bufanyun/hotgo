// Package storager
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package storager

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/s3utils"
	"mime"
	"path/filepath"
)

// MinioDrive minio对象存储驱动
type MinioDrive struct {
}

// Upload 上传到minio对象存储
func (d *MinioDrive) Upload(ctx context.Context, file *ghttp.UploadFile) (fullPath string, err error) {
	if config.MinioPath == "" {
		err = gerror.New("minio存储驱动必须配置存储路径!")
		return
	}

	client, err := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretKey, ""),
		Secure: config.MinioUseSSL == 1,
	})

	if err != nil {
		return "", err
	}

	if err = s3utils.CheckValidBucketName(config.MinioBucket); err != nil {
		return
	}

	fullPath = GenFullPath(config.MinioPath, gfile.Ext(file.Filename))
	if err = s3utils.CheckValidObjectName(fullPath); err != nil {
		return
	}

	reader, err := file.Open()
	if err != nil {
		return "", err
	}
	defer reader.Close()

	opts := minio.PutObjectOptions{
		ContentType: mime.TypeByExtension(filepath.Ext(file.Filename)),
	}
	if opts.ContentType == "" {
		opts.ContentType = "application/octet-stream"
	}

	_, err = client.PutObject(ctx, config.MinioBucket, fullPath, reader, file.Size, opts)
	return
}
