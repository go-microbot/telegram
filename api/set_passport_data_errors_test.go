package api

import (
	"context"
	"encoding/base64"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/stretchr/testify/require"
)

type setPassportDataErrors struct{}

func (h setPassportDataErrors) Test(ctx context.Context, t *testing.T) context.Context {
	adminUserID := ctx.Value(adminUserIDCtxKey)
	require.NotNil(t, adminUserID)

	err := localAPI.SetPassportDataErrors(ctx, apiModels.SetPassportDataErrorsRequest{
		UserID: adminUserID.(int32),
		Errors: []interface{}{
			models.PassportElementErrorDataField{
				DataHash:  base64.StdEncoding.EncodeToString([]byte{1, 2, 3}),
				FieldName: "first_name",
				Message:   "Test error",
				Source:    "data",
				Type:      models.PassportElementTypeDriverLicense,
			},
		},
	})
	if err != nil {
		require.Contains(t, err.Error(), "DATA_HASH_SIZE_INVALID")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
