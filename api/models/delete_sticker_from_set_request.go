package models

// DeleteStickerFromSetRequest represents `deleteStickerFromSet` request body.
type DeleteStickerFromSetRequest struct {
	// File identifier of the sticker.
	Sticker string `json:"sticker"`
}
