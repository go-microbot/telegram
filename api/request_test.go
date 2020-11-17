package api

/*type invalidMarshal struct {
	err error
}

type invalidUnmarshal struct {
	err error
}

func (i invalidMarshal) MarshalJSON() ([]byte, error) {
	return nil, i.err
}

func (i invalidUnmarshal) UnmarshalJSON([]byte) error {
	return i.err
}

func serverMock(pattern string) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc(pattern, func(rw http.ResponseWriter, r *http.Request) {
		switch r.Header.Get("x-test") {
		case "status 500":
			rw.WriteHeader(http.StatusInternalServerError)
		default:
			rw.WriteHeader(http.StatusOK)
		}
	})

	return httptest.NewServer(handler)
}

func TestTelegramAPI_NewRequest(t *testing.T) {
	api := NewTelegramAPI("123")
	req := api.NewRequest("test")
	require.NotNil(t, req)
	require.Equal(t, fmt.Sprintf("%s/bot%s/%s", api.url, "123", "test"), req.url)
	require.Equal(t, "test", req.apiMethod)
	require.NotNil(t, req.client)
	require.Equal(t, http.MethodGet, req.httpMethod)
}

func TestRequest_Method(t *testing.T) {
	api := NewTelegramAPI("123")
	req := api.NewRequest("test")
	require.NotNil(t, req)
	req.Method(http.MethodPost)
	require.Equal(t, http.MethodPost, req.httpMethod)
}

func TestRequest_Body(t *testing.T) {
	api := NewTelegramAPI("123")
	req := api.NewRequest("test")
	require.NotNil(t, req)
	req.Body("data")
	require.Equal(t, "data", req.body)
}

func TestRequest_Query(t *testing.T) {
	api := NewTelegramAPI("123")
	req := api.NewRequest("test")
	require.NotNil(t, req)
	q := map[string]string{
		"hello": "world",
	}
	req.Query(q)
	require.Equal(t, q, req.query)
}

func TestRequest_FormData(t *testing.T) {
	api := NewTelegramAPI("123")
	req := api.NewRequest("test")
	require.NotNil(t, req)
	fd := map[string]string{
		"hello": "world",
	}
	req.FormData(fd)
	require.Equal(t, fd, req.formData)
}

func TestRequest_Headers(t *testing.T) {
	api := NewTelegramAPI("123")
	req := api.NewRequest("test")
	require.NotNil(t, req)
	h := map[string]string{
		"hello": "world",
	}
	req.Headers(h)
	require.Equal(t, h, req.headers)
}

func TestRequest_CustomClient(t *testing.T) {
	api := NewTelegramAPI("123")
	req := api.NewRequest("test")
	require.NotNil(t, req)
	client := &http.Client{}
	req.CustomClient(client)
	require.Equal(t, client, req.client)
}

func TestRequest_Do(t *testing.T) {
	t.Run("encode body error (JSON)", func(t *testing.T) {
		api := NewTelegramAPI("123")
		req := api.NewRequest("test").Body(invalidMarshal{
			err: errors.New("error"),
		})
		require.NotNil(t, req)
		_, err := req.Do(context.Background())
		require.Error(t, err)
		require.True(t, errors.Is(err, ErrEncodeBody))
	})
	t.Run("prepare request error", func(t *testing.T) {
		api := NewTelegramAPI("123")
		req := api.NewRequest("test").Method(fmt.Sprintf("method%c", rune(10)))
		require.NotNil(t, req)
		_, err := req.Do(context.Background())
		require.Error(t, err)
		require.True(t, errors.Is(err, ErrPrepareReq))
	})
	t.Run("send request error", func(t *testing.T) {
		api := NewTelegramAPI("123")
		req := api.NewRequest("test").FormData(map[string]string{
			"some": "field",
		})
		require.NotNil(t, req)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := req.Do(ctx)
		require.Error(t, err)
		require.True(t, errors.Is(err, ErrSendReq))
	})
	t.Run("500 status code error", func(t *testing.T) {
		api := NewTelegramAPI("123")
		srv := serverMock("/bot123/test")
		defer srv.Close()
		api.url = srv.URL
		req := api.NewRequest("test").
			FormData(map[string]string{
				"some": "field",
			}).
			Headers(map[string]string{
				"x-test": "status 500",
			}).
			Query(map[string]string{
				"param": "test",
			})
		require.NotNil(t, req)
		_, err := req.Do(context.Background())
		require.Error(t, err)
		require.True(t, errors.Is(err, ErrResponse))
	})
	t.Run("all ok", func(t *testing.T) {
		api := NewTelegramAPI("123")
		srv := serverMock("/bot123/test")
		defer srv.Close()
		api.url = srv.URL
		req := api.NewRequest("test").
			Body("some data").
			Query(map[string]string{
				"param": "test",
			})
		require.NotNil(t, req)
		resp, err := req.Do(context.Background())
		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func TestResponse_Decode(t *testing.T) {
	t.Run("decode body error", func(t *testing.T) {
		resp := Response{
			resp: &http.Response{
				Body: http.NoBody,
			},
		}
		var result map[string]string
		err := resp.Decode(&result)
		require.Error(t, err)
		require.True(t, errors.Is(err, ErrDecodeBody))
	})
	t.Run("not OK result error", func(t *testing.T) {
		data, err := json.Marshal(apiResponse{
			badResponse: badResponse{
				ErrorCode:   100,
				Description: "error",
			},
		})
		require.NoError(t, err)
		resp := Response{
			resp: &http.Response{
				Body: ioutil.NopCloser(bytes.NewBuffer(data)),
			},
		}
		var result map[string]string
		err = resp.Decode(&result)
		require.Error(t, err)
		require.True(t, errors.Is(err, ErrResponse))
	})
	t.Run("unmarshal good response error", func(t *testing.T) {
		data, err := json.Marshal(apiResponse{
			OK: true,
		})
		require.NoError(t, err)
		resp := Response{
			resp: &http.Response{
				Body: ioutil.NopCloser(bytes.NewBuffer(data)),
			},
		}
		result := invalidUnmarshal{
			err: errors.New("error"),
		}
		err = resp.Decode(&result)
		require.Error(t, err)
		require.True(t, errors.Is(err, ErrDecodeBody))
	})
	t.Run("all ok", func(t *testing.T) {
		body := map[string]interface{}{
			"key": "value",
		}
		data, err := json.Marshal(body)
		require.NoError(t, err)
		data, err = json.Marshal(apiResponse{
			OK: true,
			goodResponse: goodResponse{
				Result: data,
			},
		})
		require.NoError(t, err)
		resp := Response{
			resp: &http.Response{
				Body: ioutil.NopCloser(bytes.NewBuffer(data)),
			},
		}
		var result map[string]interface{}
		err = resp.Decode(&result)
		require.NoError(t, err)
		require.Equal(t, body, result)
	})
}
*/
