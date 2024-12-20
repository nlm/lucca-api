package api

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type MockRoundTripper struct {
	Log      bool
	Response []byte
}

func (rt *MockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	if rt.Log {
		log.Printf("Request: %+v", r)
	}
	var body = []byte("{}")
	if rt.Response != nil {
		body = rt.Response
	}
	return &http.Response{
		Body:       io.NopCloser(bytes.NewReader(body)),
		Status:     "MOCK",
		StatusCode: 200,
	}, nil
}

func NewMockRoundTripper(log bool, response []byte) *MockRoundTripper {
	return &MockRoundTripper{
		Log:      log,
		Response: response,
	}
}
