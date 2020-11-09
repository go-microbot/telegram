package api

import (
	"context"
	"net/http"

	"github.com/go-microbot/telegram/models"
)

// TelegramAPI represents Telegram Bot API.
type TelegramAPI struct {
	client *http.Client
	token  string
}

// New returns new TelegramAPI instance.
func New(token string) TelegramAPI {
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
		httpMethod: http.MethodPost,
		response:   &me,
	}); err != nil {
		return nil, err
	}
	return &me, nil
}
