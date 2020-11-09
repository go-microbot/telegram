package models

// Video represents a video file.
// https://core.telegram.org/bots/api#video.
type Video struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`
	// Unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUID string `json:"file_unique_id"`
	// Video width as defined by sender.
	Width int32 `json:"width"`
	// Video height as defined by sender.
	Height int32 `json:"height"`
	// Duration of the video in seconds as defined by sender.
	Duration int32 `json:"duration"`
	// Optional. Video thumbnail.
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Original filename as defined by sender.
	FileName string `json:"file_name,omitempty"`
	// Optional. Mime type of a file as defined by sender.
	MIMEType string `json:"mime_type,omitempty"`
	// Optional. File size.
	FileSize int32 `json:"file_size,omitempty"`
}
