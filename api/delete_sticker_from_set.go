package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
)

// DeleteStickerFromSet represents method to delete a sticker from a set created by the bot.
// Returns True on success.
func (api *TelegramAPI) DeleteStickerFromSet(ctx context.Context, req apiModels.DeleteStickerFromSetRequest) error {
	_, err := api.NewRequest("deleteStickerFromSet").
		Body(NewJSONBody(req)).
		Method(http.MethodPost).
		Do(ctx)

	return err
}
