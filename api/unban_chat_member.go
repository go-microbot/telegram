package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// UnbanChatMember represents method to unban a previously kicked user in a supergroup or channel.
// The user will not return to the group or channel automatically, but will be able
// to join via link, etc. The bot must be an administrator for this to work.
// By default, this method guarantees that after the call the user is not a member of the chat,
// but will be able to join it. So if the user is a member of the chat they will also be removed
// from the chat. If you don't want this, use the parameter only_if_banned.
// Returns True on success.
func (api *TelegramAPI) UnbanChatMember(ctx context.Context, req apiModels.UnbanChatMemberRequest) error {
	_, err := api.NewRequest("unbanChatMember").
		Query(query.AsMap(req)).
		Body(NewJSONBody(req)).
		Method(http.MethodPost).
		Do(ctx)

	return err
}
