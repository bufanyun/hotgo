//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package com

import (
	"context"
	"time"

	"github.com/axgle/mahonia"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/bufanyun/hotgo/app/utils"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kayon/iploc"
)

// IP归属地
var Ip = new(ip)

type ip struct{}

type IpLocationData struct {
	Ip           string `json:"ip"`
	Country      string `json:"country"`
	Region       string `json:"region"`
	Province     string `json:"province"`
	ProvinceCode int    `json:"province_code"`
	City         string `json:"city"`
	CityCode     int    `json:"city_code"`
	Area         string `json:"area"`
	AreaCode     int    `json:"area_code"`
}

//
//  @Title  通过Whois接口查询IP归属地
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   ip
//  @Return  IpLocationData
//
func (component *ip) WhoisLocation(ctx context.Context, ip string) IpLocationData {

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

	if !utils.Validate.IsIp(ip) {
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

	g.Log().Print(ctx, "data:", data)
	whoisData := whoisRegionData{}
	if err := gconv.Struct(data, &whoisData); err != nil {
		err = gerror.New(err.Error())

		g.Log().Print(ctx, "err:", err)
		return IpLocationData{
			Ip: ip,
		}
	}

	g.Log().Print(ctx, "whoisData:", whoisData)

	return IpLocationData{
		Ip: whoisData.Ip,
		//Country      string `json:"country"`
		Region:       whoisData.Addr,
		Province:     whoisData.Pro,
		ProvinceCode: gconv.Int(whoisData.ProCode),
		City:         whoisData.City,
		CityCode:     gconv.Int(whoisData.CityCode),
		Area:         whoisData.Region,
		AreaCode:     gconv.Int(whoisData.RegionCode),
	}
}

//
//  @Title  通过Cz88的IP库查询IP归属地
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   ip
//  @Return  IpLocationData
//
func (component *ip) Cz88Find(ctx context.Context, ip string) IpLocationData {
	if !utils.Validate.IsIp(ip) {
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

	var (
		provinceModel *entity.SysProvinces
		cityModel     *entity.SysProvinces
		areaModel     *entity.SysProvinces
	)

	err = g.DB().Model("hg_sys_provinces").
		Where("level", 1).
		WhereLike("title", "%"+locationData.Province+"%").
		Scan(&provinceModel)

	if err != nil {
		err = gerror.New(err.Error())
		return locationData
	}

	if provinceModel != nil {
		locationData.ProvinceCode = provinceModel.Id
		locationData.Province = provinceModel.Title
	}

	if gstr.LenRune(locationData.City) == 0 {
		return locationData

		// 是否为直辖市
	} else if component.IsJurisdictionByIpTitle(locationData.City) {
		locationData.CityCode = provinceModel.Id + 100
		locationData.City = "直辖市"
	} else {

		//替换掉
		locationData.City = gstr.Replace(locationData.City, "地区", "")

		err = g.DB().Model("hg_sys_provinces").
			Where("level", 2).
			Where("pid", locationData.ProvinceCode).
			WhereLike("title", "%"+locationData.City+"%").
			Scan(&cityModel)

		if err != nil {
			err = gerror.New(err.Error())
			return locationData
		}

		if cityModel != nil {
			locationData.CityCode = cityModel.Id
			locationData.City = cityModel.Title
		}
	}

	if gstr.LenRune(locationData.Area) == 0 {
		return locationData
	}

	err = g.DB().Model("hg_sys_provinces").
		Where("level", 3).
		Where("pid", locationData.CityCode).
		WhereLike("title", "%"+locationData.Area+"%").
		Scan(&areaModel)

	if err != nil {
		err = gerror.New(err.Error())
		return locationData
	}

	if areaModel != nil {
		locationData.AreaCode = areaModel.Id
		locationData.Area = areaModel.Title
	}

	return locationData
}

//
//  @Title  判断地区名称是否为直辖市
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   title
//  @Return  bool
//
func (component *ip) IsJurisdictionByIpTitle(title string) bool {

	lists := []string{"北京市", "天津市", "重庆市", "上海市"}

	for i := 0; i < len(lists); i++ {
		if gstr.Contains(lists[i], title) {
			return true
		}
	}
	return false
}

//
//  @Title  获取IP归属地信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   ip
//  @Return  IpLocationData
//
func (component *ip) GetLocation(ctx context.Context, ip string) IpLocationData {
	method, _ := g.Cfg().Get(ctx, "hotgo.ipMethod", "cz88")

	if method.String() == "whois" {
		return component.WhoisLocation(ctx, ip)
	}
	return component.Cz88Find(ctx, ip)
}
