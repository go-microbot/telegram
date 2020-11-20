package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendVoice represents method to send audio files, if you want Telegram clients
// to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with OPUS
// (other formats may be sent as Audio or Document). On success, the sent Message is returned.
// Bots can currently send voice messages of up to 50 MB in size,
// this limit may be changed in the future.
func (api *TelegramAPI) SendVoice(ctx context.Context, req apiModels.SendVoiceRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendVoice").
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
