package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var logFile *os.File
var err error

// Function to initialize the error logger
func InitializeLogger(fileName string) error {
	logFile, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	log.SetOutput(logFile)
	return nil
}

// Function to log errors along with the line and method
func Fatal(err error) {

	fmt.Println(err)
	pc, filePath, line, ok := runtime.Caller(1)
	if !ok {
		log.Printf("Error: Unable to retrieve caller information")
		return
	}
	methodName := runtime.FuncForPC(pc).Name()

	log.Printf("Error at %s:%d - %s: %v", filePath, line, methodName, err)
	panic(err)
}
