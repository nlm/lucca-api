package api

import (
	"encoding/json"
)

type CacheKeyer interface {
	CacheKey() []byte
}

type CacheKey struct {
	Method  string
	Path    string
	Request string
}

func CacheKeyOf(method, path string, v any) CacheKey {
	var req []byte
	var err error
	if c, ok := v.(CacheKeyer); ok {
		req = c.CacheKey()
	} else {
		req, err = json.Marshal(v)
		if err != nil {
			panic(err)
		}
	}
	return CacheKey{
		Method:  method,
		Path:    path,
		Request: string(req),
	}
}

func (c *Client) SetCache(key CacheKey, body []byte) {
	if c.cache != nil {
		c.cache[key] = body
	}
}

func (c *Client) GetCache(key CacheKey) []byte {
	return c.cache[key]
}
