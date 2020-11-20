package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendAnimation struct{}

func (h sendAnimation) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	msg, err := localAPI.SendAnimation(ctx, apiModels.SendAnimationRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: form.NewPartAny(true),
		Animation:           form.NewPartFile("./test_data/test_animation.gif"),
		Caption:             form.NewPartText("test animation"),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return ctx
}
