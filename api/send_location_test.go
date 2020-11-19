package api

import (
	"context"
	"math/rand"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendLocation struct{}

func (h sendLocation) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	msg, err := localAPI.SendLocation(ctx, apiModels.SendLocationRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: true,
		HorizontalAccuracy:  rand.Float32() * 100,
		Latitude:            rand.Float32() * 80,
		Longitude:           rand.Float32() * 80,
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return ctx
}
