package base

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/websocketin"
)

// SendToTagReq 发送标签消息
type SendToTagReq struct {
	g.Meta `path:"/send/toTag" method:"post" tags:"WebSocket" summary:"发送标签消息"`
	websocketin.SendToTagInp
}

type SendToTagRes struct {
}
