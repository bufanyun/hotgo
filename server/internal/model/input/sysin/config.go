// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

type GetConfigItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// UpdateConfigInp 更新指定分组的配置
type UpdateConfigInp struct {
	Group string `json:"group"`
	List  g.Map  `json:"list"`
}

// GetConfigInp 获取指定分组的配置
type GetConfigInp struct {
	Group string `json:"group"`
}
type GetConfigModel struct {
	List g.Map `json:"list"`
}

// ConfigGetValueInp 获取指定配置键的值
type ConfigGetValueInp struct {
	Key string
}
type ConfigGetValueModel struct {
	Value string
}

// ConfigNameUniqueInp 名称是否唯一
type ConfigNameUniqueInp struct {
	Name string
	Id   int64
}

type ConfigNameUniqueModel struct {
	IsUnique bool
}

// ConfigMaxSortInp 最大排序
type ConfigMaxSortInp struct {
	Id int64
}

type ConfigMaxSortModel struct {
	Sort int
}

// ConfigEditInp 修改/新增字典数据
type ConfigEditInp struct {
	entity.SysConfig
}
type ConfigEditModel struct{}

// ConfigDeleteInp 删除字典类型
type ConfigDeleteInp struct {
	Id interface{}
}
type ConfigDeleteModel struct{}

// ConfigViewInp 获取信息
type ConfigViewInp struct {
	Id string
}

type ConfigViewModel struct {
	entity.SysConfig
}

// ConfigListInp 获取列表
type ConfigListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name      string
	Code      string
	DeptId    int
	Mobile    int
	Username  string
	Realname  string
	StartTime string
	EndTime   string
}

type ConfigListModel struct {
	entity.SysConfig
	DeptName string `json:"deptName"`
	RoleName string `json:"roleName"`
}
