// Package charset
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package charset

import (
	"bytes"
	"crypto/rand"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gutil"
	r "math/rand"
	"strings"
	"time"
)

// RandomCreateBytes 生成随机字串符
func RandomCreateBytes(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
	}
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.New(r.NewSource(time.Now().UnixNano())).Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}

// ParseErrStack 解析错误的堆栈信息
func ParseErrStack(err error) []string {
	return ParseStack(gerror.Stack(err))
}

// ParseStack 解析堆栈信息
func ParseStack(st string) []string {
	stack := gstr.Split(st, "\n")
	for i := 0; i < len(stack); i++ {
		stack[i] = gstr.Replace(stack[i], "\t", "--> ")
	}
	return stack
}

// SerializeStack 解析错误并序列化堆栈信息
func SerializeStack(err error) string {
	buffer := bytes.NewBuffer(nil)
	gutil.DumpTo(buffer, ParseErrStack(err), gutil.DumpOption{
		WithType:     false,
		ExportedOnly: false,
	})
	return buffer.String()
}

// SubstrAfter 截取指定字符后的内容
func SubstrAfter(str string, symbol string) string {
	comma := strings.Index(str, symbol)
	if comma < 0 { // -1 不存在
		return ""
	}
	pos := strings.Index(str[comma:], symbol)
	if comma+pos+1 > len(str) {
		return ""
	}
	return str[comma+pos+1:]
}
