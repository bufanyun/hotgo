// Package validate
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package validate

import (
	"github.com/gogf/gf/v2/util/gconv"
	"net"
	"net/url"
	"regexp"
	"time"
)

// IsIp 是否为ipv4
func IsIp(ip string) bool {
	if net.ParseIP(ip) != nil {
		return true
	}
	return false
}

// IsEmail 是否为邮箱地址
func IsEmail(email string) bool {
	//pattern := `\w+([-+.]\w+)@\w+([-.]\w+).\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z].){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// InSameDay 是否为同一天
func InSameDay(t1, t2 int64) bool {
	y1, m1, d1 := time.Unix(t1, 0).Date()
	y2, m2, d2 := time.Unix(t2, 0).Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

// InSameMinute 是否为同一分钟
func InSameMinute(t1, t2 int64) bool {
	d1 := time.Unix(t1, 0).Format("2006-01-02 15:04")
	d2 := time.Unix(t2, 0).Format("2006-01-02 15:04")

	return d1 == d2
}

// InSliceExistStr 判断字符或切片字符是否存在指定字符
func InSliceExistStr(elems interface{}, search string) bool {
	switch elems.(type) {
	case []string:
		elem := gconv.Strings(elems)
		for i := 0; i < len(elem); i++ {
			if gconv.String(elem[i]) == search {
				return true
			}
		}
	default:
		return gconv.String(elems) == search
	}

	return false
}

// IsURL 是否是url地址
func IsURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}
	URL, err := url.Parse(u)
	if err != nil || URL.Scheme == "" || URL.Host == "" {
		return false
	}
	return true
}
