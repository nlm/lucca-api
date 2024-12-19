package api

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type MockRoundTripper struct {
	Log bool
}

func (rt *MockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	if rt.Log {
		log.Printf("Request: %+v", r)
	}
	return &http.Response{
		Body:       io.NopCloser(strings.NewReader("{}")),
		Status:     "MOCK",
		StatusCode: 200,
	}, nil
}
