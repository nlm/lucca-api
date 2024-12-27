package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/nlm/lucca-api/api"
	"github.com/nlm/lucca-api/internal/table"
	"github.com/nlm/lucca-api/modules/leaves"
)

func init() {
	RegisterCommand("list-leave-requests", ListLeaveRequests)
}

var (
	flagsetLeaveRequests = flag.NewFlagSet("leave-requests", flag.ExitOnError)
)

func ListLeaveRequests(ctx context.Context, client *api.Client, args []string) error {
	flagsetLeaveRequests.Parse(args)

	leavesService := leaves.New(client)

	approvables, err := leavesService.ListApprovables(ctx, &leaves.ListApprovablesRequest{
		// Page: api.Page{Page: 1, PageSize: 50},
	})
	if err != nil {
		return fmt.Errorf("error listing approvables: %w", err)
	}

	for _, item := range *approvables {
		tbl := table.New("Owner", "Approver", "StartsOn", "StartsAm", "EndsOn", "EndsAm")
		tbl.AddRow(item.Owner.Name, item.ExpectedApprover.Name, item.StartsOn, item.StartsAm, item.EndsOn, item.EndsAm)
		tbl.Print()
	}
	return nil
}
