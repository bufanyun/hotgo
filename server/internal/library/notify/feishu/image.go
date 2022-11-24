package feishu

type ImageMessage struct {
	MsgType MsgType      `json:"msg_type"`
	Content ImageContent `json:"content"`
}

type ImageContent struct {
	ImageKey string `json:"image_key"`
}

func (m *ImageMessage) Body() map[string]interface{} {
	m.MsgType = MsgTypeImage
	return structToMap(m)
}

func NewImageMessage() *ImageMessage {
	return &ImageMessage{}
}

func (m *ImageMessage) SetImageKey(key string) *ImageMessage {
	m.Content.ImageKey = key
	return m
}
