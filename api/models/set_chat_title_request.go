package models

import "github.com/go-microbot/telegram/query"

// SetChatTitleRequest represents `setChatTitle` body.
type SetChatTitleRequest struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// New chat title, 1-255 characters.
	Title string `json:"title"`
}
