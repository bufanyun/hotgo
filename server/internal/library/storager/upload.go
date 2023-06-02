package storager

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
	"hotgo/utility/encrypt"
	"hotgo/utility/url"
	"hotgo/utility/validate"
	"strconv"
	"strings"
)

// UploadDrive 存储驱动
type UploadDrive interface {
	// Upload 上传
	Upload(ctx context.Context, file *ghttp.UploadFile) (fullPath string, err error)
}

// New 初始化存储驱动
func New(name ...string) UploadDrive {
	var (
		driveType = consts.UploadDriveLocal
		drive     UploadDrive
	)

	if len(name) > 0 && name[0] != "" {
		driveType = name[0]
	}

	switch driveType {
	case consts.UploadDriveLocal:
		drive = &LocalDrive{}
	case consts.UploadDriveUCloud:
		drive = &UCloudDrive{}
	case consts.UploadDriveCos:
		drive = &CosDrive{}
	case consts.UploadDriveOss:
		drive = &OssDrive{}
	case consts.UploadDriveQiNiu:
		drive = &QiNiuDrive{}
	default:
		panic(fmt.Sprintf("暂不支持的存储驱动:%v", driveType))
	}
	return drive
}

// DoUpload 上传入口
func DoUpload(ctx context.Context, file *ghttp.UploadFile, typ int) (result *entity.SysAttachment, err error) {
	if file == nil {
		err = gerror.New("文件必须!")
		return
	}

	meta, err := GetFileMeta(file)
	if err != nil {
		return
	}

	if _, err = GetFileType(meta.Ext); err != nil {
		return
	}

	switch typ {
	case consts.UploadTypeFile:
		if config.FileSize > 0 && meta.Size > config.FileSize*1024*1024 {
			err = gerror.Newf("文件大小不能超过%vMB", config.FileSize)
			return
		}
	case consts.UploadTypeImage:
		if !IsImgType(meta.Ext) {
			err = gerror.New("上传的文件不是图片")
			return
		}
		if config.ImageSize > 0 && meta.Size > config.ImageSize*1024*1024 {
			err = gerror.Newf("图片大小不能超过%vMB", config.ImageSize)
			return
		}
	case consts.UploadTypeDoc:
		if !IsDocType(meta.Ext) {
			err = gerror.New("上传的文件不是文档")
			return
		}
	case consts.UploadTypeAudio:
		if !IsAudioType(meta.Ext) {
			err = gerror.New("上传的文件不是音频")
			return
		}
	case consts.UploadTypeVideo:
		if !IsVideoType(meta.Ext) {
			err = gerror.New("上传的文件不是视频")
			return
		}
	default:
		err = gerror.Newf("无效的上传类型:%v", typ)
		return
	}

	result, err = hasFile(ctx, meta.Md5)
	if err != nil {
		return
	}

	if result != nil {
		return
	}

	// 上传到驱动
	fullPath, err := New(config.Drive).Upload(ctx, file)
	if err != nil {
		return
	}
	// 写入附件记录
	return write(ctx, meta, fullPath)
}

// LastUrl 根据驱动获取最终文件访问地址
func LastUrl(ctx context.Context, fullPath, drive string) string {
	if validate.IsURL(fullPath) {
		return fullPath
	}

	switch drive {
	case consts.UploadDriveLocal:
		return url.GetAddr(ctx) + "/" + fullPath
	case consts.UploadDriveUCloud:
		return config.UCloudEndpoint + "/" + fullPath
	case consts.UploadDriveCos:
		return config.CosBucketURL + "/" + fullPath
	case consts.UploadDriveOss:
		return config.OssBucketURL + "/" + fullPath
	case consts.UploadDriveQiNiu:
		return config.QiNiuDomain + "/" + fullPath
	default:
		return fullPath
	}
}

// GetFileMeta 获取上传文件元数据
func GetFileMeta(file *ghttp.UploadFile) (meta *FileMeta, err error) {
	meta = new(FileMeta)
	meta.Filename = file.Filename
	meta.Size = file.Size
	meta.Ext = Ext(file.Filename)
	meta.Kind = GetFileKind(meta.Ext)
	meta.MetaType, err = GetFileType(meta.Ext)
	if err != nil {
		return
	}

	// 兼容naiveUI
	naiveType := "text/plain"
	if IsImgType(Ext(file.Filename)) {
		naiveType = ""
	}
	meta.NaiveType = naiveType

	// 文件hash
	b, err := UploadFileByte(file)
	if err != nil {
		return
	}
	meta.Md5 = encrypt.Md5ToString(gconv.String(encrypt.Hash32(b)))
	return
}

// GenFullPath 根据目录和文件类型生成一个绝对地址
func GenFullPath(basePath, ext string) string {
	fileName := strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6)
	fileName = fileName + ext
	return basePath + gtime.Date() + "/" + strings.ToLower(fileName)
}

// write 写入附件记录
func write(ctx context.Context, meta *FileMeta, fullPath string) (models *entity.SysAttachment, err error) {
	models = &entity.SysAttachment{
		Id:        0,
		AppId:     contexts.GetModule(ctx),
		MemberId:  contexts.GetUserId(ctx),
		Drive:     config.Drive,
		Size:      meta.Size,
		Path:      fullPath,
		FileUrl:   fullPath,
		Name:      meta.Filename,
		Kind:      meta.Kind,
		MetaType:  meta.MetaType,
		NaiveType: meta.NaiveType,
		Ext:       meta.Ext,
		Md5:       meta.Md5,
		Status:    consts.StatusEnabled,
	}

	id, err := GetModel(ctx).Data(models).InsertAndGetId()
	if err != nil {
		return
	}
	models.Id = id
	return
}

// hasFile 检查附件是否存在
func hasFile(ctx context.Context, md5 string) (res *entity.SysAttachment, err error) {
	if err = GetModel(ctx).Where("md5", md5).Scan(&res); err != nil {
		err = gerror.Wrap(err, "检查文件hash时出现错误")
		return
	}

	if res == nil {
		return
	}

	// 只有在上传时才会检查md5值，如果附件存在则更新最后上传时间，保证上传列表更新显示在最前面
	if res.Id > 0 {
		_, _ = GetModel(ctx).WherePri(res.Id).Data(g.Map{
			"status":     consts.StatusEnabled,
			"updated_at": gtime.Now(),
		}).Update()
	}
	return
}
