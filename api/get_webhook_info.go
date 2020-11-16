package api

import (
	"context"

	"github.com/go-microbot/telegram/models"
)

// GetWebhookInfo gets current webhook status. Requires no parameters.
// On success, returns a WebhookInfo (https://core.telegram.org/bots/api#webhookinfo) object.
// If the bot is using getUpdates (https://core.telegram.org/bots/api#getupdates),
// will return an object with the url field empty.
func (api *TelegramAPI) GetWebhookInfo(ctx context.Context) (*models.WebhookInfo, error) {
	resp, err := api.NewRequest("getWebhookInfo").Do(ctx)
	if err != nil {
		return nil, err
	}

	var status models.WebhookInfo
	if err = resp.Decode(&status); err != nil {
		return nil, err
	}

	return &status, nil
}
