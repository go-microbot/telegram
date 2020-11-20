package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendAudio represents method to send audio files, if you want Telegram clients to display them
// in the music player. Your audio must be in the .MP3 or .M4A format.
// On success, the sent Message is returned.
// Bots can currently send audio files of up to 50 MB in size,
// this limit may be changed in the future.
func (api *TelegramAPI) SendAudio(ctx context.Context, req apiModels.SendAudioRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendAudio").
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
