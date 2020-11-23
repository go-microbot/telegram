package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendContact represents method to send phone contacts.
// On success, the sent Message is returned.
func (api *TelegramAPI) SendContact(ctx context.Context, req apiModels.SendContactRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendContact").
		Query(query.AsMap(req)).
		Body(NewJSONBody(req)).
		Method(http.MethodPost).
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
