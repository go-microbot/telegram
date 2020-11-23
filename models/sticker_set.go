package models

// StickerSet represents a sticker set.
// https://core.telegram.org/bots/api#stickerset.
type StickerSet struct {
	// Sticker set name.
	Name string `json:"name"`
	// Sticker set title.
	Title string `json:"title"`
	// True, if the sticker set contains animated stickers.
	IsAnimated bool `json:"is_animated"`
	// True, if the sticker set contains masks.
	ContainsMasks bool `json:"contains_masks"`
	// List of all set stickers.
	Stickers []Sticker `json:"stickers"`
	// Optional. Sticker set thumbnail in the .WEBP or .TGS format.
	Thumb *PhotoSize `json:"thumb,omitempty"`
}
