// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// EmsLogEditInp 修改/新增数据
type EmsLogEditInp struct {
	entity.SysEmsLog
}

type EmsLogEditModel struct{}

// EmsLogDeleteInp 删除
type EmsLogDeleteInp struct {
	Id interface{} `json:"id" v:"required#邮件记录ID不能为空" dc:"邮件记录ID"`
}

type EmsLogDeleteModel struct{}

// EmsLogViewInp 获取信息
type EmsLogViewInp struct {
	Id int64 `json:"id" v:"required#邮件记录ID不能为空" dc:"邮件记录ID"`
}

type EmsLogViewModel struct {
	entity.SysEmsLog
}

// EmsLogListInp 获取列表
type EmsLogListInp struct {
	form.PageReq

	form.StatusReq
	Title   string `json:"title"`
	Content string `json:"content"`
}

type EmsLogListModel struct {
	entity.SysEmsLog
}

// EmsLogStatusInp 更新状态
type EmsLogStatusInp struct {
	entity.SysSmsLog
}

type EmsLogStatusModel struct{}

// SendEmsInp 发送邮件
type SendEmsInp struct {
	Event    string `json:"event"   v:"required#邮件事件不能为空"    description:"事件"`
	Email    string `json:"email"   v:"required#邮箱地址不能为空"   description:"邮箱地址"`
	Code     string `json:"code"      description:"验证码或短信内容"`
	Content  string `json:"content"      description:"邮件内容"`
	Template string `json:"-"         description:"发信模板"`
	TplData  g.Map  `json:"-" description:"模板变量"`
}

// VerifyEmsCodeInp 效验验证码
type VerifyEmsCodeInp struct {
	Event string `json:"event"   v:"required#邮件事件不能为空"    description:"事件"`
	Email string `json:"email"   v:"required#邮箱地址不能为空"   description:"邮箱地址"`
	Code  string `json:"code"      description:"验证码"`
}
