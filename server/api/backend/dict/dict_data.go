// Package dict
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package dict

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// DataEditReq 修改/新增字典数据
type DataEditReq struct {
	entity.SysDictData
	TypeID int64 `json:"typeID"  dc:"字典类型ID"`
	g.Meta `path:"/dict_data/edit" method:"post" tags:"字典数据" summary:"修改/新增字典数据"`
}

type DataEditRes struct{}

// DataDeleteReq 删除字典数据
type DataDeleteReq struct {
	Id     interface{} `json:"id" v:"required#字典数据ID不能为空" dc:"字典数据ID"`
	g.Meta `path:"/dict_data/delete" method:"post" tags:"字典数据" summary:"删除字典数据"`
}
type DataDeleteRes struct{}

// DataListReq 查询列表
type DataListReq struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	TypeID int64  `json:"typeId" v:"required#字典类型ID不能为空" dc:"字典类型ID"` //
	Type   string `json:"type"`
	Label  string `json:"label"`
	g.Meta `path:"/dict_data/list" method:"get" tags:"字典数据" summary:"获取字典数据列表"`
}

type DataListRes struct {
	List []*sysin.DictDataListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}
