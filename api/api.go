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
	Logger     *log.Logger
	Cache      bool
}

type Client struct {
	baseUrl URL
	hc      http.Client
	logger  *log.Logger
	cache   map[CacheKey][]byte
}

func NewClient(options ClientOptions) *Client {
	u := URL{
		Scheme: "https",
		Host:   options.Host,
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
	client := &Client{
		hc: http.Client{
			Jar:       jar,
			Transport: transport,
		},
		baseUrl: u,
		logger:  options.Logger,
	}
	if options.Cache {
		client.cache = make(map[CacheKey][]byte)
	}
	return client
}

func (c *Client) Get(ctx context.Context, extraPath string, getParams any, result any) error {
	url := c.baseUrl.WithExtraPath(extraPath).WithGetParams(getParams)
	if cachedBody := c.GetCache(CacheKeyOf(http.MethodGet, extraPath, getParams)); cachedBody != nil {
		if c.logger != nil {
			c.logger.Printf("response from cache")
		}
		return json.Unmarshal(cachedBody, result)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	// log.Printf("REQ->%v\n", req)
	if c.logger != nil {
		c.logger.Printf("URL->%v\n", req.URL.String())
	}
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
	if c.logger != nil {
		log.Print("BODY->", string(body))
		log.Printf("RES->%v\n", *res)
	}
	if reflect.TypeOf(result).Kind() != reflect.Pointer {
		if c.logger != nil {
			log.Printf("warning: result is not a pointer")
		}
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	c.SetCache(CacheKeyOf(http.MethodGet, extraPath, getParams), body)
	return nil
}

func (c *Client) Post(ctx context.Context, extraPath string, postParams any, result any) error {
	url := c.baseUrl.WithExtraPath(extraPath)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	// log.Printf("REQ->%v\n", req)
	if c.logger != nil {
		c.logger.Printf("URL->%v\n", req.URL.String())
	}
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
	if c.logger != nil {
		log.Print("BODY->", string(body))
		log.Printf("RES->%v\n", *res)
	}
	if reflect.TypeOf(result).Kind() != reflect.Pointer {
		if c.logger != nil {
			log.Printf("warning: result is not a pointer")
		}
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

func Ptr[T any](v T) *T {
	return &v
}

// Get is an helper to template GET methods
func Get[REQ, RES any](c *Client, ctx context.Context, extraPath string, req *REQ) (*RES, error) {
	res := new(RES)
	err := c.Get(ctx, extraPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Post is an helper to template Post methods
func Post[REQ, RES any](c *Client, ctx context.Context, extraPath string, req *REQ) (*RES, error) {
	res := new(RES)
	err := c.Post(ctx, extraPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
