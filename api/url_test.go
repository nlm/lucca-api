package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ParamsTest struct {
	Field1 string `json:"field1,omitempty"`
	Field2 int    `json:"field2,omitempty"`
	Date   Date   `json:"date,omitempty"`
}

func TestURL(t *testing.T) {
	u := URL{
		Scheme: "https",
		Host:   "example.ilucca.net",
		Path:   "/api/v3",
	}

	// Path
	assert.Equal(t, "/api/v3", u.Path)
	u = u.WithExtraPath("/users")
	assert.Equal(t, "/api/v3/users", u.Path)

	// Query
	u = u.WithGetParams(ParamsTest{
		Field1: "demo1",
		Field2: 2,
		Date:   NewDate(2024, 12, 31),
	})
	assert.Equal(t, "date=2024-12-31&field1=demo1&field2=2", u.RawQuery)
}
