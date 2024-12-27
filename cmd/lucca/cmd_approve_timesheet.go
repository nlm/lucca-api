package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"

	"github.com/nlm/lucca-api/api"
	"github.com/nlm/lucca-api/internal/table"
	"github.com/nlm/lucca-api/modules/timesheets"
)

func init() {
	RegisterCommand("approve-timesheet", ApproveTimesheet)
}

var (
	flagsetApproveTimesheet = flag.NewFlagSet("approve-timesheet", flag.ExitOnError)
)

func ApproveTimesheet(ctx context.Context, client *api.Client, args []string) error {
	flagsetApproveTimesheet.Parse(args)
	args = flagsetApproveTimesheet.Args()

	if len(args) != 1 {
		return fmt.Errorf("you must provide exactly one id")
	}
	timesheetId, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid timesheet id '%v': %v", args[0], err)
	}

	timesheetsService := timesheets.New(client)
	res, err := timesheetsService.ApproveTimesheet(ctx, &timesheets.ApproveTimesheetRequest{
		Id: timesheetId,
	})
	if err != nil {
		return fmt.Errorf("unable to approve timesheet: %w", err)
	}
	if res == nil {
		return fmt.Errorf("nil response")
	}
	fmt.Println(Titlef("Approval of timesheet %d", timesheetId))
	tbl := table.New("Id", "Status")
	for _, item := range *res {
		tbl.AddRow(item.Id, item.Status)
	}
	tbl.Print()

	return nil
}
