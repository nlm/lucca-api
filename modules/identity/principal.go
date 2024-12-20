package identity

import (
	"context"

	"github.com/nlm/lucca-api/api"
)

type GetPrincipalRequest struct{}

type Principal struct {
	Id                  int          `json:"id"`
	FirstName           string       `json:"firstName"`
	LastName            string       `json:"lastName"`
	Fullname            string       `json:"fullName"`
	CultureCode         string       `json:"cultureCode"`
	EstablishmentId     int          `json:"establishmentId"`
	Mail                string       `json:"mail"`
	SessionExpiresAt    api.DateTime `json:"sessionExpiresAt"`
	SessionExpiresAtUTC api.DateTime `json:"sessionExpiresAtUtc"`
}

func (s *IdentityService) GetPrincipal(ctx context.Context, req *GetPrincipalRequest) (*Principal, error) {
	res := new(Principal)
	err := s.client.Get(ctx, "/identity/api/principal", req, res)
	return res, err
}
