package models

import "github.com/go-microbot/telegram/query"

// GetStickerSetRequest represents `getStickerSet` request body.
type GetStickerSetRequest struct {
	// Name of the sticker set.
	Name query.ParamString `query:"name"`
}
