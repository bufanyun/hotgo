// Package validate
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package validate

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// 是否判断

// IsDNSName 是否是域名地址
func IsDNSName(s string) bool {
	DNSName := `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	rxDNSName := regexp.MustCompile(DNSName)
	return s != "" && rxDNSName.MatchString(s)
}

// IsHTTPS 是否是https请求
func IsHTTPS(ctx context.Context) bool {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		g.Log().Info(ctx, "IsHTTPS ctx not request")
		return false
	}
	return r.TLS != nil || gstr.Equal(r.Header.Get("X-Forwarded-Proto"), "https")
}

// IsIp 是否为ipv4
func IsIp(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsPublicIp 是否是公网IP
func IsPublicIp(ip string) bool {
	i := net.ParseIP(ip)

	if i.IsLoopback() || i.IsPrivate() || i.IsMulticast() || i.IsUnspecified() || i.IsLinkLocalUnicast() || i.IsLinkLocalMulticast() {
		return false
	}

	if ip4 := i.To4(); ip4 != nil {
		return !i.Equal(net.IPv4bcast)
	}
	return true
}

// IsLocalIPAddr 检测 IP 地址字符串是否是内网地址
func IsLocalIPAddr(ip string) bool {
	if "localhost" == ip {
		return true
	}
	return HasLocalIP(net.ParseIP(ip))
}

// HasLocalIP 检测 IP 地址是否是内网地址
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

// IsMobile 是否为手机号码
func IsMobile(mobile string) bool {
	pattern := `^(1[2|3|4|5|6|7|8|9][0-9]\d{4,8})$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mobile)
}

// IsEmail 是否为邮箱地址
func IsEmail(email string) bool {
	// pattern := `\w+([-+.]\w+)@\w+([-.]\w+).\w+([-.]\w+)*` //匹配电子邮箱
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

// IsSameDay 是否为同一天
func IsSameDay(t1, t2 int64) bool {
	y1, m1, d1 := time.Unix(t1, 0).Date()
	y2, m2, d2 := time.Unix(t2, 0).Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// IsSameMinute 是否为同一分钟
func IsSameMinute(t1, t2 int64) bool {
	d1 := time.Unix(t1, 0).Format("2006-01-02 15:04")
	d2 := time.Unix(t2, 0).Format("2006-01-02 15:04")
	return d1 == d2
}

// IsMobileVisit 是否为移动端访问
func IsMobileVisit(userAgent string) bool {
	if len(userAgent) == 0 {
		return false
	}

	is := false
	mobileKeywords := []string{"Mobile", "Android", "Silk/", "Kindle", "BlackBerry", "Opera Mini", "Opera Mobi"}
	for i := 0; i < len(mobileKeywords); i++ {
		if strings.Contains(userAgent, mobileKeywords[i]) {
			is = true
			break
		}
	}
	return is
}

// IsWxBrowserVisit 是否为微信访问
func IsWxBrowserVisit(userAgent string) bool {
	if len(userAgent) == 0 {
		return false
	}

	is := false
	userAgent = strings.ToLower(userAgent)
	mobileKeywords := []string{"MicroMessenger"}
	for i := 0; i < len(mobileKeywords); i++ {
		if strings.Contains(userAgent, strings.ToLower(mobileKeywords[i])) {
			is = true
			break
		}
	}
	return is
}

// IsWxMiniProgramVisit 是否为微信小程序访问
func IsWxMiniProgramVisit(userAgent string) bool {
	if len(userAgent) == 0 {
		return false
	}

	is := false
	userAgent = strings.ToLower(userAgent)
	mobileKeywords := []string{"miniProgram"}
	for i := 0; i < len(mobileKeywords); i++ {
		if strings.Contains(userAgent, strings.ToLower(mobileKeywords[i])) {
			is = true
			break
		}
	}
	return is
}
