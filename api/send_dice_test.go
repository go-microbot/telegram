package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendDice struct{}

func (h sendDice) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	msg, err := localAPI.SendDice(ctx, apiModels.SendDiceRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: true,
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return ctx
}
