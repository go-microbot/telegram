package models

// InputMediaPhoto represents a photo to be sent.
// https://core.telegram.org/bots/api#inputmediaphoto.
type InputMediaPhoto struct {
	// Type of the result, must be photo.
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram
	// servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass "attach://<file_attach_name>" to upload a new one using
	// multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing.
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the photo caption.
	// See formatting options (https://core.telegram.org/bots/api#formatting-options)
	// for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. List of special entities that appear in the caption,
	// which can be specified instead of parse_mode.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
}
