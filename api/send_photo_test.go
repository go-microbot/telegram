package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendPhoto struct{}

func (h sendPhoto) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	/*existingPhotoID := ctx.Value(existingPhotoIDCtxKey)
	require.NotNil(t, existingPhotoID)
	existingPhotoUID := ctx.Value(existingPhotoUIDCtxKey)
	require.NotNil(t, existingPhotoUID)*/
	photoToUploadURL := ctx.Value(photoToUploadURLCtxKey)
	require.NotNil(t, photoToUploadURL)

	// send an existing photo.
	/*msg, err := localAPI.SendPhoto(ctx, apiModels.SendPhotoRequest{
		ChatID: query.NewParamAny(chatID),
		Photo:  existingPhotoID.(string),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, chatID, msg.Chat.ID)
	require.NotNil(t, msg.Photo)
	require.Equal(t, existingPhotoUID, msg.Photo[0].FileUID)

	// send external image URL.
	msg, err = localAPI.SendPhoto(ctx, apiModels.SendPhotoRequest{
		ChatID: query.NewParamAny(chatID),
		Photo:  photoToUploadURL.(string),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, chatID, msg.Chat.ID)
	require.NotNil(t, msg.Photo)*/

	// send new photo.
	msg, err := localAPI.SendPhoto(ctx, apiModels.SendPhotoRequest{
		ChatID: query.NewParamAny(chatID),
		Photo:  form.NewPartFile("./test_data/test_photo.png"),
		//Caption: form.NewPartText("text caption"),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, chatID, msg.Chat.ID)
	require.NotNil(t, msg.Photo)

	return ctx
}
