// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// SmsLogEditInp 修改/新增数据
type SmsLogEditInp struct {
	entity.SysSmsLog
}

type SmsLogEditModel struct{}

// SmsLogDeleteInp 删除
type SmsLogDeleteInp struct {
	Id interface{} `json:"id" v:"required#短信记录ID不能为空" dc:"短信记录ID"`
}

type SmsLogDeleteModel struct{}

// SmsLogViewInp 获取信息
type SmsLogViewInp struct {
	Id int64 `json:"id" v:"required#短信记录ID不能为空" dc:"短信记录ID"`
}

type SmsLogViewModel struct {
	entity.SysSmsLog
}

// SmsLogListInp 获取列表
type SmsLogListInp struct {
	form.PageReq
	form.StatusReq
	Mobile    string        `json:"mobile"    dc:"手机号"`
	Ip        string        `json:"ip"        dc:"IP"`
	Event     string        `json:"event"     dc:"事件"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *SmsLogListInp) Filter(ctx context.Context) (err error) {
	if in.Event != "" {
		if _, ok := consts.SmsTemplateEventMap[in.Event]; !ok {
			err = gerror.Newf("无效的事件类型:%v", in.Event)
			return
		}
	}
	return
}

type SmsLogListModel struct {
	entity.SysSmsLog
}

// SmsLogStatusInp 更新状态
type SmsLogStatusInp struct {
	entity.SysSmsLog
}

type SmsLogStatusModel struct{}

// SendCodeInp 发送验证码
type SendCodeInp struct {
	Event    string `json:"event"     description:"事件"`  // 必填
	Mobile   string `json:"mobile"    description:"手机号"` // 必填
	Code     string `json:"code"      description:"验证码或短信内容"`
	Template string `json:"-"         description:"发信模板 "`
}

// VerifyCodeInp 效验验证码
type VerifyCodeInp struct {
	Event  string `json:"event"     description:"事件"`       // 必填
	Mobile string `json:"mobile"    description:"手机号"`      // 必填
	Code   string `json:"code"      description:"验证码或短信内容"` // 必填
}
