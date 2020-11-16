package models

import "github.com/go-microbot/telegram/query"

// ChatID represents chat ID query model.
type ChatID struct {
	// Unique identifier for the target chat or username of the target supergroup
	// or channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id"`
}
