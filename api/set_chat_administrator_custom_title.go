package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// SetChatAdministratorCustomTitle represents method to set a custom title for an administrator
// in a supergroup promoted by the bot. Returns True on success.
func (api *TelegramAPI) SetChatAdministratorCustomTitle(ctx context.Context, req apiModels.SetChatAdminCustomTitleRequest) error {
	_, err := api.NewRequest("setChatAdministratorCustomTitle").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
