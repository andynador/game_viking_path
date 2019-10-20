package models

type Update struct {
	chatID int64
	text   *string
}

func NewUpdate(chatID int64, text *string) Update {
	return Update{
		chatID: chatID,
		text:   text,
	}
}

func (update Update) GetChatID() int64 {
	return update.chatID
}

func (update Update) GetText() *string {
	return update.text
}
