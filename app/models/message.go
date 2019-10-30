package models

type Message struct {
	chatID int64
	text   *string
}

func NewMessage(chatID int64, text *string) Message {
	return Message{
		chatID: chatID,
		text:   text,
	}
}

func (message Message) GetChatID() int64 {
	return message.chatID
}

func (message Message) GetText() *string {
	return message.text
}
