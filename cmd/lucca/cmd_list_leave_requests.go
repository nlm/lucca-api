package main

import (
	"flag"
	"fmt"

	"github.com/nlm/lucca-api/internal/table"
	"github.com/nlm/lucca-api/modules/identity"
	"github.com/nlm/lucca-api/modules/leaves"
)

func init() {
	RegisterCommand("list-leave-requests", ListLeaveRequests)
}

var (
	flagsetLeaveRequests = flag.NewFlagSet("list-leave-requests", flag.ExitOnError)
)

func ListLeaveRequests(ctx Context, args []string) error {
	flagAll := flagsetLeaveRequests.Bool("all", false, "show all approvable requests")
	flagsetLeaveRequests.Parse(args)

	client := ctx.Client()
	identityService := identity.New(client)
	leavesService := leaves.New(client)

	me, err := identityService.GetPrincipal(ctx, &identity.GetPrincipalRequest{})
	if err != nil {
		return fmt.Errorf("error getting principal info: %w", err)
	}

	approvables, err := leavesService.ListApprovables(ctx, &leaves.ListApprovablesRequest{
		// Page: api.Page{Page: 1, PageSize: 50},
	})
	if err != nil {
		return fmt.Errorf("error listing approvables: %w", err)
	}

	tbl := table.New("Owner", "Approver", "StartsOn", "StartsAm", "EndsOn", "EndsAm")
	rowCount := 0
	for _, item := range *approvables {
		if !*flagAll && item.ExpectedApprover.Id != me.Id {
			continue
		}
		tbl.AddRow(item.Owner.Name, item.ExpectedApprover.Name, item.StartsOn, item.StartsAm, item.EndsOn, item.EndsAm)
		rowCount++
	}

	if rowCount <= 0 {
		fmt.Println(Info("no leave requests to validate"))
		return nil
	}

	return nil
}
