// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.15.7
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

// NormalTreeDemoUpdateFields 修改普通树表字段过滤
type NormalTreeDemoUpdateFields struct {
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

// NormalTreeDemoInsertFields 新增普通树表字段过滤
type NormalTreeDemoInsertFields struct {
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

// NormalTreeDemoEditInp 修改/新增普通树表
type NormalTreeDemoEditInp struct {
	entity.SysGenTreeDemo
}

func (in *NormalTreeDemoEditInp) Filter(ctx context.Context) (err error) {
	// 验证标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type NormalTreeDemoEditModel struct{}

// NormalTreeDemoDeleteInp 删除普通树表
type NormalTreeDemoDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *NormalTreeDemoDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type NormalTreeDemoDeleteModel struct{}

// NormalTreeDemoViewInp 获取指定普通树表信息
type NormalTreeDemoViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *NormalTreeDemoViewInp) Filter(ctx context.Context) (err error) {
	return
}

type NormalTreeDemoViewModel struct {
	entity.SysGenTreeDemo
	CreatedBySumma *hook.MemberSumma `json:"createdBySumma" dc:"创建者摘要信息"`
}

// NormalTreeDemoListInp 获取普通树表列表
type NormalTreeDemoListInp struct {
	form.PageReq
	Title      string        `json:"title"      dc:"标题"`
	Pid        int64         `json:"pid"        dc:"上级"`
	CategoryId int64         `json:"categoryId" dc:"测试分类"`
	Status     int           `json:"status"     dc:"状态"`
	CreatedAt  []*gtime.Time `json:"createdAt"  dc:"创建时间"`
}

func (in *NormalTreeDemoListInp) Filter(ctx context.Context) (err error) {
	return
}

type NormalTreeDemoListModel struct {
	Title          string            `json:"title"          dc:"标题"`
	Id             int64             `json:"id"             dc:"ID"`
	Pid            int64             `json:"pid"            dc:"上级"`
	CategoryId     int64             `json:"categoryId"     dc:"测试分类"`
	Description    string            `json:"description"    dc:"描述"`
	Status         int               `json:"status"         dc:"状态"`
	CreatedBy      int64             `json:"createdBy"      dc:"创建者"`
	CreatedBySumma *hook.MemberSumma `json:"createdBySumma" dc:"创建者摘要信息"`
	CreatedAt      *gtime.Time       `json:"createdAt"      dc:"创建时间"`
}

// NormalTreeDemoMaxSortInp 获取普通树表最大排序
type NormalTreeDemoMaxSortInp struct{}

func (in *NormalTreeDemoMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type NormalTreeDemoMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// NormalTreeDemoTreeOption 关系树选项
type NormalTreeDemoTreeOption struct {
	Title    string      `json:"title" dc:"标题"`
	Id       int64       `json:"id"    dc:"ID"`
	Pid      int64       `json:"pid"   dc:"上级"`
	Children []tree.Node `json:"children"  dc:"子节点"`
}

// ID 获取节点ID
func (t *NormalTreeDemoTreeOption) ID() int64 {
	return t.Id
}

// PID 获取父级节点ID
func (t *NormalTreeDemoTreeOption) PID() int64 {
	return t.Pid
}

// SetChildren 设置子节点数据
func (t *NormalTreeDemoTreeOption) SetChildren(children []tree.Node) {
	t.Children = children
}
