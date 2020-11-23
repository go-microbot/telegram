package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendGame struct{}

func (h sendGame) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	gameShortName := ctx.Value(gameShortNameCtxKey)
	require.NotNil(t, gameShortName)

	msg, err := localAPI.SendGame(ctx, apiModels.SendGameRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: true,
		GameShortName:       gameShortName.(string),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return context.WithValue(ctx, TestDataKey(gameMessageIDCtxKey), msg.ID)
}
