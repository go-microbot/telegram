package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendPoll struct{}

func (h sendPoll) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	question := "Do you like Telegram?"
	msg, err := localAPI.SendPoll(ctx, apiModels.SendPollRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: true,
		Question:            question,
		Options: []string{
			"Yes",
			"No",
		},
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	ctx = context.WithValue(ctx, TestDataKey(pollMessageIDCtxKey), msg.ID)
	return context.WithValue(ctx, TestDataKey(pollQuestionCtxKey), question)
}
