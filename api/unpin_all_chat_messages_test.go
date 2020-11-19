package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type unpinAllChatMessages struct{}

func (h unpinAllChatMessages) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	err := localAPI.UnpinAllChatMessages(ctx, apiModels.ChatID{
		ChatID: query.NewParamAny(chatID.(int64)),
	})
	require.NoError(t, err)

	return ctx
}
