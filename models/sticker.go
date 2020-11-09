package models

// Sticker represents a sticker model.
type Sticker struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`
	// Unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUID string `json:"file_unique_id"`
	// Sticker width.
	Width int32 `json:"width"`
	// Sticker height.
	Height int32 `json:"height"`
	// True, if the sticker is animated.
	IsAnimated bool `json:"is_animated"`
	// Optional. Sticker thumbnail in the .WEBP or .JPG format.
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Emoji associated with the sticker.
	Emoji string `json:"emoji,omitempty"`
	// Optional. Name of the sticker set to which the sticker belongs.
	SetName string `json:"set_name,omitempty"`
	// Optional. For mask stickers, the position where the mask should be placed.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	// Optional. File size.
	FileSize int32 `json:"file_size,omitempty"`
}
