//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package apiForm

import (
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// 获取指定字典类型的属性数据
type DictAttributeReq struct {
	g.Meta `path:"/dict/attribute" method:"get" tags:"字典接口" summary:"获取指定字典类型的属性数据"`
	Type   string `json:"type" example:"sys_common_status" v:"required#字典类型不能为空"  dc:"字典类型"`
}
type DictAttributeRes []*entity.SysDictData
