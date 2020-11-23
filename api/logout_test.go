package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type logOut struct{}

func (h logOut) Test(ctx context.Context, t *testing.T) context.Context {
	prevToken := localAPI.token
	defer func() {
		localAPI.token = prevToken
	}()
	localAPI.token = "invalid"
	err := localAPI.Logout(ctx)
	require.Error(t, err)

	return ctx
}
