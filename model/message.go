package model

type Message struct {
	Id         int64  `json:"id,omitempty"`
	ChatKey    string `json:"chat_key,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}
