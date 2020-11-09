package models

// Voice represents a voice note.
// https://core.telegram.org/bots/api#voice.
type Voice struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`
	// Unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUID string `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender.
	Duration int32 `json:"duration"`
	// Optional. MIME type of the file as defined by sender.
	MIMEType string `json:"mime_type,omitempty"`
	// Optional. File size.
	FileSize int32 `json:"file_size,omitempty"`
}
