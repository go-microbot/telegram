package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// CopyMessage represents method to copy messages of any kind.
// The method is analogous to the method forwardMessages, but the copied message doesn't have
// a link to the original message. Returns the MessageId of the sent message on success.
func (api *TelegramAPI) CopyMessage(ctx context.Context, req apiModels.CopyMessageRequest) (*models.MessageID, error) {
	resp, err := api.NewRequest("copyMessage").
		Query(query.AsMap(req)).
		Method(http.MethodPost).
		Body(NewJSONBody(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var msgID models.MessageID
	if err := resp.Decode(&msgID); err != nil {
		return nil, err
	}

	return &msgID, nil
}
