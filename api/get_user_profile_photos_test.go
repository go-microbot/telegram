package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type getUserProfilePhotos struct{}

func (h getUserProfilePhotos) Test(ctx context.Context, t *testing.T) context.Context {
	botUserID := ctx.Value(botUserIDCtxKey)
	require.NotNil(t, botUserID)

	photos, err := localAPI.GetUserProfilePhotos(ctx, apiModels.GetUserProfilePhotosRequest{
		UserID: query.NewParamInt(int(botUserID.(int32))),
	})
	require.NoError(t, err)
	require.NotNil(t, photos)

	return ctx
}
