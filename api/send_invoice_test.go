package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendInvoice struct{}

func (h sendInvoice) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	msg, err := localAPI.SendInvoice(ctx, apiModels.SendInvoiceRequest{
		ChatID:         query.NewParamAny(chatID.(int64)),
		Title:          "Test invoice title",
		Description:    "Test invoice description",
		Payload:        "payload",
		ProviderToken:  "token",
		StartParameter: "start",
		Currency:       "BRL",
		Prices: []models.LabeledPrice{
			{
				Label:  "USD",
				Amount: 145,
			},
		},
		DisableNotification: true,
	})
	if err != nil {
		require.Contains(t, err.Error(), "PAYMENT_PROVIDER_INVALID")
		return ctx
	}
	require.NoError(t, err)
	require.NotNil(t, msg)

	return ctx
}
