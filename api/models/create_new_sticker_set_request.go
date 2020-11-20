package models

import "github.com/go-microbot/telegram/form"

// CreateNewStickerSetRequest represents `createNewStickerSet` request body.
type CreateNewStickerSetRequest struct {
	// User identifier of created sticker set owner.
	UserID form.PartAny `form:"user_id"`
	// Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals).
	// Can contain only english letters, digits and underscores.
	// Must begin with a letter, can't contain consecutive underscores and must end in
	// "_by_<bot username>". <bot_username> is case insensitive. 1-64 characters.
	Name form.PartText `form:"name"`
	// Sticker set title, 1-64 characters.
	Title form.PartText `form:"title"`
	// Optional. PNG image with the sticker, must be up to 512 kilobytes in size,
	// dimensions must not exceed 512px, and either width or height must be exactly 512px.
	// Pass a file_id as a String to send a file that already exists on the Telegram servers,
	// pass an HTTP URL as a String for Telegram to get a file from the Internet,
	// or upload a new one using multipart/form-data.
	PngSticker form.Part `form:"png_sticker,omitempty"`
	// Optional. TGS animation with the sticker, uploaded using multipart/form-data.
	// See https://core.telegram.org/animated_stickers#technical-requirements
	// for technical requirements.
	TgsSticker form.PartFile `form:"tgs_sticker,omitempty"`
	// One or more emoji corresponding to the sticker.
	Emojis form.PartText `form:"emojis"`
	// Optional. Pass True, if a set of mask stickers should be created.
	ContainsMasks form.PartAny `form:"contains_masks,omitempty"`
	// Optional. A JSON-serialized object for position where the mask should be placed on faces.
	MaskPosition form.PartJSON `form:"mask_position,omitempty"`
}
