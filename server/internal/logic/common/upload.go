// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package common

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	ufile "github.com/ufilesdk-dev/ufile-gosdk"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/encrypt"
	f "hotgo/utility/file"
	"hotgo/utility/format"
	"hotgo/utility/url"
	"hotgo/utility/validate"
	"strconv"
	"strings"
	"time"
)

type sCommonUpload struct{}

func NewCommonUpload() *sCommonUpload {
	return &sCommonUpload{}
}

func init() {
	service.RegisterCommonUpload(NewCommonUpload())
}

// UploadFile 上传文件
func (s *sCommonUpload) UploadFile(ctx context.Context, file *ghttp.UploadFile) (result *sysin.AttachmentListModel, err error) {
	if file == nil {
		err = gerror.New("文件必须!")
		return
	}

	meta, err := s.fileMeta(file)
	if err != nil {
		return
	}

	_, err = f.GetFileType(meta.Ext)
	if err != nil {
		return nil, err
	}

	conf, err := service.SysConfig().GetUpload(ctx)
	if err != nil {
		return
	}

	switch conf.Drive {
	case consts.UploadDriveLocal:
		return s.UploadLocal(ctx, conf, file, meta)
	case consts.UploadDriveUCloud:
		return s.UploadUCloud(ctx, conf, file, meta)
	default:
		return nil, gerror.Newf("暂不支持上传驱动:%v", conf.Drive)
	}
}

// UploadImage 上传图片
func (s *sCommonUpload) UploadImage(ctx context.Context, file *ghttp.UploadFile) (result *sysin.AttachmentListModel, err error) {
	if file == nil {
		err = gerror.New("文件必须!")
		return
	}

	meta, err := s.fileMeta(file)
	if err != nil {
		return
	}

	if !f.IsImgType(meta.Ext) {
		return nil, gerror.New("上传的文件不是图片")
	}

	if meta.Size > 2*1024*1024 {
		return nil, gerror.New("图片大小不能超过2MB")
	}

	conf, err := service.SysConfig().GetUpload(ctx)
	if err != nil {
		return
	}

	switch conf.Drive {
	case consts.UploadDriveLocal:
		return s.UploadLocal(ctx, conf, file, meta)
	case consts.UploadDriveUCloud:
		return s.UploadUCloud(ctx, conf, file, meta)
	default:
		return nil, gerror.Newf("暂不支持上传驱动:%v", conf.Drive)
	}
}

// UploadLocal 上传本地
func (s *sCommonUpload) UploadLocal(ctx context.Context, conf *model.UploadConfig, file *ghttp.UploadFile, meta *sysin.UploadFileMeta) (result *sysin.AttachmentListModel, err error) {
	result, err = dao.SysAttachment.GetMd5File(ctx, meta.Md5)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if result != nil {
		return
	}

	var (
		value   = g.Cfg().MustGet(ctx, "server.serverRoot")
		nowDate = time.Now().Format("2006-01-02")
	)

	if value.IsEmpty() {
		err = gerror.New("本地上传驱动必须配置静态路径!")
		return
	}

	if conf.LocalPath == "" {
		err = gerror.New("本地上传驱动必须配置本地存储路径!")
		return
	}

	// 包含静态文件夹的路径
	fullDirPath := strings.Trim(value.String(), "/") + "/" + conf.LocalPath + nowDate
	fileName, err := file.Save(fullDirPath, true)
	if err != nil {
		return
	}
	// 不含静态文件夹的路径
	fullPath := conf.LocalPath + nowDate + "/" + fileName

	attachment, err := service.SysAttachment().Add(ctx, meta, fullPath, consts.UploadDriveLocal)
	if err != nil {
		return nil, err
	}

	attachment.FileUrl = s.LastUrl(ctx, conf, attachment.FileUrl, attachment.Drive)
	result = &sysin.AttachmentListModel{
		SysAttachment: *attachment,
		SizeFormat:    format.FileSize(attachment.Size),
	}
	return
}

// UploadUCloud 上传UCloud对象存储
func (s *sCommonUpload) UploadUCloud(ctx context.Context, conf *model.UploadConfig, file *ghttp.UploadFile, meta *sysin.UploadFileMeta) (result *sysin.AttachmentListModel, err error) {
	result, err = dao.SysAttachment.GetMd5File(ctx, meta.Md5)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if result != nil {
		return
	}

	if conf.UCloudPath == "" {
		err = gerror.New("UCloud存储驱动必须配置存储路径!")
		return
	}

	nowDate := time.Now().Format("2006-01-02")
	fileName := gfile.Basename(file.Filename)
	fileName = strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	fileName = fileName + gfile.Ext(file.Filename)
	fullPath := conf.UCloudPath + nowDate + "/" + fileName
	config := &ufile.Config{
		PublicKey:       conf.UCloudPublicKey,
		PrivateKey:      conf.UCloudPrivateKey,
		BucketHost:      conf.UCloudBucketHost,
		BucketName:      conf.UCloudBucketName,
		FileHost:        conf.UCloudFileHost,
		Endpoint:        conf.UCloudEndpoint,
		VerifyUploadMD5: false,
	}
	req, err := ufile.NewFileRequest(config, nil)
	if err != nil {
		return nil, err
	}
	// 流式上传本地小文件
	f2, err := file.Open()
	defer func() {
		_ = f2.Close()
	}()
	if err != nil {
		return nil, err
	}

	if err = req.IOPut(f2, fullPath, ""); err != nil {
		return nil, err
	}

	g.Log().Warningf(ctx, "ras:%+v", string(req.LastResponseBody))

	attachment, err := service.SysAttachment().Add(ctx, meta, fullPath, consts.UploadDriveUCloud)
	if err != nil {
		return nil, err
	}

	attachment.FileUrl = s.LastUrl(ctx, conf, attachment.FileUrl, attachment.Drive)
	result = &sysin.AttachmentListModel{
		SysAttachment: *attachment,
		SizeFormat:    format.FileSize(attachment.Size),
	}
	return
}

// LastUrl 根据驱动获取最终文件访问地址
func (s *sCommonUpload) LastUrl(ctx context.Context, conf *model.UploadConfig, fullPath, drive string) string {
	if validate.IsURL(fullPath) {
		return fullPath
	}

	switch drive {
	case consts.UploadDriveLocal:
		return url.GetAddr(ctx) + "/" + fullPath
	case consts.UploadDriveUCloud:
		return conf.UCloudEndpoint + "/" + fullPath
	default:
		return fullPath
	}
}

// fileMeta 上传文件元数据
func (s *sCommonUpload) fileMeta(file *ghttp.UploadFile) (meta *sysin.UploadFileMeta, err error) {
	meta = new(sysin.UploadFileMeta)
	meta.Filename = file.Filename
	meta.Size = file.Size
	meta.Ext = f.Ext(file.Filename)
	meta.Kind = f.GetFileKind(meta.Ext)
	meta.MetaType, err = f.GetFileType(meta.Ext)
	if err != nil {
		return
	}

	// 兼容naiveUI
	naiveType := "text/plain"
	if f.IsImgType(f.Ext(file.Filename)) {
		naiveType = ""
	}
	meta.NaiveType = naiveType

	// 文件hash
	b, err := f.UploadFileByte(file)
	if err != nil {
		return
	}
	meta.Md5 = encrypt.Md5ToString(gconv.String(encrypt.Hash32(b)))
	return
}
