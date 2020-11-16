package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// LeaveChat represents method for your bot to leave a group,
// supergroup or channel. Returns True on success.
func (api *TelegramAPI) LeaveChat(ctx context.Context, req apiModels.ChatID) error {
	_, err := api.NewRequest("leaveChat").
		Query(query.AsMap(req)).
		Do(ctx)

	return err
}
