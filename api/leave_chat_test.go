package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type leaveChat struct{}

func (h leaveChat) Test(ctx context.Context, t *testing.T) context.Context {
	err := localAPI.LeaveChat(ctx, apiModels.ChatID{
		ChatID: query.NewParamAny(7),
	})
	require.Error(t, err)

	return ctx
}
