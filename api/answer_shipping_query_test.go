package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/stretchr/testify/require"
)

type answerShippingQuery struct{}

func (h answerShippingQuery) Test(ctx context.Context, t *testing.T) context.Context {
	err := localAPI.AnswerShippingQuery(ctx, apiModels.AnswerShippingQueryRequest{
		ShippingQueryID: "queryID",
		ShippingOptions: []models.ShippingOption{
			{
				ID: "1",
				Prices: []models.LabeledPrice{
					{
						Label:  "USD",
						Amount: 145,
					},
				},
				Title: "Test option title",
			},
		},
		OK: true,
	})
	if err != nil {
		require.Contains(t, err.Error(), "query ID is invalid")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
