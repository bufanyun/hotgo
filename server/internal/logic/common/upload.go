// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package common

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/encrypt"
	f "hotgo/utility/file"
	"hotgo/utility/format"
	"hotgo/utility/validate"
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
	return s.UploadLocal(ctx, file, meta)
}

// UploadLocal 上传本地
func (s *sCommonUpload) UploadLocal(ctx context.Context, file *ghttp.UploadFile, meta *sysin.UploadFileMeta) (result *sysin.AttachmentListModel, err error) {
	result, err = dao.SysAttachment.GetMd5File(ctx, meta.Md5)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if result != nil {
		return
	}

	var (
		value, _ = g.Cfg().Get(ctx, "server.serverRoot")
		nowDate  = time.Now().Format("2006-01-02")
	)

	if value.IsEmpty() {
		err = gerror.New("本地上传驱动必须配置静态路径!")
		return
	}

	// 包含静态文件夹的路径
	fullDirPath := strings.Trim(value.String(), "/") + "/attachment/" + nowDate
	fileName, err := file.Save(fullDirPath, true)
	if err != nil {
		return
	}
	// 不含静态文件夹的路径
	fullPath := "attachment/" + nowDate + "/" + fileName

	attachment, err := service.SysAttachment().Add(ctx, meta, fullPath, consts.UploadDriveLocal)
	if err != nil {
		return nil, err
	}

	attachment.FileUrl = s.LastUrl(ctx, attachment.FileUrl, attachment.Drive)
	result = &sysin.AttachmentListModel{
		SysAttachment: *attachment,
		SizeFormat:    format.FileSize(attachment.Size),
	}
	return
}

// LastUrl 根据驱动获取最终文件访问地址
func (s *sCommonUpload) LastUrl(ctx context.Context, fullPath, drive string) string {
	if validate.IsURL(fullPath) {
		return fullPath
	}

	if drive == consts.UploadDriveLocal {
		return fmt.Sprintf("http://%s/", g.RequestFromCtx(ctx).Host) + "/" + fullPath
	}

	return fullPath
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
