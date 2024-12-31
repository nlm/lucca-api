package worklocations

import (
	"context"

	"github.com/nlm/lucca-api/api"
)

type ListUsersWithCurrentLocationRequest struct {
	DepartmentId     []int    `json:"departmentId,omitempty"`
	EstablishmentId  int      `json:"establishmentId"`
	ExcludePrincipal bool     `json:"excludePrincipal"`
	Limit            *int     `json:"limit,omitempty"`
	Sort             []string `json:"sort,omitempty"` // sort=departmentHierarchy,lastName,firstName
	RelativeUsers    *string  `json:"relativeUsers"`  // relativeUsers=CollaboratorsLevel3
}

type UsersWithWithCurrentLocation struct {
	CurrentWorkLocationId int          `json:"currentWorkLocationId"`
	CurrentAreaId         int          `json:"currentAreaId"`
	Id                    int          `json:"id"`
	FirstName             string       `json:"firstName"`
	LastName              string       `json:"lastName"`
	EstablishmentId       *int         `json:"establishmentId,omitempty"`
	DepartmentId          *int         `json:"departmentId,omitempty"`
	DepartmentHierarchy   []Department `json:"departmentHierarchy"`
}

type Department struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Hierarchy string `json:"hierarchy"`
}

type ListUsersWithCurrentLocationResponse api.ItemsList[UsersWithWithCurrentLocation]

func (ws *WorkLocationsService) ListUsersWithCurrentLocation(ctx context.Context, req *ListUsersWithCurrentLocationRequest) (*ListUsersWithCurrentLocationResponse, error) {
	return api.Get[ListUsersWithCurrentLocationRequest, ListUsersWithCurrentLocationResponse](ws.client, ctx, "/work-locations/api/schedule/usersWithCurrentLocation", req)
}
