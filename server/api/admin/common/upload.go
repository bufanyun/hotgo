// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/sysin"
)

// UploadFileReq 上传文件
type UploadFileReq struct {
	g.Meta `path:"/upload/file" tags:"上传" method:"post" summary:"上传附件"`
}

type UploadFileRes *sysin.AttachmentListModel
