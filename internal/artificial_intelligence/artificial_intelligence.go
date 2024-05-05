package artificialintelligence

type Ai struct {
	Model                   string
	ChromedpUrl             string
	DAttributeOfWaitVisible string
	Url                     string
	SendKeys                string
	InnerHTML               string
}

type Chatgpt Ai
type Claude3 Ai

type ArtificialIntelligence interface {
	Setup() ArtificialIntelligence
	SubmitPrompt(string) (string, error)
	// setupContext() (context.Context, context.CancelFunc)
	// scrape(ctx context.Context, verbose bool, nav string, d time.Duration, question string) (string, error)
}
