package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"regexp"
	"strconv"
	"time"
)

func main() {

	var sessions_total int
	var sessions_in_use int

	var config = ReadConfig()
	c := influxDBClient(config)

	// Note to show how to enable printf debugging for chromedp traffic
	//ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))

	// Arris BGW-210 modem-router

	rootURL := `http://` + config.ArrisHost + `/`
	loginURL := `http://` + config.ArrisHost + `/cgi-bin/routerpasswd.ha`
	diagURL := `http://` + config.ArrisHost + `/cgi-bin/nattable.ha`
	selector := `//input[@name="password"]`
	button := `//input[@value="Continue"]`
	natContent := `#content-sub`
	pass := config.ArrisPass

	var result string

	for {
		fmt.Printf("Wakeup: query BGW-210...\n")
		start_time := time.Now()

		ctx, cancel := chromedp.NewContext(context.Background())
		defer cancel()

		sessions_total = 0
		sessions_in_use = 0

		login_task := chromedp.Tasks{
			chromedp.Navigate(rootURL),
			chromedp.Navigate(loginURL),
			chromedp.WaitVisible(selector),
			chromedp.SendKeys(selector, pass),
			chromedp.Click(button),
			chromedp.Sleep(2 * time.Second),
			chromedp.Navigate(diagURL),
			chromedp.Text(natContent, &result, chromedp.NodeVisible, chromedp.ByID),
		}

		err := chromedp.Run(ctx, login_task)
		if err != nil {
			log.Fatal(err)
		}

		re := regexp.MustCompile(`(?m)Total sessions available\s+(\d+)\s+Total sessions in use\s+(\d+)`)

		fmt.Printf("%s", time.Now().Truncate( 1 * time.Second ))

		if len(re.FindStringSubmatch(result)) == 3 {
			sessions_total, _ = strconv.Atoi(re.FindStringSubmatch(result)[1])
			sessions_in_use, _ = strconv.Atoi(re.FindStringSubmatch(result)[2])

			fmt.Printf(" [total: %d] [in use: %d]", sessions_total, sessions_in_use)
		}

		deliver_stats_to_influxdb(c, config, sessions_total, sessions_in_use)

		// Sleep for the delivery interval time minus the time it took to get and deliver the stats
		remaining_duration := config.DeliveryInterval - time.Since(start_time)
		fmt.Printf(" [sleep for %v]\n", remaining_duration.Round( 1 * time.Second ))

		// Don't keep thr browser running during sleep - this does mean we incur heavier startup at wake
		chromedp.Cancel(ctx)

		time.Sleep(remaining_duration)
	}
}
