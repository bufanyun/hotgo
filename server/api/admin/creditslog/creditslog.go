// Package creditslog
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.5.3
// @AutoGenerate Date 2023-04-15 15:59:58
package creditslog

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// ListReq 查询资产变动列表
type ListReq struct {
	g.Meta `path:"/creditsLog/list" method:"get" tags:"资产变动" summary:"获取资产变动列表"`
	adminin.CreditsLogListInp
}

type ListRes struct {
	form.PageRes
	List []*adminin.CreditsLogListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出资产变动列表
type ExportReq struct {
	g.Meta `path:"/creditsLog/export" method:"get" tags:"资产变动" summary:"导出资产变动列表"`
	adminin.CreditsLogListInp
}

type ExportRes struct{}
