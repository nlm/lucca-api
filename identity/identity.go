package identity

import (
	"github.com/nlm/lucca-api/api"
)

type IdentityService struct {
	client *api.Client
}

func New(client *api.Client) *IdentityService {
	return &IdentityService{
		client: client,
	}
}
