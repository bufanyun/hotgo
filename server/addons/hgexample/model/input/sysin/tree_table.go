// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// TreeTableListInp 获取列表
type TreeTableListInp struct {
	form.PageReq
	form.Sorters
	Id         int64         `json:"id"          description:""`
	Flag       *gjson.Json   `json:"flag"        description:"标签"`
	Title      string        `json:"title"       description:"标题"`
	Content    string        `json:"content"     description:"内容"`
	Price      []float64     `json:"price"       description:"价格"`
	ActivityAt *gtime.Time   `json:"activityAt"  description:"活动时间"`
	Switch     int           `json:"switch"      description:"开关"`
	Hobby      *gjson.Json   `json:"hobby"       description:"爱好"`
	Status     int           `json:"status"      description:"状态"`
	CreatedAt  []*gtime.Time `json:"createdAt"   description:"创建时间"`
	Pid        int64         `json:"pid"         description:"上级ID"`
}

type TreeTableListModel struct {
	entity.AddonHgexampleTable
	TableCategoryName        string `json:"TableCategoryName" description:"分类名称"`
	TableCategoryDescription string `json:"TableCategoryDescription" description:"分类描述"`
	TableCategoryRemark      string `json:"TableCategoryRemark" description:"分类备注"`
	SysProvincesTitle        string `json:"sysProvincesTitle" description:""`
}

// TableTree 树
type TableTree struct {
	entity.AddonHgexampleTable
	Key      int64        `json:"key"       dc:"key"`
	Label    string       `json:"label"     dc:"标签"`
	Value    int64        `json:"value"     dc:"键值"`
	Children []*TableTree `json:"children"`
}
