package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// StopMessageLiveLocation represents method to stop updating a live location message before
// live_period expires. On success, if the message was sent by the bot,
// the sent Message is returned, otherwise True is returned.
func (api *TelegramAPI) StopMessageLiveLocation(ctx context.Context, req apiModels.StopMessageLiveLocationRequest) error {
	_, err := api.NewRequest("stopMessageLiveLocation").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
