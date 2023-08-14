package validate

import (
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func TestIsEmail(t *testing.T) {
	b := IsEmail("QTT123456@163.com")
	gtest.Assert(true, b)
}
