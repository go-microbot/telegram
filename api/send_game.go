package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendGame represents method to send a game. On success, the sent Message is returned.
func (api *TelegramAPI) SendGame(ctx context.Context, req apiModels.SendGameRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendGame").
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
