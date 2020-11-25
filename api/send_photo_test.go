package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendPhoto struct{}

func (h sendPhoto) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	photoToUploadURL := ctx.Value(photoToUploadURLCtxKey)
	require.NotNil(t, photoToUploadURL)

	// send new photo.
	msg, err := localAPI.SendPhoto(ctx, apiModels.SendPhotoRequest{
		ChatID:              query.NewParamAny(chatID),
		Photo:               form.NewPartFile("./test_data/test_photo.png"),
		Caption:             form.NewPartText("new uploaded photo"),
		DisableNotification: form.NewPartAny(true),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, chatID, msg.Chat.ID)
	require.NotEmpty(t, msg.Photo)
	existingFileID := msg.Photo[0].FileID
	existingFileUID := msg.Photo[0].FileUID

	// send external image URL.
	msg, err = localAPI.SendPhoto(ctx, apiModels.SendPhotoRequest{
		ChatID: query.NewParamAny(chatID),
		Photo:  form.NewPartText(photoToUploadURL.(string)),
		CaptionEntities: form.NewPartJSON([]models.MessageEntity{
			{
				Type:   models.MessageEntityTypeTextLink,
				URL:    "https://www.google.com",
				Offset: 11,
				Length: 4,
			},
		}),
		Caption:             form.NewPartText("photo with link!"),
		DisableNotification: form.NewPartAny(true),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, chatID, msg.Chat.ID)
	require.NotEmpty(t, msg.Photo)

	// send an existing photo.
	msg, err = localAPI.SendPhoto(ctx, apiModels.SendPhotoRequest{
		ChatID:              query.NewParamAny(chatID),
		Photo:               form.NewPartText(existingFileID),
		ParseMode:           form.NewPartText("MarkdownV2"),
		Caption:             form.NewPartText("__Hello__"),
		DisableNotification: form.NewPartAny(true),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, chatID, msg.Chat.ID)
	require.NotEmpty(t, msg.Photo)
	require.Equal(t, existingFileUID, msg.Photo[0].FileUID)

	ctx = context.WithValue(ctx, TestDataKey(existingPhotoIDCtxKey), existingFileID)
	ctx = context.WithValue(ctx, TestDataKey(existingPhotoUIDCtxKey), existingFileUID)
	ctx = context.WithValue(ctx, TestDataKey(existingPhotoMessageIDCtxKey), msg.ID)

	return ctx
}
