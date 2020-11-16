package api

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

type setWebhook struct{}

func (h setWebhook) Test(ctx context.Context, t *testing.T) {
	// load certificate.
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	webhookCert, err = ioutil.ReadFile(path.Join(wd, "./test_data/telegram_test.key"))
	if err != nil {
		panic(err)
	}
}

/*func TestTelegramAPI_SetWebhook(t *testing.T) {
	// enable hook.
	cert := models.InputFile(webhookCert)
	err := localAPI.SetWebhook(context.Background(), apiModels.SetWebhookRequest{
		Certificate: &cert,
		URL:         query.NewParamString("localhost:1012"),
	})
	require.NoError(t, err)

	// disable hook.
	err = localAPI.SetWebhook(context.Background(), apiModels.SetWebhookRequest{})
	require.NoError(t, err)
}
*/
