package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type pinChatMessage struct{}

func (h pinChatMessage) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, chatID)
	messageID := ctx.Value(sentMessageIDCtxKey)
	require.NotNil(t, messageID)

	err := localAPI.PinChatMessage(ctx, apiModels.PinChatMessageRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: true,
		MessageID:           messageID.(int32),
	})
	require.NotNil(t, err)

	return ctx
}
