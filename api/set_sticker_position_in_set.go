package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
)

// SetStickerPositionInSet represents method to move a sticker in a set
// created by the bot to a specific position. Returns True on success.
func (api *TelegramAPI) SetStickerPositionInSet(ctx context.Context, req apiModels.SetStickerPositionInSetRequest) error {
	_, err := api.NewRequest("setStickerPositionInSet").
		Body(NewJSONBody(req)).
		Method(http.MethodPost).
		Do(ctx)

	return err
}
