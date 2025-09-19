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

// CreateReq 创建轮播图
type CreateReq struct {
	g.Meta `path:"/banner/create" method:"post" tags:"轮播图" summary:"创建轮播图"`
	sysin.BannerCreateInp
}

type CreateRes struct{}

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

// EditReq 修改/新增轮播图
type EditReq struct {
	g.Meta `path:"/banner/edit" method:"post" tags:"轮播图" summary:"修改/新增轮播图"`
	sysin.BannerEditInp
}

type EditRes struct{}

// DeleteReq 删除轮播图
type DeleteReq struct {
	g.Meta `path:"/banner/delete" method:"post" tags:"轮播图" summary:"删除轮播图"`
	sysin.BannerDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新轮播图状态
type StatusReq struct {
	g.Meta `path:"/banner/status" method:"post" tags:"轮播图" summary:"更新轮播图状态"`
	sysin.BannerStatusInp
}

type StatusRes struct{} 