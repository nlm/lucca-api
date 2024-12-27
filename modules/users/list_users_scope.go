package users

import (
	"github.com/nlm/lucca-api/api"
)

type ListUsersScopeRequest struct {
	AppInstanceId int      `json:"appInstanceId"` // Idk what that means
	Operations    int      `json:"operations"`    // Idk what that means
	Fields        []string `json:"fields"`        // Fields to request (ex: id)
}

type ListUsersScopeResponse struct {
	Header *Header             `json:"header,omitempty"`
	Data   api.ItemsList[User] `json:"data,omitempty"`
}

// func (us *UsersService) ListUsersScope(ctx context.Context, req *ListUsersScopeRequest) (*ListUsersScopeResponse, error) {
// 	res := new(ListUsersScopeResponse)
// 	err := us.client.Get(ctx, "")
// }
