package main

import (
	"context"
	"flag"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/nlm/lucca-api/api"
	"github.com/nlm/lucca-api/internal/table"
	"github.com/nlm/lucca-api/modules/timesheets"
)

func init() {
	RegisterCommand("list-timesheets", ListTimesheets)
}

var (
	flagsetTimesheets = flag.NewFlagSet("list-timesheets", flag.ExitOnError)
)

func axisFullName(axisSection timesheets.AxisSections) string {
	b := strings.Builder{}
	b.WriteString(axisSection.Name)
	for child := axisSection.Child; child != nil; child = child.Child {
		b.WriteString(" > ")
		b.WriteString(strings.TrimSpace(strings.Replace(child.Name, "\t", " ", -1)))
	}
	return b.String()
}

func ListTimesheets(ctx context.Context, client *api.Client, args []string) error {
	flagsetTimesheets.Parse(args)

	timesheetsService := timesheets.New(client)

	approvables, err := timesheetsService.ListApprovables(ctx, &timesheets.ListApprovablesRequest{
		Page: api.Page{Page: 1, PageSize: 50},
	})
	if err != nil {
		return fmt.Errorf("error listing approvables: %w", err)
	}

	for _, approvable := range approvables.Items {
		// get owner details
		owner, err := timesheetsService.GetOwner(ctx, &timesheets.GetOwnerRequest{Id: approvable.Owner.Id})
		if err != nil {
			return fmt.Errorf("error getting owner details: %w", err)
		}

		// get timesheet details
		tsDetails, err := timesheetsService.ListTimesheetDetails(ctx, &timesheets.ListTimesheetDetailsRequest{
			OwnerId: approvable.Owner.Id,
			From:    api.Date(approvable.StartsOn),
			Until:   api.Date(approvable.EndsOn),
		})
		if err != nil {
			return fmt.Errorf("error getting timesheets details: %w", err)
		}

		// render table
		sort.SliceStable(tsDetails.Items, func(i, j int) bool {
			a := tsDetails.Items[i]
			b := tsDetails.Items[j]
			if a.StartsAt == nil || b.StartsAt == nil {
				return false
			}
			return a.StartsAt.AsTime().Before(b.StartsAt.AsTime())
		})
		tbl := table.New("Date", "Day", "Name", "Duration")
		for _, item := range tsDetails.Items {
			startsAt := item.StartsAt.AsTime()
			addRow := func(description string) {
				tbl.AddRow(
					// fmt.Sprint(owner.FirstName, " ", owner.LastName),         // Owner
					startsAt.Format(time.DateOnly), // Date
					startsAt.Weekday(),             // Weekday
					description,                    // Description
					fmt.Sprint(item.Duration.Value, " ", item.Duration.Unit), // Duration
				)
			}
			if item.Imputation != nil {
				for _, axisSection := range item.Imputation.AxisSections {
					addRow(Work(axisFullName(axisSection)))
				}
			} else if item.Type == timesheets.ItemTypeHalfDayOff || item.Type == timesheets.ItemTypeLeave {
				if item.AccountName != nil {
					addRow(Leave(*item.AccountName))
				} else {
					addRow(Leave("half day off"))
				}
			} else {
				if item.AccountName != nil {
					addRow(*item.AccountName)
				} else {
					addRow("unknown")
				}
			}
		}
		fmt.Println(Titlef("Timesheet %d of %s %s", approvable.Id, owner.FirstName, owner.LastName))
		tbl.Print()
	}
	return nil
}
