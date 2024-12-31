package leaves

import (
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
