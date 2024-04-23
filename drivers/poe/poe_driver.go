package poe

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

var dAttributeOfReload = "M11.2 21a1 1 0 0 1-.914-.594L7.351 13.8H5.6a2.615 2.615 0 0 1-1.838-.761A2.59 2.59 0 0 1 3 11.2V5.6c0-.685.277-1.354.762-1.838A2.615 2.615 0 0 1 5.6 3h11.707a2.6 2.6 0 0 1 2.561 2.21l1.103 7.198c.057.369.031.755-.072 1.115a2.61 2.61 0 0 1-1.432 1.648 2.615 2.615 0 0 1-1.096.228H14.6v2.2c0 .907-.354 1.761-.994 2.403A3.38 3.38 0 0 1 11.2 21ZM9 12.587l2.792 6.282a1.401 1.401 0 0 0 .807-1.27v-3.2a1 1 0 0 1 1-1h4.783c.08-.009.184-.017.264-.052a.607.607 0 0 0 .33-.38.617.617 0 0 0 .018-.256l-1.104-7.2A.598.598 0 0 0 17.298 5H9v7.587ZM5.6 5a.6.6 0 0 0-.6.6v5.6a.599.599 0 0 0 .6.6H7V5H5.6Z"
var reloadSelector = `path[d="` + dAttributeOfReload + `"]`

func Prompt(question string) (string, error) {
	answer, err := setupExecutionContext(question)
	if err != nil {
		return "", err
	}
	return answer, nil
}

func setupExecutionContext(question string) (string, error) {

	verbose := flag.Bool("v", false, "verbose")
	urlstr := flag.String("url", "ws://127.0.0.1:9222", "devtools url")
	nav := flag.String("nav", "https://poe.com/Claude-3-Haiku", "nav")
	// nav := flag.String("nav", "https://poe.com", "nav")

	d := flag.Duration("d", 1*time.Second, "wait duration")
	flag.Parse()

	answer, err := run(context.Background(), *verbose, *urlstr, *nav, *d, question)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return answer, err
}

func run(ctx context.Context, verbose bool, urlstr, nav string, d time.Duration, question string) (string, error) {
	if urlstr == "" {
		return "", errors.New("invalid remote devtools url")
	}
	// create allocator context for use with creating a browser context later
	allocatorContext, _ := chromedp.NewRemoteAllocator(context.Background(), urlstr)

	var opts []chromedp.ContextOption
	if verbose {
		opts = append(opts, chromedp.WithDebugf(log.Printf))
	}

	ctx, _ = chromedp.NewContext(allocatorContext, opts...)

	var text string

	if err := chromedp.Run(ctx,
		chromedp.Navigate(nav),
		chromedp.Sleep(d),
		chromedp.SendKeys(`textarea[placeholder="Fale com Claude-3-Haiku"]`, question+"\n"),
		// chromedp.SendKeys(`textarea[placeholder="Iniciar um novo chat"]`, question+"\n"),

		chromedp.WaitVisible(reloadSelector),
		chromedp.InnerHTML("main", &text, chromedp.ByQuery),
	); err != nil {
		return "", fmt.Errorf("failed getting body of %s: %v", nav, err)
	}
	fmt.Println(text)
	// writeFile("answer.txt", text)

	return text, nil
}
