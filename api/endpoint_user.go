package api

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	Id                int    `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Url               string `json:"url,omitempty"`
	DisplayName       string `json:"displayName,omitempty"`
	ModifiedOn        string `json:"modifiedOn,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	FirstName         string `json:"firstName,omitempty"`
	Login             string `json:"login,omitempty"`
	Mail              string `json:"mail,omitempty"`
	ContractStartDate string `json:"dtContractStart,omitempty"`
	ContractEndDate   string `json:"dtContractEnd,omitempty"`
	BirthDate         string `json:"birthDate,omitempty"`
	EmployeeNumber    string `json:"employeeNumber,omitempty"`
	// ...
}

type ListUsersRequest struct {
	ContractStartDate time.Time `json:"dtContractStart,omitempty"`
	ContractStartEnd  time.Time `json:"dtContractEnd,omitempty"`
	Fields            []string  `json:"fields,omitempty"`
	Id                []int     `json:"id,omitempty"`
	Login             string    `json:"login,omitempty"`
	Mail              string    `json:"mail,omitempty"`
	ModifiedAt        time.Time `json:"modifiedAt,omitempty"`
	Paging            *Page     `json:"paging,omitempty"`
}

type ListUsersResponse DataList[User]

func (c *Client) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	res := &ListUsersResponse{}
	err := c.Get(ctx, "/users", req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetUserRequest struct {
	UserId int      `json:"-"`
	Fields []string `json:"fields,omitempty"`
}

type GetUserResponse DataItem[User]

func (c *Client) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	res := &GetUserResponse{}
	err := c.Get(ctx, fmt.Sprintf("/users/%d", req.UserId), req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
