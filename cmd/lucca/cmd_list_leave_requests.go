package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/nlm/lucca-api/api"
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

func ListLeaveRequests(ctx context.Context, client *api.Client, args []string) error {
	flagAll := flagsetLeaveRequests.Bool("all", false, "show all approvable requests")
	flagsetLeaveRequests.Parse(args)

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
	for _, item := range *approvables {
		if !*flagAll && item.ExpectedApprover.Id != me.Id {
			continue
		}
		tbl.AddRow(item.Owner.Name, item.ExpectedApprover.Name, item.StartsOn, item.StartsAm, item.EndsOn, item.EndsAm)
	}
	tbl.Print()
	return nil
}
