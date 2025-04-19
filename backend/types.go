package main

type MessageType uint32

const (
	TextMessage MessageType = iota
	JoinMessage
	LeaveMessage
	ErrorMessage 
)

type Message struct {
	sender *Client
	data   *ChatMessage
}

type ChatMessage struct {
	Content   string      `json:"content"`
	Type      MessageType `json:"type,omitempty"`
	TimeStamp string      `json:"timestamp,omitempty"`
	Username  string      `json:"username,omitempty"`
}
