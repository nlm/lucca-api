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
	"os/signal"
	"slices"
	"sync"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/nlm/lucca-api/api"
)

var (
	flagMock          = flag.Bool("mock", false, "use mock data")
	flagConfigFile    = flag.String("config", "config.toml", "config file")
	flagDebug         = flag.Bool("debug", false, "debug mode")
	flagNoCache       = flag.Bool("no-cache", false, "disable cache")
	flagThrottleQps   = flag.Int("throttle-qps", 5, "requests per second")
	flagThrottleBurst = flag.Int("throttle-burst", 5, "requests burst")
)

type Config struct {
	Host       string `toml:"host"`
	AuthCookie string `toml:"auth-cookie"`
}

type CommandFunc func(ctx context.Context, client *api.Client, args []string) error

var cliCommands = make(map[string]CommandFunc)

// RegisterCommand registers a command in the CLI handler.
func RegisterCommand(name string, fn CommandFunc, aliases ...string) {
	if _, ok := cliCommands[name]; ok {
		panic(fmt.Sprintln("command already registered:", name))
	}
	cliCommands[name] = fn
	for _, alias := range aliases {
		if _, ok := cliCommands[alias]; ok {
			panic(fmt.Sprintln("alias already registered:", alias))
		}
		cliCommands[alias] = fn
	}
}

// help displays help message.
func help(err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("available commands: ")
	fmt.Println()
	for _, v := range slices.Sorted(maps.Keys(cliCommands)) {
		fmt.Println(" ", v)
	}
	fmt.Println()
	flag.Usage()
}

var onceContext = sync.Once{}
var cliContext context.Context

// CLIContext creates a new context that gracefully handles signals.
func CLIContext() context.Context {
	// setup signal handling
	onceContext.Do(func() {
		var cancel context.CancelFunc
		cliContext, cancel = context.WithCancel(context.Background())
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			for range c {
				cancel()
			}
		}()
	})
	return cliContext
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

	// setup api client
	var transport = http.DefaultTransport
	if *flagMock {
		transport = api.NewMockRoundTripper(true, nil)
	}
	transport = api.NewHeadersRoundTripper(transport, map[string]string{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
	})
	transport = api.NewThrottleRoundTripper(transport, *flagThrottleQps, *flagThrottleBurst)
	clientOptions := api.ClientOptions{
		Host:       config.Host,
		AuthCookie: config.AuthCookie,
		Transport:  transport,
		Cache:      !*flagNoCache,
	}
	if *flagDebug {
		clientOptions.Logger = log.Default()
	}
	client := api.NewClient(clientOptions)

	// parse command line and execute
	args := flag.Args()
	if len(args) == 0 {
		help(nil)
		os.Exit(1)
	} else if cmd, ok := cliCommands[args[0]]; ok {
		err := cmd(CLIContext(), client, args[1:])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		help(fmt.Errorf("error: unknown command: %v", os.Args[1]))
		os.Exit(1)
	}
}

func PrintJson(v any) {
	b, _ := json.MarshalIndent(v, "", "  ")
	log.Println(string(b))
}
