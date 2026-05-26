// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProvinces is the golang structure for table sys_provinces.
type SysProvinces struct {
	Id        int64       `json:"id"        orm:"id"         description:"省市区ID"`
	Title     string      `json:"title"     orm:"title"      description:"栏目名称"`
	Pinyin    string      `json:"pinyin"    orm:"pinyin"     description:"拼音"`
	Lng       string      `json:"lng"       orm:"lng"        description:"经度"`
	Lat       string      `json:"lat"       orm:"lat"        description:"纬度"`
	Pid       int64       `json:"pid"       orm:"pid"        description:"父栏目"`
	Level     int         `json:"level"     orm:"level"      description:"关系树等级"`
	Tree      string      `json:"tree"      orm:"tree"       description:"关系"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`
	Status    int         `json:"status"    orm:"status"     description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
