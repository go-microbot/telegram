package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// CopyMessageRequest represents `copyMessage` request body.
type CopyMessageRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Unique identifier for the chat where the original message was sent
	// (or channel username in the format @channelusername).
	FromChatID query.ParamAny `query:"from_chat_id" json:"-"`
	// Message identifier in the chat specified in from_chat_id.
	MessageID query.ParamInt `query:"message_id" json:"-"`
	// Optional. New caption for media, 0-1024 characters after entities parsing.
	// If not specified, the original caption is kept.
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the new caption.
	// See formatting options (https://core.telegram.org/bots/api#formatting-options)
	// for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. List of special entities that appear in the new caption,
	// which can be specified instead of parse_mode.
	CaptionEntities []models.MessageEntity `json:"caption_entities,omitempty"`
	// Optional. Sends the message
	// silently (https://telegram.org/blog/channels-2-0#silent-messages).
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message.
	ReplyToMessageID *int32 `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to remove reply keyboard
	// or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
