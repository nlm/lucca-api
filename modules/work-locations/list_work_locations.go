package worklocations

import (
	"context"

	"github.com/nlm/lucca-api/api"
)

type ListWorkLocationsRequest struct{}

type ListWorkLocationsResponse []WorkLocation

type WorkLocation struct {
	Id                      int     `json:"id"`
	Name                    string  `json:"name"`
	Type                    string  `json:"type"`
	Icon                    *string `json:"icon,omitempty"`
	Color                   string  `json:"color,omitempty"`
	Office                  any     `json:"office,omitempty"`
	IsActive                bool    `json:"isActive"`
	IsOverbookingAuthorized bool    `json:"isOverbookingAuthorized"`
	CommentsEnabled         bool    `json:"commentsEnabled"`
	EstablishmentIds        []int   `json:"establishmentIds,omitempty"`
}

func (ws *WorkLocationsService) ListWorkLocations(ctx context.Context, req *ListWorkLocationsRequest) (*ListWorkLocationsResponse, error) {
	return api.Get[ListWorkLocationsRequest, ListWorkLocationsResponse](ws.client, ctx, "/work-locations/api/work-locations", req)
}
