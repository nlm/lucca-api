package users

import (
	"context"
	"fmt"

	"github.com/nlm/lucca-api/api"
)

type GetUserRequest struct {
	Id     int      `json:"-"`
	Fields []string `json:"fields,omitempty"`
}

type GetUserResponse struct {
	Header *Header `json:"header,omitempty"`
	Data   User    `json:"data,omitempty"`
}

func (us *UsersService) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return api.Get[GetUserRequest, GetUserResponse](us.client, ctx, fmt.Sprintf("api/v3/users/%d", req.Id), req)
}
