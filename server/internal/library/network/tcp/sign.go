package tcp

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/model/input/msgin"
)

type Sign interface {
	SetSign(appId, secretKey string) *msgin.RpcMsg
	SetTraceID(traceID string)
}

// PkgSign 打包签名
func PkgSign(data interface{}, appId, secretKey, traceID string) *msgin.RpcMsg {
	if c, ok := data.(Sign); ok {
		c.SetTraceID(traceID)
		return c.SetSign(appId, secretKey)
	}
	return nil
}

// VerifySign 验证签名
func VerifySign(data interface{}, appId, secretKey string) (in *msgin.RpcMsg, err error) {
	// 无密钥，无需签名
	if secretKey == "" {
		return
	}

	if err = gconv.Scan(data, &in); err != nil {
		return
	}

	if appId != in.AppId {
		err = gerror.New("appId invalid")
		return
	}

	if in.Sign != in.GetSign(secretKey) {
		err = gerror.New("sign invalid")
		return
	}
	return
}
