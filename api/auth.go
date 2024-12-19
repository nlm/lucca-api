package api

import (
	"net/http"
)

type LuccaAuthRoundTripper struct {
	token string
	rt    http.RoundTripper
}

func NewLuccaAuthRoundTripper(token string, rt http.RoundTripper) *LuccaAuthRoundTripper {
	if rt == nil {
		rt = http.DefaultTransport
	}
	return &LuccaAuthRoundTripper{
		token: token,
		rt:    rt,
	}
}

func (rt *LuccaAuthRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", "lucca application="+rt.token)
	return rt.rt.RoundTrip(r)
}
