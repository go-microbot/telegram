package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// UnpinChatMessage represents method to remove a message
// from the list of pinned messages in a chat. If the chat is not a private chat,
// the bot must be an administrator in the chat for this to work and must have
// the 'can_pin_messages' admin right in a supergroup or 'can_edit_messages'
// admin right in a channel. Returns True on success.
func (api *TelegramAPI) UnpinChatMessage(ctx context.Context, req apiModels.UnpinChatMessageRequest) error {
	_, err := api.NewRequest("unpinChatMessage").
		Query(query.AsMap(req)).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
