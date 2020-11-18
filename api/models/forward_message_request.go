package models

import "github.com/go-microbot/telegram/query"

// ForwardMessageRequest represents `forwardMessage` request body.
type ForwardMessageRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Unique identifier for the chat where the original message
	// was sent (or channel username in the format @channelusername).
	FromChatID query.ParamAny `query:"from_chat_id" json:"-"`
	// Optional. Sends the message silently (https://telegram.org/blog/channels-2-0#silent-messages).
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Message identifier in the chat specified in from_chat_id.
	MessageID int32 `json:"message_id"`
}
