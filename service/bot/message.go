package bot

type Update struct {
	chatID int64
	text string
}

func (update Update) GetChatID() int64 {
	return update.chatID
}

func NewUpdate(chatID int64, text string) Update {
	return Update{
		chatID: chatID,
		text: text,
	}
}
