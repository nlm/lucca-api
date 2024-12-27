package api

import (
	"net/http"
	"time"
)

var _ http.RoundTripper = (*ThrottleRoundTripper)(nil)

// ThrottleRoundTripper regulates the number of requests that are made on Lucca api.
type ThrottleRoundTripper struct {
	bucket chan time.Time
	next   http.RoundTripper
}

// NewThrottleRoundTripper returns a http.Roundtripper, allowing 'qps' requests per second, and a burst of 'burst'.
func NewThrottleRoundTripper(next http.RoundTripper, qps int, burst int) *ThrottleRoundTripper {
	bucket := make(chan time.Time, burst)
	go func() {
		for t := range time.Tick(1 * time.Second / time.Duration(qps)) {
			bucket <- t
		}
	}()
	return &ThrottleRoundTripper{
		bucket: bucket,
		next:   next,
	}
}

func (rt *ThrottleRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	<-rt.bucket
	return rt.next.RoundTrip(req)
}
