// Package servelog
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package servelog

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询服务日志列表
type ListReq struct {
	g.Meta `path:"/serveLog/list" method:"get" tags:"服务日志" summary:"获取服务日志列表"`
	sysin.ServeLogListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.ServeLogListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出服务日志列表
type ExportReq struct {
	g.Meta `path:"/serveLog/export" method:"get" tags:"服务日志" summary:"导出服务日志列表"`
	sysin.ServeLogListInp
}

type ExportRes struct{}

// ViewReq 获取服务日志指定信息
type ViewReq struct {
	g.Meta `path:"/serveLog/view" method:"get" tags:"服务日志" summary:"获取服务日志指定信息"`
	sysin.ServeLogViewInp
}

type ViewRes struct {
	*sysin.ServeLogViewModel
}

// DeleteReq 删除服务日志
type DeleteReq struct {
	g.Meta `path:"/serveLog/delete" method:"post" tags:"服务日志" summary:"删除服务日志"`
	sysin.ServeLogDeleteInp
}

type DeleteRes struct{}
