// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminOauth is the golang structure for table admin_oauth.
type AdminOauth struct {
	Id           int64       `json:"id"           orm:"id"            description:"主键"`
	MemberId     int64       `json:"memberId"     orm:"member_id"     description:"用户ID"`
	Unionid      string      `json:"unionid"      orm:"unionid"       description:"唯一ID"`
	OauthClient  string      `json:"oauthClient"  orm:"oauth_client"  description:"授权组别"`
	OauthOpenid  string      `json:"oauthOpenid"  orm:"oauth_openid"  description:"授权开放ID"`
	Sex          int         `json:"sex"          orm:"sex"           description:"性别"`
	Nickname     string      `json:"nickname"     orm:"nickname"      description:"昵称"`
	HeadPortrait string      `json:"headPortrait" orm:"head_portrait" description:"头像"`
	Birthday     *gtime.Time `json:"birthday"     orm:"birthday"      description:"生日"`
	Country      string      `json:"country"      orm:"country"       description:"国家"`
	Province     string      `json:"province"     orm:"province"      description:"省"`
	City         string      `json:"city"         orm:"city"          description:"市"`
	Status       int         `json:"status"       orm:"status"        description:"状态"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"修改时间"`
}
