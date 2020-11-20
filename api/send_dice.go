package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendDice represents method to send an animated emoji that will display a random value.
// On success, the sent Message is returned.
func (api *TelegramAPI) SendDice(ctx context.Context, req apiModels.SendDiceRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendDice").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
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
