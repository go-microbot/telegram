package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// ExportChatInviteLink represents method to generate a new invite link for a chat;
// any previously generated link is revoked. The bot must be an administrator in the chat
// for this to work and must have the appropriate admin rights.
// Returns the new invite link as String on success.
func (api *TelegramAPI) ExportChatInviteLink(ctx context.Context, req apiModels.ChatID) (string, error) {
	resp, err := api.NewRequest("exportChatInviteLink").
		Query(query.AsMap(req)).
		Do(ctx)
	if err != nil {
		return "", err
	}

	var link string
	if err := resp.Decode(&link); err != nil {
		return "", err
	}

	return link, nil
}
