package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// EditMessageReplyMarkup represents method to edit only the reply markup of messages.
// On success, if the edited message is not an inline message,
// the edited Message is returned, otherwise True is returned.
func (api *TelegramAPI) EditMessageReplyMarkup(ctx context.Context, req apiModels.EditMessageReplyMarkupRequest) error {
	_, err := api.NewRequest("editMessageReplyMarkup").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
