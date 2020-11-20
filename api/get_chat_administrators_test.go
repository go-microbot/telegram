package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type getChatAdministrators struct{}

func (h getChatAdministrators) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)
	botUserID := ctx.Value(botUserIDCtxKey)
	require.NotNil(t, botUserID)

	admins, err := localAPI.GetChatAdministrators(ctx, apiModels.ChatID{
		ChatID: query.NewParamAny(groupChatID.(int64)),
	})
	require.NoError(t, err)
	require.NotEmpty(t, admins)

	var found bool
	var adminID int32
	for i := range admins {
		if admins[i].User.ID == botUserID {
			found = true
			continue
		}
		adminID = admins[i].User.ID
	}
	require.True(t, found)

	return context.WithValue(ctx, TestDataKey(adminUserIDCtxKey), adminID)
}
