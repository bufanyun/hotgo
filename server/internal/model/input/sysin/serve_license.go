// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.7.6
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"
)

// ServeLicenseUpdateFields 修改服务许可证字段过滤
type ServeLicenseUpdateFields struct {
	Group       string `json:"group"       dc:"分组"`
	Name        string `json:"name"        dc:"许可名称"`
	Appid       string `json:"appid"       dc:"应用ID"`
	SecretKey   string `json:"secretKey"   dc:"应用秘钥"`
	OnlineLimit int    `json:"onlineLimit" dc:"在线数量限制，默认1"`
	//Routes      *gjson.Json `json:"routes"      dc:"路由表，空使用默认分组路由"`
	AllowedIps string      `json:"allowedIps"  dc:"白名单，*代表所有，只有允许的IP才能连接到tcp服务"`
	EndAt      *gtime.Time `json:"endAt"       dc:"授权结束时间"`
	Remark     string      `json:"remark"      dc:"备注"`
	Status     int         `json:"status"      dc:"状态"`
}

// ServeLicenseInsertFields 新增服务许可证字段过滤
type ServeLicenseInsertFields struct {
	Group       string `json:"group"       dc:"分组"`
	Name        string `json:"name"        dc:"许可名称"`
	Appid       string `json:"appid"       dc:"应用ID"`
	SecretKey   string `json:"secretKey"   dc:"应用秘钥"`
	OnlineLimit int    `json:"onlineLimit" dc:"在线数量限制，默认1"`
	//Routes      *gjson.Json `json:"routes"      dc:"路由表，空使用默认分组路由"`
	AllowedIps string      `json:"allowedIps"  dc:"白名单，*代表所有，只有允许的IP才能连接到tcp服务"`
	EndAt      *gtime.Time `json:"endAt"       dc:"授权结束时间"`
	Remark     string      `json:"remark"      dc:"备注"`
	Status     int         `json:"status"      dc:"状态"`
}

// ServeLicenseEditInp 修改/新增服务许可证
type ServeLicenseEditInp struct {
	entity.SysServeLicense
}

func (in *ServeLicenseEditInp) Filter(ctx context.Context) (err error) {
	// 验证分组
	if err := g.Validator().Rules("required").Data(in.Group).Messages("分组不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证许可名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("许可名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证应用ID
	if err := g.Validator().Rules("required").Data(in.Appid).Messages("应用ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证授权结束时间
	if err := g.Validator().Rules("required").Data(in.EndAt).Messages("授权结束时间不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type ServeLicenseEditModel struct{}

// ServeLicenseDeleteInp 删除服务许可证
type ServeLicenseDeleteInp struct {
	Id interface{} `json:"id" v:"required#许可ID不能为空" dc:"许可ID"`
}

func (in *ServeLicenseDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ServeLicenseDeleteModel struct{}

// ServeLicenseViewInp 获取指定服务许可证信息
type ServeLicenseViewInp struct {
	Id int64 `json:"id" v:"required#许可ID不能为空" dc:"许可ID"`
}

func (in *ServeLicenseViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ServeLicenseViewModel struct {
	entity.SysServeLicense
}

// ServeLicenseListInp 获取服务许可证列表
type ServeLicenseListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"许可ID"`
	Group     string        `json:"group"     dc:"分组"`
	Name      string        `json:"name"      dc:"许可名称"`
	Appid     string        `json:"appid"     dc:"应用ID"`
	EndAt     []*gtime.Time `json:"endAt"     dc:"授权结束时间"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *ServeLicenseListInp) Filter(ctx context.Context) (err error) {
	return
}

type ServeLicenseListModel struct {
	Id           int64       `json:"id"           dc:"许可ID"`
	Group        string      `json:"group"        dc:"分组"`
	Name         string      `json:"name"         dc:"许可名称"`
	Appid        string      `json:"appid"        dc:"应用ID"`
	SecretKey    string      `json:"secretKey"    dc:"应用秘钥"`
	RemoteAddr   string      `json:"remoteAddr"   dc:"最后连接地址"`
	OnlineLimit  int         `json:"onlineLimit"  dc:"在线数量限制，默认1"`
	LoginTimes   int64       `json:"loginTimes"   dc:"登录次数"`
	LastLoginAt  *gtime.Time `json:"lastLoginAt"  dc:"最后登录时间"`
	LastActiveAt *gtime.Time `json:"lastActiveAt" dc:"最后活跃时间"`
	EndAt        *gtime.Time `json:"endAt"        dc:"授权结束时间"`
	Routes       *gjson.Json `json:"routes"       dc:"路由表，空使用默认分组路由"`
	Status       int         `json:"status"       dc:"状态"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"修改时间"`
	Online       int         `json:"online"       dc:"在线"`
}

// ServeLicenseExportModel 导出服务许可证
type ServeLicenseExportModel struct {
	Id           int64       `json:"id"           dc:"许可ID"`
	Group        string      `json:"group"        dc:"分组"`
	Name         string      `json:"name"         dc:"许可名称"`
	Appid        string      `json:"appid"        dc:"应用ID"`
	SecretKey    string      `json:"secretKey"    dc:"应用秘钥"`
	RemoteAddr   string      `json:"remoteAddr"   dc:"最后连接地址"`
	OnlineLimit  int         `json:"onlineLimit"  dc:"在线数量限制，默认1"`
	LoginTimes   int64       `json:"loginTimes"   dc:"登录次数"`
	LastLoginAt  *gtime.Time `json:"lastLoginAt"  dc:"最后登录时间"`
	LastActiveAt *gtime.Time `json:"lastActiveAt" dc:"最后活跃时间"`
	EndAt        *gtime.Time `json:"endAt"        dc:"授权结束时间"`
	Status       int         `json:"status"       dc:"状态"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"修改时间"`
}

// ServeLicenseStatusInp 更新服务许可证状态
type ServeLicenseStatusInp struct {
	Id     int64 `json:"id" v:"required#许可ID不能为空" dc:"许可ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *ServeLicenseStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("许可ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type ServeLicenseStatusModel struct{}

type ServeLicenseAssignRouterInp struct {
	Id     int64       `json:"id" v:"required#许可ID不能为空" dc:"许可ID"`
	Routes *gjson.Json `json:"routes"       dc:"路由表，空使用默认分组路由"`
}

func (in *ServeLicenseAssignRouterInp) Filter(ctx context.Context) (err error) {
	return
}
