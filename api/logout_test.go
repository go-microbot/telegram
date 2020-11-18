package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type logOut struct{}

func (h logOut) Test(ctx context.Context, t *testing.T) context.Context {
	err := localAPI.Logout(ctx)
	require.NoError(t, err)

	return ctx
}
