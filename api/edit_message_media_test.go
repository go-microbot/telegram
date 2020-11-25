package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type editMessageMedia struct{}

func (h editMessageMedia) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	messageID := ctx.Value(existingPhotoMessageIDCtxKey)
	require.NotNil(t, messageID)
	photoURL := ctx.Value(photoToUploadURLCtxKey)
	require.NotNil(t, photoURL)

	err := localAPI.EditMessageMedia(ctx, apiModels.EditMessageMediaRequest{
		ChatID:    query.NewParamAny(chatID.(int64)),
		MessageID: query.NewParamInt(int(messageID.(int32))),
		Media: models.InputMediaPhoto{
			Media: photoURL.(string),
			Type:  "photo",
		},
	})
	require.NoError(t, err)

	return ctx
}
