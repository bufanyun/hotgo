package feishu

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"strconv"
	"time"

	"hotgo/internal/library/notify/feishu/internal/security"
)

const feishuAPI = "https://open.feishu.cn/open-apis/bot/v2/hook/"

// Client feishu client
type Client struct {
	AccessToken string
	Secret      string
}

// NewClient new client
func NewClient(accessToken, secret string) *Client {
	return &Client{
		AccessToken: accessToken,
		Secret:      secret,
	}
}

// Response response struct
type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`

	Extra         interface{} `json:"Extra"`
	StatusCode    int64       `json:"StatusCode"`
	StatusMessage string      `json:"StatusMessage"`
}

// Send send message
func (d *Client) Send(message Message) (string, *Response, error) {
	res := &Response{}

	if len(d.AccessToken) < 1 {
		return "", res, fmt.Errorf("accessToken is empty")
	}

	timestamp := time.Now().Unix()
	sign, err := security.GenSign(d.Secret, timestamp)
	if err != nil {
		return "", res, err
	}

	body := message.Body()
	body["timestamp"] = strconv.FormatInt(timestamp, 10)
	body["sign"] = sign

	reqBytes, err := json.Marshal(body)
	if err != nil {
		return "", res, err
	}

	var (
		result    *Response
		URL       = fmt.Sprintf("%v%v", feishuAPI, d.AccessToken)
		reqString = string(reqBytes)
	)

	g.Client().
		Retry(3, time.Second).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		PostVar(gctx.New(), URL, &result)
	if result.Code != 0 {
		return reqString, result, fmt.Errorf("send message to feishu error = %s", result.Msg)
	}
	return reqString, result, nil
}
