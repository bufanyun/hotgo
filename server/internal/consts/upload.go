// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

// 上传类型
const (
	UploadTypeFile  = 1 // 文件
	UploadTypeImage = 2 // 图片
	UploadTypeDoc   = 3 // 文档
	UploadTypeAudio = 4 // 音频
	UploadTypeVideo = 5 // 视频
)

// 上传存储驱动
const (
	UploadDriveLocal  = "local"  // 本地驱动
	UploadDriveUCloud = "ucloud" // ucloud对象存储
	UploadDriveCos    = "cos"    // 腾讯云cos
	UploadDriveOss    = "oss"    // 阿里云oss
	UploadDriveQiNiu  = "qiniu"  // 七牛云对象存储
)
