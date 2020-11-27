package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
)

// AnswerPreCheckoutQuery represents method to respond to such pre-checkout queries.
// Once the user has confirmed their payment and shipping details,
// the Bot API sends the final confirmation in the form of an Update
// with the field pre_checkout_query. On success, True is returned.
// Note: The Bot API must receive an answer within 10 seconds
// after the pre-checkout query was sent.
func (api *TelegramAPI) AnswerPreCheckoutQuery(ctx context.Context, req apiModels.AnswerPreCheckoutQueryRequest) error {
	_, err := api.NewRequest("answerPreCheckoutQuery").
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
