package input

import "github.com/bufanyun/hotgo/app/model/entity"

// 名称是否唯一
type AdminNoticeNameUniqueInp struct {
	Title string
	Id    int64
}

type AdminNoticeNameUniqueModel struct {
	IsUnique bool
}

// 最大排序
type AdminNoticeMaxSortInp struct {
	Id int64
}

type AdminNoticeMaxSortModel struct {
	Sort int
}

//  修改/新增字典数据
type AdminNoticeEditInp struct {
	entity.AdminNotice
}
type AdminNoticeEditModel struct{}

//  删除字典类型
type AdminNoticeDeleteInp struct {
	Id interface{}
}
type AdminNoticeDeleteModel struct{}

// 获取信息
type AdminNoticeViewInp struct {
	Id string
}

type AdminNoticeViewModel struct {
	entity.AdminNotice
}

//  获取列表
type AdminNoticeListInp struct {
	Page      int
	Limit     int
	Name      string
	Code      string
	DeptId    int
	Mobile    int
	Username  string
	Realname  string
	StartTime string
	EndTime   string
	Status    int
}

type AdminNoticeListModel struct {
	entity.AdminNotice
	DeptName string `json:"dept_name"`
	RoleName string `json:"role_name"`
}
