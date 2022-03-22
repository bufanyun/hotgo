//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package apiForm

import (
	"github.com/bufanyun/hotgo/app/com"
	"github.com/gogf/gf/v2/frame/g"
)

//  获取lang信息
type BaseLangReq struct {
	g.Meta `path:"/base/lang" method:"get" tags:"基础接口" summary:"获取lang信息"`
	L      string `json:"l" v:"required#语言不能为空" dc:"语言"`
}
type BaseLangRes struct {
}

//  获取登录验证码
type IpLocationReq struct {
	g.Meta `path:"/base/ip_location" method:"get" tags:"基础接口" summary:"获取IP归属地信息"`
	Ip     string `json:"ip" v:"required#ip不能为空" dc:"ipv4地址"`
}
type IpLocationRes struct {
	com.IpLocationData
}

type ExportReq struct {
	g.Meta `path:"/base/export" method:"get" tags:"字典接口" summary:"导出字典类型"`
}
type ExportRes struct{}
