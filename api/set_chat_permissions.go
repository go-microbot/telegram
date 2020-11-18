package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// SetChatPermissions represents method to set default chat permissions for all members.
// The bot must be an administrator in the group or a supergroup for this to work
// and must have the can_restrict_members admin rights. Returns True on success.
func (api *TelegramAPI) SetChatPermissions(ctx context.Context, req apiModels.SetChatPermissionsRequest) error {
	_, err := api.NewRequest("setChatPermissions").
		Query(query.AsMap(req)).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
