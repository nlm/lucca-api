package timesheets

import (
	"context"

	"github.com/nlm/lucca-api/api"
)

type ListTimesheetDetailsRequest struct {
	OwnerId int      `json:"ownerId"`
	From    api.Date `json:"from,omitempty"`
	Until   api.Date `json:"until,omitempty"`
}

type Duration struct {
	Value float32 `json:"value,omitempty"`
	Unit  string  `json:"unit,omitempty"`
}

type Imputation struct {
	Id           string         `json:"id,omitempty"`
	AxisSections []AxisSections `json:"axisSections,omitempty"`
}

type AxisSections struct {
	Id    int        `json:"id"`
	Name  string     `json:"name"`
	Code  string     `json:"code"`
	Axis  *Axis      `json:"axis,omitempty"`
	Child *AxisChild `json:"child,omitempty"`
}

type Axis struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type AxisChild struct {
	Id       int        `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	Code     string     `json:"code,omitempty"`
	IsActive bool       `json:"isActive"`
	Child    *AxisChild `json:"child"`
}

type ListTimesheetDetailsResponse api.ItemsList[ListTimesheetDetailsItem]

const (
	ItemTypeLeave      = "leave"
	ItemTypeHalfDayOff = "halfDayOff"
	ItemTypeTimeEntry  = "timeEntry"
)

type ListTimesheetDetailsItem struct {
	Id          int           `json:"id"`
	OwnerId     int           `json:"ownerId,omitempty"`
	StartsAt    *api.DateTime `json:"startsAt,omitempty"` // FIXME omitzero + remove pointer
	EndsAt      *api.DateTime `json:"endsAt,omitempty"`   // FIXME omitzero + remove pointer
	Duration    *Duration     `json:"duration,omitempty"`
	Imputation  *Imputation   `json:"imputation,omitempty"`
	AccountName *string       `json:"accountName,omitempty"`
	Position    string        `json:"position,omitempty"`
	Type        string        `json:"type"` // `leave` or `halfDayOff` or `timeEntry`
	// {
	// 	"id": 314837,
	// 	"ownerId": 368,
	// 	"startsAt": "2024-10-28T00:00:00",
	// 	"endsAt": "2024-10-29T00:00:00",
	// 	"duration": {
	// 		"value": 1,
	// 		"iso": "P1D",
	// 		"unit": "day"
	// 	},
	// 	"imputation": {
	// 		"id": "22-147",
	// 		"axisSections": [
	// 			{
	// 				"child": {
	// 					"child": null,
	// 					"id": 147,
	// 					"axis": {
	// 						"id": 20,
	// 						"name": "LOTS",
	// 						"isVisible": true,
	// 						"isEditable": true,
	// 						"selectedDisplayOption": "code"
	// 					},
	// 					"name": "LOT 0 Websites",
	// 					"code": "LOT 0 Websites",
	// 					"ownerId": null,
	// 					"isActive": true,
	// 					"hasAvailableChildren": false,
	// 					"startsOn": null,
	// 					"endsOn": null
	// 				},
	// 				"id": 22,
	// 				"axis": {
	// 					"id": 19,
	// 					"name": "Engine",
	// 					"isVisible": true,
	// 					"isEditable": true,
	// 					"selectedDisplayOption": "codeName"
	// 				},
	// 				"name": "BUILD",
	// 				"code": "BUILD",
	// 				"ownerId": null,
	// 				"isActive": true,
	// 				"hasAvailableChildren": false,
	// 				"startsOn": null,
	// 				"endsOn": null
	// 			}
	// 		],
	// 		"isActive": true,
	// 		"isStarred": false,
	// 		"startsOn": null,
	// 		"endsOn": null
	// 	},
	// 	"comment": null,
	// 	"unit": "day",
	// 	"timeType": null,
	// 	"validationErrors": [],
	// 	"type": "timeEntry",
	// 	"creationSource": "manual"
	// },
}

func (ts *TimesheetsService) ListTimesheetDetails(ctx context.Context, req *ListTimesheetDetailsRequest) (*ListTimesheetDetailsResponse, error) {
	res := new(ListTimesheetDetailsResponse)
	err := ts.client.Get(ctx, "/timmi-timesheet/services/timesheet-days/details", req, res)
	return res, err
}
