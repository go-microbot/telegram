package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// ForwardMessage represents method to forward messages of any kind.
// On success, the sent Message is returned.
func (api *TelegramAPI) ForwardMessage(ctx context.Context, req apiModels.ForwardMessageRequest) (*models.Message, error) {
	resp, err := api.NewRequest("forwardMessage").
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
