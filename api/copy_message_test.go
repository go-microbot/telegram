package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type copyMessage struct{}

func (h copyMessage) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	msgID := ctx.Value(sentMessageIDCtxKey)
	require.NotNil(t, msgID)
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)

	msgID, err := localAPI.CopyMessage(ctx, apiModels.CopyMessageRequest{
		ChatID:     query.NewParamAny(groupChatID.(int64)),
		FromChatID: query.NewParamAny(chatID.(int64)),
		MessageID:  query.NewParamInt(int(msgID.(int32))),
	})
	require.NoError(t, err)
	require.NotNil(t, msgID)

	return ctx
}
