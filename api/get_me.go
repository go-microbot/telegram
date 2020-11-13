package api

import (
	"context"
	"net/http"

	"github.com/go-microbot/telegram/models"
)

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
