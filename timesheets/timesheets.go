package timesheets

import (
	"context"

	"github.com/nlm/lucca-api/api"
)

type TimesheetsService struct {
	client *api.Client
}

func New(client *api.Client) *TimesheetsService {
	return &TimesheetsService{
		client: client,
	}
}

type ListApprovablesRequest struct {
	api.Page
	LoadPayrollInfo   bool `json:"loadPayrollInfo"`
	IncludeOnlyOthers bool `json:"includeOnlyOthers"`
}

type ListApprovablesResponse api.ItemsList[Approvable]

type Owner struct {
	Id int `json:"id"`
	// Picture   *string `json:"picture"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Approvable struct {
	Id       int          `json:"id"`
	Uid      string       `json:"uid"`
	Owner    Owner        `json:"Owner"`
	StartsOn api.DateTime `json:"startsOn"`
	EndsOn   api.DateTime `json:"endsOn"`
	// "id": 12496,
	// "uid": "12496-360-20231018-20241104",
	// "owner": {
	// 	"picture": null,
	// 	"jobTitle": null,
	// 	"id": 368,
	// 	"firstName": "Prenom",
	// 	"lastName": "NOM",
	// 	"dtContractEnd": null,
	// 	"departmentId": 0,
	// 	"establishmentId": 0
	// },
	// "startsOn": "2024-10-28T00:00:00",
	// "endsOn": "2024-11-04T00:00:00",
	// "hasComments": false,
	// "commentsCount": 0,
	// "hasAlerts": false,
	// "alertsCount": 0,
	// "hasNonCriticalAlerts": false,
	// "nonCriticalAlertsCount": 0,
	// "hasCriticalAlerts": false,
	// "criticalAlertsCount": 0,
	// "hasPayrollVariables": false,
	// "payrollVariablesCount": 0,
	// "attendance": null,
	// "leaveDuration": null,
	// "theoreticalWorkingTime": {
	// 	"value": 1,
	// 	"iso": "P1D",
	// 	"unit": "day"
	// }
}

func (ts *TimesheetsService) ListApprovables(ctx context.Context, req *ListApprovablesRequest) (*ListApprovablesResponse, error) {
	res := new(ListApprovablesResponse)
	err := ts.client.Get(ctx, "/timmi-timesheet/api/timesheets/approvables", req, res)
	return res, err
}

type GetTimesheetDetailsRequest struct {
	OwnerId int      `json:"ownerId"`
	From    api.Date `json:"from,omitempty"`
	Until   api.Date `json:"until,omitempty"`
}

type Duration struct {
	Value int    `json:"value,omitempty"`
	Unit  string `json:"unit,omitempty"`
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
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type GetTimesheetDetailsResponse struct {
	Owner      int          `json:"ownerId,omitempty"`
	StartsAt   api.DateTime `json:"startsAt,omitempty"`
	EndsAt     api.DateTime `json:"endsAt,omitempty"`
	Duration   *Duration    `json:"duration,omitempty"`
	Imputation *Imputation  `json:"imputation,omitempty"`
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

func (ts *TimesheetsService) GetTimesheetDetails(ctx context.Context, req *GetTimesheetDetailsRequest) (*GetTimesheetDetailsResponse, error) {
	res := new(GetTimesheetDetailsResponse)
	err := ts.client.Get(ctx, "/timmi-timesheet/services/timesheet-days/details", req, res)
	return res, err
}
