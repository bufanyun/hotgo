// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminin

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// PostListInp 获取列表
type PostListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name string
	Code string
}

type PostListModel struct {
	entity.AdminPost
}

// PostViewInp 获取信息
type PostViewInp struct {
	Id string
}

type PostViewModel struct {
	entity.AdminPost
}

// PostCodeUniqueInp 编码是否唯一
type PostCodeUniqueInp struct {
	Code string
	Id   int64
}

type PostCodeUniqueModel struct {
	IsUnique bool
}

// PostNameUniqueInp 名称是否唯一
type PostNameUniqueInp struct {
	Name string
	Id   int64
}

type PostNameUniqueModel struct {
	IsUnique bool
}

// PostMaxSortInp 最大排序
type PostMaxSortInp struct {
	Id int64
}

type PostMaxSortModel struct {
	Sort int
}

// PostEditInp 修改/新增字典数据
type PostEditInp struct {
	entity.AdminPost
}
type PostEditModel struct{}

// PostDeleteInp 删除字典类型
type PostDeleteInp struct {
	Id interface{}
}
type PostDeleteModel struct{}

// PostStatusInp 更新状态
type PostStatusInp struct {
	entity.AdminPost
}
type PostStatusModel struct{}
