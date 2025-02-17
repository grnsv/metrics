package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleUpdateMetric(t *testing.T) {
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	type req struct {
		method      string
		url         string
		contentType string
	}
	type want struct {
		statusCode  int
		body        string
		contentType string
	}
	tests := []struct {
		name string
		req  req
		want want
	}{
		{
			"valid gauge",
			req{
				method:      http.MethodPost,
				url:         "/update/gauge/Alloc/123.45",
				contentType: "text/plain",
			},
			want{
				statusCode:  http.StatusOK,
				body:        "",
				contentType: "",
			},
		},
		{
			"valid counter",
			req{
				method:      http.MethodPost,
				url:         "/update/counter/PollCount/1",
				contentType: "text/plain",
			},
			want{
				statusCode:  http.StatusOK,
				body:        "",
				contentType: "",
			},
		},
		{
			"invalid method",
			req{
				method:      http.MethodPut,
				url:         "/update/gauge/Alloc/123.45",
				contentType: "text/plain",
			},
			want{
				statusCode: http.StatusMethodNotAllowed,
			},
		},
		{
			"invalid path",
			req{
				method:      http.MethodPost,
				url:         "/update/gauge/Alloc",
				contentType: "text/plain",
			},
			want{
				statusCode:  http.StatusNotFound,
				body:        "404 page not found\n",
				contentType: "text/plain",
			},
		},
		{
			"invalid metric name",
			req{
				method:      http.MethodPost,
				url:         "/update/gauge//123.45",
				contentType: "text/plain",
			},
			want{
				statusCode:  http.StatusNotFound,
				body:        "Metric name is required\n",
				contentType: "text/plain",
			},
		},
		{
			"invalid metric type",
			req{
				method:      http.MethodPost,
				url:         "/update/invalid/Alloc/123.45",
				contentType: "text/plain",
			},
			want{
				statusCode:  http.StatusBadRequest,
				body:        "Invalid metric type\n",
				contentType: "text/plain",
			},
		},
		{
			"invalid gauge",
			req{
				method:      http.MethodPost,
				url:         "/update/gauge/Alloc/abc",
				contentType: "text/plain",
			},
			want{
				statusCode:  http.StatusBadRequest,
				body:        "Invalid gauge value\n",
				contentType: "text/plain",
			},
		},
		{
			"invalid counter",
			req{
				method:      http.MethodPost,
				url:         "/update/counter/PollCount/abc",
				contentType: "text/plain",
			},
			want{
				statusCode:  http.StatusBadRequest,
				body:        "Invalid counter value\n",
				contentType: "text/plain",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, err := http.NewRequest(tt.req.method, ts.URL+tt.req.url, nil)
			require.NoError(t, err, tt.name)

			request.Header.Add("Content-Type", tt.req.contentType)

			res, err := ts.Client().Do(request)
			require.NoError(t, err, tt.name)

			assert.Equal(t, tt.want.statusCode, res.StatusCode, tt.name)
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err, tt.name)
			assert.Equal(t, tt.want.body, string(resBody), tt.name)
			assert.Contains(t, res.Header.Get("Content-Type"), tt.want.contentType, tt.name)
		})
	}
}
