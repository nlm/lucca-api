package leaves

import (
	"context"

	"github.com/nlm/lucca-api/api"
)

type LeavesService struct {
	client *api.Client
}

func New(client *api.Client) *LeavesService {
	return &LeavesService{
		client: client,
	}
}

type ListApprovablesRequest struct {
}

type ListApprovablesResponse []LeaveApproval

type LeaveApproval struct {
	Id                             int            `json:"id"`
	ExpectedApprover               *Person        `json:"expectedApprover,omitempty"`
	Owner                          *Person        `json:"owner,omitempty"`
	IsCancellationExpectedApproval bool           `json:"isCancellationExpectedApproval"`
	StartsOn                       *api.DateTime  `json:"startsOn,omitempty"`
	StartsAm                       bool           `json:"startsAm"`
	EndsOn                         *api.DateTime  `json:"endsOn,omitempty"`
	EndsAm                         bool           `json:"endsAm"`
	IsHalfDay                      bool           `json:"isHalfDay"`
	RelateToOneDay                 bool           `json:"relateToOneDay"`
	Unit                           int            `json:"unit"`
	TotalDuration                  float64        `json:"totalDuration"`
	LeaveType                      string         `json:"leaveType"`
	Accounts                       []LeaveAccount `json:"accounts,omitempty"`
	HasDelegation                  bool           `json:"hasDelegation"`
	HasTransfer                    bool           `json:"hasTransfer"`
	// Warnings []string ?
}

type LeaveAccount struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Unit        int    `json:"unit"`
	IsActive    bool   `json:"isActive"`
	IsRecurring bool   `json:"isRecurring"`
}

type Person struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	DepartmentId    int    `json:"departmentId"`
	EstablishmentId int    `json:"establishmentId"`
	CultureId       int    `json:"cultureId"`
}

func (svc *LeavesService) ListApprovables(ctx context.Context, req *ListApprovablesRequest) (*ListApprovablesResponse, error) {
	res := new(ListApprovablesResponse)
	err := svc.client.Get(ctx, "/timmi-absences/api/leaveRequests/v1.0/toApprove", req, res)
	return res, err
}
