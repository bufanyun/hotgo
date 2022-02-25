package input

import "github.com/bufanyun/hotgo/app/model/entity"

// 名称是否唯一
type AdminDeptNameUniqueInp struct {
	Name string
	Id   int64
}

type AdminDeptNameUniqueModel struct {
	IsUnique bool
}

// 最大排序
type AdminDeptMaxSortInp struct {
	Id int64
}

type AdminDeptMaxSortModel struct {
	Sort int
}

//  修改/新增字典数据
type AdminDeptEditInp struct {
	entity.AdminDept
}
type AdminDeptEditModel struct{}

//  删除字典类型
type AdminDeptDeleteInp struct {
	Id interface{}
}
type AdminDeptDeleteModel struct{}

// 获取信息
type AdminDeptViewInp struct {
	Id int64
}

type AdminDeptViewModel struct {
	entity.AdminDept
}

//  获取列表
type AdminDeptListInp struct {
	Name string
}

//  树
type AdminDeptTreeDept struct {
	entity.AdminDept
	Children []*AdminDeptTreeDept `json:"children"`
}

type AdminDeptListModel AdminDeptTreeDept

//  获取列表树
type AdminDeptListTreeInp struct {
	Name string
}

//  树
type AdminDeptListTreeDept struct {
	Id       int64                    `json:"id" `
	Key      int64                    `json:"key" `
	Pid      int64                    `json:"pid"  `
	Label    string                   `json:"label"`
	Title    string                   `json:"title"`
	Name     string                   `json:"name"`
	Type     string                   `json:"type"`
	Children []*AdminDeptListTreeDept `json:"children"`
}

type AdminDeptListTreeModel AdminDeptListTreeDept
