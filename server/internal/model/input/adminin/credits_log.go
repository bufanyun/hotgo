// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.5.3
// @AutoGenerate Date 2023-04-15 15:59:58
package adminin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/location"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// CreditsLogSaveBalanceInp 更新余额
type CreditsLogSaveBalanceInp struct {
	MemberId    int64   `json:"memberId"    dc:"管理员ID"`
	AppId       string  `json:"appId"       dc:"应用id"`
	AddonsName  string  `json:"addonsName"  dc:"插件名称"`
	CreditGroup string  `json:"creditGroup" dc:"变动的组别"`
	Num         float64 `json:"num"         dc:"变动数据"`
	Ip          string  `json:"ip"          dc:"操作人IP"`
	MapId       int64   `json:"mapId"       dc:"关联ID"`
	Remark      string  `json:"remark"      dc:"备注"`
}

func (in *CreditsLogSaveBalanceInp) Filter(ctx context.Context) (err error) {
	if in.Num == 0 {
		err = gerror.New("更新余额不能为0")
	}

	if in.AppId == "" {
		in.AppId = contexts.GetModule(ctx)
	}

	if in.CreditGroup == "" {
		if in.Num > 0 {
			in.CreditGroup = consts.CreditGroupIncr
		} else {
			in.CreditGroup = consts.CreditGroupDecr
		}
	}

	if in.Ip == "" {
		in.Ip = location.GetClientIp(ghttp.RequestFromCtx(ctx))
	}
	return
}

type CreditsLogSaveBalanceModel struct {
}

// CreditsLogSaveIntegralInp 更新积分
type CreditsLogSaveIntegralInp struct {
	MemberId    int64   `json:"memberId"    dc:"管理员ID"`
	AppId       string  `json:"appId"       dc:"应用id"`
	AddonsName  string  `json:"addonsName"  dc:"插件名称"`
	CreditGroup string  `json:"creditGroup" dc:"变动的组别"`
	Num         float64 `json:"num"         dc:"变动数据"`
	Ip          string  `json:"ip"          dc:"操作人IP"`
	MapId       int64   `json:"mapId"       dc:"关联ID"`
	Remark      string  `json:"remark"      dc:"备注"`
}

func (in *CreditsLogSaveIntegralInp) Filter(ctx context.Context) (err error) {
	if in.Num == 0 {
		err = gerror.New("更新积分不能为0")
	}

	if in.AppId == "" {
		in.AppId = contexts.GetModule(ctx)
	}

	if in.CreditGroup == "" {
		if in.Num > 0 {
			in.CreditGroup = consts.CreditGroupIncr
		} else {
			in.CreditGroup = consts.CreditGroupDecr
		}
	}

	if in.Ip == "" {
		in.Ip = location.GetClientIp(ghttp.RequestFromCtx(ctx))
	}
	return
}

type CreditsLogSaveIntegralModel struct {
}

// CreditsLogListInp 获取资产变动列表
type CreditsLogListInp struct {
	form.PageReq
	Id          int64         `json:"id"          dc:"变动ID"`
	MemberId    int64         `json:"memberId"    dc:"管理员ID"`
	AppId       string        `json:"appId"       dc:"应用id"`
	CreditType  string        `json:"creditType"  dc:"变动类型"`
	CreditGroup string        `json:"creditGroup" dc:"变动的组别"`
	Remark      string        `json:"remark"      dc:"备注"`
	Ip          string        `json:"ip"          dc:"操作人IP"`
	Status      int           `json:"status"      dc:"状态"`
	CreatedAt   []*gtime.Time `json:"createdAt"   dc:"创建时间"`
}

func (in *CreditsLogListInp) Filter(ctx context.Context) (err error) {
	return
}

type CreditsLogListModel struct {
	Id          int64       `json:"id"          dc:"变动ID"`
	MemberId    int64       `json:"memberId"    dc:"管理员ID"`
	AppId       string      `json:"appId"       dc:"应用id"`
	AddonsName  string      `json:"addonsName"  dc:"插件名称"`
	CreditType  string      `json:"creditType"  dc:"变动类型"`
	CreditGroup string      `json:"creditGroup" dc:"变动的组别"`
	BeforeNum   float64     `json:"beforeNum"   dc:"变动前"`
	Num         float64     `json:"num"         dc:"变动数据"`
	AfterNum    float64     `json:"afterNum"    dc:"变动后"`
	Remark      string      `json:"remark"      dc:"备注"`
	Ip          string      `json:"ip"          dc:"操作人IP"`
	MapId       int64       `json:"mapId"       dc:"关联ID"`
	Status      int         `json:"status"      dc:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}

// CreditsLogExportModel 导出资产变动
type CreditsLogExportModel struct {
	Id          int64       `json:"id"          dc:"变动ID"`
	MemberId    int64       `json:"memberId"    dc:"管理员ID"`
	AppId       string      `json:"appId"       dc:"应用id"`
	AddonsName  string      `json:"addonsName"  dc:"插件名称"`
	CreditType  string      `json:"creditType"  dc:"变动类型"`
	CreditGroup string      `json:"creditGroup" dc:"变动的组别"`
	BeforeNum   float64     `json:"beforeNum"   dc:"变动前"`
	Num         float64     `json:"num"         dc:"变动数据"`
	AfterNum    float64     `json:"afterNum"    dc:"变动后"`
	Remark      string      `json:"remark"      dc:"备注"`
	Ip          string      `json:"ip"          dc:"操作人IP"`
	MapId       int64       `json:"mapId"       dc:"关联ID"`
	Status      int         `json:"status"      dc:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}
