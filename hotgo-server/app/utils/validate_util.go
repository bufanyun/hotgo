//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package utils

import (
	"net"
)

// 验证类
var Validate = new(validate)

type validate struct{}

//
//  @Title  是否为ipv4
//  @Description 
//  @Author  Ms <133814250@qq.com>
//  @Param   ip
//  @Return  bool
//
func (util *validate) IsIp(ip string) bool {
	if net.ParseIP(ip) != nil {
		return true
	}
	return false
}
