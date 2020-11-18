package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/go-microbot/telegram/form"
)

const (
	contentTypeHeader = "Content-Type"
)

// Response represents base API response.
type Response struct {
	resp *http.Response
}

// Request represents API request configuration.
type Request struct {
	url        string
	httpMethod string
	apiMethod  string
	body       *RequestBody
	query      map[string]string
	headers    map[string]string
	client     *http.Client
}

// RequestBody represents model for body request.
type RequestBody struct {
	m    BodyMarshaler
	body interface{}
}

// BodyMarshaler represents body marshaler func.
type BodyMarshaler func(v interface{}, ct *string) ([]byte, error)

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

// NewJSONBody returns new RequestBody with JSON marshaler.
func NewJSONBody(body interface{}) *RequestBody {
	return &RequestBody{
		m:    jsonMarsaler,
		body: body,
	}
}

// NewFormDataBody returns new RequestBody with Form Data marshaler.
func NewFormDataBody(body interface{}) *RequestBody {
	return &RequestBody{
		m:    form.Marshal,
		body: body,
	}
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

// Marshal marshal request body.
func (b *RequestBody) Marshal(contentType *string) ([]byte, error) {
	return b.m(b.body, contentType)
}

// Method sets HTTP method, like GET, POST, etc.
// GET by default if not provided.
func (req *Request) Method(method string) *Request {
	req.httpMethod = method
	return req
}

// Body sets request body.
func (req *Request) Body(body *RequestBody) *Request {
	req.body = body
	return req
}

// Query sets query params.
func (req *Request) Query(q map[string]string) *Request {
	req.query = q
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
		if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
			return nil, newErr(ErrReqTimeout, err)
		}

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
	if req.body == nil {
		return nil, nil
	}

	var cType string
	data, err := req.body.Marshal(&cType)
	if err != nil {
		return nil, err
	}

	// detect content-type.
	if cType == "" {
		cType = http.DetectContentType(data)
	}

	if len(req.headers) == 0 {
		req.headers = make(map[string]string)
	}

	if _, ok := req.headers[contentTypeHeader]; !ok {
		req.headers[contentTypeHeader] = cType
	}

	return bytes.NewBuffer(data), nil
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

func jsonMarsaler(v interface{}, ct *string) ([]byte, error) {
	if ct != nil {
		*ct = "application/json"
	}

	return json.Marshal(v)
}
