package ierror

import "github.com/evandrojr/expert-ai/ilog"

func PanicOnError(error error) {
	if error != nil {
		ilog.Fatal(error)
	}
}
