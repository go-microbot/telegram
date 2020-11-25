package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// EditMessageCaptionRequest represents `editMessageCaption` request body.
type EditMessageCaptionRequest struct {
	// Optional. Required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id,omitempty" json:"-"`
	// Optional. Required if inline_message_id is not specified.
	// Identifier of the message to edit.
	MessageID query.ParamInt `query:"message_id,omitempty" json:"-"`
	// Optional. Required if chat_id and message_id are not specified.
	// Identifier of the inline message.
	InlineMessageID query.ParamString `query:"inline_message_id,omitempty" json:"-"`
	// Optional. New caption of the message, 0-1024 characters after entities parsing.
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the message caption.
	// See formatting options (https://core.telegram.org/bots/api#formatting-options)
	// for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. List of special entities that appear in the caption,
	// which can be specified instead of parse_mode.
	CaptionEntities []models.MessageEntity `json:"caption_entities,omitempty"`
	// Optional. A JSON-serialized object for an
	// inline keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	ReplyMarkup *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
