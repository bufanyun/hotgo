package excel

import (
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestSafeSheetName(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"超长中文截到 31 个字符", "索引条件共3718行,共372页,当前导出是第1页,本页共10行", "索引条件共3718行,共372页,当前导出是第1页,本页共10"},
		{"非法字符替换", `a:b\c/d?e*f[g]h`, "a_b_c_d_e_f_g_h"},
		{"空串", "", "Sheet1"},
		{"首尾单引号", "'报表'", "报表"},
		{"只有单引号", "''", "Sheet1"},
		{"emoji 按字符计数", "😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀", "😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀"},
		{"短名原样", "登录日志", "登录日志"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := safeSheetName(c.in)
			if got != c.want {
				t.Fatalf("safeSheetName(%q) = %q, want %q", c.in, got, c.want)
			}
			// 以 excelize 自身校验为准
			if err := excelize.NewFile().SetSheetName("Sheet1", got); err != nil {
				t.Fatalf("SetSheetName(%q) err: %v", got, err)
			}
		})
	}
}

func TestSanitizeCellValue(t *testing.T) {
	cases := []struct {
		name string
		in   interface{}
		want interface{}
	}{
		{"等号公式", "=1+1", "'=1+1"},
		{"加号开头", "+1", "'+1"},
		{"减号开头", "-1", "'-1"},
		{"at 开头", "@x", "'@x"},
		{"tab 开头", "\tx", "'\tx"},
		{"回车开头", "\rx", "'\rx"},
		{"普通字符串", "abc", "abc"},
		{"空串", "", ""},
		{"整数", 1, 1},
		{"负整数", -1, -1},
		{"浮点数", 1.5, 1.5},
		{"nil", nil, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := sanitizeCellValue(c.in); got != c.want {
				t.Fatalf("sanitizeCellValue(%#v) = %#v, want %#v", c.in, got, c.want)
			}
		})
	}
}
