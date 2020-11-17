package models

import (
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
)

// SendPhotoRequest represents `sendPhoto` body.
type SendPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id"`
	// Photo to send. Pass a file_id as String to send a photo that exists
	// on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram
	// to get a photo from the Internet, or upload a new photo using multipart/form-data.
	// Initialize with InputFile object to send a new (non existing) photo.
	Photo form.PartFile `form:"photo"`
	// Optional. Photo caption (may also be used when resending photos by file_id),
	// 0-1024 characters after entities parsing.
	/*Caption form.PartText `form:"caption,omitempty"`
	// Optional. Mode for parsing entities in the photo caption.
	// See formatting options (https://core.telegram.org/bots/api#formatting-options)
	// for more details.
	ParseMode form.PartText `json:"parse_mode,omitempty"`
	// Optional. List of special entities that appear in the caption,
	// which can be specified instead of parse_mode.
	CaptionEntities []models.MessageEntity `json:"caption_entities,omitempty"` // TODO:
	// Optional. Sends the message
	// silently (https://telegram.org/blog/channels-2-0#silent-messages).
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message.
	ReplyToMessageID *int32 `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent
	// even if the specified replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to remove reply keyboard
	// or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`*/
}
