package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// GetChat represents method to get up to date information about the chat
// (current name of the user for one-on-one conversations,
// current username of a user, group or channel, etc.). Returns a Chat object on success.
func (api *TelegramAPI) GetChat(ctx context.Context, req apiModels.ChatID) (*models.Chat, error) {
	resp, err := api.NewRequest("getChat").
		Query(query.AsMap(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var chat models.Chat
	if err := resp.Decode(&chat); err != nil {
		return nil, err
	}

	return &chat, nil
}
