package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type getStickerSet struct{}

func (h getStickerSet) Test(ctx context.Context, t *testing.T) context.Context {
	name := ctx.Value(stickerSetNameCtxKey)
	require.NotNil(t, name)

	set, err := localAPI.GetStickerSet(ctx, apiModels.GetStickerSetRequest{
		Name: query.NewParamString(name.(string)),
	})
	require.NoError(t, err)
	require.NotNil(t, set)
	require.Equal(t, name, set.Name)

	return ctx
}
