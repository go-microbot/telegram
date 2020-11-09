package models

// Document represents a general file (as opposed to photos, voice messages and audio files).
// https://core.telegram.org/bots/api#document.
type Document struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`
	// Unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUID string `json:"file_unique_id"`
	// Optional. Document thumbnail as defined by sender.
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Original filename as defined by sender.
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file as defined by sender.
	MIMEType string `json:"mime_type,omitempty"`
	// Optional. File size.
	FileSize int32 `json:"file_size"`
}
