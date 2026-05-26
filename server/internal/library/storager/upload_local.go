// Package storager
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package storager

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"os"
	"path/filepath"
	"sort"
	"strconv"
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

// CreateMultipart 创建分片事件
func (d *LocalDrive) CreateMultipart(ctx context.Context, in *CheckMultipartParams) (mp *MultipartProgress, err error) {
	mp = new(MultipartProgress)
	mp.UploadId = GenUploadId(ctx, in.Md5)
	mp.Meta = in.meta
	mp.ShardCount = in.ShardCount
	mp.UploadedIndex = make([]int, 0)
	mp.CreatedAt = gtime.Now()
	if err = CreateMultipartProgress(ctx, mp); err != nil {
		return nil, err
	}
	return
}

// UploadPart 上传分片
func (d *LocalDrive) UploadPart(ctx context.Context, in *UploadPartParams) (res *UploadPartModel, err error) {
	sp := g.Cfg().MustGet(ctx, "server.serverRoot")
	if sp.IsEmpty() {
		err = gerror.New("本地上传驱动必须配置静态路径!")
		return
	}

	spStr := strings.Trim(sp.String(), "/") + "/"

	if config.LocalPath == "" {
		err = gerror.New("本地上传驱动必须配置本地存储路径!")
		return
	}

	// 分片文件存放路径
	partFilePath := spStr + config.LocalPath + "tmp/" + in.Md5

	// 写入文件
	in.File.Filename = gconv.String(in.Index)
	if _, err = in.File.Save(partFilePath, false); err != nil {
		return
	}

	// 更新上传进度
	in.mp.UploadedIndex = append(in.mp.UploadedIndex, in.Index)
	if err = UpdateMultipartProgress(ctx, in.mp); err != nil {
		return nil, err
	}

	res = new(UploadPartModel)

	// 已全部上传完毕
	if len(in.mp.UploadedIndex) == in.mp.ShardCount {
		// 删除进度统计
		if err = DelMultipartProgress(ctx, in.mp); err != nil {
			return nil, err
		}

		// 合并文件
		finalDirPath := GenFullPath(config.LocalPath, gfile.Ext(in.mp.Meta.Filename))
		if err = MergePartFile(partFilePath, spStr+finalDirPath); err != nil {
			err = gerror.Newf("合并分片文件出错:%v", err.Error())
			return nil, err
		}

		// 删除临时分片
		if err = os.RemoveAll(partFilePath); err != nil {
			err = gerror.Newf("删除临时分片文件出错:%v", err.Error())
			return nil, err
		}

		// 写入附件记录
		attachment, err := write(ctx, in.mp.Meta, finalDirPath)
		if err != nil {
			return nil, err
		}

		res.Finish = true
		res.Progress = 100
		res.Attachment = attachment
		return res, nil
	}

	// 计算上传进度
	res.Progress = CalcUploadProgress(in.mp.UploadedIndex, in.mp.ShardCount)
	return
}

// MergePartFile 合并分片文件
func MergePartFile(srcPath, dstPath string) (err error) {
	dir, err := os.ReadDir(srcPath)
	if err != nil {
		return err
	}
	sort.Slice(dir, func(i, j int) bool {
		fiIndex, _ := strconv.Atoi(dir[i].Name())
		fjIndex, _ := strconv.Atoi(dir[j].Name())
		return fiIndex < fjIndex
	})
	for _, file := range dir {
		filePath := filepath.Join(srcPath, file.Name())
		if err = gfile.PutBytesAppend(dstPath, gfile.GetBytes(filePath)); err != nil {
			return err
		}
	}
	return
}
