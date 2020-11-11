package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const baseURL = "https://api.telegram.org"

type apiResponse struct {
	OK bool `json:"ok"`
	apiBadResponse
	apiGoodResponse
}

type apiBadResponse struct {
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}

type apiGoodResponse struct {
	Result json.RawMessage `json:"result,omitempty"`
}

type apiRequest struct {
	httpMethod string
	apiMethod  string
	request    interface{}
	response   interface{}
	query      map[string]string
}

func (api *TelegramAPI) newRequest(ctx context.Context, req apiRequest) error {
	// prepare URL.
	url := fmt.Sprintf("%s/bot%s/%s", baseURL, api.token, req.apiMethod)

	// prepare request data.
	var body io.Reader
	if req.request != nil {
		data, err := json.Marshal(req.request)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(data)
	}

	request, err := http.NewRequestWithContext(ctx, req.httpMethod, url, body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	query := request.URL.Query()
	for k, v := range req.query {
		query.Add(k, v)
	}

	// make request.
	result, err := api.client.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		if err := result.Body.Close(); err != nil {
			log.Print(err)
		}
	}()

	// parse response.
	var apiResp apiResponse
	if err := json.NewDecoder(result.Body).Decode(&apiResp); err != nil {
		return err
	}

	// check status code.
	if result.StatusCode > http.StatusIMUsed || !apiResp.OK {
		return fmt.Errorf("%s: status %d", apiResp.Description, apiResp.ErrorCode)
	}

	// create response.
	if req.response == nil {
		return nil
	}

	if err := json.Unmarshal(apiResp.Result, req.response); err != nil {
		return err
	}

	return nil
}
