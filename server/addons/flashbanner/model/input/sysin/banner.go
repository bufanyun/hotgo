// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"context"
	"hotgo/addons/flashbanner/model/entity"
	"hotgo/internal/model/input/form"
	"github.com/gogf/gf/v2/errors/gerror"
)

// BannerCreateInp 创建轮播图
type BannerCreateInp struct {
	Name  string `json:"name" v:"required#轮播图名称不能为空" dc:"轮播图名称"`
	Cover string `json:"cover" v:"required#轮播图封面不能为空" dc:"轮播图封面"`
	Link  string `json:"link" dc:"轮播图链接"`
	Type  int    `json:"type" dc:"轮播图类型"`
	Sort  int    `json:"sort" dc:"排序，数字越大越靠前"`
}

func (in *BannerCreateInp) Filter(ctx context.Context) (err error) {
	return
}

type BannerCreateModel struct{}

// BannerEditInp 修改/新增轮播图
type BannerEditInp struct {
	entity.Banner
}

func (in *BannerEditInp) Filter(ctx context.Context) (err error) {
	if in.Name == "" {
		err = gerror.New("轮播图名称不能为空")
		return
	}
	if in.Cover == "" {
		err = gerror.New("轮播图封面不能为空")
		return
	}
	return
}

type BannerEditModel struct{}

// BannerDeleteInp 删除轮播图
type BannerDeleteInp struct {
	Id interface{} `json:"id" v:"required#轮播图ID不能为空" dc:"轮播图ID"`
}

func (in *BannerDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type BannerDeleteModel struct{}

// BannerViewInp 获取指定轮播图信息
type BannerViewInp struct {
	Id int64 `json:"id" v:"required#轮播图ID不能为空" dc:"轮播图ID"`
}

func (in *BannerViewInp) Filter(ctx context.Context) (err error) {
	return
}

type BannerViewModel struct {
	entity.Banner
}

// BannerListInp 获取轮播图列表
type BannerListInp struct {
	form.PageReq
	Name string `json:"name" dc:"轮播图名称"`
	Type int    `json:"type" dc:"轮播图类型"`
}

func (in *BannerListInp) Filter(ctx context.Context) (err error) {
	return
}

type BannerListModel struct {
	entity.Banner
}

// BannerStatusInp 更新轮播图状态
type BannerStatusInp struct {
	Id     int64 `json:"id" v:"required#轮播图ID不能为空" dc:"轮播图ID"`
	Status int   `json:"status" v:"required#状态不能为空" dc:"状态"`
}

func (in *BannerStatusInp) Filter(ctx context.Context) (err error) {
	return
}

type BannerStatusModel struct{} 