package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendMediaGroup represents method to send a group of photos, videos, documents or audios
// as an album. Documents and audio files can be only grouped in an album with messages
// of the same type. On success, an array of Messages that were sent is returned.
func (api *TelegramAPI) SendMediaGroup(ctx context.Context, req apiModels.SendMediaGroupRequest) ([]models.Message, error) {
	resp, err := api.NewRequest("sendMediaGroup").
		Query(query.AsMap(req)).
		Body(NewJSONBody(req)).
		Method(http.MethodPost).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var messages []models.Message
	if err := resp.Decode(&messages); err != nil {
		return nil, err
	}

	return messages, nil
}
