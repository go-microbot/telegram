package models

import "github.com/go-microbot/telegram/query"

// KickChatMemberRequest represents `kickChatMember` request body.
type KickChatMemberRequest struct {
	// Unique identifier for the target group or username of the target supergroup or
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Unique identifier of the target user.
	UserID int32 `json:"user_id"`
	// Optional. Date when the user will be unbanned, unix time.
	// If user is banned for more than 366 days or less than 30 seconds
	// from the current time they are considered to be banned forever.
	UntilDate *int32 `json:"until_date,omitempty"`
}
