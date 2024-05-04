package artificialintelligence

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/evandrojr/expert-ai/tools"
)

func (ai Claude3) Setup() ArtificialIntelligence {
	ai = Claude3{ChromedpUrl: `ws://127.0.0.1:9222`,
		DAttributeOfWaitVisible: "M11.2 21a1 1 0 0 1-.914-.594L7.351 13.8H5.6a2.615 2.615 0 0 1-1.838-.761A2.59 2.59 0 0 1 3 11.2V5.6c0-.685.277-1.354.762-1.838A2.615 2.615 0 0 1 5.6 3h11.707a2.6 2.6 0 0 1 2.561 2.21l1.103 7.198c.057.369.031.755-.072 1.115a2.61 2.61 0 0 1-1.432 1.648 2.615 2.615 0 0 1-1.096.228H14.6v2.2c0 .907-.354 1.761-.994 2.403A3.38 3.38 0 0 1 11.2 21ZM9 12.587l2.792 6.282a1.401 1.401 0 0 0 .807-1.27v-3.2a1 1 0 0 1 1-1h4.783c.08-.009.184-.017.264-.052a.607.607 0 0 0 .33-.38.617.617 0 0 0 .018-.256l-1.104-7.2A.598.598 0 0 0 17.298 5H9v7.587ZM5.6 5a.6.6 0 0 0-.6.6v5.6a.599.599 0 0 0 .6.6H7V5H5.6Z",
		Url:                     "https://poe.com/Claude-3-Haiku",
		SendKeys:                `textarea[placeholder="Talk to Claude-3-Haiku or @ a bot"]`,
		InnerHTML:               `main`,
	}
	return ai
}

func (ai Claude3) SubmitPrompt(prompt string) (string, error) {
	ctx, _ := ai.setupContext()
	// defer cancel()

	answer, err := ai.scrape(ctx, true, ai.Url, 2*time.Second, prompt)
	if err != nil {
		return "", err
	}

	answer = ai.cleanAnswer(answer, prompt)

	return answer, nil
}

func (ai Claude3) cleanAnswer(dirtAnswer string, prompt string) string {
	text := tools.HTMLToText(dirtAnswer)
	text = tools.RemoveRubbishFromBeginning(text, prompt)
	text = tools.CleanPoeAnswer(text)
	return text
}

func (ai Claude3) setupContext() (context.Context, context.CancelFunc) {
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), ai.ChromedpUrl)
	return ctx, cancel
}

func (ai Claude3) scrape(ctx context.Context, verbose bool, nav string, d time.Duration, question string) (string, error) {
	var opts []chromedp.ContextOption
	var answer string
	var activeElementClass string
	waitVisibleSelector := `path[d="` + ai.DAttributeOfWaitVisible + `"]`

	if verbose {
		opts = append(opts, chromedp.WithDebugf(log.Printf))
	}
	ctx, _ = chromedp.NewContext(ctx, opts...)

	if err := chromedp.Run(ctx,
		chromedp.Navigate(nav),
		chromedp.Sleep(d),
		chromedp.Evaluate(`document.activeElement.className`, &activeElementClass),
	); err != nil {
		return "", fmt.Errorf("1st run failed %s: %v activeElementClass %v", nav, err, activeElementClass)
	}
	fmt.Println(activeElementClass)
	fmt.Println(answer)

	if err := chromedp.Run(ctx,
		// chromedp.Sleep(d),
		chromedp.SetValue(`.`+activeElementClass, question, chromedp.ByQuery),
		chromedp.SendKeys(`.`+activeElementClass, "\n"),
		chromedp.WaitVisible(waitVisibleSelector),
		chromedp.InnerHTML(ai.InnerHTML, &answer, chromedp.ByQuery),
	); err != nil {
		return "", fmt.Errorf("2nd run failed %s: %v activeElementClass %v", nav, err, activeElementClass)
	}

	return answer, nil
}
