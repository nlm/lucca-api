package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"reflect"
)

type ClientOptions struct {
	Host       string
	AuthCookie string
	AppToken   string
	Transport  http.RoundTripper
}

type Client struct {
	baseUrl URL
	hc      http.Client
}

func NewClient(options ClientOptions) *Client {
	u := URL{
		Scheme: "https",
		Host:   options.Host,
		// Path:   "",
	}
	jar, _ := cookiejar.New(nil)
	// Handle Cookie Authentication
	if options.AuthCookie != "" {
		jar.SetCookies((*url.URL)(&u), []*http.Cookie{
			{
				Name:   "authToken",
				Value:  options.AuthCookie,
				Secure: true,
			},
		})
	}
	// Handle Application Token Authentication
	transport := options.Transport
	if options.AppToken != "" {
		transport = NewLuccaAuthRoundTripper(options.AppToken, transport)
	}
	return &Client{
		hc: http.Client{
			Jar:       jar,
			Transport: transport,
		},
		baseUrl: u,
	}
}

func (c *Client) Get(ctx context.Context, extraPath string, getParams any, result any) error {
	url := c.baseUrl.WithExtraPath(extraPath).WithGetParams(getParams)
	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	// log.Printf("REQ->%v\n", req)
	log.Printf("URL->%v\n", req.URL.String())
	res, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode/100 != 2 {
		return fmt.Errorf("error: %s", res.Status)
	}
	body, err := io.ReadAll(io.LimitReader(res.Body, 1000000000))
	if err != nil {
		return err
	}
	log.Print("BODY->", string(body))
	log.Printf("RES->%v\n", *res)
	if reflect.TypeOf(result).Kind() != reflect.Pointer {
		log.Printf("warning: result is not a pointer")
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	return nil
}

// HTTPClient returns the inner http.Client contained in Client.
func (c *Client) HTTPClient() *http.Client {
	return &c.hc
}
