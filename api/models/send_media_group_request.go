package models

import (
	"github.com/go-microbot/telegram/query"
)

// SendMediaGroupRequest represents `sendMediaGroup` request body.
type SendMediaGroupRequest struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// A JSON-serialized array describing messages to be sent, must include 2-10 items.
	Media interface{} `json:"media"`
	// Optional. Sends messages
	// silently (https://telegram.org/blog/channels-2-0#silent-messages).
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the messages are a reply, ID of the original message.
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
}
