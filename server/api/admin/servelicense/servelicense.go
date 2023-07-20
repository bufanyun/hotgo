// Package servelicense
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.7.6
package servelicense

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询服务许可证列表
type ListReq struct {
	g.Meta `path:"/serveLicense/list" method:"get" tags:"服务许可证" summary:"获取服务许可证列表"`
	sysin.ServeLicenseListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.ServeLicenseListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出服务许可证列表
type ExportReq struct {
	g.Meta `path:"/serveLicense/export" method:"get" tags:"服务许可证" summary:"导出服务许可证列表"`
	sysin.ServeLicenseListInp
}

type ExportRes struct{}

// ViewReq 获取服务许可证指定信息
type ViewReq struct {
	g.Meta `path:"/serveLicense/view" method:"get" tags:"服务许可证" summary:"获取服务许可证指定信息"`
	sysin.ServeLicenseViewInp
}

type ViewRes struct {
	*sysin.ServeLicenseViewModel
}

// EditReq 修改/新增服务许可证
type EditReq struct {
	g.Meta `path:"/serveLicense/edit" method:"post" tags:"服务许可证" summary:"修改/新增服务许可证"`
	sysin.ServeLicenseEditInp
}

type EditRes struct{}

// DeleteReq 删除服务许可证
type DeleteReq struct {
	g.Meta `path:"/serveLicense/delete" method:"post" tags:"服务许可证" summary:"删除服务许可证"`
	sysin.ServeLicenseDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新服务许可证状态
type StatusReq struct {
	g.Meta `path:"/serveLicense/status" method:"post" tags:"服务许可证" summary:"更新服务许可证状态"`
	sysin.ServeLicenseStatusInp
}

type StatusRes struct{}

// AssignRouterReq 分配服务许可证路由
type AssignRouterReq struct {
	g.Meta `path:"/serveLicense/assignRouter" method:"post" tags:"服务许可证" summary:"分配服务许可证路由"`
	sysin.ServeLicenseAssignRouterInp
}

type AssignRouterRes struct{}
