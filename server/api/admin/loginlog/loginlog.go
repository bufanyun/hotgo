// Package loginlog
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package loginlog

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询登录日志列表
type ListReq struct {
	g.Meta `path:"/loginLog/list" method:"get" tags:"登录日志" summary:"获取登录日志列表"`
	sysin.LoginLogListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.LoginLogListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出登录日志列表
type ExportReq struct {
	g.Meta `path:"/loginLog/export" method:"get" tags:"登录日志" summary:"导出登录日志列表"`
	sysin.LoginLogListInp
}

type ExportRes struct{}

// ViewReq 获取登录日志指定信息
type ViewReq struct {
	g.Meta `path:"/loginLog/view" method:"get" tags:"登录日志" summary:"获取登录日志指定信息"`
	sysin.LoginLogViewInp
}

type ViewRes struct {
	*sysin.LoginLogViewModel
}

// DeleteReq 删除登录日志
type DeleteReq struct {
	g.Meta `path:"/loginLog/delete" method:"post" tags:"登录日志" summary:"删除登录日志"`
	sysin.LoginLogDeleteInp
}

type DeleteRes struct{}
