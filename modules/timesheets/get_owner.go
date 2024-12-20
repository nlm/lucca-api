package timesheets

import (
	"context"
	"fmt"

	"github.com/nlm/lucca-api/api"
)

type GetOwnerRequest struct {
	Id int `json:"-"`
}

type GetOwnerResponse struct {
	EmployeeNumber    string        `json:"employeeNumber,omitempty"`
	DepartmentName    string        `json:"departmentName"`
	LegalEntityName   *string       `json:"legalEntityName"`
	ContractStartDate *api.DateTime `json:"dtContractStart,omitempty"`
	ContractEndDate   *api.DateTime `json:"dtContractEnd,omitempty"`
	Picture           any           `json:"picture,omitempty"`
	Id                int           `json:"id,omitempty"`
	FirstName         string        `json:"firstName,omitempty"`
	LastName          string        `json:"lastName,omitempty"`
	DepartmentId      int           `json:"departmentId,omitempty"`
	EstablishmentId   int           `json:"establishmentId,omitempty"`
}

func (ts *TimesheetsService) GetOwner(ctx context.Context, req *GetOwnerRequest) (*GetOwnerResponse, error) {
	res := new(GetOwnerResponse)
	err := ts.client.Get(ctx, fmt.Sprintf("/timmi-timesheet/api/timesheets/owners/%d", req.Id), req, res)
	return res, err
}
