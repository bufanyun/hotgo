// Package location
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package location

import (
	"context"
	"github.com/axgle/mahonia"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kayon/iploc"
	"hotgo/utility/validate"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type IpLocationData struct {
	Ip           string `json:"ip"`
	Country      string `json:"country"`
	Region       string `json:"region"`
	Province     string `json:"province"`
	ProvinceCode int64  `json:"province_code"`
	City         string `json:"city"`
	CityCode     int64  `json:"city_code"`
	Area         string `json:"area"`
	AreaCode     int64  `json:"area_code"`
}

// WhoisLocation 通过Whois接口查询IP归属地
func WhoisLocation(ctx context.Context, ip string) IpLocationData {

	type whoisRegionData struct {
		Ip         string `json:"ip"`
		Pro        string `json:"pro" `
		ProCode    string `json:"proCode" `
		City       string `json:"city" `
		CityCode   string `json:"cityCode"`
		Region     string `json:"region"`
		RegionCode string `json:"regionCode"`
		Addr       string `json:"addr"`
		Err        string `json:"err"`
	}

	if !validate.IsIp(ip) {
		return IpLocationData{}
	}

	response, err := g.Client().Timeout(10*time.Second).Get(ctx, "http://whois.pconline.com.cn/ipJson.jsp?ip="+ip+"&json=true")
	if err != nil {
		err = gerror.New(err.Error())
		return IpLocationData{
			Ip: ip,
		}
	}

	defer response.Close()

	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	data := enc.ConvertString(response.ReadAllString())
	whoisData := whoisRegionData{}
	if err := gconv.Struct(data, &whoisData); err != nil {
		err = gerror.New(err.Error())

		g.Log().Print(ctx, "err:", err)
		return IpLocationData{
			Ip: ip,
		}
	}

	return IpLocationData{
		Ip: whoisData.Ip,
		//Country      string `json:"country"`
		Region:       whoisData.Addr,
		Province:     whoisData.Pro,
		ProvinceCode: gconv.Int64(whoisData.ProCode),
		City:         whoisData.City,
		CityCode:     gconv.Int64(whoisData.CityCode),
		Area:         whoisData.Region,
		AreaCode:     gconv.Int64(whoisData.RegionCode),
	}
}

// Cz88Find 通过Cz88的IP库查询IP归属地
func Cz88Find(ctx context.Context, ip string) IpLocationData {
	if !validate.IsIp(ip) {
		g.Log().Print(ctx, "ip格式错误:", ip)
		return IpLocationData{}
	}

	loc, err := iploc.OpenWithoutIndexes("./storage/ip/qqwry-utf8.dat")
	if err != nil {
		err = gerror.New(err.Error())
		return IpLocationData{
			Ip: ip,
		}
	}

	detail := loc.Find(ip)
	if detail == nil {
		return IpLocationData{
			Ip: ip,
		}
	}

	locationData := IpLocationData{
		Ip:       ip,
		Country:  detail.Country,
		Region:   detail.Region,
		Province: detail.Province,
		City:     detail.City,
		Area:     detail.County,
	}

	if gstr.LenRune(locationData.Province) == 0 {
		return locationData
	}

	return locationData
}

// IsJurisByIpTitle 判断地区名称是否为直辖市
func IsJurisByIpTitle(title string) bool {

	lists := []string{"北京市", "天津市", "重庆市", "上海市"}

	for i := 0; i < len(lists); i++ {
		if gstr.Contains(lists[i], title) {
			return true
		}
	}
	return false
}

// GetLocation 获取IP归属地信息
func GetLocation(ctx context.Context, ip string) IpLocationData {
	method, _ := g.Cfg().Get(ctx, "hotgo.ipMethod", "cz88")

	if method.String() == "whois" {
		return WhoisLocation(ctx, ip)
	}
	return Cz88Find(ctx, ip)
}

// GetPublicIP 获取公网IP
func GetPublicIP() (ip string, err error) {
	response, err := http.Get("http://members.3322.org/dyndns/getip")
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	ip = string(body)
	ip = strings.ReplaceAll(ip, "\n", "")
	return
}

// GetLocalIP 获取服务器内网IP
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

// GetClientIp 获取客户端IP
func GetClientIp(r *ghttp.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.GetClientIp()
	}
	return ip
}
