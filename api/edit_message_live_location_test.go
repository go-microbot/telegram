package api

import (
	"context"
	"math/rand"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type editMessageLiveLocation struct{}

func (h editMessageLiveLocation) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)
	messageID := ctx.Value(sentMessageIDCtxKey)
	require.NotNil(t, messageID)

	err := localAPI.EditMessageLiveLocation(ctx, apiModels.EditMessageLiveLocationRequest{
		ChatID:    query.NewParamAny(groupChatID.(int64)),
		MessageID: query.NewParamInt(int(messageID.(int32))),
		Latitude:  rand.Float32() * 80,
		Longitude: rand.Float32() * 80,
	})
	if err != nil {
		require.Contains(t, err.Error(), "message to edit not found")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
