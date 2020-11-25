package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/stretchr/testify/require"
)

type deleteStickerFromSet struct{}

func (h deleteStickerFromSet) Test(ctx context.Context, t *testing.T) context.Context {
	stickerFileID := ctx.Value(uploadedStickerFileIDCtxKey)
	require.NotNil(t, stickerFileID)

	err := localAPI.DeleteStickerFromSet(ctx, apiModels.DeleteStickerFromSetRequest{
		Sticker: stickerFileID.(string),
	})
	if err != nil {
		require.Contains(t, err.Error(), "STICKER_INVALID")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
