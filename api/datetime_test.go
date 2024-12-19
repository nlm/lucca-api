package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDateMarshallJSON(t *testing.T) {
	dt := NewDate(2024, 12, 19)
	b, err := json.Marshal(dt)
	require.NoError(t, err)
	assert.Equal(t, `"2024-12-19"`, string(b))
}
