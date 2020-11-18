package api

import (
	"context"

	"github.com/go-microbot/telegram/models"
)

// GetMyCommands represents method to get the current list of the bot's commands.
// Requires no parameters. Returns Array of BotCommand on success.
func (api *TelegramAPI) GetMyCommands(ctx context.Context) ([]models.BotCommand, error) {
	resp, err := api.NewRequest("getMyCommands").Do(ctx)
	if err != nil {
		return nil, err
	}

	var commands []models.BotCommand
	if err := resp.Decode(&commands); err != nil {
		return nil, err
	}

	return commands, nil
}
