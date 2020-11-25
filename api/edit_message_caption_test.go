package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type editMessageCaption struct{}

func (h editMessageCaption) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	messageID := ctx.Value(existingPhotoMessageIDCtxKey)
	require.NotNil(t, messageID)

	err := localAPI.EditMessageCaption(ctx, apiModels.EditMessageCaptionRequest{
		ChatID:    query.NewParamAny(chatID.(int64)),
		MessageID: query.NewParamInt(int(messageID.(int32))),
		Caption:   "New test caption",
	})
	require.NoError(t, err)

	return ctx
}
