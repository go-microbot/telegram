package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendVideoNote represents method to send video messages.
// On success, the sent Message is returned.
// support rounded square mp4 videos of up to 1 minute long
func (api *TelegramAPI) SendVideoNote(ctx context.Context, req apiModels.SendVideoNoteRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendVideoNote").
		Query(query.AsMap(req)).
		Body(NewFormDataBody(req)).
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
