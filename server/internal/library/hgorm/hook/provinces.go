package hook

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/library/location"
)

// CityLabel 城市地区标签
var CityLabel = gdb.HookHandler{
	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
		result, err = in.Next(ctx)
		if err != nil {
			return
		}

		parse := func(id int64, index int) {
			cityLabel, err := location.ParseSimpleRegion(ctx, id)
			if err != nil {
				g.Log().Warningf(ctx, "hook.CityLabel parse err:%+v", err)
			}
			result[index]["cityLabel"] = gvar.New(cityLabel)
			return
		}

		for i, record := range result {
			// 优先级： 区 > 市 > 省

			cityId, ok := record["city_id"]
			if ok && !cityId.IsEmpty() {
				parse(cityId.Int64(), i)
				continue
			}

			provinceId, ok := record["province_id"]
			if ok && !provinceId.IsEmpty() {
				parse(cityId.Int64(), i)
				continue
			}

			// 以下是默认关联表 省市区字段

			sysLogCityId, ok := record["sysLogCityId"]
			if ok && !sysLogCityId.IsEmpty() {
				parse(sysLogCityId.Int64(), i)
				continue
			}

			sysLogProvinceId, ok := record["sysLogProvinceId"]
			if ok && !sysLogProvinceId.IsEmpty() {
				parse(cityId.Int64(), i)
				continue
			}
		}
		return
	},
}
