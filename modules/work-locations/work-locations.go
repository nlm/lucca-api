package worklocations

import (
	"github.com/nlm/lucca-api/api"
)

type WorkLocationsService struct {
	client *api.Client
}

func New(client *api.Client) *WorkLocationsService {
	return &WorkLocationsService{
		client: client,
	}
}
