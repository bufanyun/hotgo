// Package location
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package location

import (
	"context"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kayon/iploc"
	"hotgo/utility/validate"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	whoisApi = "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip="
	dyndns   = "http://members.3322.org/dyndns/getip"
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

type WhoisRegionData struct {
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

var cacheMap *gmap.Map

func init() {
	cacheMap = gmap.New(true)
}

// WhoisLocation 通过Whois接口查询IP归属地
func WhoisLocation(ctx context.Context, ip string) (*IpLocationData, error) {
	response, err := g.Client().Timeout(10*time.Second).Get(ctx, whoisApi+ip)
	if err != nil {
		return nil, err
	}

	defer response.Close()

	var (
		whoisData *WhoisRegionData
		enc       = mahonia.NewDecoder("gbk")
		data      = enc.ConvertString(response.ReadAllString())
	)

	if err = gconv.Struct(data, &whoisData); err != nil {
		return nil, err
	}

	return &IpLocationData{
		Ip: whoisData.Ip,
		//Country      string `json:"country"`
		Region:       whoisData.Addr,
		Province:     whoisData.Pro,
		ProvinceCode: gconv.Int64(whoisData.ProCode),
		City:         whoisData.City,
		CityCode:     gconv.Int64(whoisData.CityCode),
		Area:         whoisData.Region,
		AreaCode:     gconv.Int64(whoisData.RegionCode),
	}, nil
}

// Cz88Find 通过Cz88的IP库查询IP归属地
func Cz88Find(ctx context.Context, ip string) (*IpLocationData, error) {
	loc, err := iploc.OpenWithoutIndexes("./resource/ip/qqwry-utf8.dat")
	if err != nil {
		return nil, fmt.Errorf("%v for help, please go to: https://github.com/kayon/iploc", err.Error())
	}

	detail := loc.Find(ip)
	if detail == nil {
		return nil, fmt.Errorf("no ip data is queried. procedure:%v", ip)
	}

	locationData := &IpLocationData{
		Ip:       ip,
		Country:  detail.Country,
		Region:   detail.Region,
		Province: detail.Province,
		City:     detail.City,
		Area:     detail.County,
	}

	return locationData, nil
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
func GetLocation(ctx context.Context, ip string) (data *IpLocationData, err error) {
	if !validate.IsIp(ip) {
		return nil, fmt.Errorf("invalid input ip:%v", ip)
	}

	if validate.IsLocalIPAddr(ip) {
		return nil, fmt.Errorf("must be a public ip:%v", ip)
	}

	if cacheMap.Contains(ip) {
		value := cacheMap.Get(ip)
		data1, ok := value.(*IpLocationData)
		if !ok {
			cacheMap.Remove(ip)
			err = fmt.Errorf("data assertion failed in the cache ip:%v", ip)
			return
		}
		return data1, nil
	}

	mode := g.Cfg().MustGet(ctx, "hotgo.ipMethod", "cz88").String()
	switch mode {
	case "whois":
		data, err = WhoisLocation(ctx, ip)
	default:
		data, err = Cz88Find(ctx, ip)
	}

	if err == nil && data != nil {
		if cacheMap.Size() > 20000 {
			cacheMap.Clear()
		}
		cacheMap.Set(ip, data)
	}

	return
}

// GetPublicIP 获取公网IP
func GetPublicIP(ctx context.Context) (ip string, err error) {
	var data *WhoisRegionData
	err = g.Client().Timeout(10*time.Second).GetVar(ctx, whoisApi).Scan(&data)
	if err != nil {
		g.Log().Info(ctx, "GetPublicIP fail, alternatives are being tried.")
		return GetPublicIP2()
	}

	if data == nil {
		g.Log().Info(ctx, "publicIP address Parsing failure, check the network and firewall blocking.")
		return "0.0.0.0", nil
	}
	return data.Ip, nil
}

func GetPublicIP2() (ip string, err error) {
	response, err := http.Get(dyndns)
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	ip = strings.ReplaceAll(string(body), "\n", "")
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
	if r == nil {
		return ""
	}
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.GetClientIp()
	}

	// 兼容部分云厂商CDN，如果存在多个，默认取第一个
	if gstr.Contains(ip, ",") {
		ip = gstr.StrTillEx(ip, ",")
	}

	if gstr.Contains(ip, ", ") {
		ip = gstr.StrTillEx(ip, ", ")
	}

	return ip
}
