package models

import (
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
)

// SendDocumentRequest represents `sendDocument` request body.
type SendDocumentRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id"`
	// File to send. Pass a file_id as String to send a file that exists
	// on the Telegram servers (recommended), pass an HTTP URL as a String
	// for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	Document form.Part `form:"document"`
	// Optional. Thumbnail of the file sent;
	// can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file,
	// so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded
	// using multipart/form-data under <file_attach_name>.
	Thumb form.PartAny `form:"thumb,omitempty"`
	// Optional. Document caption (may also be used when resending documents by file_id),
	// 0-1024 characters after entities parsing.
	Caption form.PartText `form:"caption,omitempty"`
	// Optional. Mode for parsing entities in the document caption.
	// See formatting options (https://core.telegram.org/bots/api#formatting-options)
	// for more details.
	ParseMode form.PartText `form:"parse_mode,omitempty"`
	// Optional. List of special entities that appear in the caption,
	// which can be specified instead of parse_mode.
	CaptionEntries form.PartJSON `form:"caption_entities,omitempty"`
	// Optional. Disables automatic server-side content type detection
	// for files uploaded using multipart/form-data.
	DisableContentTypeDetection form.PartAny `form:"disable_content_type_detection,omitempty"`
	// Optional. Sends the message
	// silently (https://telegram.org/blog/channels-2-0#silent-messages).
	// Users will receive a notification with no sound.
	DisableNotification form.PartAny `form:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message.
	ReplyToMessageID form.PartAny `form:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even
	// if the specified replied-to message is not found.
	AllowSendingWithoutReply form.PartAny `form:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to remove reply keyboard
	// or to force a reply from the user.
	ReplyMarkup form.PartJSON `form:"reply_markup,omitempty"`
}
