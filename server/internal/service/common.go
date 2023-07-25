// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/model/input/commonin"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ICommonUpload interface {
		// UploadFile 上传文件
		UploadFile(ctx context.Context, uploadType string, file *ghttp.UploadFile) (res *sysin.AttachmentListModel, err error)
	}
	ICommonWechat interface {
		// Authorize 微信用户授权
		Authorize(ctx context.Context, in *commonin.WechatAuthorizeInp) (res *commonin.WechatAuthorizeModel, err error)
		AuthorizeCall(ctx context.Context, in *commonin.WechatAuthorizeCallInp) (res *commonin.WechatAuthorizeCallModel, err error)
		// GetOpenId 从缓存中获取临时openid
		GetOpenId(ctx context.Context) (openId string, err error)
		GetCacheKey(typ, ak string) string
		// CleanTempMap 清理临时map
		CleanTempMap(ctx context.Context)
	}
)

var (
	localCommonUpload ICommonUpload
	localCommonWechat ICommonWechat
)

func CommonUpload() ICommonUpload {
	if localCommonUpload == nil {
		panic("implement not found for interface ICommonUpload, forgot register?")
	}
	return localCommonUpload
}

func RegisterCommonUpload(i ICommonUpload) {
	localCommonUpload = i
}

func CommonWechat() ICommonWechat {
	if localCommonWechat == nil {
		panic("implement not found for interface ICommonWechat, forgot register?")
	}
	return localCommonWechat
}

func RegisterCommonWechat(i ICommonWechat) {
	localCommonWechat = i
}
