package models

import "github.com/go-microbot/telegram/query"

// GetGameHighScoresRequest represents `getGameHighScores` request body.
type GetGameHighScoresRequest struct {
	// Target user id.
	UserID int32 `json:"user_id"`
	// Optional. Required if inline_message_id is not specified.
	// Unique identifier for the target chat.
	ChatID query.ParamAny `query:"chat_id,omitempty" json:"-"`
	// Optional. Required if inline_message_id is not specified.
	// Identifier of the sent message.
	MessageID query.ParamInt `query:"message_id,omitempty" json:"-"`
	// Optional. Required if chat_id and message_id are not specified.
	// Identifier of the inline message.
	InlineMessageID string `json:"inline_message_id,omitempty"`
}
