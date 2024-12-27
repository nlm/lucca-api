package timesheets

import (
	"context"
	"fmt"

	"github.com/nlm/lucca-api/api"
)

// /timmi-timesheet/api/timesheets/13815/approve

type ApproveTimesheetRequest struct {
	Id        int   `json:"-"`
	Transfers []any `json:"transfers"` // idk what's inside, but the field is present
}

type ApproveTimesheetResponse []ApproveTimesheetResponseItem

type ApproveTimesheetResponseItem struct {
	Id               int           `json:"id"`
	ExceptionMessage string        `json:"exceptionMessage"`
	StartsOn         *api.DateTime `json:"startsOn"`
	EndsOn           *api.DateTime `json:"endsOn"`
	Status           string        `json:"status"`
	// Owner *Person
	// ExpectedNextActor *Person
}

func (ts *TimesheetsService) ApproveTimesheet(ctx context.Context, req *ApproveTimesheetRequest) (*ApproveTimesheetResponse, error) {
	res := new(ApproveTimesheetResponse)
	err := ts.client.Post(ctx, fmt.Sprintf("/timmi-timesheet/api/timesheets/%d/approve", req.Id), req, res)
	return res, err
}
