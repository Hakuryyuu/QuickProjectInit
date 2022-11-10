package main

import (
	"fmt"
	"time"
)

// Console Log Colors
const RESET = "\033[0m"
const COLOR_RED = "\033[31m"
const COLOR_GREEN = "\033[32m"
const COLOR_YELLOW = "\033[33m"
const COLOR_BLUE = "\033[34m"
const COLOR_PURPLE = "\033[35m"
const COLOR_CYAN = "\033[36m"
const COLOR_GRAY = "\033[37m"
const COLOR_WHITE = "\033[97m"

type level int

// Log-Level
const (
	INFO level = iota // Default Value
	ERROR
	WARNING
	DEBUG
)

const DEBUG_MODE = false // Enable/Disable Debug Logging

func log(bSaveToFile bool, sInfo string, lvLevel level) {

	/*if runtime.GOOS == "windows" { // Color Codes do not work for standard Windows CMD so we remove it so no weird strings will be displayed.
		RESET = ""
		COLOR_RED = ""
		COLOR_GREEN = ""
		COLOR_YELLOW = ""
		COLOR_BLUE = ""
		COLOR_PURPLE = ""
		COLOR_CYAN = ""
		COLOR_GRAY = ""
		COLOR_WHITE = ""
	}*/
	switch lvLevel {
	case 1: // Error
		logError(bSaveToFile, sInfo, getCurrentTime())
		break
	case 2: // Warning
		logWarning(bSaveToFile, sInfo, getCurrentTime())
		break
	case 3: // Debug
		break
	default: // If defined by '0' or none this case is used (INFO Level)
		logInfo(bSaveToFile, sInfo, getCurrentTime())
		break
	}
}

func logInfo(bSaveToFile bool, sInfo string, sTime string) {
	fmt.Println(sTime + " | " + COLOR_GREEN + "INFO" + RESET + " | " + sInfo)
	if bSaveToFile {
		appendToLogFile(sTime + " | INFO | " + sInfo)
	}
}

func logError(bSaveToFile bool, sInfo string, sTime string) {
	fmt.Println(sTime + " | " + COLOR_RED + "ERROR" + RESET + " | " + sInfo)
	if bSaveToFile {
		appendToLogFile(sTime + " | ERROR | " + sInfo)
	}
}

func logWarning(bSaveToFile bool, sInfo string, sTime string) {
	fmt.Println(sTime + " | " + COLOR_YELLOW + "WARNING" + RESET + " | " + sInfo)
	if bSaveToFile {
		appendToLogFile(sTime + " | WARNING | " + sInfo)
	}
}

func logDebug(bSaveToFile bool, sInfo string, sTime string) {
	if DEBUG_MODE {
		fmt.Println(sTime + " | " + COLOR_CYAN + "DEBUG" + RESET + " | " + sInfo)
		if bSaveToFile {
			appendToLogFile(sTime + " | DEBUG | " + sInfo)
		}
	}
}

func getCurrentTime() string {
	tTime := time.Now()
	tTimeFormatted := tTime.Format("2006.01.02 15:04:05") // Must be the string used here otherwise Go won't recognize it as Format.
	return tTimeFormatted                                 // For me it would've been 11.11.2016 ;)
}
