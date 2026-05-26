// Package tenantorder
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.13.1
package tenantorder

import (
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询多租户功能演示列表
type ListReq struct {
	g.Meta `path:"/tenantOrder/list" method:"get" tags:"多租户功能演示" summary:"获取多租户功能演示列表"`
	sysin.TenantOrderListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.TenantOrderListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出多租户功能演示列表
type ExportReq struct {
	g.Meta `path:"/tenantOrder/export" method:"get" tags:"多租户功能演示" summary:"导出多租户功能演示列表"`
	sysin.TenantOrderListInp
}

type ExportRes struct{}

// ViewReq 获取多租户功能演示指定信息
type ViewReq struct {
	g.Meta `path:"/tenantOrder/view" method:"get" tags:"多租户功能演示" summary:"获取多租户功能演示指定信息"`
	sysin.TenantOrderViewInp
}

type ViewRes struct {
	*sysin.TenantOrderViewModel
}

// EditReq 修改/新增多租户功能演示
type EditReq struct {
	g.Meta `path:"/tenantOrder/edit" method:"post" tags:"多租户功能演示" summary:"修改/新增多租户功能演示"`
	sysin.TenantOrderEditInp
}

type EditRes struct{}

// DeleteReq 删除多租户功能演示
type DeleteReq struct {
	g.Meta `path:"/tenantOrder/delete" method:"post" tags:"多租户功能演示" summary:"删除多租户功能演示"`
	sysin.TenantOrderDeleteInp
}

type DeleteRes struct{}