package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// DeleteChatStickerSet represents method to delete a group sticker set from a supergroup.
// The bot must be an administrator in the chat for this to work
// and must have the appropriate admin rights. Use the field can_set_sticker_set optionally
// returned in getChat requests to check if the bot can use this method. Returns True on success.
func (api *TelegramAPI) DeleteChatStickerSet(ctx context.Context, req apiModels.ChatID) error {
	_, err := api.NewRequest("deleteChatStickerSet").
		Query(query.AsMap(req)).
		Do(ctx)

	return err
}
