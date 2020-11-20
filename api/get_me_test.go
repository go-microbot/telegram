package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type getMe struct{}

func (h getMe) Test(ctx context.Context, t *testing.T) context.Context {
	me, err := localAPI.GetMe(ctx)
	require.NoError(t, err)
	require.NotNil(t, me)
	require.True(t, me.IsBot)
	require.NotNil(t, me.Username)
	ctx = context.WithValue(ctx, TestDataKey(botUserIDCtxKey), me.ID)
	ctx = context.WithValue(ctx, TestDataKey(botUsernameCtxKey), me.Username)
	return ctx
}
