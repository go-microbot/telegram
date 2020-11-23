package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type getGameHighScores struct{}

func (h getGameHighScores) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	gameScore := ctx.Value(gameScoreCtxKey)
	require.NotNil(t, gameScore)
	adminUserID := ctx.Value(adminUserIDCtxKey)
	require.NotNil(t, adminUserID)
	messageID := ctx.Value(gameMessageIDCtxKey)
	require.NotNil(t, messageID)

	scores, err := localAPI.GetGameHighScores(ctx, apiModels.GetGameHighScoresRequest{
		ChatID:    query.NewParamAny(chatID.(int64)),
		UserID:    adminUserID.(int32),
		MessageID: query.NewParamInt(int(messageID.(int32))),
	})
	require.NoError(t, err)
	require.NotEmpty(t, scores)
	var found bool
	for i := range scores {
		if scores[i].Score == gameScore {
			require.Equal(t, adminUserID, scores[i].User.ID)
			found = true
			break
		}
	}
	require.True(t, found)

	return ctx
}
