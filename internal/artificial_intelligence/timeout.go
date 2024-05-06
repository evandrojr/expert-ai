package artificialintelligence

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func Timeout() {
	// options := append(chromedp.DefaultExecAllocatorOptions[:],
	// 	// block all images
	// 	chromedp.Flag("blink-settings", "imagesEnabled=false"),
	// )
	// allocatorCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	allocatorCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), `ws://127.0.0.1:9222`)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocatorCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	var activeElementId string

	err := chromedp.Run(ctx,
		// block the urls that take very long to load
		// network.SetBlockedURLS([]string{"https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"}),
		chromedp.Navigate("https://chat.openai.com/chat"),
		chromedp.Sleep(1*time.Second),
		chromedp.Evaluate(`document.activeElement.id`, &activeElementId),

		// chromedp.Text("question-header", &example, chromedp.ByID),
	)
	log.Printf("Got text: %s", activeElementId)
	if err != nil {
		log.Fatal(err)
	}

	question := "Quem descobriu o Brasil"
	DAttributeOfWaitVisible := "M7 11L12 6L17 11M12 18V7"
	var answer = ""
	waitVisibleSelector := `path[d="` + DAttributeOfWaitVisible + `"]`

	InnerHTML := `.prose`

	if err := chromedp.Run(ctx,
		// chromedp.Sleep(d),
		chromedp.SetValue(`#`+activeElementId, question, chromedp.ByQuery),
		chromedp.SendKeys(`#`+activeElementId, "\n"),
		chromedp.WaitVisible(waitVisibleSelector),
		chromedp.InnerHTML(InnerHTML, &answer, chromedp.ByQuery),
	); err != nil {
		fmt.Println("2nd run failed   activeElementClass ", err, activeElementId)
	}
	fmt.Println(answer)
}
