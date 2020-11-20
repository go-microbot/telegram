package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type deleteMessage struct{}

func (h deleteMessage) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	msgID := ctx.Value(sentMessageIDCtxKey)
	require.NotNil(t, msgID)

	err := localAPI.DeleteMessage(ctx, apiModels.DeleteMessageRequest{
		ChatID:    query.NewParamAny(chatID.(int64)),
		MessageID: query.NewParamInt(int(msgID.(int32))),
	})
	require.NoError(t, err)

	return ctx
}
