package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// GetUpdates uses to receive incoming updates using long polling
// (https://en.wikipedia.org/wiki/Push_technology#Long_polling).
// An Array of Update objects is returned.
func (api *TelegramAPI) GetUpdates(ctx context.Context, req apiModels.GetUpdatesRequest) ([]models.Update, error) {
	resp, err := api.NewRequest("getUpdates").
		Query(query.AsMap(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var updates []models.Update
	if err := resp.Decode(&updates); err != nil {
		return nil, err
	}

	return updates, nil
}

// GetPollUpdates uses to receive incoming updates using long polling
// (https://en.wikipedia.org/wiki/Push_technology#Long_polling).
// Uses custom http client to get messages.
// An Array of Update objects is returned.
func (api *TelegramAPI) GetPollUpdates(ctx context.Context, req apiModels.GetUpdatesRequest, client *http.Client) ([]models.Update, error) {
	resp, err := api.NewRequest("getUpdates").
		Query(query.AsMap(req)).
		CustomClient(client).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var updates []models.Update
	if err := resp.Decode(&updates); err != nil {
		return nil, err
	}

	return updates, nil
}
