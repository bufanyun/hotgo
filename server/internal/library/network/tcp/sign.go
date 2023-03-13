package tcp

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/model/input/msgin"
)

type Sign interface {
	SetSign(traceID, appId, secretKey string)
}

// SetSign 设置签名
func SetSign(data interface{}, traceID, appId, secretKey string) {
	if c, ok := data.(Sign); ok {
		c.SetSign(traceID, appId, secretKey)
		return
	}
}

// VerifySign 验证签名
func VerifySign(data interface{}, appId, secretKey string) (err error) {
	// 无密钥，无需签名
	if secretKey == "" {
		return
	}

	var in *msgin.Request
	if err = gconv.Scan(data, &in); err != nil {
		return
	}

	if appId != in.AppId {
		return gerror.New("appId invalid")
	}

	if in.Sign != in.GetSign(secretKey) {
		return gerror.New("sign invalid")
	}
	return
}
