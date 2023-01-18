package location

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type AMapGeocodeAddressRes struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Count    string `json:"count"`
	Geocodes []struct {
		FormattedAddress string        `json:"formatted_address"`
		Country          string        `json:"country"`
		Province         string        `json:"province"`
		Citycode         string        `json:"citycode"`
		City             string        `json:"city"`
		District         string        `json:"district"`
		Township         []interface{} `json:"township"`
		Neighborhood     struct {
			Name []interface{} `json:"name"`
			Type []interface{} `json:"type"`
		} `json:"neighborhood"`
		Building struct {
			Name []interface{} `json:"name"`
			Type []interface{} `json:"type"`
		} `json:"building"`
		Adcode   string        `json:"adcode"`
		Street   []interface{} `json:"street"`
		Number   []interface{} `json:"number"`
		Location string        `json:"location"`
		Level    string        `json:"level"`
	} `json:"geocodes"`
}

type AddressRegion struct {
	ProvinceName string `json:"provinceName"`
	CityName     string `json:"cityName"`
	CountyName   string `json:"countyName"`
	ProvinceCode string `json:"provinceCode"`
	CityCode     string `json:"cityCode"`
	CountyCode   string `json:"countyCode"`
}

// AnalysisAddress 将地址解析出省市区编码
func AnalysisAddress(ctx context.Context, address, key string) (region *AddressRegion, err error) {
	var (
		url         = fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?address=%v&output=JSON&key=%v", address, key)
		responseMap = make(g.Map)
		response    *AMapGeocodeAddressRes
	)

	err = g.Client().GetVar(ctx, url).Scan(&responseMap)
	if err != nil {
		return nil, gerror.Newf("AMap AnalysisAddress err:%+v", err)
	}

	err = gconv.Scan(responseMap, &response)
	if err != nil {
		return nil, err
	}
	// 异常状态码
	if response.Status != "1" {
		return nil, gerror.Newf("AMap AnalysisAddress 错误码：%+v, 错误提示：%+v", response.Status, response.Info)
	}

	if len(response.Geocodes) == 0 {
		return nil, gerror.New("AMap AnalysisAddress 没有解析到地区信息")
	}

	region = new(AddressRegion)
	region.ProvinceName = response.Geocodes[0].Province
	region.CityName = response.Geocodes[0].City
	region.CountyName = response.Geocodes[0].District

	// 有效区域编码
	if len(response.Geocodes[0].Adcode) == 6 {
		code := gconv.Int64(response.Geocodes[0].Adcode)
		if code > 0 {
			region.ProvinceCode = gconv.String(code / 10000 * 10000)
			region.CityCode = gconv.String(code / 100 * 100)
			region.CountyCode = response.Geocodes[0].Adcode
		}
	}
	return
}
