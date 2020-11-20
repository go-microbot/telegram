package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendAnimation represents method to send animation files
// (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned.
// Bots can currently send animation files of up to 50 MB in size,
// this limit may be changed in the future.
func (api *TelegramAPI) SendAnimation(ctx context.Context, req apiModels.SendAnimationRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendAnimation").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewFormDataBody(req)).
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
