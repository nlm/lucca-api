package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/nlm/lucca-api/api"
	"github.com/nlm/lucca-api/timesheets"
)

var (
	flagConfigFile = flag.String("config", "config.toml", "config file")
)

type Config struct {
	Host       string `toml:"host"`
	AuthCookie string `toml:"auth-cookie"`
}

func main() {
	flag.Parse()

	// read config
	var config Config
	md, err := toml.DecodeFile(*flagConfigFile, &config)
	if len(md.Undecoded()) > 0 {
		log.Fatal("extra config keys:", md.Undecoded())
	}
	if err != nil {
		log.Fatal(err)
	}

	// Client
	client := api.NewClient(api.ClientOptions{
		Host:       config.Host,
		AuthCookie: config.AuthCookie,
	})

	// identityService := identity.New(client)
	// pri, err := identityService.GetPrincipal(context.Background(), &identity.GetPrincipalRequest{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(pri)

	timeSheetsService := timesheets.New(client)
	// res, err := timeSheetsService.ListApprovables(context.Background(), &timesheets.ListApprovablesRequest{
	// 	Page: api.Page{
	// 		Page:     1,
	// 		PageSize: 50,
	// 	},
	// 	LoadPayrollInfo:   true,
	// 	IncludeOnlyOthers: false,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client.HTTPClient().Transport = &api.MockRoundTripper{}
	res, err := timeSheetsService.GetTimesheetDetails(context.Background(), &timesheets.GetTimesheetDetailsRequest{
		OwnerId: 368,
		From:    api.NewDate(2024, 10, 28),
		Until:   api.NewDate(2024, 11, 3),
	})
	if err != nil {
		log.Fatal(err)
	}
	PrintJson(res)
}

func PrintJson(v any) {
	b, _ := json.MarshalIndent(v, "", "  ")
	log.Println(string(b))
}
