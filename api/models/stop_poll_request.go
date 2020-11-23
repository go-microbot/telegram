package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// StopPollRequest represents `stopPoll` request body.
type StopPollRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Identifier of the original message with the poll.
	MessageID query.ParamInt `query:"message_id" json:"-"`
	// Optional. A JSON-serialized object for a new message
	// inline keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	ReplyMarkup *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
