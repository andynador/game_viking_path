package models

const MESSAGE_WITH_KEYBOARD = "messageWithKeyboard"

type KeyboardButton struct {
	text string
}

func NewKeyboardButton(text string) KeyboardButton {
	return KeyboardButton{
		text: text,
	}
}

func (KeyboardButton KeyboardButton) GetText() string  {
	return KeyboardButton.text
}

func NewKeyboardButtonRow(buttons ...KeyboardButton) []KeyboardButton {
	var row []KeyboardButton

	row = append(row, buttons...)

	return row
}

