package timesheets

import (
	"encoding/json"
	"testing"

	"github.com/nlm/lucca-api/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTimesheetDetailsRequestMarshallJSON(t *testing.T) {
	gtdr := ListTimesheetDetailsRequest{
		OwnerId: 0,
		From:    api.NewDate(2024, 12, 01),
		Until:   api.NewDate(2024, 12, 31),
	}
	b, err := json.Marshal(gtdr)
	require.NoError(t, err)
	assert.Equal(t, "prout", string(b))
}
