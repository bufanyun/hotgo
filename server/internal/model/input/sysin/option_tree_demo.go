// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sysin

import (
	"context"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/tree"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OptionTreeDemoUpdateFields 修改选项树表字段过滤
type OptionTreeDemoUpdateFields struct {
	Title       string `json:"title"       dc:"标题"`
	Pid         int64  `json:"pid"         dc:"上级"`
	Level       int    `json:"level"       dc:"关系树级别"`
	Tree        string `json:"tree"        dc:"关系树"`
	CategoryId  int64  `json:"categoryId"  dc:"测试分类"`
	Description string `json:"description" dc:"描述"`
	Sort        int    `json:"sort"        dc:"排序"`
	Status      int    `json:"status"      dc:"状态"`
	UpdatedBy   int64  `json:"updatedBy"   dc:"更新者"`
}

// OptionTreeDemoInsertFields 新增选项树表字段过滤
type OptionTreeDemoInsertFields struct {
	Title       string `json:"title"       dc:"标题"`
	Pid         int64  `json:"pid"         dc:"上级"`
	Level       int    `json:"level"       dc:"关系树级别"`
	Tree        string `json:"tree"        dc:"关系树"`
	CategoryId  int64  `json:"categoryId"  dc:"测试分类"`
	Description string `json:"description" dc:"描述"`
	Sort        int    `json:"sort"        dc:"排序"`
	Status      int    `json:"status"      dc:"状态"`
	CreatedBy   int64  `json:"createdBy"   dc:"创建者"`
}

// OptionTreeDemoEditInp 修改/新增选项树表
type OptionTreeDemoEditInp struct {
	entity.SysGenTreeDemo
}

func (in *OptionTreeDemoEditInp) Filter(ctx context.Context) (err error) {
	// 验证标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type OptionTreeDemoEditModel struct{}

// OptionTreeDemoDeleteInp 删除选项树表
type OptionTreeDemoDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *OptionTreeDemoDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type OptionTreeDemoDeleteModel struct{}

// OptionTreeDemoViewInp 获取指定选项树表信息
type OptionTreeDemoViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *OptionTreeDemoViewInp) Filter(ctx context.Context) (err error) {
	return
}

type OptionTreeDemoViewModel struct {
	entity.SysGenTreeDemo
	CreatedBySumma *hook.MemberSumma `json:"createdBySumma" dc:"创建者摘要信息"`
}

// OptionTreeDemoListInp 获取选项树表列表
type OptionTreeDemoListInp struct {
	form.PageReq
	Title      string        `json:"title"      dc:"标题"`
	Pid        int64         `json:"pid"        dc:"上级"`
	CategoryId int64         `json:"categoryId" dc:"测试分类"`
	Status     int           `json:"status"     dc:"状态"`
	CreatedAt  []*gtime.Time `json:"createdAt"  dc:"创建时间"`
}

func (in *OptionTreeDemoListInp) Filter(ctx context.Context) (err error) {
	return
}

type OptionTreeDemoListModel struct {
	Title          string            `json:"title"          dc:"标题"`
	Id             int64             `json:"id"             dc:"ID"`
	Pid            int64             `json:"pid"            dc:"上级"`
	CategoryId     int64             `json:"categoryId"     dc:"测试分类"`
	Status         int               `json:"status"         dc:"状态"`
	CreatedBy      int64             `json:"createdBy"      dc:"创建者"`
	CreatedBySumma *hook.MemberSumma `json:"createdBySumma" dc:"创建者摘要信息"`
	CreatedAt      *gtime.Time       `json:"createdAt"      dc:"创建时间"`
}

// OptionTreeDemoMaxSortInp 获取选项树表最大排序
type OptionTreeDemoMaxSortInp struct{}

func (in *OptionTreeDemoMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type OptionTreeDemoMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// OptionTreeDemoTreeOption 关系树选项
type OptionTreeDemoTreeOption struct {
	Title string `json:"title" dc:"标题"`
	Id    int64  `json:"id"    dc:"ID"`
	Pid   int64  `json:"pid"   dc:"上级"`

	Children []tree.Node `json:"children"  dc:"子节点"`
}

// ID 获取节点ID
func (t *OptionTreeDemoTreeOption) ID() int64 {
	return t.Id
}

// PID 获取父级节点ID
func (t *OptionTreeDemoTreeOption) PID() int64 {
	return t.Pid
}

// SetChildren 设置子节点数据
func (t *OptionTreeDemoTreeOption) SetChildren(children []tree.Node) {
	t.Children = children
}