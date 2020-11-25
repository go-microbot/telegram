package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SetGameScore represents method to set the score of the specified user in a game.
// On success, if the message was sent by the bot, returns the edited Message,
// otherwise returns True. Returns an error, if the new score is not greater than the user's
// current score in the chat and force is False.
func (api *TelegramAPI) SetGameScore(ctx context.Context, req apiModels.SetGameScoreRequest) (*models.Message, error) {
	resp, err := api.NewRequest("setGameScore").
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
