package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/nlm/lucca-api/api"
	"github.com/nlm/lucca-api/internal/table"
	worklocations "github.com/nlm/lucca-api/modules/work-locations"
)

func init() {
	// RegisterCommand("list-working-today", ListWorkingToday)
}

var (
	flagsetListWorkingToday = flag.NewFlagSet("list-working-today", flag.ExitOnError)
)

func ListWorkingToday(ctx context.Context, client *api.Client, args []string) error {
	flagsetListWorkingToday.Parse(args)

	worklocationsService := worklocations.New(client)

	// resolve locations
	// wlres, err := worklocationsService.ListWorkLocations(ctx, &worklocations.ListWorkLocationsRequest{})
	// if err != nil {
	// 	return fmt.Errorf("error retrieving work locations: %w", err)
	// }
	locations := make(map[int]worklocations.WorkLocation)
	// for _, loc := range *wlres {
	// 	locations[loc.Id] = loc
	// }

	// check working users
	res, err := worklocationsService.ListUsersWithCurrentLocation(ctx, &worklocations.ListUsersWithCurrentLocationRequest{
		ExcludePrincipal: true,
		Limit:            api.Ptr(50),
		Sort:             []string{"departmentHierarchy", "lastName", "firstName"},
		RelativeUsers:    api.Ptr("CollaboratorsLevel3"),
	})
	if err != nil {
		return err
	}

	fmt.Println(Titlef("People working today"))
	tbl := table.New("Name")
	for _, item := range res.Items {
		tbl.AddRow(
			fmt.Sprint(item.FirstName, " ", item.LastName),
			locations[item.CurrentWorkLocationId].Name,
		)
	}
	tbl.Print()
	return nil
}
