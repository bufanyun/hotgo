//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

//  获取指定配置键的值
type ConfigGetValueReq struct {
	g.Meta `path:"/config/get_value" method:"get" tags:"配置" summary:"获取指定配置键的值"`
	Key    string `json:"key" v:"required#配置键不能为空"  dc:"配置键"`
}
type ConfigGetValueRes struct {
	Value string `json:"value" dc:"配置值"`
}

//  名称是否唯一
type ConfigNameUniqueReq struct {
	g.Meta `path:"/config/name_unique" method:"get" tags:"配置" summary:"配置名称是否唯一"`
	Name   string `json:"name" v:"required#配置名称不能为空"  dc:"配置名称"`
	Id     int64  `json:"id" dc:"配置ID"`
}
type ConfigNameUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

//  查询列表
type ConfigListReq struct {
	g.Meta `path:"/config/list" method:"get" tags:"配置" summary:"获取配置列表"`
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name string `json:"name"   dc:"配置名称"`
}

type ConfigListRes struct {
	List []*input.SysConfigListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

//  获取指定信息
type ConfigViewReq struct {
	g.Meta `path:"/config/view" method:"get" tags:"配置" summary:"获取指定信息"`
	Id     string `json:"id" v:"required#配置ID不能为空" dc:"配置ID"`
}
type ConfigViewRes struct {
	*input.SysConfigViewModel
}

//  修改/新增
type ConfigEditReq struct {
	g.Meta `path:"/config/edit" method:"post" tags:"配置" summary:"修改/新增配置"`
	entity.SysConfig
}
type ConfigEditRes struct{}

//  删除
type ConfigDeleteReq struct {
	g.Meta `path:"/config/delete" method:"post" tags:"配置" summary:"删除配置"`
	Id     interface{} `json:"id" v:"required#配置ID不能为空" dc:"配置ID"`
}
type ConfigDeleteRes struct{}

//  最大排序
type ConfigMaxSortReq struct {
	g.Meta `path:"/config/max_sort" method:"get" tags:"配置" summary:"配置最大排序"`
	Id     int64 `json:"id" dc:"配置ID"`
}
type ConfigMaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}
