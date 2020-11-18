package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendMessageRequest represents `sendMessage` body.
type SendMessageRequest struct {
	// Unique identifier for the target chat.
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Text of the message to be sent, 1-4096 characters after entities parsing.
	Text string `json:"text"`
	// Optional. Mode for parsing entities in the message text. See formatting options
	// (https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. List of special entities that appear in message text,
	// which can be specified instead of parse_mode.
	Entities []models.MessageEntity `json:"entities,omitempty"`
	// Optional. Disables link previews for links in this message.
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message.
	ReplyToMessageID *int32 `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent
	// even if the specified replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options.
	// A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
