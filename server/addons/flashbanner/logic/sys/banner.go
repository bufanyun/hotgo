// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/addons/flashbanner/dao"
	"hotgo/addons/flashbanner/model/entity"
	"hotgo/addons/flashbanner/model/input/sysin"
	"hotgo/addons/flashbanner/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sSysBanner struct{}

func NewSysBanner() *sSysBanner {
	return &sSysBanner{}
}

func init() {
	service.RegisterSysBanner(NewSysBanner())
}

// GetMaxSort 获取最大排序值
func (s *sSysBanner) GetMaxSort(ctx context.Context) (maxSort int, err error) {
	var result struct {
		MaxSort int `json:"maxSort"`
	}
	err = dao.Banner.Ctx(ctx).Fields("MAX(sort) as maxSort").Scan(&result)
	if err != nil {
		return 0, gerror.Wrap(err, "获取最大排序值失败")
	}
	return result.MaxSort, nil
}

// Create 创建轮播图
func (s *sSysBanner) Create(ctx context.Context, in *sysin.BannerCreateInp) (err error) {
	// 如果没有设置排序值，自动设置为最大值+1
	if in.Sort == 0 {
		maxSort, err := s.GetMaxSort(ctx)
		if err != nil {
			return err
		}
		in.Sort = maxSort + 1
	}
	
	_, err = dao.Banner.Ctx(ctx).Data(g.Map{
		dao.Banner.Columns().Name:      in.Name,
		dao.Banner.Columns().Cover:     in.Cover,
		dao.Banner.Columns().Link:      in.Link,
		dao.Banner.Columns().Type:      in.Type,
		dao.Banner.Columns().Status:    1, // 默认启用
		dao.Banner.Columns().Sort:      in.Sort,
		dao.Banner.Columns().CreatedAt: gtime.Now(),
		dao.Banner.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	if err != nil {
		err = gerror.Wrap(err, "创建轮播图失败")
		return
	}
	return
}

// Edit 修改/新增轮播图
func (s *sSysBanner) Edit(ctx context.Context, in *sysin.BannerEditInp) (err error) {
	if in.Id > 0 {
		// 修改
		_, err = dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().Id, in.Id).Data(g.Map{
			dao.Banner.Columns().Name:      in.Name,
			dao.Banner.Columns().Cover:     in.Cover,
			dao.Banner.Columns().Link:      in.Link,
			dao.Banner.Columns().Type:      in.Type,
			dao.Banner.Columns().Status:    in.Status,
			dao.Banner.Columns().Sort:      in.Sort,
			dao.Banner.Columns().UpdatedAt: gtime.Now(),
		}).Update()
		if err != nil {
			err = gerror.Wrap(err, "修改轮播图失败")
			return
		}
	} else {
		// 新增
		_, err = dao.Banner.Ctx(ctx).Data(g.Map{
			dao.Banner.Columns().Name:      in.Name,
			dao.Banner.Columns().Cover:     in.Cover,
			dao.Banner.Columns().Link:      in.Link,
			dao.Banner.Columns().Type:      in.Type,
			dao.Banner.Columns().Status:    in.Status,
			dao.Banner.Columns().Sort:      in.Sort,
			dao.Banner.Columns().CreatedAt: gtime.Now(),
			dao.Banner.Columns().UpdatedAt: gtime.Now(),
		}).Insert()
		if err != nil {
			err = gerror.Wrap(err, "新增轮播图失败")
			return
		}
	}
	return
}

// Delete 删除轮播图
func (s *sSysBanner) Delete(ctx context.Context, in *sysin.BannerDeleteInp) (err error) {
	_, err = dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().Id, in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, "删除轮播图失败")
		return
	}
	return
}

// View 获取指定轮播图信息
func (s *sSysBanner) View(ctx context.Context, in *sysin.BannerViewInp) (res *sysin.BannerViewModel, err error) {
	var banner entity.Banner
	err = dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().Id, in.Id).Scan(&banner)
	if err != nil {
		err = gerror.Wrap(err, "获取轮播图信息失败")
		return
	}
	if banner.Id == 0 {
		err = gerror.New("轮播图不存在")
		return
	}
	res = &sysin.BannerViewModel{Banner: banner}
	return
}

// List 获取轮播图列表
func (s *sSysBanner) List(ctx context.Context, in *sysin.BannerListInp) (list []*sysin.BannerListModel, totalCount int, err error) {
	m := dao.Banner.Ctx(ctx)
	
	// 条件查询
	if in.Name != "" {
		m = m.WhereLike(dao.Banner.Columns().Name, "%"+in.Name+"%")
	}
	if in.Type > 0 {
		m = m.Where(dao.Banner.Columns().Type, in.Type)
	}
	
	var banners []*entity.Banner
	if in.Page > 0 && in.PerPage > 0 {
		err = m.Page(in.Page, in.PerPage).OrderDesc(dao.Banner.Columns().Sort).OrderDesc(dao.Banner.Columns().Id).ScanAndCount(&banners, &totalCount, false)
	} else {
		err = m.OrderDesc(dao.Banner.Columns().Sort).OrderDesc(dao.Banner.Columns().Id).ScanAndCount(&banners, &totalCount, false)
	}
	if err != nil {
		err = gerror.Wrap(err, "获取轮播图列表失败")
		return
	}
	
	list = make([]*sysin.BannerListModel, len(banners))
	for i, banner := range banners {
		list[i] = &sysin.BannerListModel{Banner: *banner}
	}
	return
}

// Status 更新轮播图状态
func (s *sSysBanner) Status(ctx context.Context, in *sysin.BannerStatusInp) (err error) {
	_, err = dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().Id, in.Id).Data(g.Map{
		dao.Banner.Columns().Status:    in.Status,
		dao.Banner.Columns().UpdatedAt: gtime.Now(),
	}).Update()
	if err != nil {
		err = gerror.Wrap(err, "更新轮播图状态失败")
		return
	}
	return
} 