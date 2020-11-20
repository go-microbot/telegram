package models

import "github.com/go-microbot/telegram/query"

// DeleteMessageRequest represents `deleteMessage` request body.
type DeleteMessageRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id"`
	// Identifier of the message to delete.
	MessageID query.ParamInt `query:"message_id"`
}
