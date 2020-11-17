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
	contentTypeHeader = "Contet-Type"
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
	bodyData   []byte
	client     *http.Client
}

// RequestBody represents model for body request.
type RequestBody struct {
	m    BodyMarshaler
	body interface{}
}

// BodyMarshaler represents body marshaler func.
type BodyMarshaler func(v interface{}) ([]byte, error)

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
		m:    json.Marshal,
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
func (b *RequestBody) Marshal() ([]byte, error) {
	return b.m(b.body)
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

	//panic(body)

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
		///
		var r apiResponse
		json.NewDecoder(result.Body).Decode(&r)
		panic(r.Description)
		///
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

	data, err := req.body.Marshal()
	if err != nil {
		return nil, err
	}
	req.bodyData = data

	return bytes.NewBuffer(data), nil

	/*var body io.Reader

	if len(req.headers) == 0 {
		req.headers = make(map[string]string)
	}

	// create metadata.
	var metaData []byte
	var err error
	if req.body != nil {
		metaData, err = json.Marshal(req.body)
		if err != nil {
			return nil, err
		}

		if len(req.formData) == 0 {
			body = bytes.NewBuffer(metaData)
			req.headers["Content-Type"] = "application/json"
			return body, nil
		}
	}

	// create metadata form part.
	formBody := &bytes.Buffer{}
	writer := multipart.NewWriter(formBody)

	if len(metaData) != 0 {
		/*metadataHeader := textproto.MIMEHeader{}
		metadataHeader.Set("Content-Type", "application/json")
		metadataHeader.Set("Content-ID", "body")
		part, err := writer.CreatePart(metadataHeader)
		if err != nil {
			return nil, err
		}
		_, err = part.Write([]byte(metaData))
		if err != nil {
			return nil, err
		}

		err = writer.WriteField("", string(metaData))
		if err != nil {
			return nil, err
		}
	}

	///
	/*mediaData, err := ioutil.ReadFile("./test_data/test_photo.png")
	if err != nil {
		return nil, err
	}
	mediaHeader := textproto.MIMEHeader{}
	mediaHeader.Set("Content-Disposition", fmt.Sprintf("attachment; name=\"photo\"; filename=\"%v\".", "./test_data/test_photo.png"))
	mediaHeader.Set("Content-ID", "photo")
	mediaHeader.Set("Content-Filename", "photo.png")

	mediaPart, err := writer.CreatePart(mediaHeader)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(mediaPart, bytes.NewReader(mediaData))
	if err != nil {
		return nil, err
	}

	file, err := os.Open("./test_data/test_photo.png")
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	part, err := writer.CreateFormFile("photo", fi.Name())
	if err != nil {
		return nil, err
	}
	if _, err := part.Write(fileContents); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	body = formBody
	req.headers["Content-Type"] = writer.FormDataContentType() // fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary())

	/*switch {
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
	}*/
}

func (req *Request) setHeaders(request *http.Request) {
	if len(req.headers) == 0 {
		req.headers = make(map[string]string)
	}

	// detect content-type.
	if _, ok := req.headers[contentTypeHeader]; !ok && len(req.bodyData) != 0 {
		req.headers[contentTypeHeader] = form.CT //http.DetectContentType(req.bodyData)
	}

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
