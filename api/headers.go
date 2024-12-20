package api

import (
	"net/http"
)

type HeadersRoundTripper struct {
	next    http.RoundTripper
	Headers map[string]string
}

func (rt HeadersRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	for k, v := range rt.Headers {
		r.Header.Set(k, v)
	}
	return rt.next.RoundTrip(r)
}

func NewHeadersRoundTripper(rt http.RoundTripper, headers map[string]string) *HeadersRoundTripper {
	return &HeadersRoundTripper{
		next:    rt,
		Headers: headers,
	}
}
