package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// PinChatMessage represents method to add a message to the list of pinned messages in a chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this
// to work and must have the 'can_pin_messages' admin right in a supergroup
// or 'can_edit_messages' admin right in a channel. Returns True on success.
func (api *TelegramAPI) PinChatMessage(ctx context.Context, req apiModels.PinChatMessageRequest) error {
	_, err := api.NewRequest("pinChatMessage").
		Query(query.AsMap(req)).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
