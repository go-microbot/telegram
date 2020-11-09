package models

// InlineKeyboardMarkup represents an inline keyboard
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating)
// that appears right next to the message it belongs to.
// https://core.telegram.org/bots/api#inlinekeyboardmarkup.
type InlineKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of InlineKeyboardButton objects.
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}
