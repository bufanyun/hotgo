// Package log
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package log

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ClearReq 清空日志
type ClearReq struct {
	g.Meta `path:"/log/clear" method:"post" tags:"日志" summary:"清空日志"`
}

type ClearRes struct{}

// ExportReq 导出
type ExportReq struct {
	g.Meta `path:"/log/export" method:"get" tags:"日志" summary:"导出日志"`
	sysin.LogListInp
}

type ExportRes struct{}

// ListReq 获取菜单列表
type ListReq struct {
	g.Meta `path:"/log/list" method:"get" tags:"日志" summary:"获取日志列表"`
	sysin.LogListInp
}

type ListRes struct {
	List []*sysin.LogListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// DeleteReq 删除
type DeleteReq struct {
	g.Meta `path:"/log/delete" method:"post" tags:"日志" summary:"删除日志"`
	sysin.LogDeleteInp
}

type DeleteRes struct{}

// ViewReq 获取指定信息
type ViewReq struct {
	g.Meta `path:"/log/view" method:"get" tags:"日志" summary:"获取指定信息"`
	sysin.LogViewInp
}

type ViewRes struct {
	*sysin.LogViewModel
}
