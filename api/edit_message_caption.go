package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// EditMessageCaption represents method to edit captions of messages.
// On success, if the edited message is not an inline message,
// the edited Message is returned, otherwise True is returned.
func (api *TelegramAPI) EditMessageCaption(ctx context.Context, req apiModels.EditMessageCaptionRequest) error {
	_, err := api.NewRequest("editMessageCaption").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
