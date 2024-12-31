package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ParamsTest struct {
	Field1 string   `json:"field1,omitempty"`
	Field2 int      `json:"field2,omitempty"`
	Date   Date     `json:"date,omitempty"`
	Field3 *int     `json:"field3,omitempty"`
	Field4 []int    `json:"field4,omitempty"`
	Field5 []string `json:"field5,omitempty"`
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
	pt := ParamsTest{
		Field1: "demo1",
		Field2: 2,
		Date:   NewDate(2024, 12, 31),
		Field3: Ptr(42),
		Field4: []int{1, 2, 3, 4},
		Field5: []string{"x", " "},
	}
	u = u.WithGetParams(pt)
	assert.Equal(t, "date=2024-12-31&field1=demo1&field2=2&field3=42&field4=1,2,3,4&field5=x,%20", u.RawQuery)
}
