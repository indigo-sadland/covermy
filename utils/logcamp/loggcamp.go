package logcamp

import (
	"log"
	"os"
)

var ErrorLogger *log.Logger
var InfoLogger *log.Logger

// Init creates logger.
func Init() {

	var logFile string

	dir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	logFile = dir + "/.local/share/covermy/" + "logs.txt"

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
