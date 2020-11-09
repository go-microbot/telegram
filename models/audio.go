package models

// Audio represents an audio file to be treated as music by the Telegram clients.
// https://core.telegram.org/bots/api#audio.
type Audio struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`
	// Unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUID string `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender.
	Duration int32 `json:"duration"`
	// Optional. Performer of the audio as defined by sender or by audio tags.
	Performer string `json:"performer,omitempty"`
	// Optional. Title of the audio as defined by sender or by audio tags.
	Title string `json:"title,omitempty"`
	// Optional. Original filename as defined by sender.
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file as defined by sender.
	MIMEType string `json:"mime_type,omitempty"`
	// Optional. File size.
	FileSize int32 `json:"file_size,omitempty"`
	// Optional. Thumbnail of the album cover to which the music file belongs.
	Thumb *PhotoSize `json:"thumb,omitempty"`
}
