package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendVideo represents method to send video files, Telegram clients support mp4 videos
// (other formats may be sent as Document). On success, the sent Message is returned.
// Bots can currently send video files of up to 50 MB in size,
// this limit may be changed in the future.
func (api *TelegramAPI) SendVideo(ctx context.Context, req apiModels.SendVideoRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendVideo").
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
