package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// SetWebhook sets a webhook.
// Use this method to specify a url and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request
// to the specified url, containing a JSON-serialized Update
// (https://core.telegram.org/bots/api#update). In case of an unsuccessful request,
// we will give up after a reasonable amount of attempts. Returns True on success.
//
// If you'd like to make sure that the Webhook request comes from Telegram,
// we recommend using a secret path in the URL, e.g. https://www.example.com/<token>.
// Since nobody else knows your bot's token, you can be pretty sure it's us.
func (api *TelegramAPI) SetWebhook(ctx context.Context, req apiModels.SetWebhookRequest) error {
	request := api.NewRequest("setWebhook").
		Method(http.MethodPost).
		Query(query.AsMap(req)).
		Body(NewFormDataBody(req))

	_, err := request.Do(ctx)
	return err
}
