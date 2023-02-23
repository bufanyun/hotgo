// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"context"
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
	LoginAt  []*gtime.Time `json:"loginAt" dc:"登录时间"`
	SysLogIp string        `json:"sysLogIp"  dc:"IP地址"`
}

type LoginLogListModel struct {
	Id               int64       `json:"id"               dc:"日志ID"`
	ReqId            string      `json:"reqId"            dc:"请求ID"`
	MemberId         int64       `json:"memberId"         dc:"用户ID"`
	Username         string      `json:"username"         dc:"用户名"`
	LoginAt          *gtime.Time `json:"loginAt"          dc:"登录时间"`
	ErrMsg           string      `json:"errMsg"           dc:"错误提示"`
	Status           int         `json:"status"           dc:"状态"`
	CreatedAt        *gtime.Time `json:"createdAt"        dc:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        dc:"修改时间"`
	SysLogId         int64       `json:"sysLogId"         dc:"日志ID"`
	SysLogIp         string      `json:"sysLogIp"         dc:"IP地址"`
	SysLogProvinceId int64       `json:"sysLogProvinceId" dc:"省编码"`
	SysLogCityId     int64       `json:"sysLogCityId"     dc:"市编码"`
	SysLogErrorCode  int         `json:"sysLogErrorCode"  dc:"报错code"`
	SysLogUserAgent  string      `json:"sysLogUserAgent"  dc:"UA信息"`
	CityLabel        string      `json:"cityLabel"        dc:"城市标签"`
	Os               string      `json:"os"               dc:"系统信息"`
	Browser          string      `json:"browser"          dc:"浏览器信息"`
}

func (in *LoginLogListInp) Filter(ctx context.Context) (err error) {
	return
}

// LoginLogExportModel 导出登录日志
type LoginLogExportModel struct {
	Id               int64       `json:"id"               dc:"日志ID"`
	ReqId            string      `json:"reqId"            dc:"请求ID"`
	MemberId         int64       `json:"memberId"         dc:"用户ID"`
	Username         string      `json:"username"         dc:"用户名"`
	LoginAt          int64       `json:"loginAt"          dc:"登录时间"`
	ErrMsg           string      `json:"errMsg"           dc:"错误提示"`
	Status           int         `json:"status"           dc:"状态"`
	CreatedAt        *gtime.Time `json:"createdAt"        dc:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        dc:"修改时间"`
	SysLogIp         string      `json:"sysLogIp"         dc:"IP地址"`
	SysLogProvinceId int64       `json:"sysLogProvinceId" dc:"省编码"`
	SysLogCityId     int64       `json:"sysLogCityId"     dc:"市编码"`
	SysLogErrorCode  int         `json:"sysLogErrorCode"  dc:"报错code"`
	SysLogUserAgent  string      `json:"sysLogUserAgent"  dc:"UA信息"`
}

// LoginLogPushInp 解推送登录日志
type LoginLogPushInp struct {
	Input    adminin.MemberLoginInp
	Response *adminin.MemberLoginModel
	Err      error
}
