package apiForm

import (
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// 获取指定字典类型的属性数据
type DictAttributeReq struct {
	Type   string `json:"type" example:"sys_common_status" v:"required#字典类型不能为空"  description:"字典类型"`
	g.Meta `path:"/dict/attribute" method:"get" tags:"字典" summary:"获取指定字典类型的属性数据"`
}
type DictAttributeRes []*entity.SysDictData
