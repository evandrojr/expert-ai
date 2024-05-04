package log

import "log"

func Fatal(error string) {
	log.Println(error)
	panic(error)
}
