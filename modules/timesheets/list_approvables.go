package timesheets

import (
	"context"

	"github.com/nlm/lucca-api/api"
)

type ListApprovablesRequest struct {
	api.Page
	LoadPayrollInfo   bool `json:"loadPayrollInfo"`
	IncludeOnlyOthers bool `json:"includeOnlyOthers"`
}

type ListApprovablesResponse api.ItemsList[ListApprovablesItem]

type ListApprovablesItem struct {
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

type Owner struct {
	Id int `json:"id"`
	// Picture   *string `json:"picture"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (ts *TimesheetsService) ListApprovables(ctx context.Context, req *ListApprovablesRequest) (*ListApprovablesResponse, error) {
	res := new(ListApprovablesResponse)
	err := ts.client.Get(ctx, "/timmi-timesheet/api/timesheets/approvables", req, res)
	return res, err
}
