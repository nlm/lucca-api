package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"maps"
	"net/http"
	"os"
	"slices"

	"github.com/BurntSushi/toml"
	"github.com/nlm/lucca-api/api"
)

var (
	flagMock       = flag.Bool("mock", false, "use mock data")
	flagConfigFile = flag.String("config", "config.toml", "config file")
	flagDebug      = flag.Bool("debug", false, "debug mode")
	flagNoCache    = flag.Bool("no-cache", false, "disable cache")
)

type Config struct {
	Host       string `toml:"host"`
	AuthCookie string `toml:"auth-cookie"`
}

var Commands = map[string]func(context.Context, *api.Client, []string) error{
	"list-timesheets": ListTimesheets,
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

	transport := api.NewHeadersRoundTripper(http.DefaultTransport, map[string]string{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
	})

	// Client
	client := api.NewClient(api.ClientOptions{
		Host:       config.Host,
		AuthCookie: config.AuthCookie,
		Transport:  transport,
		Cache:      !*flagNoCache,
	})

	ctx := context.Background()

	args := flag.Args()
	if cmd, ok := Commands[args[0]]; ok {
		err := cmd(ctx, client, args[1:])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("error: unknown command:", os.Args[1])
		fmt.Println()
		fmt.Println("available commands: ")
		fmt.Println()
		for _, v := range slices.Sorted(maps.Keys(Commands)) {
			fmt.Println(" ", v)
		}
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	}

}

func PrintJson(v any) {
	b, _ := json.MarshalIndent(v, "", "  ")
	log.Println(string(b))
}
