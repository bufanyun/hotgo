//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

/************************ 字典数据 *****************/

//  数据键值是否唯一
type DictDataUniqueReq struct {
	g.Meta `path:"/dict_data/unique" method:"get" tags:"字典" summary:"数据键值是否唯一"`
	Value  string `json:"value" v:"required#数据键值不能为空"  dc:"数据键值"`
	Type   string `json:"type" example:"sys_common_status" v:"required#字典类型不能为空"  dc:"字典类型"`
	Id     int64  `json:"id" dc:"字典数据ID"`
}
type DictDataUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

//  查询字典数据最大排序
type DictDataMaxSortReq struct {
	g.Meta `path:"/dict_data/max_sort" method:"get" tags:"字典" summary:"查询字典数据最大排序"`
	Type   string `json:"type" example:"sys_common_status" v:"required#字典类型不能为空"  dc:"字典类型"`
}
type DictDataMaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

//  修改/新增字典数据
type DictDataEditReq struct {
	g.Meta `path:"/dict_data/edit" method:"post" tags:"字典" summary:"修改/新增字典数据"`
	entity.SysDictData
}
type DictDataEditRes struct{}

//  删除字典类型
type DictDataDeleteReq struct {
	g.Meta `path:"/dict_data/delete" method:"post" tags:"字典" summary:"删除字典数据"`
	Id     interface{} `json:"id" v:"required#字典数据ID不能为空" dc:"字典数据ID"`
}
type DictDataDeleteRes struct{}

//  获取指定字典数据信息
type DictDataViewReq struct {
	g.Meta `path:"/dict_data/view" method:"get" tags:"字典" summary:"获取指定字典数据信息"`
	Id     string `json:"id" v:"required#字典数据ID不能为空" dc:"字典数据ID"`
}
type DictDataViewRes struct {
	*entity.SysDictData
}

//  获取字典数据列表
type DictDataListReq struct {
	g.Meta `path:"/dict_data/list" method:"get" tags:"字典" summary:"获取字典数据列表"`
	form.PageReq
	Type string `json:"type" example:"sys_common_status" v:"required#字典类型不能为空"  dc:"字典类型"`
}
type DictDataListRes struct {
	List []*entity.SysDictData `json:"list"   dc:"数据列表"`
	form.PageRes
}

// 获取指定字典类型的属性数据
type DictAttributeReq struct {
	g.Meta `path:"/dict/attribute" method:"get" tags:"字典" summary:"获取指定字典类型的属性数据"`
	Type   string `json:"type" example:"sys_common_status" v:"required#字典类型不能为空"  dc:"字典类型"`
}
type DictAttributeRes []*entity.SysDictData

/************************ 字典类型 *****************/

//  修改/新增字典类型
type DictTypeExportReq struct {
	g.Meta `path:"/dict_type/export" method:"get" tags:"字典" summary:"导出字典类型"`
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name string `json:"name" dc:"字典名称"`
	Type string `json:"type" dc:"字典类型"`
}
type DictTypeExportRes struct{}

//  刷新字典缓存
type DictTypeRefreshCacheReq struct {
	g.Meta `path:"/dict_type/refresh_cache" method:"get" tags:"字典" summary:"刷新字典缓存"`
}
type DictTypeRefreshCacheRes struct{}

//  获取字典类型列表
type DictTypeListReq struct {
	g.Meta `path:"/dict_type/list" method:"get" tags:"字典" summary:"获取字典类型列表"`
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name string `json:"name" dc:"字典名称"`
	Type string `json:"type" dc:"字典类型"`
}
type DictTypeListRes struct {
	List []*entity.SysDictType `json:"list"   dc:"数据列表"`
	form.PageRes
}

//  修改/新增字典类型
type DictTypeEditReq struct {
	g.Meta `path:"/dict_type/edit" method:"post" tags:"字典" summary:"修改/新增字典类型"`
	entity.SysDictType
}
type DictTypeEditRes struct{}

//  删除字典类型
type DictTypeDeleteReq struct {
	g.Meta `path:"/dict_type/delete" method:"post" tags:"字典" summary:"删除字典类型"`
	Id     interface{} `json:"id" v:"required#字典类型ID不能为空" dc:"字典类型ID"`
}
type DictTypeDeleteRes struct{}

//  获取指定字典类型信息
type DictTypeViewReq struct {
	g.Meta `path:"/dict_type/view" method:"get" tags:"字典" summary:"获取指定字典类型信息"`
	Id     string `json:"id" v:"required#字典类型ID不能为空" dc:"字典类型ID"`
}
type DictTypeViewRes struct {
	*entity.SysDictType
}

//  类型是否唯一
type DictTypeUniqueReq struct {
	g.Meta `path:"/dict_type/unique" method:"get" tags:"字典" summary:"类型是否唯一"`
	Type   string `json:"type" example:"sys_common_status" v:"required#字典类型不能为空"  dc:"字典类型"`
	Id     int64  `json:"id" dc:"字典类型ID"`
}
type DictTypeUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}
