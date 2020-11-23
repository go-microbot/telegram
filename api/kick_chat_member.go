package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// KickChatMember represents method to kick a user from a group, a supergroup or a channel.
// In the case of supergroups and channels, the user will not be able to return to the group
// on their own using invite links, etc., unless unbanned
// (https://core.telegram.org/bots/api#unbanchatmember) first.
// The bot must be an administrator in the chat for this to work and must have the
// appropriate admin rights. Returns True on success.
func (api *TelegramAPI) KickChatMember(ctx context.Context, req apiModels.KickChatMemberRequest) error {
	_, err := api.NewRequest("kickChatMember").
		Query(query.AsMap(req)).
		Body(NewJSONBody(req)).
		Method(http.MethodPost).
		Do(ctx)

	return err
}
