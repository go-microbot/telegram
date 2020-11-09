package models

// PhotoSize represents one size of a photo or a file/sticker thumbnail.
// https://core.telegram.org/bots/api#photosize.
type PhotoSize struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`
	// Unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUID string `json:"file_unique_id"`
	// Photo width.
	Width int32 `json:"width"`
	// Photo height.
	Height int32 `json:"height"`
	// Optional. File size.
	FileSize int32 `json:"file_size,omitempty"`
}
