package directory

import (
	"context"

	"github.com/nlm/lucca-api/api"
)

type DirectoryService struct {
	client *api.Client
}

func New(client *api.Client) *DirectoryService {
	return &DirectoryService{
		client: client,
	}
}

type ListDiplomaRequest struct {
	OwnerId int `json:"ownerId,omitempty"`
}

type ListDiplomaResponse api.ItemsList[Diploma]

type Diploma struct {
	Id      int           `json:"id"`
	OwnerId int           `json:"ownerId"`
	File    string        `json:"e_Fichier"`
	Date    *api.DateTime `json:"e_Date-d-obtention"`
	Level   int           `json:"e_Niveau_Diplome"`
}

func (ds *DirectoryService) ListDiploma(ctx context.Context, req *ListDiplomaRequest) (*ListDiplomaResponse, error) {
	return api.Get[ListDiplomaRequest, ListDiplomaResponse](ds.client, ctx, "/directory/api/custom-resources/e_diplome-et-certification", req)
}
