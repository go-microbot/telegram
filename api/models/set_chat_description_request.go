package models

import "github.com/go-microbot/telegram/query"

// SetChatDescriptionRequest represents `setChatDescription` request body.
type SetChatDescriptionRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Optional. New chat description, 0-255 characters.
	Description string `json:"description,omitempty"`
}
