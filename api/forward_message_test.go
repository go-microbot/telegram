package api

import (
	"context"
	"testing"

	"github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type forwardMessage struct{}

func (h forwardMessage) Test(ctx context.Context, t *testing.T) context.Context {
	sentMessageID := ctx.Value(sentMessageIDCtxKey)
	require.NotNil(t, sentMessageID)
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)

	msg, err := localAPI.ForwardMessage(ctx, models.ForwardMessageRequest{
		ChatID:              query.NewParamAny(groupChatID.(int64)),
		DisableNotification: true,
		MessageID:           sentMessageID.(int32),
		FromChatID:          query.NewParamAny(chatID.(int64)),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, msg.Chat.ID, groupChatID)

	return ctx
}
