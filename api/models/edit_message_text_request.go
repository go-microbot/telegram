package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// EditMessageTextRequest represents `editMessageText` request body.
type EditMessageTextRequest struct {
	// Optional. Required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id,omitempty" json:"-"`
	// Optional. Required if inline_message_id is not specified.
	// Identifier of the message to edit.
	MessageID query.ParamAny `query:"message_id,omitempty" json:"-"`
	// Optional. Required if chat_id and message_id are not specified.
	// Identifier of the inline message.
	InlineMessageID query.ParamString `query:"inline_message_id,omitempty" json:"-"`
	// New text of the message, 1-4096 characters after entities parsing.
	Text string `json:"text"`
	// Optional. Mode for parsing entities in the message text.
	// See formatting options (https://core.telegram.org/bots/api#formatting-options)
	// for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. List of special entities that appear in message text,
	// which can be specified instead of parse_mode
	Entities []models.MessageEntity `json:"entities,omitempty"`
	// Optional. Disables link previews for links in this message.
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	// Optional. A JSON-serialized object for an
	// inline keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	ReplyMarkup *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
