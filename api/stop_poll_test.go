package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type stopPoll struct{}

func (h stopPoll) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	pollQuestion := ctx.Value(pollQuestionCtxKey)
	require.NotNil(t, pollQuestion)
	pollMessageID := ctx.Value(pollMessageIDCtxKey)
	require.NotNil(t, pollMessageID)

	poll, err := localAPI.StopPoll(ctx, apiModels.StopPollRequest{
		ChatID:    query.NewParamAny(chatID.(int64)),
		MessageID: query.NewParamInt(int(pollMessageID.(int32))),
	})
	require.NoError(t, err)
	require.NotNil(t, poll)
	require.Equal(t, pollQuestion, poll.Question)

	return ctx
}
