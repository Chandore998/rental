package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func createLogger(logType string) *log.Logger {

	// get current path location
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// get current log file location from env
	logFileLocation := os.Getenv("LOG_FILE_LOCATION")
	if logFileLocation == "" {
		logFileLocation = "/temp/logs/"
	}

	logFileName := fmt.Sprintf("%s%s%s_%s.log", path, logFileLocation, "rental", logType)
	flag.Parse()

	var file *os.File

	// Check if the log file exists at the provided file path. If it does, open the file and append a new log entry.
	// If it does not exist, create the file and add the log entry.

	if _, err := os.Stat(logFileName); os.IsNotExist(err) {
		file, err = os.Create(logFileName)
		if err != nil {
			panic(err)
		}
	} else {
		file, err = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
	}
	return log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func init() {
	InfoLog = createLogger("info")
	ErrorLog = createLogger("error")
}
