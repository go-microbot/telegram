package models

import "github.com/go-microbot/telegram/query"

// UnbanChatMemberRequest represents `unbanChatMember` request body.
type UnbanChatMemberRequest struct {
	// Unique identifier for the target group or username of the target supergroup or
	// channel (in the format @username).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Unique identifier of the target user.
	UserID int32 `json:"user_id"`
	// Optional. Do nothing if the user is not banned.
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
}
