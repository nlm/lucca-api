package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"

	"github.com/nlm/lucca-api/api"
	"github.com/nlm/lucca-api/internal/table"
	"github.com/nlm/lucca-api/modules/leaves"
)

const nameApproveLeaveRequest = "approve-leave-request"

func init() {
	RegisterCommand(nameApproveLeaveRequest, ApproveLeaveRequest)
}

func ApproveLeaveRequest(ctx context.Context, client *api.Client, args []string) error {
	fset := flag.NewFlagSet(nameApproveLeaveRequest, flag.ExitOnError)
	if err := fset.Parse(args); err != nil {
		return err
	}
	args = fset.Args()
	if len(args) != 1 {
		return fmt.Errorf("one arg only")
	}
	leaveRequestId, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	leavesService := leaves.New(client)
	res, err := leavesService.ApproveLeave(ctx, &leaves.ApproveLeaveRequest{
		Id: leaveRequestId,
	})
	if err != nil {
		return err
	}

	tbl := table.New("Leave Id")
	tbl.AddRow(res.Data.Id)
	tbl.Print()
	return nil
}
