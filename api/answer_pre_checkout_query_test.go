package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/stretchr/testify/require"
)

type answerPreCheckoutQuery struct{}

func (h answerPreCheckoutQuery) Test(ctx context.Context, t *testing.T) context.Context {
	err := localAPI.AnswerPreCheckoutQuery(ctx, apiModels.AnswerPreCheckoutQueryRequest{
		PreCheckoutQueryID: "queryID",
		OK:                 true,
	})
	if err != nil {
		require.Contains(t, err.Error(), "query ID is invalid")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
