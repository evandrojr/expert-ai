package poe

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/evandrojr/expert-ai/tools"
)

var chromedpUrl = `ws://127.0.0.1:9222`
var dAttributeOfWaitVisible = "M11.2 21a1 1 0 0 1-.914-.594L7.351 13.8H5.6a2.615 2.615 0 0 1-1.838-.761A2.59 2.59 0 0 1 3 11.2V5.6c0-.685.277-1.354.762-1.838A2.615 2.615 0 0 1 5.6 3h11.707a2.6 2.6 0 0 1 2.561 2.21l1.103 7.198c.057.369.031.755-.072 1.115a2.61 2.61 0 0 1-1.432 1.648 2.615 2.615 0 0 1-1.096.228H14.6v2.2c0 .907-.354 1.761-.994 2.403A3.38 3.38 0 0 1 11.2 21ZM9 12.587l2.792 6.282a1.401 1.401 0 0 0 .807-1.27v-3.2a1 1 0 0 1 1-1h4.783c.08-.009.184-.017.264-.052a.607.607 0 0 0 .33-.38.617.617 0 0 0 .018-.256l-1.104-7.2A.598.598 0 0 0 17.298 5H9v7.587ZM5.6 5a.6.6 0 0 0-.6.6v5.6a.599.599 0 0 0 .6.6H7V5H5.6Z"
var waitVisibleSelector = `path[d="` + dAttributeOfWaitVisible + `"]`
var url = "https://poe.com/Claude-3-Haiku"
var sendKeys = `textarea[placeholder="Fale com Claude-3-Haiku"]`
var innerHTML = `main`

func Prompt(question string) (string, error) {
	ctx, _ := setupContext()
	// defer cancel()

	answer, err := ask(ctx, true, url, 1*time.Second, question)
	if err != nil {
		return "", err
	}
	text := tools.HTMLToText(answer)
	text = tools.RemoveLine(text, 3)
	return text, nil
}

func setupContext() (context.Context, context.CancelFunc) {
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), chromedpUrl)
	return ctx, cancel
}

func ask(ctx context.Context, verbose bool, nav string, d time.Duration, question string) (string, error) {
	var opts []chromedp.ContextOption
	var answer string

	if verbose {
		opts = append(opts, chromedp.WithDebugf(log.Printf))
	}
	ctx, _ = chromedp.NewContext(ctx, opts...)

	if err := chromedp.Run(ctx,
		chromedp.Navigate(nav),
		chromedp.Sleep(d),
		chromedp.SendKeys(sendKeys, question+"\n"),
		// chromedp.SendKeys(`textarea[placeholder="Iniciar um novo chat"]`, question+"\n"),
		chromedp.WaitVisible(waitVisibleSelector),
		chromedp.InnerHTML(innerHTML, &answer, chromedp.ByQuery),
	); err != nil {
		return "", fmt.Errorf("failed getting body of %s: %v", nav, err)
	}
	fmt.Println(answer)
	return answer, nil
}
