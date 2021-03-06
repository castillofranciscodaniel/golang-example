package main

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpointsRestyPost(t *testing.T) {
	r := Routes
	ts := httptest.NewServer(r)
	defer ts.Close()

	testcases := map[string]struct {
		method   string
		path     string
		body     string
		header   http.Header
		wantCode int
		wantBody string
	}{
		"POST /message": {
			method: http.MethodPost,
			path:   "/message",
			body:   `"{did":"holi","msisdn":"5491159513122"}`,
			header: map[string][]string{
				"Content-Type": {"application/json"},
			},
			wantCode: http.StatusCreated,
			wantBody: `{"did":"holi","msisdn":"5491159513122"}`,
		},
	}

	for name, test := range testcases {
		for i := 0; i < 1000; i++ {
			t.Run(name, func(t *testing.T) {
				body := bytes.NewReader([]byte(test.body))
				gotResponse, gotBody := testRequest(t, ts, test.method, test.path, body, test.header)
				assert.Equal(t, test.wantCode, gotResponse.StatusCode)
				if test.wantBody != "" {
					assert.Equal(t, test.wantBody, gotBody, "body did not match")
				}
			})
		}
	}
}

func TestEndpointsRestyGet(t *testing.T) {
	r := Routes
	ts := httptest.NewServer(r)
	defer ts.Close()

	testcases := map[string]struct {
		method   string
		path     string
		body     string
		header   http.Header
		wantCode int
		wantBody string
	}{
		"GET /message": {
			method: http.MethodPost,
			path:   "/message",
			//body:   `{"name":"Skittles","price":1.99}`,
			header: map[string][]string{
				"Content-Type": {"application/json"},
			},
			wantCode: http.StatusCreated,
			wantBody: `[{"did":"edit did","msisdn":"5491159513122"},{"did":"edit did","msisdn":"5491159513122"}]`,
		},
	}

	for name, test := range testcases {
		for i := 0; i < 1; i++ {
			t.Run(name, func(t *testing.T) {
				body := bytes.NewReader([]byte(test.body))
				gotResponse, gotBody := testRequest(t, ts, test.method, test.path, body, test.header)
				assert.Equal(t, test.wantCode, gotResponse.StatusCode)
				if test.wantBody != "" {
					assert.Equal(t, test.wantBody, gotBody, "body did not match")
				}
			})
		}
	}
}

// testRequest is a helper function to exectute the http request against the server
func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader, header http.Header) (*http.Response, string) {
	t.Helper()
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	respBody = bytes.TrimSpace(respBody)

	return resp, string(respBody)
}
