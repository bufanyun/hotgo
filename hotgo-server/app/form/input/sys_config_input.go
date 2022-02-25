package input

import "github.com/bufanyun/hotgo/app/model/entity"

//  获取指定配置键的值
type SysConfigGetValueInp struct {
	Key string
}
type SysConfigGetValueModel struct {
	Value string
}

// 名称是否唯一
type SysConfigNameUniqueInp struct {
	Name string
	Id   int64
}

type SysConfigNameUniqueModel struct {
	IsUnique bool
}

// 最大排序
type SysConfigMaxSortInp struct {
	Id int64
}

type SysConfigMaxSortModel struct {
	Sort int
}

//  修改/新增字典数据
type SysConfigEditInp struct {
	entity.SysConfig
}
type SysConfigEditModel struct{}

//  删除字典类型
type SysConfigDeleteInp struct {
	Id interface{}
}
type SysConfigDeleteModel struct{}

// 获取信息
type SysConfigViewInp struct {
	Id string
}

type SysConfigViewModel struct {
	entity.SysConfig
}

//  获取列表
type SysConfigListInp struct {
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

type SysConfigListModel struct {
	entity.SysConfig
	DeptName string `json:"dept_name"`
	RoleName string `json:"role_name"`
}
