package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type deleteChatStickerSet struct{}

func (h deleteChatStickerSet) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)

	err := localAPI.DeleteChatStickerSet(ctx, apiModels.ChatID{
		ChatID: query.NewParamAny(groupChatID.(int64)),
	})
	if err != nil {
		require.Contains(t, err.Error(), "Bad Request")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
