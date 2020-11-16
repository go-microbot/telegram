package api

import (
	"context"
)

// DeleteWebhook removes webhook integration if you decide
// to switch back to getUpdates. Returns True on success.
func (api *TelegramAPI) DeleteWebhook(ctx context.Context) error {
	_, err := api.NewRequest("deleteWebhook").Do(ctx)
	return err
}
