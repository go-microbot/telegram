package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// UnpinAllChatMessages represents method to clear the list of pinned messages in a chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work
// and must have the 'can_pin_messages' admin right in a supergroup or 'can_edit_messages' admin
// right in a channel. Returns True on success.
func (api *TelegramAPI) UnpinAllChatMessages(ctx context.Context, req apiModels.ChatID) error {
	_, err := api.NewRequest("unpinAllChatMessages").
		Query(query.AsMap(req)).
		Do(ctx)

	return err
}
