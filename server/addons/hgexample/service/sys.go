// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/addons/hgexample/model"
	"hotgo/addons/hgexample/model/input/sysin"
	"hotgo/internal/library/hgorm/handler"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysConfig interface {
		// GetBasic 获取基础配置
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		// GetConfigByGroup 获取指定分组配置
		GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error)
		// UpdateConfigByGroup 更新指定分组的配置
		UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateConfigInp) error
	}
	ISysIndex interface {
		// Test 测试
		Test(ctx context.Context, in *sysin.IndexTestInp) (res *sysin.IndexTestModel, err error)
	}
	ISysTable interface {
		// Model Orm模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取列表
		List(ctx context.Context, in *sysin.TableListInp) (list []*sysin.TableListModel, totalCount int, err error)
		// Export 导出
		Export(ctx context.Context, in *sysin.TableListInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.TableEditInp) (err error)
		// Delete 删除
		Delete(ctx context.Context, in *sysin.TableDeleteInp) (err error)
		// Status 更新状态
		Status(ctx context.Context, in *sysin.TableStatusInp) (err error)
		// Switch 更新开关状态
		Switch(ctx context.Context, in *sysin.TableSwitchInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *sysin.TableMaxSortInp) (res *sysin.TableMaxSortModel, err error)
		// View 获取指定信息
		View(ctx context.Context, in *sysin.TableViewInp) (res *sysin.TableViewModel, err error)
	}
	ISysTenantOrder interface {
		// Model 多租户功能演示ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取多租户功能演示列表
		List(ctx context.Context, in *sysin.TenantOrderListInp) (list []*sysin.TenantOrderListModel, totalCount int, err error)
		// Export 导出多租户功能演示
		Export(ctx context.Context, in *sysin.TenantOrderListInp) (err error)
		// Edit 修改/新增多租户功能演示
		Edit(ctx context.Context, in *sysin.TenantOrderEditInp) (err error)
		// Delete 删除多租户功能演示
		Delete(ctx context.Context, in *sysin.TenantOrderDeleteInp) (err error)
		// View 获取多租户功能演示指定信息
		View(ctx context.Context, in *sysin.TenantOrderViewInp) (res *sysin.TenantOrderViewModel, err error)
	}
	ISysTreeTable interface {
		// Model Orm模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取列表
		List(ctx context.Context, in *sysin.TreeTableListInp) (list []*sysin.TreeTableListModel, totalCount int, err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.TableEditInp) (err error)
		// Delete 删除
		Delete(ctx context.Context, in *sysin.TableDeleteInp) (err error)
		// Select 关系树选项列表
		Select(ctx context.Context) (list []*sysin.TableTree, err error)
	}
)

var (
	localSysConfig      ISysConfig
	localSysIndex       ISysIndex
	localSysTable       ISysTable
	localSysTenantOrder ISysTenantOrder
	localSysTreeTable   ISysTreeTable
)

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}

func SysIndex() ISysIndex {
	if localSysIndex == nil {
		panic("implement not found for interface ISysIndex, forgot register?")
	}
	return localSysIndex
}

func RegisterSysIndex(i ISysIndex) {
	localSysIndex = i
}

func SysTable() ISysTable {
	if localSysTable == nil {
		panic("implement not found for interface ISysTable, forgot register?")
	}
	return localSysTable
}

func RegisterSysTable(i ISysTable) {
	localSysTable = i
}

func SysTenantOrder() ISysTenantOrder {
	if localSysTenantOrder == nil {
		panic("implement not found for interface ISysTenantOrder, forgot register?")
	}
	return localSysTenantOrder
}

func RegisterSysTenantOrder(i ISysTenantOrder) {
	localSysTenantOrder = i
}

func SysTreeTable() ISysTreeTable {
	if localSysTreeTable == nil {
		panic("implement not found for interface ISysTreeTable, forgot register?")
	}
	return localSysTreeTable
}

func RegisterSysTreeTable(i ISysTreeTable) {
	localSysTreeTable = i
}
