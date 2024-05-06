package artificialintelligence

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func PromptBrowserChatGpt(prompt string) (string, error) {

	var answer string
	chromedpUrl := `ws://127.0.0.1:9222`
	timeoutTime := 10 * time.Second
	sleepTime := 1 * time.Second

	AiUrl := `https://chat.openai.com/chat`
	DAttributeOfWaitVisible := "M7 11L12 6L17 11M12 18V7"
	waitVisibleSelector := `path[d="` + DAttributeOfWaitVisible + `"]`

	var activeElementId string
	InnerHTML := `.prose`
	jsEvaluateElement := `document.activeElement.id`

	allocatorCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), chromedpUrl)
	// defer cancel()
	fmt.Println(cancel)

	ctx, cancel := chromedp.NewContext(allocatorCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, timeoutTime)
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate(AiUrl),
		chromedp.Sleep(sleepTime),
		chromedp.Evaluate(jsEvaluateElement, &activeElementId),
	); err != nil {
		log.Printf("1st chromedp.Run failed! err: '%s' AiUrl: %s activeElementId: '%s'", err.Error(), AiUrl, activeElementId)
		return "", err
	}

	if err := chromedp.Run(ctx,
		// chromedp.Sleep(d),
		chromedp.SetValue(`#`+activeElementId, prompt, chromedp.ByQuery),
		chromedp.SendKeys(`#`+activeElementId, "\n"),
		chromedp.WaitVisible(waitVisibleSelector),
		chromedp.InnerHTML(InnerHTML, &answer, chromedp.ByQuery),
	); err != nil {
		log.Printf("2nd chromedp.Run failed! err: '%s' AiUrl: %s activeElementId '%s' waitVisibleSelector: '%s'", err.Error(), AiUrl, activeElementId, waitVisibleSelector)
		return "", err
	}
	fmt.Println(answer)
	return answer, nil
}
