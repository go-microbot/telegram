package models

import "github.com/go-microbot/telegram/query"

// SetChatAdminCustomTitleRequest represents `setChatAdministratorCustomTitle` request body.
type SetChatAdminCustomTitleRequest struct {
	// Unique identifier for the target chat or username of the target
	// supergroup (in the format @supergroupusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Unique identifier of the target user.
	UserID int32 `json:"user_id"`
	// New custom title for the administrator; 0-16 characters, emoji are not allowed.
	CustomTitle string `json:"custom_title"`
}
