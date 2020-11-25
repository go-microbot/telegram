package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
)

// AnswerCallbackQuery represents method to send answers to callback queries
// sent from inline keyboards. The answer will be displayed to the user as a notification
// at the top of the chat screen or as an alert. On success, True is returned.
func (api *TelegramAPI) AnswerCallbackQuery(ctx context.Context, req apiModels.AnswerCallbackQueryRequest) error {
	_, err := api.NewRequest("answerCallbackQuery").
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)

	return err
}
