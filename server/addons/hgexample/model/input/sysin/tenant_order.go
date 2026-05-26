// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.13.1
package sysin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TenantOrderUpdateFields 修改多租户功能演示字段过滤
type TenantOrderUpdateFields struct {
	TenantId    int64   `json:"tenantId"    dc:"租户ID"`
	MerchantId  int64   `json:"merchantId"  dc:"商户ID"`
	UserId      int64   `json:"userId"      dc:"用户ID"`
	ProductName string  `json:"productName" dc:"购买产品"`
	OrderSn     string  `json:"orderSn"     dc:"关联订单号"`
	Money       float64 `json:"money"       dc:"充值金额"`
	Remark      string  `json:"remark"      dc:"备注"`
	Status      int     `json:"status"      dc:"支付状态"`
}

// TenantOrderInsertFields 新增多租户功能演示字段过滤
type TenantOrderInsertFields struct {
	TenantId    int64   `json:"tenantId"    dc:"租户ID"`
	MerchantId  int64   `json:"merchantId"  dc:"商户ID"`
	UserId      int64   `json:"userId"      dc:"用户ID"`
	ProductName string  `json:"productName" dc:"购买产品"`
	OrderSn     string  `json:"orderSn"     dc:"关联订单号"`
	Money       float64 `json:"money"       dc:"充值金额"`
	Remark      string  `json:"remark"      dc:"备注"`
	Status      int     `json:"status"      dc:"支付状态"`
}

// TenantOrderEditInp 修改/新增多租户功能演示
type TenantOrderEditInp struct {
	entity.AddonHgexampleTenantOrder
}

func (in *TenantOrderEditInp) Filter(ctx context.Context) (err error) {
	// 验证充值金额
	if err := g.Validator().Rules("required").Data(in.Money).Messages("充值金额不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type TenantOrderEditModel struct{}

// TenantOrderDeleteInp 删除多租户功能演示
type TenantOrderDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *TenantOrderDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TenantOrderDeleteModel struct{}

// TenantOrderViewInp 获取指定多租户功能演示信息
type TenantOrderViewInp struct {
	Id int64 `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *TenantOrderViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TenantOrderViewModel struct {
	entity.AddonHgexampleTenantOrder
}

// TenantOrderListInp 获取多租户功能演示列表
type TenantOrderListInp struct {
	form.PageReq
	Id         int64         `json:"id"         dc:"主键"`
	TenantId   int64         `json:"tenantId"   dc:"租户ID"`
	MerchantId int64         `json:"merchantId" dc:"商户ID"`
	UserId     int64         `json:"userId"     dc:"用户ID"`
	OrderSn    string        `json:"orderSn"    dc:"关联订单号"`
	Status     int           `json:"status"     dc:"支付状态"`
	CreatedAt  []*gtime.Time `json:"createdAt"  dc:"创建时间"`
}

func (in *TenantOrderListInp) Filter(ctx context.Context) (err error) {
	return
}

type TenantOrderListModel struct {
	Id          int64       `json:"id"          dc:"主键"`
	TenantId    int64       `json:"tenantId"    dc:"租户ID"`
	MerchantId  int64       `json:"merchantId"  dc:"商户ID"`
	UserId      int64       `json:"userId"      dc:"用户ID"`
	ProductName string      `json:"productName" dc:"购买产品"`
	OrderSn     string      `json:"orderSn"     dc:"关联订单号"`
	Money       float64     `json:"money"       dc:"充值金额"`
	Remark      string      `json:"remark"      dc:"备注"`
	Status      int         `json:"status"      dc:"支付状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
}

// TenantOrderExportModel 导出多租户功能演示
type TenantOrderExportModel struct {
	Id          int64       `json:"id"          dc:"主键"`
	TenantId    int64       `json:"tenantId"    dc:"租户ID"`
	MerchantId  int64       `json:"merchantId"  dc:"商户ID"`
	UserId      int64       `json:"userId"      dc:"用户ID"`
	ProductName string      `json:"productName" dc:"购买产品"`
	OrderSn     string      `json:"orderSn"     dc:"关联订单号"`
	Money       float64     `json:"money"       dc:"充值金额"`
	Remark      string      `json:"remark"      dc:"备注"`
	Status      int         `json:"status"      dc:"支付状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
}
