package users

import (
	"context"
	"time"

	"github.com/nlm/lucca-api/api"
)

type ListUsersRequest struct {
	ContractStartDate time.Time `json:"dtContractStart,omitempty"`
	ContractStartEnd  time.Time `json:"dtContractEnd,omitempty"`
	Fields            []string  `json:"fields,omitempty"`
	Id                []int     `json:"id,omitempty"`
	Login             string    `json:"login,omitempty"`
	Mail              string    `json:"mail,omitempty"`
	ModifiedAt        time.Time `json:"modifiedAt,omitempty"`
	Paging            *api.Page `json:"paging,omitempty"`
}

type ListUsersResponse struct {
	Header *Header             `json:"header,omitempty"`
	Data   api.ItemsList[User] `json:"data,omitempty"`
}

func (us *UsersService) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	// res := new(ListUsersResponse)
	// err := us.client.Get(ctx, "/api/v3/users", req, res)
	// return res, err
	return api.Get[ListUsersRequest, ListUsersResponse](us.client, ctx, "/api/v3/users", req)
}
