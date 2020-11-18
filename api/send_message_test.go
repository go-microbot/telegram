package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendMessage struct{}

func (h sendMessage) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	req := apiModels.SendMessageRequest{
		ChatID: query.NewParamAny(chatID.(int64)),
		Text:   "This is a test message!",
	}
	msg, err := localAPI.SendMessage(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, msg.Text, req.Text)
	require.Equal(t, chatID, msg.Chat.ID)

	return ctx
}
