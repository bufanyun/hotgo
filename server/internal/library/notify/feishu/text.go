package feishu

type TextMessage struct {
	MsgType MsgType `json:"msg_type"`
	Content Content `json:"content"`
}

type Content struct {
	Text string `json:"text"`
}

func (m *TextMessage) Body() map[string]interface{} {
	m.MsgType = MsgTypeText
	return structToMap(m)
}

func NewTextMessage() *TextMessage {
	return &TextMessage{}
}

func (m *TextMessage) SetText(text string) *TextMessage {
	m.Content.Text = text
	return m
}
