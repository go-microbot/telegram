package models

import (
	"github.com/go-microbot/telegram/form"
)

// UploadStickerFileRequest represents `uploadStickerFile` request body.
type UploadStickerFileRequest struct {
	// User identifier of sticker file owner.
	UserID form.PartAny `form:"user_id"`
	// PNG image with the sticker, must be up to 512 kilobytes in size,
	// dimensions must not exceed 512px, and either width or height must be exactly 512px.
	PngSticker form.PartFile `form:"png_sticker"`
}
