// Package charset
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package charset

import (
	"crypto/rand"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/utility/convert"
	r "math/rand"
	"strings"
	"time"
)

// SplitMemberIds 从截取字串符中读取用户ID
func SplitMemberIds(str, pos string) (memberIds []int64) {
	receiver := strings.Split(strings.TrimSpace(str), pos)
	if len(receiver) == 0 {
		return memberIds
	}
	if len(receiver) == 1 && strings.TrimSpace(receiver[0]) == "" {
		return memberIds
	}

	for _, memberId := range receiver {
		memberIds = append(memberIds, gconv.Int64(strings.TrimSpace(memberId)))
	}

	return convert.UniqueSliceInt64(memberIds)
}

// GetMapKeysByString 获取map的所有key，字串符类型
func GetMapKeysByString(m map[string]string) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

// RandomCreateBytes 生成随机字串符
func RandomCreateBytes(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
	}
	var bytes = make([]byte, n)
	var randBy bool
	r.Seed(time.Now().UnixNano())
	if num, err := rand.Read(bytes); num != n || err != nil {
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}

// GetStack 格式化错误的堆栈信息
func GetStack(err error) []string {
	stackList := gstr.Split(gerror.Stack(err), "\n")
	for i := 0; i < len(stackList); i++ {
		stackList[i] = gstr.Replace(stackList[i], "\t", "--> ")
	}

	return stackList
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
