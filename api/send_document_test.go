package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendDocument struct{}

func (h sendDocument) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	msg, err := localAPI.SendDocument(ctx, apiModels.SendDocumentRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: form.NewPartAny(true),
		Document:            form.NewPartFile("./test_data/test_document.pdf"),
		Caption:             form.NewPartText("test document"),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return ctx
}
