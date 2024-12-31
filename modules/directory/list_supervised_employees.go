package directory

import (
	"context"
	"fmt"

	"github.com/nlm/lucca-api/api"
)

type ListSupervisedEmployeesRequest struct {
	Id int `json:"-"` // Your employee id
}

type ListSupervisedEmployeesResponse []Employee

type Employee struct {
	Id        int       `json:"id,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Status    string    `json:"status,omitempty"`
	JobTitle  string    `json:"jobTitle,omitempty"`
	EndsOn    *api.Date `json:"endsOn,omitempty"`
}

func (ds *DirectoryService) ListSupervisedEmployees(ctx context.Context, req *ListSupervisedEmployeesRequest) (*ListSupervisedEmployeesResponse, error) {
	return api.Get[ListSupervisedEmployeesRequest, ListSupervisedEmployeesResponse](ds.client, ctx, fmt.Sprintf("/directory/api/employees/%d/supervised-employees/", req.Id), req)
}
