package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendPoll represents method to send a native poll. On success, the sent Message is returned.
func (api *TelegramAPI) SendPoll(ctx context.Context, req apiModels.SendPollRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendPoll").
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
