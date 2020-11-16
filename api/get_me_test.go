package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type getMe struct{}

func (h getMe) Test(ctx context.Context, t *testing.T) {
	me, err := localAPI.GetMe(context.Background())
	require.NoError(t, err)
	require.NotNil(t, me)
	require.True(t, me.IsBot)
	require.NotNil(t, me.Username)
}
