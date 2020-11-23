package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/stretchr/testify/require"
)

type uploadStickerFile struct{}

func (h uploadStickerFile) Test(ctx context.Context, t *testing.T) context.Context {
	adminUserID := ctx.Value(adminUserIDCtxKey)
	require.NotNil(t, adminUserID)

	file, err := localAPI.UploadStickerFile(ctx, apiModels.UploadStickerFileRequest{
		UserID:     form.NewPartAny(adminUserID.(int32)),
		PngSticker: form.NewPartFile("./test_data/stickers/sticker1.png"),
	})
	require.NoError(t, err)
	require.NotNil(t, file)

	return context.WithValue(ctx, TestDataKey(uploadedStickerFileIDCtxKey), file.FileID)
}
