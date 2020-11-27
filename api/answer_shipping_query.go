package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
)

// AnswerShippingQuery represents method to reply to shipping queries.
// If you sent an invoice requesting a shipping address and the parameter
// is_flexible was specified, the Bot API will send an Update
// with a shipping_query field to the bot. On success, True is returned.
func (api *TelegramAPI) AnswerShippingQuery(ctx context.Context, req apiModels.AnswerShippingQueryRequest) error {
	_, err := api.NewRequest("answerShippingQuery").
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
