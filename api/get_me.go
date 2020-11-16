package api

import (
	"context"

	"github.com/go-microbot/telegram/models"
)

// GetMe returns basic information about the bot.
func (api *TelegramAPI) GetMe(ctx context.Context) (*models.User, error) {
	resp, err := api.NewRequest("getMe").Do(ctx)
	if err != nil {
		return nil, err
	}

	var me models.User
	if err := resp.Decode(&me); err != nil {
		return nil, err
	}

	return &me, nil
}
