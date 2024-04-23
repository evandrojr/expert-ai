package chatgpt

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

var dAttributeOfReload = "M7 11L12 6L17 11M12 18V7"
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
	nav := flag.String("nav", "https://chat.openai.com/chat", "nav")
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
		chromedp.SendKeys(`#prompt-textarea`, question+"\n"),
		chromedp.WaitVisible(reloadSelector),
		chromedp.InnerHTML(".prose", &text, chromedp.ByQuery),
	); err != nil {
		return "", fmt.Errorf("failed getting body of %s: %v", nav, err)
	}
	fmt.Println(text)
	// writeFile("answer.txt", text)

	return text, nil
}
