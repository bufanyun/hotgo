// Package storager
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package storager

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/consts"
	"hotgo/internal/library/cache"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
	"hotgo/utility/convert"
	"hotgo/utility/format"
	"hotgo/utility/url"
	"hotgo/utility/validate"
	"strconv"
	"strings"
	"time"
)

// UploadDrive 存储驱动
type UploadDrive interface {
	// Upload 上传
	Upload(ctx context.Context, file *ghttp.UploadFile) (fullPath string, err error)
	// CreateMultipart 创建分片事件
	CreateMultipart(ctx context.Context, in *CheckMultipartParams) (res *MultipartProgress, err error)
	// UploadPart 上传分片
	UploadPart(ctx context.Context, in *UploadPartParams) (res *UploadPartModel, err error)
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
	case consts.UploadDriveMinio:
		drive = &MinioDrive{}
	default:
		panic(fmt.Sprintf("暂不支持的存储驱动:%v", driveType))
	}
	return drive
}

// DoUpload 上传入口
func DoUpload(ctx context.Context, typ string, file *ghttp.UploadFile) (result *entity.SysAttachment, err error) {
	if file == nil {
		err = gerror.New("文件必须!")
		return
	}

	meta, err := GetFileMeta(file)
	if err != nil {
		return
	}

	if err = ValidateFileMeta(typ, meta); err != nil {
		return
	}

	result, err = HasFile(ctx, meta.Md5)
	if err != nil {
		return
	}

	// 相同存储相同身份才复用
	if result != nil && result.Drive == config.Drive && result.MemberId == contexts.GetUserId(ctx) && result.AppId == contexts.GetModule(ctx) {
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

// ValidateFileMeta 验证文件元数据
func ValidateFileMeta(typ string, meta *FileMeta) (err error) {
	if _, err = GetFileMimeType(meta.Ext); err != nil {
		return
	}

	switch typ {
	case KindImg:
		if !IsImgType(meta.Ext) {
			err = gerror.New("上传的文件不是图片")
			return
		}
		if config.ImageSize > 0 && meta.Size > config.ImageSize*1024*1024 {
			err = gerror.Newf("图片大小不能超过%vMB", config.ImageSize)
			return
		}

		if len(config.ImageType) > 0 && !validate.InSlice(strings.Split(config.ImageType, `,`), meta.Ext) {
			err = gerror.New("上传图片类型未经允许")
			return
		}
	case KindDoc:
		if !IsDocType(meta.Ext) {
			err = gerror.New("上传的文件不是文档")
			return
		}
	case KindAudio:
		if !IsAudioType(meta.Ext) {
			err = gerror.New("上传的文件不是音频")
			return
		}
	case KindVideo:
		if !IsVideoType(meta.Ext) {
			err = gerror.New("上传的文件不是视频")
			return
		}
	case KindZip:
		if !IsZipType(meta.Ext) {
			err = gerror.New("上传的文件不是压缩文件")
			return
		}
	case KindOther:
		fallthrough
	default:
		// 默认为通用的文件上传
		if config.FileSize > 0 && meta.Size > config.FileSize*1024*1024 {
			err = gerror.Newf("文件大小不能超过%vMB", config.FileSize)
			return
		}

		if len(config.FileType) > 0 && !validate.InSlice(strings.Split(config.FileType, `,`), meta.Ext) {
			err = gerror.New("上传文件类型未经允许")
			return
		}
	}
	return
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
	case consts.UploadDriveMinio:
		return fmt.Sprintf("%s/%s/%s", config.MinioDomain, config.MinioBucket, fullPath)
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
	meta.MimeType, err = GetFileMimeType(meta.Ext)
	if err != nil {
		return
	}

	// 兼容naiveUI
	naiveType := "text/plain"
	if IsImgType(Ext(file.Filename)) {
		naiveType = ""
	}
	meta.NaiveType = naiveType

	// 计算md5值
	meta.Md5, err = CalcFileMd5(file)
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
		MimeType:  meta.MimeType,
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

// HasFile 检查附件是否存在
func HasFile(ctx context.Context, md5 string) (res *entity.SysAttachment, err error) {
	if err = GetModel(ctx).Where("md5", md5).Scan(&res); err != nil {
		err = gerror.Wrap(err, "检查文件hash时出现错误")
		return
	}

	if res == nil {
		return
	}

	// 只有在上传时才会检查md5值，如果附件存在则更新最后上传时间，保证上传列表更新显示在最前面
	if res.Id > 0 {
		update := g.Map{
			"status":     consts.StatusEnabled,
			"updated_at": gtime.Now(),
		}
		_, _ = GetModel(ctx).WherePri(res.Id).Data(update).Update()
	}
	return
}

// CheckMultipart 检查文件分片
func CheckMultipart(ctx context.Context, in *CheckMultipartParams) (res *CheckMultipartModel, err error) {
	res = new(CheckMultipartModel)

	meta := new(FileMeta)
	meta.Filename = in.FileName
	meta.Size = in.Size
	meta.Ext = Ext(in.FileName)
	meta.Kind = GetFileKind(meta.Ext)
	meta.MimeType, err = GetFileMimeType(meta.Ext)
	if err != nil {
		return
	}

	// 兼容naiveUI
	naiveType := "text/plain"
	if IsImgType(Ext(in.FileName)) {
		naiveType = ""
	}
	meta.NaiveType = naiveType
	meta.Md5 = in.Md5

	if err = ValidateFileMeta(in.UploadType, meta); err != nil {
		return
	}

	result, err := HasFile(ctx, in.Md5)
	if err != nil {
		return nil, err
	}

	// 文件已存在，直接返回。相同存储相同身份才复用
	if result != nil && result.Drive == config.Drive && result.MemberId == contexts.GetUserId(ctx) && result.AppId == contexts.GetModule(ctx) {
		res.Attachment = result
		return
	}

	for i := 0; i < in.ShardCount; i++ {
		res.WaitUploadIndex = append(res.WaitUploadIndex, i+1)
	}

	in.meta = meta
	progress, err := GetOrCreateMultipartProgress(ctx, in)
	if err != nil {
		return nil, err
	}

	if len(progress.UploadedIndex) > 0 {
		res.WaitUploadIndex = convert.DifferenceSlice(progress.UploadedIndex, res.WaitUploadIndex)
	}

	if len(res.WaitUploadIndex) == 0 {
		res.WaitUploadIndex = make([]int, 0)
	}
	res.UploadId = progress.UploadId
	res.Progress = CalcUploadProgress(progress.UploadedIndex, progress.ShardCount)
	res.SizeFormat = format.FileSize(progress.Meta.Size)
	return
}

// CalcUploadProgress 计算上传进度
func CalcUploadProgress(uploadedIndex []int, shardCount int) float64 {
	return format.Round2Float64(float64(len(uploadedIndex)) / float64(shardCount) * 100)
}

// GenUploadId 生成上传ID
func GenUploadId(ctx context.Context, md5 string) string {
	return fmt.Sprintf("%v:%v:%v@%v", md5, contexts.GetUserId(ctx), contexts.GetModule(ctx), config.Drive)
}

// GetOrCreateMultipartProgress 获取或创建分片上传事件进度
func GetOrCreateMultipartProgress(ctx context.Context, in *CheckMultipartParams) (res *MultipartProgress, err error) {
	uploadId := GenUploadId(ctx, in.Md5)
	res, err = GetMultipartProgress(ctx, uploadId)
	if err != nil {
		return nil, err
	}
	if res != nil {
		return res, nil
	}
	return New(config.Drive).CreateMultipart(ctx, in)
}

// GetMultipartProgress 获取分片上传事件进度
func GetMultipartProgress(ctx context.Context, uploadId string) (res *MultipartProgress, err error) {
	key := fmt.Sprintf("%v:%v", consts.CacheMultipartUpload, uploadId)
	get, err := cache.Instance().Get(ctx, key)
	if err != nil {
		return nil, err
	}
	err = get.Scan(&res)
	return
}

// CreateMultipartProgress 创建分片上传事件进度
func CreateMultipartProgress(ctx context.Context, in *MultipartProgress) (err error) {
	key := fmt.Sprintf("%v:%v", consts.CacheMultipartUpload, in.UploadId)
	return cache.Instance().Set(ctx, key, in, time.Hour*24*7)
}

// UpdateMultipartProgress 更新分片上传事件进度
func UpdateMultipartProgress(ctx context.Context, in *MultipartProgress) (err error) {
	key := fmt.Sprintf("%v:%v", consts.CacheMultipartUpload, in.UploadId)
	return cache.Instance().Set(ctx, key, in, time.Hour*24*7)
}

// DelMultipartProgress 删除分片上传事件进度
func DelMultipartProgress(ctx context.Context, in *MultipartProgress) (err error) {
	key := fmt.Sprintf("%v:%v", consts.CacheMultipartUpload, in.UploadId)
	_, err = cache.Instance().Remove(ctx, key)
	return
}

// UploadPart 上传分片
func UploadPart(ctx context.Context, in *UploadPartParams) (res *UploadPartModel, err error) {
	in.mp, err = GetMultipartProgress(ctx, in.UploadId)
	if err != nil {
		return nil, err
	}
	if in.mp == nil {
		err = gerror.New("分片事件不存在，请重新上传！")
		return
	}

	if validate.InSlice(in.mp.UploadedIndex, in.Index) {
		err = gerror.New("该分片已上传过了")
		return
	}

	res, err = New(config.Drive).UploadPart(ctx, in)
	if err != nil {
		return nil, err
	}
	return
}
