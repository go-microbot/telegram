package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendMediaGroup struct{}

func (h sendMediaGroup) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	photoURL := ctx.Value(photoToUploadURLCtxKey)
	require.NotNil(t, photoURL)
	audioURL := ctx.Value(audioToUploadURLCtxKey)
	require.NotNil(t, audioURL)
	documentURL := ctx.Value(documentToUploadURLCtxKey)
	require.NotNil(t, documentURL)
	videoURL := ctx.Value(videoToUploadURLCtxKey)
	require.NotNil(t, videoURL)

	allMessages := make([]models.Message, 0, 4)

	// audio.
	messages, err := localAPI.SendMediaGroup(ctx, apiModels.SendMediaGroupRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: true,
		Media: []interface{}{
			models.InputMediaAudio{
				Type:  "audio",
				Media: audioURL.(string),
			},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, messages)
	allMessages = append(allMessages, messages...)

	// photo.
	messages, err = localAPI.SendMediaGroup(ctx, apiModels.SendMediaGroupRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: true,
		Media: []interface{}{
			models.InputMediaPhoto{
				Type:  "photo",
				Media: photoURL.(string),
			},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, messages)
	allMessages = append(allMessages, messages...)

	// document.
	messages, err = localAPI.SendMediaGroup(ctx, apiModels.SendMediaGroupRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: true,
		Media: []interface{}{
			models.InputMediaDocument{
				Type:  "document",
				Media: documentURL.(string),
			},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, messages)
	allMessages = append(allMessages, messages...)

	// video.
	messages, err = localAPI.SendMediaGroup(ctx, apiModels.SendMediaGroupRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: true,
		Media: []interface{}{
			models.InputMediaVideo{
				Type:  "video",
				Media: videoURL.(string),
			},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, messages)
	allMessages = append(allMessages, messages...)

	require.Equal(t, 4, len(allMessages))

	return ctx
}
