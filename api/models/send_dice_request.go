package models

import "github.com/go-microbot/telegram/query"

// SendDiceRequest represents `sendDice` request body.
type SendDiceRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Optional. Emoji on which the dice throw animation is based.
	// Currently, must be one of "🎲", “🎯", “🏀", “⚽", or “🎰".
	// Dice can have values 1-6 for “🎲" and “🎯", values 1-5 for “🏀" and
	// “⚽", and values 1-64 for “🎰". Defaults to “🎲".
	Emoji string `json:"emoji,omitempty"`
	// Optional. Sends the message
	// silently (https://telegram.org/blog/channels-2-0#silent-messages).
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message.
	ReplyToMessageID *int32 `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even
	// if the specified replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to remove reply keyboard
	// or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
