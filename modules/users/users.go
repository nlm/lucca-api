package users

import "github.com/nlm/lucca-api/api"

type UsersService struct {
	client *api.Client
}

func New(client *api.Client) *UsersService {
	return &UsersService{
		client: client,
	}
}

type Header struct {
	Generated *api.DateTime `json:"generated,omitempty"`
	Principal string        `json:"principal"`
	ProcessId int           `json:"processId"`
}

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
