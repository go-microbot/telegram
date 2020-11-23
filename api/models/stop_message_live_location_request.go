package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// StopMessageLiveLocationRequest represents `stopMessageLiveLocation` request body.
type StopMessageLiveLocationRequest struct {
	// Optional. Required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id,omitempty" json:"-"`
	// Optional. Required if inline_message_id is not specified.
	// Identifier of the message with live location to stop.
	MessageID query.ParamInt `query:"message_id,omitempty" json:"-"`
	// Optional. Required if chat_id and message_id are not specified.
	// Identifier of the inline message.
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// Optional. A JSON-serialized object for a new
	// inline keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	ReplyMarkup *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
