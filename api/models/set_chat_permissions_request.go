package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SetChatPermissionsRequest represents `setChatPermissions` request body.
type SetChatPermissionsRequest struct {
	// Unique identifier for the target chat or username of the target
	// supergroup (in the format @supergroupusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// New default chat permissions.
	Permissions models.ChatPermissions `json:"permissions"`
}
