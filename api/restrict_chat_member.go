package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// RestrictChatMember represents method to restrict a user in a supergroup.
// The bot must be an administrator in the supergroup for this to work and must have
// the appropriate admin rights.
// Pass True for all permissions to lift restrictions from a user. Returns True on success.
func (api *TelegramAPI) RestrictChatMember(ctx context.Context, req apiModels.RestrictChatMemberRequest) error {
	_, err := api.NewRequest("restrictChatMember").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
