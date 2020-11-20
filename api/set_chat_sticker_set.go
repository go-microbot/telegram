package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// SetChatStickerSet represents method to set a new group sticker set for a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate
// admin rights. Use the field can_set_sticker_set optionally returned in getChat requests
// to check if the bot can use this method. Returns True on success.
func (api *TelegramAPI) SetChatStickerSet(ctx context.Context, req apiModels.SetChatStickerSetRequest) error {
	_, err := api.NewRequest("setChatStickerSet").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
