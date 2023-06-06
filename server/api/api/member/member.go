// Package member
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package member

import "github.com/gogf/gf/v2/frame/g"

// GetIdByCodeReq 通过邀请码获取用户ID
type GetIdByCodeReq struct {
	g.Meta `path:"/member/getIdByCode" method:"post" tags:"用户" summary:"通过邀请码获取用户ID"`
	Code   string `json:"code"   dc:"邀请码"`
}

type GetIdByCodeRes struct{}
