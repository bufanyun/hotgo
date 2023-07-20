// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCreditsLog is the golang structure for table admin_credits_log.
type AdminCreditsLog struct {
	Id          int64       `json:"id"          description:"变动ID"`
	MemberId    int64       `json:"memberId"    description:"管理员ID"`
	AppId       string      `json:"appId"       description:"应用id"`
	AddonsName  string      `json:"addonsName"  description:"插件名称"`
	CreditType  string      `json:"creditType"  description:"变动类型"`
	CreditGroup string      `json:"creditGroup" description:"变动组别"`
	BeforeNum   float64     `json:"beforeNum"   description:"变动前"`
	Num         float64     `json:"num"         description:"变动数据"`
	AfterNum    float64     `json:"afterNum"    description:"变动后"`
	Remark      string      `json:"remark"      description:"备注"`
	Ip          string      `json:"ip"          description:"操作人IP"`
	MapId       int64       `json:"mapId"       description:"关联ID"`
	Status      int         `json:"status"      description:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"修改时间"`
}
