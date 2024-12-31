package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/nlm/lucca-api/api"
	"github.com/nlm/lucca-api/internal/table"
	"github.com/nlm/lucca-api/modules/directory"
	"github.com/nlm/lucca-api/modules/identity"
)

const nameListDiploma = "list-diploma"

func init() {
	RegisterCommand(nameListDiploma, ListDiploma)
}

func ListDiploma(ctx context.Context, client *api.Client, args []string) error {
	fset := flag.NewFlagSet(nameListDiploma, flag.ExitOnError)
	if err := fset.Parse(args); err != nil {
		return err
	}
	args = fset.Args()

	identityService := identity.New(client)
	directoryService := directory.New(client)

	principal, err := identityService.GetPrincipal(ctx, &identity.GetPrincipalRequest{})
	if err != nil {
		return err
	}

	diploma, err := directoryService.ListDiploma(ctx, &directory.ListDiplomaRequest{
		OwnerId: principal.Id,
	})
	if err != nil {
		return err
	}

	if len(diploma.Items) == 0 {
		fmt.Println("no diploma found")
		return nil
	}

	tbl := table.New("Diploma", "Date", "Level")
	for _, d := range diploma.Items {
		tbl.AddRow(d.Id, d.Date, d.Level)
	}
	fmt.Println(Titlef("Diploma for %v", principal.Fullname))
	tbl.Print()
	return nil
}
