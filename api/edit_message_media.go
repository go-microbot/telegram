package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// EditMessageMedia represents method to edit animation, audio, document, photo, or video messages.
// If a message is part of a message album, then it can be edited only to an audio for
// audio albums, only to a document for document albums and to a photo or a video otherwise.
// When an inline message is edited, a new file can't be uploaded.
// Use a previously uploaded file via its file_id or specify a URL.
// On success, if the edited message was sent by the bot,
// the edited Message is returned, otherwise True is returned.
func (api *TelegramAPI) EditMessageMedia(ctx context.Context, req apiModels.EditMessageMediaRequest) error {
	_, err := api.NewRequest("editMessageMedia").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
