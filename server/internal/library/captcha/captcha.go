// Package captcha
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package captcha

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mojocn/base64Captcha"
)

// GetVerifyImgString 生成验证码
func GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string) {
	driver := &base64Captcha.DriverString{
		Height: 80,
		Width:  240,
		//NoiseCount:      50,
		//ShowLineOptions: 20,
		Length: 4,
		Source: "abcdefghjkmnpqrstuvwxyz23456789",
		Fonts:  []string{"chromohv.ttf"},
	}
	driver = driver.ConvertFonts()
	store := base64Captcha.DefaultMemStore
	c := base64Captcha.NewCaptcha(driver, store)
	idKeyC, base64stringC, err := c.Generate()
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

// VerifyString 验证输入的验证码是否正确
func VerifyString(id, answer string) bool {
	driver := new(base64Captcha.DriverString)
	store := base64Captcha.DefaultMemStore
	c := base64Captcha.NewCaptcha(driver, store)
	answer = gstr.ToLower(answer)
	return c.Verify(id, answer, true)
}
