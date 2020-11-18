package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// SetChatTitle represents method to change the title of a chat.
// Titles can't be changed for private chats.
// The bot must be an administrator in the chat for this to work
// and must have the appropriate admin rights. Returns True on success.
func (api *TelegramAPI) SetChatTitle(ctx context.Context, req apiModels.SetChatTitleRequest) error {
	_, err := api.NewRequest("setChatTitle").
		Method(http.MethodPost).
		Query(query.AsMap(req)).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
