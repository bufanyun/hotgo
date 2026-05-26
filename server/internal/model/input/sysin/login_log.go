// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// LoginLogDeleteInp 删除登录日志
type LoginLogDeleteInp struct {
	Id interface{} `json:"id" v:"required#日志ID不能为空" dc:"日志ID"`
}

type LoginLogDeleteModel struct{}

// LoginLogViewInp 获取指定登录日志信息
type LoginLogViewInp struct {
	Id int64 `json:"id" v:"required#日志ID不能为空" dc:"日志ID"`
}

type LoginLogViewModel struct {
	entity.SysLoginLog
}

// LoginLogListInp 获取登录日志列表
type LoginLogListInp struct {
	form.PageReq
	Username string        `json:"username"  dc:"用户名"`
	Status   int           `json:"status"    dc:"状态"`
	LoginAt  []*gtime.Time `json:"loginAt"   dc:"登录时间"`
	LoginIp  string        `json:"loginIp"   dc:"登录IP"`
}

type LoginLogListModel struct {
	entity.SysLoginLog
	Os        string    `json:"os"`
	Browser   string    `json:"browser"`
	CityLabel string    `json:"cityLabel"`
	SysLogId  gdb.Value `json:"sysLogId"`
}

func (in *LoginLogListInp) Filter(ctx context.Context) (err error) {
	return
}

// LoginLogExportModel 导出登录日志
type LoginLogExportModel struct {
	Id         int64       `json:"id"         orm:"id"          description:"日志ID"`
	ReqId      string      `json:"reqId"      orm:"req_id"      description:"请求ID"`
	MemberId   int64       `json:"memberId"   orm:"member_id"   description:"用户ID"`
	Username   string      `json:"username"   orm:"username"    description:"用户名"`
	Response   *gjson.Json `json:"response"   orm:"response"    description:"响应数据"`
	LoginAt    *gtime.Time `json:"loginAt"    orm:"login_at"    description:"登录时间"`
	LoginIp    string      `json:"loginIp"    orm:"login_ip"    description:"登录IP"`
	ProvinceId int64       `json:"provinceId" orm:"province_id" description:"省编码"`
	CityId     int64       `json:"cityId"     orm:"city_id"     description:"市编码"`
	UserAgent  string      `json:"userAgent"  orm:"user_agent"  description:"UA信息"`
	ErrMsg     string      `json:"errMsg"     orm:"err_msg"     description:"错误提示"`
	Status     int         `json:"status"     orm:"status"      description:"状态"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"修改时间"`
	Os         string      `json:"os"`
	Browser    string      `json:"browser"`
}

// LoginLogPushInp 解推送登录日志
type LoginLogPushInp struct {
	Response *adminin.LoginModel
	Err      error
}
