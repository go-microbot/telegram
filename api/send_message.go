package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
)

// SendMessage sends text messages.
// On success, the sent Message (https://core.telegram.org/bots/api#message) is returned.
func (api *TelegramAPI) SendMessage(ctx context.Context, req apiModels.SendMessageRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendMessage").
		Method(http.MethodPost).
		Body(req).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err = resp.Decode(&message); err != nil {
		return nil, err
	}

	return &message, nil
}
