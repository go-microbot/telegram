package models

import "github.com/go-microbot/telegram/query"

// PinChatMessageRequest represents `pinChatMessage` request body.
type PinChatMessageRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Identifier of a message to pin.
	MessageID int32 `json:"message_id"`
	// Optional. Pass True, if it is not necessary to send a notification
	// to all chat members about the new pinned message.
	// Notifications are always disabled in channels and private chats.
	DisableNotification bool `json:"disable_notification,omitempty"`
}
