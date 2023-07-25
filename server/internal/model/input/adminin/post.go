// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package adminin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"
)

// PostListInp 获取列表
type PostListInp struct {
	form.PageReq
	form.StatusReq
	Name      string        `json:"name"             dc:"岗位名称"`
	Code      string        `json:"code"             dc:"岗位编码"`
	CreatedAt []*gtime.Time `json:"createdAt"        dc:"创建时间"`
}

func (in *PostListInp) Filter(ctx context.Context) (err error) {
	return
}

type PostListModel struct {
	entity.AdminPost
}

// PostViewInp 获取信息
type PostViewInp struct {
	Id string `json:"id" v:"required#岗位ID不能为空" description:"岗位ID"`
}

type PostViewModel struct {
	entity.AdminPost
}

// PostMaxSortInp 最大排序
type PostMaxSortInp struct {
	Id int64 `json:"id" description:"岗位ID"`
}

type PostMaxSortModel struct {
	Sort int `json:"sort" description:"排序"`
}

// PostEditInp 修改/新增字典数据
type PostEditInp struct {
	entity.AdminPost
}

func (in *PostEditInp) Filter(ctx context.Context) (err error) {
	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return
	}
	if in.Code == "" {
		err = gerror.New("编码不能为空")
		return
	}
	return
}

type PostEditModel struct{}

// PostDeleteInp 删除字典类型
type PostDeleteInp struct {
	Id interface{} `json:"id" v:"required#岗位ID不能为空" description:"岗位ID"`
}

type PostDeleteModel struct{}

// PostStatusInp 更新状态
type PostStatusInp struct {
	entity.AdminPost
}

func (in *PostStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type PostStatusModel struct{}
