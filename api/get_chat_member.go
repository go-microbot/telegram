package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// GetChatMember represents method to get information about a member of a chat.
// Returns a ChatMember object on success.
func (api *TelegramAPI) GetChatMember(ctx context.Context, req apiModels.GetChatMemberRequest) (*models.ChatMember, error) {
	resp, err := api.NewRequest("getChatMember").
		Query(query.AsMap(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var member models.ChatMember
	if err := resp.Decode(&member); err != nil {
		return nil, err
	}

	return &member, nil
}
