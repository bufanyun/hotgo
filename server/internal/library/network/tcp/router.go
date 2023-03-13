package tcp

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/util/gconv"
)

type RouterHandler func(args ...interface{})

// Message 路由消息
type Message struct {
	Router string      `json:"router"`
	Data   interface{} `json:"data"`
}

func SendPkg(conn *gtcp.Conn, message *Message) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return conn.SendPkg(b)
}

func RecvPkg(conn *gtcp.Conn) (*Message, error) {
	if data, err := conn.RecvPkg(); err != nil {
		return nil, err
	} else {
		var msg = new(Message)
		if err = gconv.Scan(data, &msg); err != nil {
			return nil, gerror.Newf("invalid package structure: %s", err.Error())
		}
		if msg.Router == "" {
			return nil, gerror.Newf("message is not routed: %+v", msg)
		}
		return msg, err
	}
}
