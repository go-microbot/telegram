package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendContact struct{}

func (h sendContact) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	msg, err := localAPI.SendContact(ctx, apiModels.SendContactRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: true,
		FirstName:           "Test",
		LastName:            "User",
		PhoneNumber:         "+375259400000",
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return ctx
}
