// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonHgexampleTable is the golang structure for table addon_hgexample_table.
type AddonHgexampleTable struct {
	Id          int64       `json:"id"          orm:"id"          description:"ID"`
	Pid         int64       `json:"pid"         orm:"pid"         description:"上级ID"`
	Level       int         `json:"level"       orm:"level"       description:"树等级"`
	Tree        string      `json:"tree"        orm:"tree"        description:"关系树"`
	CategoryId  int64       `json:"categoryId"  orm:"category_id" description:"分类ID"`
	Flag        *gjson.Json `json:"flag"        orm:"flag"        description:"标签"`
	Title       string      `json:"title"       orm:"title"       description:"标题"`
	Description string      `json:"description" orm:"description" description:"描述"`
	Content     string      `json:"content"     orm:"content"     description:"内容"`
	Image       string      `json:"image"       orm:"image"       description:"单图"`
	Images      *gjson.Json `json:"images"      orm:"images"      description:"多图"`
	Attachfile  string      `json:"attachfile"  orm:"attachfile"  description:"附件"`
	Attachfiles *gjson.Json `json:"attachfiles" orm:"attachfiles" description:"多附件"`
	Map         *gjson.Json `json:"map"         orm:"map"         description:"动态键值对"`
	Star        float64     `json:"star"        orm:"star"        description:"推荐星"`
	Price       float64     `json:"price"       orm:"price"       description:"价格"`
	Views       int64       `json:"views"       orm:"views"       description:"浏览次数"`
	ActivityAt  *gtime.Time `json:"activityAt"  orm:"activity_at" description:"活动时间"`
	StartAt     *gtime.Time `json:"startAt"     orm:"start_at"    description:"开启时间"`
	EndAt       *gtime.Time `json:"endAt"       orm:"end_at"      description:"结束时间"`
	Switch      int         `json:"switch"      orm:"switch"      description:"开关"`
	Sort        int         `json:"sort"        orm:"sort"        description:"排序"`
	Avatar      string      `json:"avatar"      orm:"avatar"      description:"头像"`
	Sex         int         `json:"sex"         orm:"sex"         description:"性别"`
	Qq          string      `json:"qq"          orm:"qq"          description:"qq"`
	Email       string      `json:"email"       orm:"email"       description:"邮箱"`
	Mobile      string      `json:"mobile"      orm:"mobile"      description:"手机号码"`
	Hobby       *gjson.Json `json:"hobby"       orm:"hobby"       description:"爱好"`
	Channel     int         `json:"channel"     orm:"channel"     description:"渠道"`
	CityId      int64       `json:"cityId"      orm:"city_id"     description:"所在城市"`
	Remark      string      `json:"remark"      orm:"remark"      description:"备注"`
	Status      int         `json:"status"      orm:"status"      description:"状态"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"  description:"创建者"`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"  description:"更新者"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"修改时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`
}
