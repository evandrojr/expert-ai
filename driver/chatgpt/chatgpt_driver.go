package chatgpt

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/evandrojr/expert-ai/tool"
)

var chromedpUrl = `ws://127.0.0.1:9222`
var dAttributeOfWaitVisible = "M7 11L12 6L17 11M12 18V7"
var waitVisibleSelector = `path[d="` + dAttributeOfWaitVisible + `"]`
var url = "https://chat.openai.com/chat"
var sendKeys = `#prompt-textarea`
var innerHTML = `.prose`

func Prompt(question string) (string, error) {
	ctx, cancel := setupContext()
	defer cancel()

	answer, err := ask(ctx, true, url, 1*time.Second, question)
	if err != nil {
		return "", err
	}

	return tool.HTMLToText(answer), nil
}

func setupContext() (context.Context, context.CancelFunc) {
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), chromedpUrl)
	return ctx, cancel
}

func ask(ctx context.Context, verbose bool, nav string, d time.Duration, question string) (string, error) {

	var opts []chromedp.ContextOption

	if verbose {
		opts = append(opts, chromedp.WithDebugf(log.Printf))
	}

	ctx, _ = chromedp.NewContext(ctx, opts...)
	var answer string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(nav),
		chromedp.Sleep(d),
		chromedp.SendKeys(sendKeys, question+"\n"),
		chromedp.WaitVisible(waitVisibleSelector),
		chromedp.InnerHTML(innerHTML, &answer, chromedp.ByQuery),
	); err != nil {
		return "", fmt.Errorf("failed getting body of %s: %v", nav, err)
	}
	fmt.Println(answer)
	return answer, nil
}
