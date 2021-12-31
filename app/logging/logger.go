package logs

import (
	"log"
	"os"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "nitd-results", log.LstdFlags)
}

func Write(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func Fatal(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}
