package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// StopPoll represents method to stop a poll which was sent by the bot.
// On success, the stopped Poll with the final results is returned.
func (api *TelegramAPI) StopPoll(ctx context.Context, req apiModels.StopPollRequest) (*models.Poll, error) {
	resp, err := api.NewRequest("stopPoll").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var poll models.Poll
	if err := resp.Decode(&poll); err != nil {
		return nil, err
	}

	return &poll, nil
}
