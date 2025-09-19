// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Banner is the golang structure for table banner.
type Banner struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Name      string      `json:"name"      orm:"name"       description:"轮播图名称"`
	Cover     string      `json:"cover"     orm:"cover"      description:"图片URL"`
	Link      string      `json:"link"      orm:"link"       description:"跳转链接，小程序内用相对地址"`
	Type      int         `json:"type"      orm:"type"       description:"类型默认不传"`
	Status    uint        `json:"status"    orm:"status"     description:"1可用"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序，数字越大越靠前"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
