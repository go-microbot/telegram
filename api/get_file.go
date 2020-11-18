package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// GetFile represents method to get basic info about a file and prepare it for downloading.
// For the moment, bots can download files of up to 20MB in size.
// On success, a File object is returned.
// The file can then be downloaded via the link
// https://api.telegram.org/file/bot<token>/<file_path>,
// where <file_path> is taken from the response.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling
// getFile (https://core.telegram.org/bots/api#getfile) again.
func (api *TelegramAPI) GetFile(ctx context.Context, req apiModels.FileID) (*models.File, error) {
	resp, err := api.NewRequest("getFile").
		Query(query.AsMap(req)).
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
