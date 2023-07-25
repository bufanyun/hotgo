package feishu

type ShareChatMessage struct {
	MsgType MsgType          `json:"msg_type"`
	Content ShareChatContent `json:"content"`
}

type ShareChatContent struct {
	ShareChatID string `json:"share_chat_id"`
}

func (m *ShareChatMessage) Body() map[string]interface{} {
	m.MsgType = MsgTypeShareChat
	return structToMap(m)
}

func NewShareChatMessage() *ShareChatMessage {
	return &ShareChatMessage{}
}

func (m *ShareChatMessage) SetShareChatID(id string) *ShareChatMessage {
	m.Content.ShareChatID = id
	return m
}
