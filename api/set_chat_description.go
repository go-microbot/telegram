package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// SetChatDescription represents method to change the description of a group,
// a supergroup or a channel. The bot must be an administrator in the chat
// for this to work and must have the appropriate admin rights. Returns True on success.
func (api *TelegramAPI) SetChatDescription(ctx context.Context, req apiModels.SetChatDescriptionRequest) error {
	_, err := api.NewRequest("setChatDescription").
		Query(query.AsMap(req)).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
