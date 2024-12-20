package timesheets

import (
	"github.com/nlm/lucca-api/api"
)

type TimesheetsService struct {
	client *api.Client
}

func New(client *api.Client) *TimesheetsService {
	return &TimesheetsService{
		client: client,
	}
}
