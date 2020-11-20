package api

import (
	"context"
	"math/rand"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendVenue struct{}

func (h sendVenue) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	msg, err := localAPI.SendVenue(ctx, apiModels.SendVenueRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: true,
		Address:             "test address",
		Latitude:            rand.Float32() * 80,
		Longitude:           rand.Float32() * 80,
		Title:               "test title",
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return ctx
}
