package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
)

// AddStickerToSet represents method to add a new sticker to a set created by the bot.
// You must use exactly one of the fields png_sticker or tgs_sticker.
// Animated stickers can be added to animated sticker sets and only to them.
// Animated sticker sets can have up to 50 stickers.
// Static sticker sets can have up to 120 stickers. Returns True on success.
func (api *TelegramAPI) AddStickerToSet(ctx context.Context, req apiModels.AddStickerToSetRequest) error {
	_, err := api.NewRequest("addStickerToSet").
		Method(http.MethodPost).
		Body(NewFormDataBody(req)).
		Do(ctx)

	return err
}
