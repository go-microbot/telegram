package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// Response represents base API response.
type Response struct {
	resp *http.Response
}

type apiResponse struct {
	OK bool `json:"ok"`
	badResponse
	goodResponse
}

type badResponse struct {
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}

type goodResponse struct {
	Result json.RawMessage `json:"result,omitempty"`
}

// Request represents API request configuration.
type Request struct {
	url        string
	httpMethod string
	apiMethod  string
	body       interface{}
	query      map[string]string
	formData   map[string]string
	headers    map[string]string
	client     *http.Client
}

// NewRequest returns new Request instance.
func (api *TelegramAPI) NewRequest(apiMethod string) *Request {
	return &Request{
		url:        fmt.Sprintf("%s/bot%s/%s", api.url, api.token, apiMethod),
		apiMethod:  apiMethod,
		client:     api.client,
		httpMethod: http.MethodGet,
	}
}

// Method sets HTTP method, like GET, POST, etc.
// GET by default if not provided.
func (req *Request) Method(method string) *Request {
	req.httpMethod = method
	return req
}

// Body sets request body.
// Will be converted to JSON before send.
func (req *Request) Body(body interface{}) *Request {
	req.body = body
	return req
}

// Query sets query params.
func (req *Request) Query(q map[string]string) *Request {
	req.query = q
	return req
}

// FormData sets form data.
func (req *Request) FormData(data map[string]string) *Request {
	req.formData = data
	return req
}

// Headers sets request custom headers.
func (req *Request) Headers(h map[string]string) *Request {
	req.headers = h
	return req
}

// CustomClient sets custom HTTP client to send request.
func (req *Request) CustomClient(client *http.Client) *Request {
	req.client = client
	return req
}

// Do sends a HTTP request.
func (req *Request) Do(ctx context.Context) (*Response, error) {
	// prepare body.
	body, err := req.prepareBody()
	if err != nil {
		return nil, newErr(ErrEncodeBody, err)
	}

	// create request.
	request, err := http.NewRequestWithContext(ctx, req.httpMethod, req.url, body)
	if err != nil {
		return nil, newErr(ErrPrepareReq, err)
	}

	// set headers.
	req.setHeaders(request)

	// set query params.
	request.URL.RawQuery = req.encodeQuery(request)

	// make request.
	result, err := req.client.Do(request)
	if err != nil {
		return nil, newErr(ErrSendReq, err)
	}

	// check response status code.
	if !isValidStatusCode(result.StatusCode) {
		return nil, newErr(ErrResponse,
			fmt.Errorf("status %d: %s", result.StatusCode, result.Status))
	}

	return &Response{
		resp: result,
	}, nil
}

// Decode unmarshal response body.
func (res *Response) Decode(resp interface{}) error {
	var apiResp apiResponse
	if err := json.NewDecoder(res.resp.Body).Decode(&apiResp); err != nil {
		return newErr(ErrDecodeBody, err)
	}

	if !apiResp.OK {
		return newErr(ErrResponse,
			fmt.Errorf("status %d: %s", apiResp.ErrorCode, apiResp.Description))
	}

	if err := json.Unmarshal(apiResp.Result, resp); err != nil {
		return newErr(ErrDecodeBody, err)
	}

	return nil
}

func (req *Request) prepareBody() (io.Reader, error) {
	var body io.Reader

	if len(req.headers) == 0 {
		req.headers = make(map[string]string)
	}

	switch {
	case req.body != nil:
		data, err := json.Marshal(req.body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(data)
		req.headers["Content-Type"] = "application/json"
	case len(req.formData) != 0:
		formData := new(bytes.Buffer)
		writer := multipart.NewWriter(formData)
		for k, v := range req.formData {
			if err := writer.WriteField(k, v); err != nil {
				return nil, err
			}
		}
		err := writer.Close()
		if err != nil {
			return nil, err
		}
		body = formData
		req.headers["Content-Type"] = writer.FormDataContentType()
	}

	return body, nil
}

func (req *Request) setHeaders(request *http.Request) {
	for k, v := range req.headers {
		request.Header.Set(k, v)
	}
}

func (req *Request) encodeQuery(request *http.Request) string {
	query := request.URL.Query()
	for k, v := range req.query {
		query.Add(k, v)
	}

	return query.Encode()
}

func isValidStatusCode(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices
}
