// Package banner
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package banner

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/flashbanner/model/input/sysin"
	"hotgo/internal/model/input/form"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/banner/list" method:"get" tags:"轮播图" summary:"获取轮播图列表"`
	sysin.BannerListInp
}

type ListRes struct {
	List []*sysin.BannerListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取指定信息
type ViewReq struct {
	g.Meta `path:"/banner/view" method:"get" tags:"轮播图" summary:"获取指定轮播图信息"`
	sysin.BannerViewInp
}

type ViewRes struct {
	*sysin.BannerViewModel
}
