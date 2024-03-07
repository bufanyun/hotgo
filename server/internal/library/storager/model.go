// Package storager
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package storager

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/model/entity"
)

// FileMeta 文件元数据
type FileMeta struct {
	Filename  string `json:"filename"`  // 文件名称
	Size      int64  `json:"size"`      // 文件大小
	Kind      string `json:"kind"`      // 文件上传类型
	MimeType  string `json:"mimeType"`  // 文件扩展类型
	NaiveType string `json:"naiveType"` // NaiveUI类型
	Ext       string `json:"ext"`       // 文件扩展名
	Md5       string `json:"md5"`       // 文件hash
}

// MultipartProgress 分片进度
type MultipartProgress struct {
	UploadId      string      `json:"uploadId"`      // 上传事件ID
	ThirdUploadId string      `json:"thirdUploadId"` // 第三方上传事件ID
	Meta          *FileMeta   `json:"meta"`          // 文件元数据
	ShardCount    int         `json:"shardCount"`    // 分片数量
	UploadedIndex []int       `json:"uploadedIndex"` // 已上传的分片索引
	CreatedAt     *gtime.Time `json:"createdAt"`     // 创建时间
}

// CheckMultipartParams 检查文件分片
type CheckMultipartParams struct {
	UploadType string `json:"uploadType"  dc:"文件类型"`
	FileName   string `json:"fileName"    dc:"文件名称"`
	Size       int64  `json:"size"        dc:"文件大小"`
	Md5        string `json:"md5"         dc:"文件md5值"`
	ShardCount int    `json:"shardCount"  dc:"分片数量"`
	meta       *FileMeta
}

type CheckMultipartModel struct {
	UploadId        string                `json:"uploadId"        dc:"上传事件ID"`
	Attachment      *entity.SysAttachment `json:"attachment"      dc:"附件"`
	WaitUploadIndex []int                 `json:"waitUploadIndex" dc:"等待上传的分片索引"`
	Progress        float64               `json:"progress"        dc:"上传进度"`
	SizeFormat      string                `json:"sizeFormat"      dc:"文件大小"`
}

// UploadPartParams 分片上传
type UploadPartParams struct {
	UploadId   string            `json:"uploadId"    dc:"上传事件ID"`
	UploadType string            `json:"uploadType"  dc:"文件类型"`
	FileName   string            `json:"fileName"    dc:"文件名称"`
	Size       int64             `json:"size"        dc:"文件大小"`
	Md5        string            `json:"md5"         dc:"文件md5值"`
	Index      int               `json:"index"       dc:"分片索引"`
	File       *ghttp.UploadFile `json:"file" type:"file" dc:"分片文件"`
	mp         *MultipartProgress
}

type UploadPartModel struct {
	Attachment *entity.SysAttachment `json:"attachment" dc:"附件"`
	Progress   float64               `json:"progress"   dc:"上传进度"`
	Finish     bool                  `json:"finish"     dc:"是否完成"`
}
