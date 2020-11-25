package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/stretchr/testify/require"
)

type answerCallbackQuery struct{}

func (h answerCallbackQuery) Test(ctx context.Context, t *testing.T) context.Context {
	err := localAPI.AnswerCallbackQuery(ctx, apiModels.AnswerCallbackQueryRequest{
		CallbackQueryID: "query",
		Text:            "test callback text",
	})
	if err != nil {
		require.Contains(t, err.Error(), "query ID is invalid")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
