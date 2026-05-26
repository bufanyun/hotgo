// Package sms
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sms

import (
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/sysin"
)

type AliYunDrive struct{}

// SendCode 发送验证码
func (d *AliYunDrive) SendCode(ctx context.Context, in *sysin.SendCodeInp) (err error) {
	client, err := d.CreateClient(tea.String(config.AliYunAccessKeyID), tea.String(config.AliYunAccessKeySecret))
	if err != nil {
		return err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(in.Mobile),
		SignName:      tea.String(config.AliYunSign),
		TemplateCode:  tea.String(in.Template),
		TemplateParam: tea.String(fmt.Sprintf("{\"code\":\"%v\"}", in.Code)),
	}

	return func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		// 复制代码运行请自行打印 API 的返回值
		response, err := client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if err != nil {
			return err
		}

		g.Log().Debugf(ctx, "aliyun.sendCode response:%+v", response.GoString())

		return nil
	}()
}

// CreateClient 使用AK&SK初始化账号Client
func (d *AliYunDrive) CreateClient(accessKeyId *string, accessKeySecret *string) (result *dysmsapi20170525.Client, err error) {
	conf := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	conf.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	return dysmsapi20170525.NewClient(conf)
}
