package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type exportChatInviteLink struct{}

func (h exportChatInviteLink) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)

	link, err := localAPI.ExportChatInviteLink(ctx, apiModels.ChatID{
		ChatID: query.NewParamAny(groupChatID.(int64)),
	})
	require.NoError(t, err)
	require.NotEmpty(t, link)

	return ctx
}
