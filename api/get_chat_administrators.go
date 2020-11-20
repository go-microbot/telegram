package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// GetChatAdministrators represents method to get a list of administrators in a chat.
// On success, returns an Array of ChatMember objects that contains information about all
// chat administrators except other bots. If the chat is a group or a supergroup
// and no administrators were appointed, only the creator will be returned.
func (api *TelegramAPI) GetChatAdministrators(ctx context.Context, req apiModels.ChatID) ([]models.ChatMember, error) {
	resp, err := api.NewRequest("getChatAdministrators").
		Query(query.AsMap(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var admins []models.ChatMember
	if err := resp.Decode(&admins); err != nil {
		return nil, err
	}

	return admins, nil
}
