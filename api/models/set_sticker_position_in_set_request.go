package models

// SetStickerPositionInSetRequest represents `setStickerPositionInSet` request body.
type SetStickerPositionInSetRequest struct {
	// File identifier of the sticker.
	Sticker string `json:"sticker"`
	// New sticker position in the set, zero-based.
	Position int32 `json:"position"`
}
