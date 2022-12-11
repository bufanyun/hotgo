// Package form
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package form

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	page     int64
	pageSize int64
)

// DefaultPageSize 列表分页默认加载页码
func DefaultPageSize(ctx context.Context) int64 {
	if pageSize > 0 {
		return pageSize
	}
	defaultPageSize, _ := g.Cfg().Get(ctx, "hotgo.admin.defaultPageSize", 10)
	pageSize = defaultPageSize.Int64()
	if pageSize <= 0 {
		pageSize = 10
	}
	return pageSize
}

// DefaultPage 列表分页默认加载数量
func DefaultPage(ctx context.Context) int64 {
	if page > 0 {
		return page
	}
	defaultPage, _ := g.Cfg().Get(ctx, "hotgo.admin.defaultPage", 1)
	page = defaultPage.Int64()
	if page <= 0 {
		page = 10
	}
	return page
}

// PageReq 分页
type PageReq struct {
	Page    int64 `json:"page" example:"10" d:"1" v:"min:1#页码最小值不能低于1"  dc:"当前页码"`
	PerPage int64 `json:"pageSize" example:"1" d:"10" v:"min:1|max:100#|每页数量最小值不能低于1|最大值不能大于100" dc:"每页数量"`
}
type PageRes struct {
	PageReq
	PageCount int64 `json:"pageCount" example:"0" dc:"全部数据量"`
}

// RangeDateReq 时间查询
type RangeDateReq struct {
	StartTime string `json:"start_time" v:"date#开始日期格式不正确"  dc:"开始日期"`
	EndTime   string `json:"end_time" v:"date#结束日期格式不正确" dc:"结束日期"`
}

// StatusReq 状态查询
type StatusReq struct {
	Status int `json:"status"  v:"in:-1,0,1,2,3#输入的状态是无效的" dc:"状态"`
}

// CalPage 解析分页
func CalPage(ctx context.Context, page, perPage int64) (newPage, newPerPage int64, offset int64) {
	if page <= 0 {
		newPage = DefaultPage(ctx)
	} else {
		newPage = page
	}
	if perPage <= 0 {
		newPerPage = DefaultPageSize(ctx)
	} else {
		newPerPage = perPage
	}

	offset = (newPage - 1) * newPerPage
	return
}

func CalPageCount(totalCount int64, perPage int64) int64 {
	return (totalCount + perPage - 1) / perPage
}
