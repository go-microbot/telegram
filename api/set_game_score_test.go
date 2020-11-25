package api

import (
	"context"
	"math/rand"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type setGameScore struct{}

func (h setGameScore) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	gameShortName := ctx.Value(gameShortNameCtxKey)
	require.NotNil(t, gameShortName)
	adminUserID := ctx.Value(adminUserIDCtxKey)
	require.NotNil(t, adminUserID)
	messageID := ctx.Value(gameMessageIDCtxKey)
	require.NotNil(t, messageID)

	score := int32(rand.Intn(150))
	msg, err := localAPI.SetGameScore(ctx, apiModels.SetGameScoreRequest{
		ChatID:    query.NewParamAny(chatID.(int64)),
		Score:     score,
		UserID:    adminUserID.(int32),
		MessageID: query.NewParamInt(int(messageID.(int32))),
		Force:     true,
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return context.WithValue(ctx, TestDataKey(gameScoreCtxKey), score)
}
