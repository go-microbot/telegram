package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type getFile struct{}

func (h getFile) Test(ctx context.Context, t *testing.T) context.Context {
	fileID := ctx.Value(existingPhotoIDCtxKey)
	require.NotNil(t, fileID)
	fileUID := ctx.Value(existingPhotoUIDCtxKey)
	require.NotNil(t, fileUID)

	file, err := localAPI.GetFile(ctx, apiModels.FileID{
		FileID: query.NewParamString(fileID.(string)),
	})
	require.NoError(t, err)
	require.NotNil(t, file)
	require.Equal(t, fileUID, fileUID)

	return ctx
}
