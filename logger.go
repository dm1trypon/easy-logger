package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

// Time output format in logs
const timeFormat = "2006/01/02 - 15:04:05"

// List of color codes for displaying logger messages
const (
	fgBoldBlack   string = "\x1b[30;1m"
	fgBoldRed     string = "\x1b[31;1m"
	fgBoldGreen   string = "\x1b[32;1m"
	fgBoldYellow  string = "\x1b[33;1m"
	fgBoldBlue    string = "\x1b[34;1m"
	fgBoldMagenta string = "\x1b[35;1m"
	fgBoldCyan    string = "\x1b[36;1m"
	fgBoldWhite   string = "\x1b[37;1m"
	fbBlack       string = "\x1b[30m"
	fgRed         string = "\x1b[31m"
	fgGreen       string = "\x1b[32m"
	fgYellow      string = "\x1b[33m"
	fgBlue        string = "\x1b[34m"
	fgMagenta     string = "\x1b[35m"
	fgCyan        string = "\x1b[36m"
	fgWhite       string = "\x1b[37m"
	fgNormal      string = "\x1b[0m"
)

// Array of matching the color of the log message to its level
var levelColors = map[string]string{
	"DEBUG":    fgWhite,
	"INFO":     fgGreen,
	"WARNING":  fgYellow,
	"ERROR":    fgRed,
	"CRITICAL": fgBoldRed,
}

// Logger settings
var loggerCfg cfg

/*
cfg - a structure that stores the settings of the logging module
*/
type cfg struct {
	appName string // name of the program.
	logPath string // path of the file for writing logs to it.
}

/*
Debug - category of logs, used for debugging code.
	lc <string> - logging category
	text <string> - text of the log's message
*/
func Debug(lc string, text string) {
	fmt.Println(makeLogString("DEBUG", lc, text))
}

/*
Info - the category of logs used for information messages in the algorithm.
	lc <string> - logging category
	text <string> - text of the log's message
*/
func Info(lc string, text string) {
	fmt.Println(makeLogString("INFO", lc, text))
}

/*
Warning - the category of logs used to display warnings of the program logic.
	lc <string> - logging category
	text <string> - text of the log's message
*/
func Warning(lc string, text string) {
	fmt.Println(makeLogString("WARNING", lc, text))
}

/*
Error - a category of logs used to display errors in the program logic.
	lc <string> - logging category
	text <string> - text of the log's message
*/
func Error(lc string, text string) {
	fmt.Println(makeLogString("ERROR", lc, text))
}

/*
Critical - a category of logs used to display fatal errors in the event of which the program terminates.
	lc <string> - logging category
	text <string> - text of the log's message
*/
func Critical(lc string, text string) {
	fmt.Println(makeLogString("CRITICAL", lc, text))
}

func makeLogString(level string, lc string, message string) string {
	now := time.Now().Format(timeFormat)
	logPos := getLogPosition()
	return fmt.Sprint(levelColors[level],
		"[", loggerCfg.appName, "]", "[", now, "]", "[", lc, "]", "[", logPos, "]", "[", level, "]  ▶  ", message)
}

func getLogPosition() string {
	_, file, line, _ := runtime.Caller(3)

	return fmt.Sprint(filepath.Base(file), ":", line)
}

/*
SetConfig - configures the logging module for further work.
	appName - name of the program.
	logPath - path of the file for writing logs to it. If the variable is empty, then logs are not written to the file.
*/
func SetConfig(appName string, logPath string) {
	loggerCfg.appName = appName
	loggerCfg.logPath = logPath
}
