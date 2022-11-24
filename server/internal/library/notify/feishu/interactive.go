package feishu

type InteractiveMessage struct {
	MsgType MsgType `json:"msg_type"`
	Card    string  `json:"card"`
}

func (m *InteractiveMessage) Body() map[string]interface{} {
	m.MsgType = MsgTypeInteractive
	return structToMap(m)
}

func NewInteractiveMessage() *InteractiveMessage {
	return &InteractiveMessage{}
}

// SetCard set card with cardbuilder https://open.feishu.cn/tool/cardbuilder?from=custom_bot_doc
func (m *InteractiveMessage) SetCard(card string) *InteractiveMessage {
	m.Card = card
	return m
}
