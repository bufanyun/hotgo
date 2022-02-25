package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

//  获取指定配置键的值
type ConfigGetValueReq struct {
	Key    string `json:"key" v:"required#配置键不能为空"  description:"配置键"`
	g.Meta `path:"/config/get_value" method:"get" tags:"配置" summary:"获取指定配置键的值"`
}
type ConfigGetValueRes struct {
	Value string `json:"value" description:"配置值"`
}

//  名称是否唯一
type ConfigNameUniqueReq struct {
	Name   string `json:"name" v:"required#配置名称不能为空"  description:"配置名称"`
	Id     int64  `json:"id" description:"配置ID"`
	g.Meta `path:"/config/name_unique" method:"get" tags:"配置" summary:"配置名称是否唯一"`
}
type ConfigNameUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}

//  查询列表
type ConfigListReq struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name   string `json:"name"   description:"配置名称"`
	g.Meta `path:"/config/list" method:"get" tags:"配置" summary:"获取配置列表"`
}

type ConfigListRes struct {
	List []*input.SysConfigListModel `json:"list"   description:"数据列表"`
	form.PageRes
}

//  获取指定信息
type ConfigViewReq struct {
	Id     string `json:"id" v:"required#配置ID不能为空" description:"配置ID"`
	g.Meta `path:"/config/view" method:"get" tags:"配置" summary:"获取指定信息"`
}
type ConfigViewRes struct {
	*input.SysConfigViewModel
}

//  修改/新增
type ConfigEditReq struct {
	entity.SysConfig
	g.Meta `path:"/config/edit" method:"post" tags:"配置" summary:"修改/新增配置"`
}
type ConfigEditRes struct{}

//  删除
type ConfigDeleteReq struct {
	Id     interface{} `json:"id" v:"required#配置ID不能为空" description:"配置ID"`
	g.Meta `path:"/config/delete" method:"post" tags:"配置" summary:"删除配置"`
}
type ConfigDeleteRes struct{}

//  最大排序
type ConfigMaxSortReq struct {
	Id     int64 `json:"id" description:"配置ID"`
	g.Meta `path:"/config/max_sort" method:"get" tags:"配置" summary:"配置最大排序"`
}
type ConfigMaxSortRes struct {
	Sort int `json:"sort" description:"排序"`
}
