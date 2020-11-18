package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
)

// SetMyCommands represents method to change the list of the bot's commands.
// Returns True on success.
func (api *TelegramAPI) SetMyCommands(ctx context.Context, req apiModels.SetMyCommandsRequest) error {
	_, err := api.NewRequest("setMyCommands").
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
