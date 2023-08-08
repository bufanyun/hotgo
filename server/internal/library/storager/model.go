// Package storager
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package storager

// FileMeta 文件元数据
type FileMeta struct {
	Filename  string // 文件名称
	Size      int64  // 文件大小
	Kind      string // 文件上传类型
	MimeType  string // 文件扩展类型
	NaiveType string // NaiveUI类型
	Ext       string // 文件扩展名
	Md5       string // 文件hash
}
