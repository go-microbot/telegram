package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/stretchr/testify/require"
)

type answerInlineQuery struct{}

func (h answerInlineQuery) Test(ctx context.Context, t *testing.T) context.Context {
	err := localAPI.AnswerInlineQuery(ctx, apiModels.AnswerInlineQueryRequest{
		InlineQueryID:     "query",
		SwitchPmParameter: "test",
	})
	if err != nil {
		require.Contains(t, err.Error(), "query ID is invalid")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
