package models

const MESSAGE_SIMPLE = "messageSimple"

type Update struct {
	updateType string
	chatID     int64
	text       string
	keyboard   [][]KeyboardButton
}

func NewUpdate(chatID int64) *Update {
	return &Update{
		chatID: chatID,
	}
}

func (update *Update) SetText(text string) *Update {
	update.text = text

	return update
}

func (update *Update) SetUpdateType(updateType string) *Update {
	update.updateType = updateType

	return update
}

func (update *Update) AddKeyboardRows(rows ...[]KeyboardButton) *Update {
	update.keyboard = append(update.keyboard, rows...)

	return update
}

func (update *Update) GetKeyboardRows() [][]KeyboardButton {
	return update.keyboard
}

func (update *Update) GetChatID() int64 {
	return update.chatID
}

func (update *Update) GetText() string {
	return update.text
}

func (update *Update) IsUpdateType(updateType string) bool {
	return update.updateType == updateType
}
