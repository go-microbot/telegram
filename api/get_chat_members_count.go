package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// GetChatMembersCount represents method to get the number of members in a chat.
// Returns Int on success.
func (api *TelegramAPI) GetChatMembersCount(ctx context.Context, req apiModels.ChatID) (int32, error) {
	resp, err := api.NewRequest("getChatMembersCount").
		Query(query.AsMap(req)).
		Do(ctx)
	if err != nil {
		return 0, err
	}

	var count int32
	if err := resp.Decode(&count); err != nil {
		return 0, err
	}

	return count, nil
}
