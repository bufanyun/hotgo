// Package validate_test
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package validate_test

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/test/gtest"
	"hotgo/utility/validate"
	"testing"
)

// MockFilter 是 Filter 接口的模拟实现。
type MockFilter struct {
	Foo string
	Bar int
}

func (f *MockFilter) Filter(ctx context.Context) error {
	// 模拟过滤逻辑

	// 过滤出错的例子
	if f.Foo == "" {
		return gerror.New("Foo 字段是必需的")
	}

	// 过滤操作的例子
	f.Bar += 10
	return nil
}

func TestPreFilter(t *testing.T) {
	ctx := context.Background()
	input := &MockFilter{
		Foo: "test",
		Bar: 5,
	}

	err := validate.PreFilter(ctx, input)
	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(err)
	})

	t.Logf("input:%+v", input)

	// 验证过滤结果
	expectedBar := 15
	gtest.C(t, func(t *gtest.T) {
		t.Assert(input.Bar, expectedBar)
	})
}

func TestPreFilter_Error(t *testing.T) {
	ctx := context.Background()
	input := &MockFilter{
		Foo: "",
		Bar: 5,
	}

	err := validate.PreFilter(ctx, input)
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(err, nil)
	})

	expectedError := "Foo 字段是必需的"
	gtest.C(t, func(t *gtest.T) {
		t.Assert(err.Error(), expectedError)
	})
}
