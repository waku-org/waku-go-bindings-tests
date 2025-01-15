package utilities

import (
	"fmt"
	"log"
	"os"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
)

var (
	debugLogger *log.Logger
	errorLogger *log.Logger
	infoLogger  *log.Logger
)

func init() {
	// Open file in write mode to overwrite content on each run using os.O_TRUNC
	file, err := os.OpenFile("waku_node.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	debugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// logs debug messages in yellow.
func LogDebug(message string) {
	coloredMessage := fmt.Sprintf("%s%s%s", ColorYellow, message, ColorReset)
	fmt.Println(coloredMessage)
	debugLogger.Println(message)
}

// logs error messages in red.
func LogError(message string) {
	coloredMessage := fmt.Sprintf("%s%s%s", ColorRed, message, ColorReset)
	fmt.Println(coloredMessage)
	errorLogger.Println(message)
}

// logs informational messages.
func LogInfo(message string) {
	fmt.Println(message)
	infoLogger.Println(message)
}
