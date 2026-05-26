// Package handler
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package handler_test

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"testing"
)

type SorterInput struct {
	form.Sorters
}

// TestSorterDefault 默认排序
func TestSorterDefault(t *testing.T) {
	in := &SorterInput{} // 不存在排序条件，默认使用主表主键降序排序

	_, err := dao.SysGenCurdDemo.Ctx(gctx.New()).Handler(handler.Sorter(in)).All()
	if err != nil {
		t.Error(err)
		return
	}
}

// TestSorter 多字段排序
func TestSorter(t *testing.T) {
	in := &SorterInput{
		Sorters: form.Sorters{
			Sorters: []*form.Sorter{
				{
					ColumnKey: "id",
					Order:     "descend", // 降序
				},
				{
					ColumnKey: "categoryId", // 自动转换为下划线。categoryId -> category_id
					Order:     false,        // 不参与排序
				},
				{
					ColumnKey: "created_at",
					Order:     "descend", // 降序
				},
			},
		},
	}

	_, err := dao.SysGenCurdDemo.Ctx(gctx.New()).Handler(handler.Sorter(in)).All()
	if err != nil {
		t.Error(err)
		return
	}
}

// TestSorterJoinTable 关联表多字段排序
func TestSorterJoinTable(t *testing.T) {
	in := &SorterInput{
		Sorters: form.Sorters{
			Sorters: []*form.Sorter{
				{
					ColumnKey: "id",
					Order:     "descend", // 降序
				},
				{
					ColumnKey: "categoryId", // 自动转换为下划线。categoryId -> category_id
					Order:     false,        // 不参与排序
				},
				{
					ColumnKey: "created_at",
					Order:     "descend", // 降序
				},
				{
					ColumnKey: "testCategoryName", // 自动识别关联表别名。 testCategoryName -> testCategory.name
					Order:     "ascend",           // 升序
				},
			},
		},
	}

	_, err := dao.SysGenCurdDemo.Ctx(gctx.New()).
		LeftJoin(hgorm.GenJoinOnRelation(
			dao.SysGenCurdDemo.Table(), dao.SysGenCurdDemo.Columns().CategoryId, // 主表表名,关联字段
			dao.TestCategory.Table(), "testCategory", dao.TestCategory.Columns().Id, // 关联表表名,别名,关联字段
		)...).
		Handler(handler.Sorter(in)).All()
	if err != nil {
		t.Error(err)
		return
	}
}
