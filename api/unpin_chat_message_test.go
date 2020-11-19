package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type unpinChatMessage struct{}

func (h unpinChatMessage) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	messageID := ctx.Value(sentMessageIDCtxKey)
	require.NotNil(t, messageID)

	msgID := messageID.(int32)
	err := localAPI.UnpinChatMessage(ctx, apiModels.UnpinChatMessageRequest{
		ChatID:    query.NewParamAny(chatID.(int64)),
		MessageID: &msgID,
	})
	require.NoError(t, err)

	return ctx
}
