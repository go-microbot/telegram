package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// PromoteChatMember represents method to promote or demote a user in a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have
// the appropriate admin rights.
// Pass False for all boolean parameters to demote a user. Returns True on success.
func (api *TelegramAPI) PromoteChatMember(ctx context.Context, req apiModels.PromoteChatMemberRequest) error {
	_, err := api.NewRequest("promoteChatMember").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
