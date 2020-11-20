package models

import "github.com/go-microbot/telegram/query"

// SendChatActionRequest represents `sendChatAction` request body.
type SendChatActionRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Type of action to broadcast.
	// Choose one, depending on what the user is about to receive:
	// typing for text messages, upload_photo for photos, record_video or upload_video for videos,
	// record_voice or upload_voice for voice notes, upload_document for general files,
	// find_location for location data, record_video_note or upload_video_note for video notes.
	//
	// Example: The ImageBot needs some time to process a request and upload the image.
	// Instead of sending a text message along the lines of "Retrieving image, please wait…",
	// the bot may use sendChatAction with action = upload_photo.
	// The user will see a “sending photo” status for the bot.
	Action string `json:"action"`
}
