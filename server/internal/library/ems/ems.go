// Package ems
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package ems

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/model"
	"hotgo/utility/validate"
	"net/smtp"
	"strings"
)

// Send 发送邮件入口
func Send(config *model.EmailConfig, to string, subject string, body string) error {
	return sendToMail(config, to, subject, body, "html")
}

func sendToMail(config *model.EmailConfig, to, subject, body, mailType string) error {
	if config == nil {
		return gerror.New("邮件配置不能为空")
	}

	var (
		contentType string
		auth        = smtp.PlainAuth("", config.User, config.Password, config.Host)
		sendTo      = strings.Split(to, ";")
	)

	if len(sendTo) == 0 {
		return gerror.New("收件人不能为空")
	}

	for _, em := range sendTo {
		if !validate.IsEmail(em) {
			return gerror.Newf("邮件格式不正确，请检查：%v", em)
		}
	}

	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + config.SendName + "<" + config.User + ">" + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	return smtp.SendMail(config.Addr, auth, config.User, sendTo, msg)
}
