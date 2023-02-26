package aliyun

import (
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"hotgo/internal/model"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

// SendCode 发送验证码
func SendCode(ctx context.Context, in sysin.SendCodeInp, config *model.SmsConfig) (err error) {
	if config == nil {
		config, err = service.SysConfig().GetSms(ctx)
		if err != nil {
			return err
		}
	}

	client, err := CreateClient(tea.String(config.SmsAliyunAccessKeyID), tea.String(config.SmsAliyunAccessKeySecret))
	if err != nil {
		return err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(in.Mobile),
		SignName:      tea.String(config.SmsAliyunSign),
		TemplateCode:  tea.String(in.Template),
		TemplateParam: tea.String(fmt.Sprintf("{\"code\":\"%v\"}", in.Code)),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		// 复制代码运行请自行打印 API 的返回值
		_, err = client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if err != nil {
			return err
		}

		return nil
	}()
	return tryErr
}

// CreateClient 使用AK&SK初始化账号Client
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func TestSend(accessKeyId string, accessKeySecret string) (_err error) {
	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String("15303830571"),
		SignName:      tea.String("布帆云"),
		TemplateCode:  tea.String("SMS_198921686"),
		TemplateParam: tea.String("{\"code\":\"1234\"}"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var err = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			err = _t
		} else {
			err.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(err.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}
