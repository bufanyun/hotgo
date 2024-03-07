// Package validate_test
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package validate_test

import (
	"github.com/gogf/gf/v2/test/gtest"
	"hotgo/utility/validate"
	"testing"
)

func TestIsEmail(t *testing.T) {
	b := validate.IsEmail("QTT123456@163.com")
	gtest.Assert(true, b)
}
