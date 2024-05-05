package ierror

import "github.com/evandrojr/expert-ai/internal/ilog"

func PanicOnError(error error) {
	if error != nil {
		ilog.Fatal(error)
	}
}
