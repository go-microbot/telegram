package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendPhoto represents method to send photos. On success, the sent Message is returned.
func (api *TelegramAPI) SendPhoto(ctx context.Context, req apiModels.SendPhotoRequest) (*models.Message, error) {
	request := api.NewRequest("sendPhoto").
		Method(http.MethodPost).
		Query(query.AsMap(req))
	//Body(req)

	/*if photo, ok := req.Photo.(models.InputFile); ok {
		request.FormData(map[string]string{
			"photo": string(photo),
		})
	}*/

	request.FormData(map[string]string{
		"photo": req.Photo.(string),
	})

	resp, err := request.Do(ctx)
	if err != nil {
		return nil, err
	}

	var msg models.Message
	if err := resp.Decode(&msg); err != nil {
		return nil, err
	}

	return &msg, err
}
