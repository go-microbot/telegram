package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type deleteChatPhoto struct{}

func (h deleteChatPhoto) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)

	err := localAPI.DeleteChatPhoto(ctx, apiModels.ChatID{
		ChatID: query.NewParamAny(groupChatID.(int64)),
	})
	require.NoError(t, err)

	return ctx
}
