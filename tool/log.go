package tool

import "log"

func LogFatal(error string) {
	log.Println(error)
	panic(error)
}
