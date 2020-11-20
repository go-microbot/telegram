package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendDocument represents method to send general files. On success, the sent Message is returned.
// Bots can currently send files of any type of up to 50 MB in size,
// this limit may be changed in the future.
func (api *TelegramAPI) SendDocument(ctx context.Context, req apiModels.SendDocumentRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendDocument").
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
