// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminin

import (
	"hotgo/internal/model/entity"
)

// DeptNameUniqueInp 名称是否唯一
type DeptNameUniqueInp struct {
	Name string
	Id   int64
}

type DeptNameUniqueModel struct {
	IsUnique bool
}

// DeptMaxSortInp 最大排序
type DeptMaxSortInp struct {
	Id int64
}

type DeptMaxSortModel struct {
	Sort int
}

// DeptEditInp 修改/新增字典数据
type DeptEditInp struct {
	entity.AdminDept
}
type DeptEditModel struct{}

// DeptDeleteInp 删除字典类型
type DeptDeleteInp struct {
	Id interface{}
}
type DeptDeleteModel struct{}

// DeptViewInp 获取信息
type DeptViewInp struct {
	Id int64
}

type DeptViewModel struct {
	entity.AdminDept
}

// DeptListInp 获取列表
type DeptListInp struct {
	Name string
	Code string
}

// DeptTree 树
type DeptTree struct {
	entity.AdminDept
	Label    string      `json:"label"     dc:"标签"`
	Value    int64       `json:"value"     dc:"键值"`
	Children []*DeptTree `json:"children"`
}

type DeptListModel struct {
	List []*DeptTree `json:"list"`
}

// DeptStatusInp 更新部门状态
type DeptStatusInp struct {
	entity.AdminDept
}
type DeptStatusModel struct{}
