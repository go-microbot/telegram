package models

import "github.com/go-microbot/telegram/query"

// UnpinChatMessageRequest represents `unpinChatMessage` request body.
type UnpinChatMessageRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Optional. Identifier of a message to unpin.
	// If not specified, the most recent pinned message (by sending date) will be unpinned.
	MessageID *int32 `json:"message_id,omitempty"`
}
