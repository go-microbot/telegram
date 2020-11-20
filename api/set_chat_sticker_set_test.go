package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type setChatStickerSet struct{}

func (h setChatStickerSet) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)
	stickerSetName := ctx.Value(stickerSetNameCtxKey)
	require.NotNil(t, stickerSetName)

	err := localAPI.SetChatStickerSet(ctx, apiModels.SetChatStickerSetRequest{
		ChatID:         query.NewParamAny(groupChatID.(int64)),
		StickerSetName: stickerSetName.(string),
	})
	if err != nil {
		require.Contains(t, err.Error(), "Bad Request")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
