package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/nlm/lucca-api/internal/table"
	"github.com/nlm/lucca-api/modules/identity"
	"github.com/nlm/lucca-api/modules/talent"
)

func init() {
	RegisterCommand("list-goals", ListGoals)
}

var (
	flagsetListGoals = flag.NewFlagSet("list-goals", flag.ExitOnError)
)

func ListGoals(ctx Context, args []string) error {
	flagsetListGoals.Parse(args)

	identityService := identity.New(ctx.Client())
	talentService := talent.New(ctx.Client())

	me, err := identityService.GetPrincipal(ctx, &identity.GetPrincipalRequest{})
	if err != nil {
		return fmt.Errorf("error getting principal: %w", err)
	}

	periods, err := talentService.ListPeriods(ctx, &talent.ListPeriodsRequest{
		Year:    time.Now().Year(),
		OwnerId: me.Id,
	})
	if err != nil {
		return fmt.Errorf("error listing periods: %w", err)
	}

	fmt.Println(Title(fmt.Sprint("Goals for ", me.Fullname)))
	tbl := table.New("Period", "Type", "Name", "Completion", "Weight", "Status")
	for _, period := range periods.Data.Items {
		goalsRes, err := talentService.ListGoals(ctx, &talent.ListGoalsRequest{
			// Fields:          []string{},
			PeriodStartDate: *period.Period.StartDate,
			PeriodEndDate:   *period.Period.EndDate,
			PeriodType:      2,
			OwnerID:         me.Id,
		})
		if err != nil {
			return fmt.Errorf("error listing goals: %w", err)
		}

		// fmt.Println(Title(fmt.Sprint("Goals for ", period.ShortName)))
		for _, item := range goalsRes.Data.Items {
			for _, goal := range item.Goals {
				tbl.AddRow(period.ShortName, talent.GoalTypeText(goal.Type), goal.Name, fmt.Sprint(goal.CompletionLevel, " %"), goal.Weight, goal.Status)
			}
		}
	}
	tbl.Print()
	return nil
}
