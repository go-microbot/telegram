package models

import (
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
)

// SendVideoNoteRequest represents `sendVideoNote` request body.
type SendVideoNoteRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id"`
	// Video note to send. Pass a file_id as String to send a video note that exists on the
	// Telegram servers (recommended) or upload a new video using multipart/form-data.
	// Sending video notes by a URL is currently unsupported.
	VideoNote form.Part `form:"video_note"`
	// Optional. Duration of sent video in seconds.
	Duration form.PartAny `form:"duration,omitempty"`
	// Optional. Video width and height, i.e. diameter of the video message.
	Length form.PartAny `form:"length,omitempty"`
	// Optional. Thumbnail of the file sent;
	// can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file,
	// so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded
	// using multipart/form-data under <file_attach_name>.
	Thumb form.PartAny `form:"thumb,omitempty"`
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
