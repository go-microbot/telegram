package api

import (
	"context"
	"net/http"

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
func (api *TelegramAPI) GetUpdates(ctx context.Context, req GetUpdatesRequest) ([]models.Update, error) {
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
