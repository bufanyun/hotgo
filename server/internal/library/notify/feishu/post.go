package feishu

import (
	"encoding/json"
	"log"
)

type PostMessage struct {
	MsgType MsgType     `json:"msg_type"`
	Content PostContent `json:"content"`
}

func NewPostMessage() *PostMessage {
	return &PostMessage{}
}

func (m *PostMessage) Body() map[string]interface{} {
	m.MsgType = MsgTypePost
	return structToMap(m)
}

func (m *PostMessage) SetZH(u PostUnit) *PostMessage {
	m.Content.Post.ZH = u
	return m
}

func (m *PostMessage) SetZHTitle(t string) *PostMessage {
	m.Content.Post.ZH.Title = t
	return m
}

func (m *PostMessage) AppendZHContent(i []PostItem) *PostMessage {
	m.Content.Post.ZH.Content = append(m.Content.Post.ZH.Content, i)
	return m
}

func (m *PostMessage) SetJA(u PostUnit) *PostMessage {
	m.Content.Post.JA = u
	return m
}

func (m *PostMessage) SetJATitle(t string) *PostMessage {
	m.Content.Post.JA.Title = t
	return m
}

func (m *PostMessage) AppendJAContent(i []PostItem) *PostMessage {
	m.Content.Post.JA.Content = append(m.Content.Post.JA.Content, i)
	return m
}

func (m *PostMessage) SetEN(u PostUnit) *PostMessage {
	m.Content.Post.EN = u
	return m
}

func (m *PostMessage) SetENTitle(t string) *PostMessage {
	m.Content.Post.EN.Title = t
	return m
}

func (m *PostMessage) AppendENContent(i []PostItem) *PostMessage {
	m.Content.Post.EN.Content = append(m.Content.Post.EN.Content, i)
	return m
}

type PostContent struct {
	Post PostBody `json:"post"`
}

type PostBody struct {
	ZH PostUnit `json:"zh_cn,omitempty"`
	JA PostUnit `json:"ja_jp,omitempty"`
	EN PostUnit `json:"en_us,omitempty"`
}

type PostUnit struct {
	Title   string       `json:"title,omitempty"`
	Content [][]PostItem `json:"content"`
}

type PostItem interface{}

type Text struct {
	Tag      string `json:"tag"`
	Text     string `json:"text"`
	UnEscape bool   `json:"un_escape,omitempty"`
}

func NewText(text string) Text {
	t := Text{
		Tag:  "text",
		Text: text,
	}
	return t
}

type A struct {
	Tag      string `json:"tag"`
	Text     string `json:"text"`
	Href     string `json:"href"`
	UnEscape bool   `json:"un_escape,omitempty"`
}

func NewA(text, href string) A {
	t := A{
		Tag:  "a",
		Text: text,
		Href: href,
	}
	return t
}

type AT struct {
	Tag    string `json:"tag"`
	UserID string `json:"user_id"`
}

func NewAT(userID string) AT {
	t := AT{
		Tag:    "at",
		UserID: userID,
	}
	return t
}

type Image struct {
	Tag      string `json:"tag"`
	ImageKey string `json:"image_key"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

func NewImage(imageKey string, height, width int) Image {
	t := Image{
		Tag:      "image",
		ImageKey: imageKey,
		Height:   height,
		Width:    width,
	}
	return t
}

type PostCMDMessage struct {
	MsgType MsgType        `json:"msg_type"`
	Content PostCMDContent `json:"content"`
}

func (m *PostCMDMessage) Body() map[string]interface{} {
	m.MsgType = MsgTypePost
	return structToMap(m)
}

type PostCMDContent struct {
	Post map[string]interface{} `json:"post"`
}

func NewPostCMDMessage() *PostCMDMessage {
	return &PostCMDMessage{}
}

func (m *PostCMDMessage) SetPost(post string) *PostCMDMessage {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(post), &result)
	if err != nil {
		log.Print("SetPost err: ", err)
	}
	m.Content.Post = result
	return m
}
