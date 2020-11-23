package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// EditMessageLiveLocation represents method to edit live location messages.
// A location can be edited until its live_period expires or editing is explicitly
// disabled by a call to stopMessageLiveLocation.
// On success, if the edited message is not an inline message,
// the edited Message is returned, otherwise True is returned.
func (api *TelegramAPI) EditMessageLiveLocation(ctx context.Context, req apiModels.EditMessageLiveLocationRequest) error {
	_, err := api.NewRequest("editMessageLiveLocation").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
