// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package utils_test

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"hotgo/internal/library/hggen/internal/utility/utils"
)

func Test_GetModPath(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		goModPath := utils.GetModPath()
		fmt.Println(goModPath)
	})
}
