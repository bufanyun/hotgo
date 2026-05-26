// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCreditsLog is the golang structure for table admin_credits_log.
type AdminCreditsLog struct {
	Id          int64       `json:"id"          orm:"id"           description:"变动ID"`
	MemberId    int64       `json:"memberId"    orm:"member_id"    description:"管理员ID"`
	AppId       string      `json:"appId"       orm:"app_id"       description:"应用id"`
	AddonsName  string      `json:"addonsName"  orm:"addons_name"  description:"插件名称"`
	CreditType  string      `json:"creditType"  orm:"credit_type"  description:"变动类型"`
	CreditGroup string      `json:"creditGroup" orm:"credit_group" description:"变动组别"`
	BeforeNum   float64     `json:"beforeNum"   orm:"before_num"   description:"变动前"`
	Num         float64     `json:"num"         orm:"num"          description:"变动数据"`
	AfterNum    float64     `json:"afterNum"    orm:"after_num"    description:"变动后"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
	Ip          string      `json:"ip"          orm:"ip"           description:"操作人IP"`
	MapId       int64       `json:"mapId"       orm:"map_id"       description:"关联ID"`
	Status      int         `json:"status"      orm:"status"       description:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"修改时间"`
}
