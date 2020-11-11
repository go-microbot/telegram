package models

// CallbackQuery represents an incoming callback query from a callback button
// in an inline keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
// If the button that originated the query was attached to a message sent by the bot,
// the field message will be present. If the button was attached to a message sent via the bot
// (in inline mode) (https://core.telegram.org/bots/api#inline-mode),
// the field inline_message_id will be present.
// Exactly one of the fields data or game_short_name will be present.
// https://core.telegram.org/bots/api#callbackquery.
type CallbackQuery struct {
	// Unique identifier for this query.
	ID string `json:"id"`
	// Sender.
	From User `json:"from"`
	// Optional. Message with the callback button that originated the query.
	// Note that message content and message date will not be available if the message is too old.
	Message *Message `json:"message,omitempty"`
	// Optional. Identifier of the message sent via the bot in inline mode,
	// that originated the query.
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// Global identifier, uniquely corresponding to the chat to which the message with
	// the callback button was sent.
	// Useful for high scores in games (https://core.telegram.org/bots/api#games).
	ChatInstance string `json:"chat_instance,omitempty"`
	// Optional. Data associated with the callback button.
	// Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data,omitempty"`
	// Optional. Short name of a Game (https://core.telegram.org/bots/api#games) to be returned,
	// serves as the unique identifier for the game.
	GameShortName string `json:"game_short_name,omitempty"`
}
