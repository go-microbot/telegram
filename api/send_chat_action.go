package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// SendChatAction represents method when you need to tell the user that something
// is happening on the bot's side. The status is set for 5 seconds or less
// (when a message arrives from your bot, Telegram clients clear its typing status).
// Returns True on success.
//
// Example: The ImageBot needs some time to process a request and upload the image.
// Instead of sending a text message along the lines of "Retrieving image, please wait…",
// the bot may use sendChatAction with action = upload_photo.
// The user will see a “sending photo” status for the bot.
func (api *TelegramAPI) SendChatAction(ctx context.Context, req apiModels.SendChatActionRequest) error {
	_, err := api.NewRequest("sendChatAction").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
