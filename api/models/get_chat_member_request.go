package models

import "github.com/go-microbot/telegram/query"

// GetChatMemberRequest represents `getChatMember` request body.
type GetChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id"`
	// Unique identifier of the target user.
	UserID query.ParamInt `query:"user_id"`
}
