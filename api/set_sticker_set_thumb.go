package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
)

// SetStickerSetThumb represents method to set the thumbnail of a sticker set.
// Animated thumbnails can be set for animated sticker sets only. Returns True on success.
func (api *TelegramAPI) SetStickerSetThumb(ctx context.Context, req apiModels.SetStickerSetThumbRequest) error {
	_, err := api.NewRequest("setStickerSetThumb").
		Body(NewFormDataBody(req)).
		Method(http.MethodPost).
		Do(ctx)

	return err
}
