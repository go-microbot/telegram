package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// SetChatPhoto represents method to set a new profile photo for the chat.
// Photos can't be changed for private chats.
// The bot must be an administrator in the chat for this to work
// and must have the appropriate admin rights. Returns True on success.
func (api *TelegramAPI) SetChatPhoto(ctx context.Context, req apiModels.SetChatPhotoRequest) error {
	_, err := api.NewRequest("setChatPhoto").
		Method(http.MethodPost).
		Query(query.AsMap(req)).
		Body(NewFormDataBody(req)).
		Do(ctx)

	return err
}
