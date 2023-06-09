package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger

	logFileName string = time.Now().Format("2006-01-02_15.04.05")
)

func Enable(logPath string) {
	logFile, err := os.Create(fmt.Sprintf("%s/%s.log", logPath, logFileName))
	if err != nil {
		log.Fatalln("Error creating log file:", err)
	}

	InfoLogger = log.New(logFile, "INFO: ", log.Lmicroseconds)
	WarningLogger = log.New(logFile, "WARNING: ", log.Lmicroseconds)
	ErrorLogger = log.New(logFile, "ERROR: ", log.Lmicroseconds)
}

func Disable() {
	InfoLogger.SetOutput(io.Discard)
	ErrorLogger.SetOutput(io.Discard)
	WarningLogger.SetOutput(io.Discard)
}
