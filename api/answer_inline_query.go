package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
)

// AnswerInlineQuery represents method to send answers to an inline query.
// On success, True is returned.
// No more than 50 results per query are allowed.
func (api *TelegramAPI) AnswerInlineQuery(ctx context.Context, req apiModels.AnswerInlineQueryRequest) error {
	_, err := api.NewRequest("answerInlineQuery").
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
