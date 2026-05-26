// Package form
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package form

import (
	"hotgo/internal/consts"
)

type ReqPageFunc interface {
	GetPage() int
	GetPerPage() int
}

// PageReq 分页请求
type PageReq struct {
	Page       int  `json:"page" example:"10" d:"1" v:"page@min:1#页码最小值不能低于1"  dc:"当前页码"`
	PerPage    int  `json:"pageSize" example:"1" d:"10" v:"pageSize@min:1|max:200#每页数量最小值不能低于1|每页数量最大值不能大于200" dc:"每页数量"`
	Pagination bool `json:"pagination" d:"true" dc:"是否需要进行分页"`
}

// GetPage 获取当前页码
func (req *PageReq) GetPage() int {
	return req.Page
}

// GetPerPage 获取每页数量
func (req *PageReq) GetPerPage() int {
	return req.PerPage
}

// PageRes 分页响应
type PageRes struct {
	PageReq
	PageCount  int `json:"pageCount" example:"0"  dc:"分页个数"`
	TotalCount int `json:"totalCount" example:"0" dc:"数据总行数"`
}

// Pack 打包分页数据
func (res *PageRes) Pack(req ReqPageFunc, totalCount int) {
	res.TotalCount = totalCount
	res.PageCount = CalPageCount(totalCount, req.GetPerPage())
	res.Page = req.GetPage()
	res.PerPage = req.GetPerPage()
}

func CalPageCount(totalCount int, perPage int) int {
	return (totalCount + perPage - 1) / perPage
}

// CalPage 计算分页偏移量
func CalPage(page, perPage int) (newPage, newPerPage int, offset int) {
	if page <= 0 {
		newPage = consts.DefaultPage
	} else {
		newPage = page
	}
	if perPage <= 0 {
		newPerPage = consts.DefaultPageSize
	} else {
		newPerPage = perPage
	}

	offset = (newPage - 1) * newPerPage
	return
}
