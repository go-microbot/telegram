package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// DeleteChatPhoto represents method to delete a chat photo.
// Photos can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate
// admin rights. Returns True on success.
func (api *TelegramAPI) DeleteChatPhoto(ctx context.Context, req apiModels.ChatID) error {
	_, err := api.NewRequest("deleteChatPhoto").
		Query(query.AsMap(req)).
		Do(ctx)

	return err
}
