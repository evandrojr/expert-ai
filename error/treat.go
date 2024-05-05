package error

import "github.com/evandrojr/expert-ai/log"

func PanicOnError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}
