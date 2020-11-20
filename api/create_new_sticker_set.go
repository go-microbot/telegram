package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
)

// CreateNewStickerSet represents method to create a new sticker set owned by a user.
// The bot will be able to edit the sticker set thus created.
// You must use exactly one of the fields png_sticker or tgs_sticker. Returns True on success.
func (api *TelegramAPI) CreateNewStickerSet(ctx context.Context, req apiModels.CreateNewStickerSetRequest) error {
	_, err := api.NewRequest("createNewStickerSet").
		Method(http.MethodPost).
		Body(NewFormDataBody(req)).
		Do(ctx)

	return err
}
