// Package form
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package form

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	page     int
	pageSize int
)

// DefaultPageSize 列表分页默认加载页码
func DefaultPageSize(ctx context.Context) int {
	if pageSize > 0 {
		return pageSize
	}
	pageSize = g.Cfg().MustGet(ctx, "hotgo.admin.defaultPageSize", 10).Int()
	if pageSize <= 0 {
		pageSize = 10
	}
	return pageSize
}

// DefaultPage 列表分页默认加载数量
func DefaultPage(ctx context.Context) int {
	if page > 0 {
		return page
	}
	page = g.Cfg().MustGet(ctx, "hotgo.admin.defaultPage", 1).Int()
	if page <= 0 {
		page = 1
	}
	return page
}

// PageReq 分页
type PageReq struct {
	Page    int `json:"page" example:"10" d:"1" v:"min:1#页码最小值不能低于1"  dc:"当前页码"`
	PerPage int `json:"pageSize" example:"1" d:"10" v:"min:1|max:200#每页数量最小值不能低于1|最大值不能大于200" dc:"每页数量"`
}
type PageRes struct {
	PageReq
	PageCount int `json:"pageCount" example:"0" dc:"全部数据量"`
}

// RangeDateReq 时间范围查询
type RangeDateReq struct {
	StartTime string `json:"start_time" v:"date#开始日期格式不正确"  dc:"开始日期"`
	EndTime   string `json:"end_time" v:"date#结束日期格式不正确" dc:"结束日期"`
}

// StatusReq 通用状态查询
type StatusReq struct {
	Status int `json:"status"  v:"in:-1,0,1,2,3#输入的状态是无效的" dc:"状态"`
}

// SwitchReq 更新开关状态
type SwitchReq struct {
	Key   string `json:"key" v:"required#测试ID不能为空" dc:"开关字段"`
	Value int    `json:"value" v:"in:1,2#输入的开关值是无效的" dc:"更新值"`
}

// CalPage 解析分页
func CalPage(ctx context.Context, page, perPage int) (newPage, newPerPage int, offset int) {
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

func CalPageCount(totalCount int, perPage int) int {
	return (totalCount + perPage - 1) / perPage
}

// Selects 选项
type Selects []*Select

type Select struct {
	Value interface{} `json:"value"`
	Label string      `json:"label"`
	Name  string      `json:"name"`
}

func (p Selects) Len() int {
	return len(p)
}

func (p Selects) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Selects) Less(i, j int) bool {
	return gconv.Int64(p[j].Value) > gconv.Int64(p[i].Value)
}

type SelectInt64s []*SelectInt64
type SelectInt64 struct {
	Value int64  `json:"value"`
	Label string `json:"label"`
	Name  string `json:"name"`
}

// DefaultMaxSort 默认最大排序
func DefaultMaxSort(ctx context.Context, baseSort int) int {
	return baseSort + g.Cfg().MustGet(ctx, "hotgo.admin.maxSortIncrement", 10).Int()
}

// AvatarGroup 头像组
type AvatarGroup struct {
	Name string `json:"name" dc:"姓名"`
	Src  string `json:"src" dc:"头像地址"`
}
