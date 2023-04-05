// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package common

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/tencentyun/cos-go-sdk-v5"
	ufile "github.com/ufilesdk-dev/ufile-gosdk"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/encrypt"
	f "hotgo/utility/file"
	"hotgo/utility/format"
	utilityurl "hotgo/utility/url"
	"hotgo/utility/validate"
	"net/http"
	"net/url"
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
	case consts.UploadDriveCos:
		return s.UploadCOS(ctx, conf, file, meta)
	case consts.UploadDriveOss:
		return s.UploadOSS(ctx, conf, file, meta)
	case consts.UploadDriveQiNiu:
		return s.UploadQiNiu(ctx, conf, file, meta)
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
	case consts.UploadDriveCos:
		return s.UploadCOS(ctx, conf, file, meta)
	case consts.UploadDriveOss:
		return s.UploadOSS(ctx, conf, file, meta)
	case consts.UploadDriveQiNiu:
		return s.UploadQiNiu(ctx, conf, file, meta)
	default:
		return nil, gerror.Newf("暂不支持上传驱动:%v", conf.Drive)
	}
}

// UploadLocal 上传本地
func (s *sCommonUpload) UploadLocal(ctx context.Context, conf *model.UploadConfig, file *ghttp.UploadFile, meta *sysin.UploadFileMeta) (result *sysin.AttachmentListModel, err error) {
	if ok, err1 := s.HasFile(ctx, meta.Md5); ok || err1 != nil {
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
	if ok, err1 := s.HasFile(ctx, meta.Md5); ok || err1 != nil {
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

// UploadCOS 上传腾讯云对象存储
func (s *sCommonUpload) UploadCOS(ctx context.Context, conf *model.UploadConfig, file *ghttp.UploadFile, meta *sysin.UploadFileMeta) (result *sysin.AttachmentListModel, err error) {
	if ok, err1 := s.HasFile(ctx, meta.Md5); ok || err1 != nil {
		return
	}

	if conf.CosPath == "" {
		err = gerror.New("COS存储驱动必须配置存储路径!")
		return
	}

	nowDate := time.Now().Format("2006-01-02")
	fileName := gfile.Basename(file.Filename)
	fileName = strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	fileName = fileName + gfile.Ext(file.Filename)
	fullPath := conf.CosPath + nowDate + "/" + fileName

	// 流式上传本地小文件
	f2, err := file.Open()
	defer func() {
		_ = f2.Close()
	}()
	if err != nil {
		return nil, err
	}

	u, _ := url.Parse(conf.CosBucketURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.CosSecretId,
			SecretKey: conf.CosSecretKey,
		},
	})

	_, err = c.Object.Put(ctx, fullPath, f2, nil)
	if err != nil {
		return nil, err
	}

	attachment, err := service.SysAttachment().Add(ctx, meta, fullPath, consts.UploadDriveCos)
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

// UploadOSS 上传阿里云云对象存储
func (s *sCommonUpload) UploadOSS(ctx context.Context, conf *model.UploadConfig, file *ghttp.UploadFile, meta *sysin.UploadFileMeta) (result *sysin.AttachmentListModel, err error) {
	if ok, err1 := s.HasFile(ctx, meta.Md5); ok || err1 != nil {
		return
	}

	if conf.OssPath == "" {
		err = gerror.New("OSS存储驱动必须配置存储路径!")
		return
	}

	nowDate := time.Now().Format("2006-01-02")
	fileName := gfile.Basename(file.Filename)
	fileName = strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	fileName = fileName + gfile.Ext(file.Filename)
	fullPath := conf.OssPath + nowDate + "/" + fileName

	// 流式上传本地小文件
	f2, err := file.Open()
	defer func() {
		_ = f2.Close()
	}()
	if err != nil {
		return nil, err
	}

	client, err := oss.New(conf.OssEndpoint, conf.OssSecretId, conf.OssSecretKey)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(conf.OssBucket)
	if err != nil {
		return nil, err
	}

	if err = bucket.PutObject(fullPath, f2); err != nil {
		return nil, err
	}

	attachment, err := service.SysAttachment().Add(ctx, meta, fullPath, consts.UploadDriveOss)
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

// UploadQiNiu 上传七牛云对象存储
func (s *sCommonUpload) UploadQiNiu(ctx context.Context, conf *model.UploadConfig, file *ghttp.UploadFile, meta *sysin.UploadFileMeta) (result *sysin.AttachmentListModel, err error) {
	if ok, err1 := s.HasFile(ctx, meta.Md5); ok || err1 != nil {
		return
	}

	if conf.QiNiuPath == "" {
		err = gerror.New("七牛云存储驱动必须配置存储路径!")
		return
	}

	nowDate := time.Now().Format("2006-01-02")
	fileName := gfile.Basename(file.Filename)
	fileName = strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	fileName = fileName + gfile.Ext(file.Filename)
	fullPath := conf.QiNiuPath + nowDate + "/" + fileName

	// 流式上传本地小文件
	f2, err := file.Open()
	defer func() {
		_ = f2.Close()
	}()
	if err != nil {
		return nil, err
	}

	putPolicy := storage.PutPolicy{
		Scope: conf.QiNiuBucket,
	}
	token := putPolicy.UploadToken(qbox.NewMac(conf.QiNiuAccessKey, conf.QiNiuSecretKey))

	cfg := storage.Config{}
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 空间对应的机房
	cfg.Region, err = storage.GetRegion(conf.QiNiuAccessKey, conf.QiNiuBucket)
	if err != nil {
		return
	}

	if err = storage.NewFormUploader(&cfg).Put(ctx, &storage.PutRet{}, token, fullPath, f2, file.Size, &storage.PutExtra{}); err != nil {
		return
	}

	attachment, err := service.SysAttachment().Add(ctx, meta, fullPath, consts.UploadDriveQiNiu)
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
		return utilityurl.GetAddr(ctx) + "/" + fullPath
	case consts.UploadDriveUCloud:
		return conf.UCloudEndpoint + "/" + fullPath
	case consts.UploadDriveCos:
		return conf.CosBucketURL + "/" + fullPath
	case consts.UploadDriveOss:
		return conf.OssBucketURL + "/" + fullPath
	case consts.UploadDriveQiNiu:
		return conf.QiNiuDomain + "/" + fullPath
	default:
		return fullPath
	}
}

// HasFile 文件是否存在
func (s *sCommonUpload) HasFile(ctx context.Context, md5 string) (bool, error) {
	result, err := dao.SysAttachment.GetMd5File(ctx, md5)
	if err != nil {
		return false, err
	}

	return result != nil, nil
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
