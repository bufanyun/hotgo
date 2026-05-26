// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminDept is the golang structure for table admin_dept.
type AdminDept struct {
	Id        int64       `json:"id"        orm:"id"         description:"部门ID"`
	Pid       int64       `json:"pid"       orm:"pid"        description:"父部门ID"`
	Name      string      `json:"name"      orm:"name"       description:"部门名称"`
	Code      string      `json:"code"      orm:"code"       description:"部门编码"`
	Type      string      `json:"type"      orm:"type"       description:"部门类型"`
	Leader    string      `json:"leader"    orm:"leader"     description:"负责人"`
	Phone     string      `json:"phone"     orm:"phone"      description:"联系电话"`
	Email     string      `json:"email"     orm:"email"      description:"邮箱"`
	Level     int         `json:"level"     orm:"level"      description:"关系树等级"`
	Tree      string      `json:"tree"      orm:"tree"       description:"关系树"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`
	Status    int         `json:"status"    orm:"status"     description:"部门状态"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
