//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysService

import (
	"context"
	"github.com/bufanyun/hotgo/app/com"
)

var Provinces = new(provinces)

type provinces struct{}

//
//  @Title  获取地区中的省市编码
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   location
//
func (service *provinces) GetLocationCode(ctx context.Context, location com.IpLocationData)  {

	return
}

