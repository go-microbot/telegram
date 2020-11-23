package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// GetGameHighScores represents method to get data for high score tables.
// Will return the score of the specified user and several of their neighbors in a game.
// On success, returns an Array of GameHighScore objects.
func (api *TelegramAPI) GetGameHighScores(ctx context.Context, req apiModels.GetGameHighScoresRequest) ([]models.GameHighScore, error) {
	resp, err := api.NewRequest("getGameHighScores").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var highScores []models.GameHighScore
	if err := resp.Decode(&highScores); err != nil {
		return nil, err
	}

	return highScores, nil
}
