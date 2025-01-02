package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"fyne.io/systray"
	"github.com/nlm/lucca-api/api"

	"github.com/nlm/lucca-api/cmd/lucca/config"
	"github.com/nlm/lucca-api/cmd/menubar/images"
	"github.com/nlm/lucca-api/modules/identity"
	"github.com/nlm/lucca-api/modules/leaves"
	"github.com/nlm/lucca-api/modules/timesheets"
)

type Tray struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (s *Tray) Run(ctx context.Context, fn func(ctx context.Context, tray *Tray)) {
	s.ctx, s.cancel = context.WithCancel(ctx)
	systray.Run(func() {
		systray.SetIcon(images.LargeBlueDiamond)
		systray.SetTitle("Lucca Initializing")
		fn(ctx, s)
	}, nil)
}

func (s *Tray) SetError(err error) {
	log.Println(err)
	systray.SetIcon(images.Warning)
	systray.SetTooltip(err.Error())
}

func (s *Tray) SetOK() {
	systray.SetTemplateIcon(images.LargeBlueDiamond, images.LargeBlueDiamond)
	systray.SetTitle("")
}

func (s *Tray) SetTODO() {
	systray.SetIcon(images.LargeOrangeDiamond)
	systray.SetTitle("Lucca TODO")
}

func (s Tray) PlayAnimation(ctx context.Context, frames [][]byte, interval time.Duration) {
	frameIdx := 0
	tick := time.Tick(interval)
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick:
			systray.SetIcon(frames[frameIdx])
			frameIdx = (frameIdx + 1) % len(frames)
		}
	}
}

func (s Tray) Context() context.Context {
	return s.ctx
}

func (s Tray) Cancel() {
	s.cancel()
}

func main() {
	var flagConfigFile = flag.String("config", "config.toml", "config file")
	flag.Parse()

	conf, err := config.ReadConfig(*flagConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	tray := Tray{}
	tray.Run(context.TODO(), onReady(conf))
	select {}
}

func onReady(conf *config.Config) func(context.Context, *Tray) {
	return func(ctx context.Context, tray *Tray) {
		log.Println("starting...")

		transport := http.DefaultTransport
		transport = api.NewHeadersRoundTripper(transport, map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
		})
		transport = api.NewThrottleRoundTripper(transport, 5, 5)

		client := api.NewClient(api.ClientOptions{
			Host:       conf.Host,
			AuthCookie: conf.AuthCookie,
			Transport:  transport,
			// Cache:      false,
		})

		tick := time.Tick(245 * time.Second)
		for {
			log.Println("loop")
			evLsn := UpdateMenu(ctx, conf, client, tray)
			selectCases := make([]reflect.SelectCase, 0)
			for ev := range evLsn {
				selectCases = append(selectCases, reflect.SelectCase{
					Chan: reflect.ValueOf(ev),
					Dir:  reflect.SelectRecv,
				})
			}
			selectCases = append(selectCases, reflect.SelectCase{
				Chan: reflect.ValueOf(tick),
				Dir:  reflect.SelectRecv,
			})
			if chosen, _, ok := reflect.Select(selectCases); ok {
				c := selectCases[chosen].Chan.Interface()
				switch v := c.(type) {
				case <-chan struct{}:
					if fn := evLsn[v]; fn != nil {
						fn()
					}
				case <-chan time.Time:
					log.Println("timeout")
				}
			}
		}
	}
}

func UpdateMenu(ctx context.Context, conf *config.Config, client *api.Client, tray *Tray) map[<-chan struct{}]func() {
	// setup menu
	eventListeners := make(map[<-chan struct{}]func(), 0)
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	systray.ResetMenu()
	todo := false
	defer func() {
		systray.AddSeparator()
		eventListeners[systray.AddMenuItem("Quit", "").ClickedCh] = func() {
			systray.Quit()
		}
	}()

	// detect my identity
	identityService := identity.New(client)
	me, err := identityService.GetPrincipal(tray.Context(), &identity.GetPrincipalRequest{})
	if err != nil {
		tray.SetError(err)
		return eventListeners
	}
	mMe := systray.AddMenuItem(fmt.Sprintf("%s (%d)", me.Fullname, me.Id), fmt.Sprintf("Session Expire: %s", me.SessionExpiresAt))
	mMe.Disable()
	systray.AddSeparator()

	// approvable leaves
	leavesService := leaves.New(client)
	leavesRes, err := leavesService.ListApprovables(ctx, &leaves.ListApprovablesRequest{})
	if err != nil {
		tray.SetError(err)
		return eventListeners
	}
	leavesToApprove := 0
	for _, res := range *leavesRes {
		if res.ExpectedApprover != nil && res.ExpectedApprover.Id == me.Id {
			leavesToApprove++
		}
	}
	mLeaves := systray.AddMenuItem(fmt.Sprintf("Leaves to approve: %d", leavesToApprove), "")
	eventListeners[mLeaves.ClickedCh] = func() {
		err := openURL(fmt.Sprintf("https://%s/timmi-absences/approvals", conf.Host))
		if err != nil {
			log.Println(err)
		}
	}
	if leavesToApprove > 0 {
		todo = true
	}

	// timesheets
	timesheetsService := timesheets.New(client)
	timesheetsRes, err := timesheetsService.ListApprovables(ctx, &timesheets.ListApprovablesRequest{})
	if err != nil {
		tray.SetError(err)
		return eventListeners
	}
	timesheetsToApprove := len(timesheetsRes.Items)
	mTimesheets := systray.AddMenuItem(fmt.Sprintf("Timesheets to approve: %d", timesheetsToApprove), "")
	eventListeners[mTimesheets.ClickedCh] = func() {
		err := openURL(fmt.Sprintf("https://%s/timmi-timesheet/approval/me", conf.Host))
		if err != nil {
			log.Println(err)
		}
	}
	if timesheetsToApprove > 0 {
		todo = true
	}

	// finish
	if todo {
		tray.SetTODO()
	} else {
		tray.SetOK()
	}
	return eventListeners
}
