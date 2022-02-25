package input

import (
	"github.com/bufanyun/hotgo/app/model/entity"
)

//  获取列表
type AdminPostListInp struct {
	Page   int
	Limit  int
	Name   string
	Code   string
	Status int
}

type AdminPostListModel struct {
	entity.AdminPost
}

// 获取信息
type AdminPostViewInp struct {
	Id string
}

type AdminPostViewModel struct {
	entity.AdminPost
}

// 编码是否唯一
type AdminPostCodeUniqueInp struct {
	Code string
	Id   int64
}

type AdminPostCodeUniqueModel struct {
	IsUnique bool
}

// 名称是否唯一
type AdminPostNameUniqueInp struct {
	Name string
	Id   int64
}

type AdminPostNameUniqueModel struct {
	IsUnique bool
}

// 最大排序
type AdminPostMaxSortInp struct {
	Id int64
}

type AdminPostMaxSortModel struct {
	Sort int
}

//  修改/新增字典数据
type AdminPostEditInp struct {
	entity.AdminPost
}
type AdminPostEditModel struct{}

//  删除字典类型
type AdminPostDeleteInp struct {
	Id interface{}
}
type AdminPostDeleteModel struct{}
