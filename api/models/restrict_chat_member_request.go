package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// RestrictChatMemberRequest represents `restrictChatMember` request body.
type RestrictChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target
	// supergroup (in the format @supergroupusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Unique identifier of the target user.
	UserID query.ParamInt `query:"user_id" json:"-"`
	// A JSON-serialized object for new user permissions.
	Permissions models.ChatPermissions `json:"permissions"`
	// Optional. Date when restrictions will be lifted for the user, unix time.
	// If user is restricted for more than 366 days or less than 30 seconds
	// from the current time, they are considered to be restricted forever.
	UntilDate int32 `json:"until_date,omitempty"`
}
