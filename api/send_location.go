package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendLocation represents method to send point on the map.
// On success, the sent Message is returned.
func (api *TelegramAPI) SendLocation(ctx context.Context, req apiModels.SendLocationRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendLocation").
		Query(query.AsMap(req)).
		Body(NewJSONBody(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var msg models.Message
	if err := resp.Decode(&msg); err != nil {
		return nil, err
	}

	return &msg, nil
}
