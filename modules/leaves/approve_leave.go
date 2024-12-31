package leaves

import (
	"context"
	"fmt"

	"github.com/nlm/lucca-api/api"
)

type ApproveLeaveRequest struct {
	Id int `json:"-"`
}

type Header struct {
	// Generated *api.DateTime
	// Principal *Principal
}

type Leave struct {
	Id int `json:"id"`
	// Name string `json:"name,omitempty"` // also useless as this is the id
	// URL  string `json:"url,omitempty"` // this is useless
}

type ApproveLeaveResponse api.HeaderData[struct{}, Leave]

func (ls *LeavesService) ApproveLeave(ctx context.Context, req *ApproveLeaveRequest) (*ApproveLeaveResponse, error) {
	return api.Post[ApproveLeaveRequest, ApproveLeaveResponse](ls.client, ctx, fmt.Sprintf("/timmi-absences/api/leaveRequests/v1.0/%d/approvals", req.Id), req)
}
