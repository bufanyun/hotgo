// Package validate
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package validate

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"net"
	"net/url"
	"regexp"
)

func IsDNSName(s string) bool {
	DNSName := `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	rxDNSName := regexp.MustCompile(DNSName)
	return s != "" && rxDNSName.MatchString(s)
}

func IsHTTPS(ctx context.Context) bool {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		g.Log().Warningf(ctx, "IsHTTPS ctx not request")
		return false
	}
	var (
		proto = r.Header.Get("X-Forwarded-Proto")
	)

	if r.TLS != nil || gstr.Equal(proto, "https") {
		return true
	}
	return false
}

// IsIp 是否为ipv4
func IsIp(ip string) bool {
	if net.ParseIP(ip) != nil {
		return true
	}
	return false
}

// IsPublicIp 是否是公网IP
func IsPublicIp(Ip string) bool {
	ip := net.ParseIP(Ip)

	if ip.IsLoopback() || ip.IsPrivate() || ip.IsMulticast() || ip.IsUnspecified() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return false
	}

	if ip4 := ip.To4(); ip4 != nil {
		return !ip.Equal(net.IPv4bcast)
	}

	return true

}

// IsMobile 是否为手机号码
func IsMobile(mobile string) bool {
	pattern := `^(1[2|3|4|5|6|7|8|9][0-9]\d{4,8})$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mobile)
}

// IsEmail 是否为邮箱地址
func IsEmail(email string) bool {
	//pattern := `\w+([-+.]\w+)@\w+([-.]\w+).\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z].){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
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

// IsIDCard 是否为身份证
func IsIDCard(idCard string) bool {
	sz := len(idCard)
	if sz != 18 {
		return false
	}
	weight := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	validate := []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	sum := 0
	for i := 0; i < len(weight); i++ {
		sum += weight[i] * int(byte(idCard[i])-'0')
	}
	m := sum % 11
	return validate[m] == idCard[sz-1]
}
