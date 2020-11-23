package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
)

// UploadStickerFile represents method to upload a .PNG file with a sticker for later use
// in createNewStickerSet and addStickerToSet methods (can be used multiple times).
// Returns the uploaded File on success.
func (api *TelegramAPI) UploadStickerFile(ctx context.Context, req apiModels.UploadStickerFileRequest) (*models.File, error) {
	resp, err := api.NewRequest("uploadStickerFile").
		Method(http.MethodPost).
		Body(NewFormDataBody(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var file models.File
	if err := resp.Decode(&file); err != nil {
		return nil, err
	}

	return &file, nil
}
