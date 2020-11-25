package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// EditMessageText represents method to edit text and game messages.
// On success, if the edited message is not an inline message,
// the edited Message is returned, otherwise True is returned.
func (api *TelegramAPI) EditMessageText(ctx context.Context, req apiModels.EditMessageTextRequest) error {
	_, err := api.NewRequest("editMessageText").
		Query(query.AsMap(req)).
		Body(NewJSONBody(req)).
		Method(http.MethodPost).
		Do(ctx)

	return err
}
