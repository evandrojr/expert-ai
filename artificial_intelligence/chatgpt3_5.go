package artificialintelligence

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/evandrojr/expert-ai/tool"
)

func (ai Chatgpt) Setup() ArtificialIntelligence {
	ai = Chatgpt{
		Model:                   "chatgpt",
		ChromedpUrl:             `ws://127.0.0.1:9222`,
		DAttributeOfWaitVisible: "M7 11L12 6L17 11M12 18V7",
		Url:                     "https://chat.openai.com/chat",
		SendKeys:                `#prompt-textarea`,
		InnerHTML:               `.prose`,
	}
	return ai
}

func (ai Chatgpt) SubmitPrompt(prompt string) (string, error) {

	ctx, _ := ai.setupContext()
	// defer cancel()

	answer, err := ai.scrape(ctx, true, ai.Url, 1*time.Second, prompt)
	if err != nil {
		return "", err
	}

	return tool.HTMLToText(answer), nil
}

func (ai Chatgpt) setupContext() (context.Context, context.CancelFunc) {
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), ai.ChromedpUrl)
	return ctx, cancel
}

func (ai Chatgpt) scrape(ctx context.Context, verbose bool, nav string, d time.Duration, question string) (string, error) {

	var opts []chromedp.ContextOption
	var activeElementId string

	if verbose {
		opts = append(opts, chromedp.WithDebugf(log.Printf))
	}

	ctx, _ = chromedp.NewContext(ctx, opts...)
	var answer string
	waitVisibleSelector := `path[d="` + ai.DAttributeOfWaitVisible + `"]`

	if err := chromedp.Run(ctx,
		chromedp.Navigate(nav),
		chromedp.Sleep(d),
		chromedp.Evaluate(`document.activeElement.id`, &activeElementId),
	); err != nil {
		return "", fmt.Errorf("1st run failed %s: %v activeElementClass %v", nav, err, activeElementId)
	}
	fmt.Println(activeElementId)
	fmt.Println(answer)

	if err := chromedp.Run(ctx,
		// chromedp.Sleep(d),
		chromedp.SetValue(`#`+activeElementId, question, chromedp.ByQuery),
		chromedp.SendKeys(ai.SendKeys, "\n"),
		chromedp.WaitVisible(waitVisibleSelector),
		chromedp.InnerHTML(ai.InnerHTML, &answer, chromedp.ByQuery),
	); err != nil {
		return "", fmt.Errorf("2nd run failed %s: %v activeElementClass %v", nav, err, activeElementId)
	}
	return answer, nil
}
