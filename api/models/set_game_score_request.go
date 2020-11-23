package models

import "github.com/go-microbot/telegram/query"

// SetGameScoreRequest represents `setGameScore` request body.
type SetGameScoreRequest struct {
	// User identifier.
	UserID int32 `json:"user_id"`
	// New score, must be non-negative.
	Score int32 `json:"score"`
	// Optional. Pass True, if the high score is allowed to decrease.
	// This can be useful when fixing mistakes or banning cheaters.
	Force bool `json:"force,omitempty"`
	// Optional. Pass True, if the game message should not be automatically
	// edited to include the current scoreboard.
	DisableEditMessage bool `json:"disable_edit_message,omitempty"`
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
