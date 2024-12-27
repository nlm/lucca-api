package worklocations

import (
	"context"

	"github.com/nlm/lucca-api/api"
)

type ListUserLocationsRequest struct {
	Start   api.Date `json:"start"`
	End     api.Date `json:"end"`
	UserIds []int    `json:"userIds"`
}

type ListUserLocationsResponse []UserLocation

type UserLocation struct {
	Id             int       `json:"id"`
	OwnerId        int       `json:"ownerId"`
	WorkLocationId int       `json:"workLocationId"`
	AreaId         int       `json:"areaId"`
	Position       string    `json:"position"`
	Date           *api.Date `json:"date"`
	AssigneeId     *int      `json:"assigneeId"`
	Origin         string    `json:"origin"`
	Status         string    `json:"status"`
	IsUrgent       *bool     `json:"isUrgent"`
}

func (ws *WorkLocationsService) ListUserLocations(ctx context.Context, req *ListUserLocationsRequest) (*ListUserLocationsResponse, error) {
	// res := new(ListUserLocationsResponse)
	// err := ws.client.Get(ctx, "/work-locations/api/user-locations", req, res)
	// return res, err
	return api.Get[ListUserLocationsRequest, ListUserLocationsResponse](ws.client, ctx, "/work-locations/api/user-locations", req)
}
