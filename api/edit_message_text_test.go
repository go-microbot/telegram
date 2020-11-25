package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type editMessageText struct{}

func (h editMessageText) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	messageID := ctx.Value(sentMessageIDCtxKey)
	require.NotNil(t, messageID)

	err := localAPI.EditMessageText(ctx, apiModels.EditMessageTextRequest{
		ChatID:    query.NewParamAny(chatID.(int64)),
		MessageID: query.NewParamAny(messageID.(int32)),
		Text:      "New test message text",
	})
	require.NoError(t, err)

	return ctx
}
