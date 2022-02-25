package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

/************************ 字典数据 *****************/

//  数据键值是否唯一
type DictDataUniqueReq struct {
	Value  string `json:"value" v:"required#数据键值不能为空"  description:"数据键值"`
	Type   string `json:"type" example:"sys_common_status" v:"required#字典类型不能为空"  description:"字典类型"`
	Id     int64  `json:"id" description:"字典数据ID"`
	g.Meta `path:"/dict_data/unique" method:"get" tags:"字典" summary:"数据键值是否唯一"`
}
type DictDataUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}

//  查询字典数据最大排序
type DictDataMaxSortReq struct {
	Type   string `json:"type" example:"sys_common_status" v:"required#字典类型不能为空"  description:"字典类型"`
	g.Meta `path:"/dict_data/max_sort" method:"get" tags:"字典" summary:"查询字典数据最大排序"`
}
type DictDataMaxSortRes struct {
	Sort int `json:"sort" description:"排序"`
}

//  修改/新增字典数据
type DictDataEditReq struct {
	entity.SysDictData
	g.Meta `path:"/dict_data/edit" method:"post" tags:"字典" summary:"修改/新增字典数据"`
}
type DictDataEditRes struct{}

//  删除字典类型
type DictDataDeleteReq struct {
	Id     interface{} `json:"id" v:"required#字典数据ID不能为空" description:"字典数据ID"`
	g.Meta `path:"/dict_data/delete" method:"post" tags:"字典" summary:"删除字典数据"`
}
type DictDataDeleteRes struct{}

//  获取指定字典数据信息
type DictDataViewReq struct {
	Id     string `json:"id" v:"required#字典数据ID不能为空" description:"字典数据ID"`
	g.Meta `path:"/dict_data/view" method:"get" tags:"字典" summary:"获取指定字典数据信息"`
}
type DictDataViewRes struct {
	*entity.SysDictData
}

//  获取字典数据列表
type DictDataListReq struct {
	form.PageReq
	Type   string `json:"type" example:"sys_common_status" v:"required#字典类型不能为空"  description:"字典类型"`
	g.Meta `path:"/dict_data/list" method:"get" tags:"字典" summary:"获取字典数据列表"`
}
type DictDataListRes struct {
	List []*entity.SysDictData `json:"list"   description:"数据列表"`
	form.PageRes
}

// 获取指定字典类型的属性数据
type DictAttributeReq struct {
	Type   string `json:"type" example:"sys_common_status" v:"required#字典类型不能为空"  description:"字典类型"`
	g.Meta `path:"/dict/attribute" method:"get" tags:"字典" summary:"获取指定字典类型的属性数据"`
}
type DictAttributeRes []*entity.SysDictData

/************************ 字典类型 *****************/

//  修改/新增字典类型
type DictTypeExportReq struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name   string `json:"name" description:"字典名称"`
	Type   string `json:"type" description:"字典类型"`
	g.Meta `path:"/dict_type/export" method:"get" tags:"字典" summary:"导出字典类型"`
}
type DictTypeExportRes struct{}

//  刷新字典缓存
type DictTypeRefreshCacheReq struct {
	g.Meta `path:"/dict_type/refresh_cache" method:"get" tags:"字典" summary:"刷新字典缓存"`
}
type DictTypeRefreshCacheRes struct{}

//  获取字典类型列表
type DictTypeListReq struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Name   string `json:"name" description:"字典名称"`
	Type   string `json:"type" description:"字典类型"`
	g.Meta `path:"/dict_type/list" method:"get" tags:"字典" summary:"获取字典类型列表"`
}
type DictTypeListRes struct {
	List []*entity.SysDictType `json:"list"   description:"数据列表"`
	form.PageRes
}

//  修改/新增字典类型
type DictTypeEditReq struct {
	entity.SysDictType
	g.Meta `path:"/dict_type/edit" method:"post" tags:"字典" summary:"修改/新增字典类型"`
}
type DictTypeEditRes struct{}

//  删除字典类型
type DictTypeDeleteReq struct {
	Id     interface{} `json:"id" v:"required#字典类型ID不能为空" description:"字典类型ID"`
	g.Meta `path:"/dict_type/delete" method:"post" tags:"字典" summary:"删除字典类型"`
}
type DictTypeDeleteRes struct{}

//  获取指定字典类型信息
type DictTypeViewReq struct {
	Id     string `json:"id" v:"required#字典类型ID不能为空" description:"字典类型ID"`
	g.Meta `path:"/dict_type/view" method:"get" tags:"字典" summary:"获取指定字典类型信息"`
}
type DictTypeViewRes struct {
	*entity.SysDictType
}

//  类型是否唯一
type DictTypeUniqueReq struct {
	Type   string `json:"type" example:"sys_common_status" v:"required#字典类型不能为空"  description:"字典类型"`
	Id     int64  `json:"id" description:"字典类型ID"`
	g.Meta `path:"/dict_type/unique" method:"get" tags:"字典" summary:"类型是否唯一"`
}
type DictTypeUniqueRes struct {
	IsUnique bool `json:"is_unique" description:"是否唯一"`
}
