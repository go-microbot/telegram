package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// TelegramAPI represents Telegram Bot API.
type TelegramAPI struct {
	client *http.Client
	token  string
}

// NewTelegramAPI returns new TelegramAPI instance.
func NewTelegramAPI(token string) TelegramAPI {
	return TelegramAPI{
		token:  token,
		client: http.DefaultClient,
	}
}

// GetMe returns basic information about the bot.
func (api *TelegramAPI) GetMe(ctx context.Context) (*models.User, error) {
	var me models.User
	if err := api.newRequest(ctx, apiRequest{
		apiMethod:  "getMe",
		httpMethod: http.MethodGet,
		response:   &me,
	}); err != nil {
		return nil, err
	}
	return &me, nil
}

// GetUpdates uses to receive incoming updates using long polling
// (https://en.wikipedia.org/wiki/Push_technology#Long_polling).
// An Array of Update objects is returned.
func (api *TelegramAPI) GetUpdates(ctx context.Context, req apiModels.GetUpdatesRequest) ([]models.Update, error) {
	var updates []models.Update
	if err := api.newRequest(ctx, apiRequest{
		apiMethod:  "getUpdates",
		httpMethod: http.MethodGet,
		response:   &updates,
		query:      query.AsMap(req),
	}); err != nil {
		return nil, err
	}
	return updates, nil
}

// GetPollUpdates uses to receive incoming updates using long polling
// (https://en.wikipedia.org/wiki/Push_technology#Long_polling).
// Uses custom http client to get messages.
// An Array of Update objects is returned.
func (api *TelegramAPI) GetPollUpdates(ctx context.Context, req apiModels.GetUpdatesRequest, client *http.Client) ([]models.Update, error) {
	var updates []models.Update
	if err := api.newRequest(ctx, apiRequest{
		apiMethod:    "getUpdates",
		httpMethod:   http.MethodGet,
		response:     &updates,
		query:        query.AsMap(req),
		customClient: client,
	}); err != nil {
		return nil, err
	}
	return updates, nil
}

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
	request := apiRequest{
		apiMethod:  "setWebhook",
		httpMethod: http.MethodPost,
		query:      query.AsMap(req),
	}
	if req.Certificate != nil {
		request.formData = map[string]string{
			"certificate": string(*req.Certificate),
		}
	}

	return api.newRequest(ctx, request)
}

// GetWebhookInfo gets current webhook status. Requires no parameters.
// On success, returns a WebhookInfo (https://core.telegram.org/bots/api#webhookinfo) object.
// If the bot is using getUpdates (https://core.telegram.org/bots/api#getupdates),
// will return an object with the url field empty.
func (api *TelegramAPI) GetWebhookInfo(ctx context.Context) (*models.WebhookInfo, error) {
	var status models.WebhookInfo
	if err := api.newRequest(ctx, apiRequest{
		apiMethod:  "getWebhookInfo",
		httpMethod: http.MethodGet,
		response:   &status,
	}); err != nil {
		return nil, err
	}
	return &status, nil
}

// DeleteWebhook removes webhook integration if you decide
// to switch back to getUpdates. Returns True on success.
func (api *TelegramAPI) DeleteWebhook(ctx context.Context) error {
	return api.newRequest(ctx, apiRequest{
		apiMethod:  "deleteWebhook",
		httpMethod: http.MethodGet,
	})
}

// SendMessage sends text messages.
// On success, the sent Message (https://core.telegram.org/bots/api#message) is returned.
func (api *TelegramAPI) SendMessage(ctx context.Context, req apiModels.SendMessageRequest) (*models.Message, error) {
	var message models.Message
	if err := api.newRequest(ctx, apiRequest{
		apiMethod:  "sendMessage",
		httpMethod: http.MethodPost,
		response:   &message,
		request:    req,
	}); err != nil {
		return nil, err
	}
	return &message, nil
}
