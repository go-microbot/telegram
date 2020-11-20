package models

import "github.com/go-microbot/telegram/query"

// SetChatStickerSetRequest represents `setChatStickerSet` body request.
type SetChatStickerSetRequest struct {
	// Unique identifier for the target chat or username of the target
	// supergroup (in the format @supergroupusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Name of the sticker set to be set as the group sticker set.
	StickerSetName string `json:"sticker_set_name"`
}
