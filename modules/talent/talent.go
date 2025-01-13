package talent

import (
	"context"

	"github.com/nlm/lucca-api/api"
)

type TalentService struct {
	client *api.Client
}

func New(client *api.Client) *TalentService {
	return &TalentService{
		client: client,
	}
}

type ListGoalsRequest struct {
	Fields          []string     `json:"fields,omitempty"`
	PeriodStartDate api.DateTime `json:"period.startDate"`
	PeriodEndDate   api.DateTime `json:"period.endDate"`
	PeriodType      int          `json:"period.periodType"`
	OwnerID         int          `json:"ownerId"`
	// SupervisedLevel string `json:"supervisedLevel"` // FIXME, currently "undefined"
	// Paging string `json:"paging"` // FIXME 0,1000
}

type ListGoalsResponse api.HeaderData[api.Header, ListGoalsData]

type ListGoalsData api.ItemsList[ListGoalsItem]

type ListGoalsItem struct {
	IsSnapshot        bool    `json:"isSnapshot"`
	Owner             Owner   `json:"owner"`
	OwnerId           int     `json:"ownerId"`
	Period            Period  `json:"period,omitempty"`
	TotalGoalsCount   int     `json:"totalGoalsCount"`
	AverageCompletion float64 `json:"averageCompletion"`
	Goals             []Goal  `json:"goals"`
}

type Owner struct {
	Id        int      `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Name      string   `json:"name"`
	Picture   *Picture `json:"picture"`
	ManagerId int      `json:"managerId"`
}

type Picture struct {
	Href string `json:"href"`
}

type Goal struct {
	Type            int    `json:"type"` // 0 for individual, 1 for team
	ParentId        string `json:"parentId"`
	Parent          any    `json:"parent"` // FIXME
	AutoCalculation bool   `json:"autoCalculation"`
	// Children        []Goal `json:"children"`
	// Ancestors // FIXME
	OwnerId         int    `json:"ownerId"`
	Owner           Owner  `json:"owner"`
	CompletionLevel int    `json:"completionLevel"`
	IsConfirmed     bool   `json:"isConfirmed"`
	Period          Period `json:"period"`
	Status          int    `json:"status"`
	Weight          int    `json:"weight"`
	Id              string `json:"id"`
	Name            string `json:"name"`
}

func GoalTypeText(id int) string {
	switch id {
	case 0:
		return "individual"
	case 1:
		return "team"
	default:
		return ""
	}
}

func (ts *TalentService) ListGoals(ctx context.Context, req *ListGoalsRequest) (*ListGoalsResponse, error) {
	return api.Get[ListGoalsRequest, ListGoalsResponse](ts.client, ctx, "/popleetalent/services/goalsummaries", req)
}

type ListPeriodsRequest struct {
	Year    int `json:"year,omitempty"`
	OwnerId int `json:"ownerId,omitempty"`
}

type ListPeriodsResponse api.HeaderData[api.Header, ListPeriodsData]

type ListPeriodsData api.ItemsList[PeriodItem]

type PeriodItem struct {
	Period                 Period   `json:"period"`
	ShortName              string   `json:"shortName"`
	AverageCompletionLevel *float64 `json:"averageCompletionLevel"`
	GoalsCount             int      `json:"goalsCount"`
	GoalType               any      `json:"goalType"`
}

type Period struct {
	Name           string        `json:"name"`
	ShortName      string        `json:"shortName"`
	NavigationName string        `json:"navigationName"`
	Year           int           `json:"year"`
	StartDate      *api.DateTime `json:"startDate"`
	PeriodType     int           `json:"periodType"`
	EndDate        *api.DateTime `json:"endDate"`
}

func (ts *TalentService) ListPeriods(ctx context.Context, req *ListPeriodsRequest) (*ListPeriodsResponse, error) {
	return api.Get[ListPeriodsRequest, ListPeriodsResponse](ts.client, ctx, "/popleetalent/api/goals/periodNavigationsInfos", req)
	//?year=2025&ownerId=29")
}
