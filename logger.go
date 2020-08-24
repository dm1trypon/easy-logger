package logger

import (
	"bufio"
	"fmt"
	"os"
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
var loggerCfg = cfg{
	appName: "APP",
	logPath: "",
	level:   0,
}

// bWriter - used to output logs to a file
var bWriter *bufio.Writer

/*
cfg - a structure that stores the settings of the logging module
*/
type cfg struct {
	appName string // name of the program.
	logPath string // path of the file for writing logs to it.
	level   int    // logger's level: 0 - DEBUG, 1 - INFO, 2 - WARNING, 3 - ERROR, 4 - CRITICAL
}

/*
Debug - category of logs, used for debugging code.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func Debug(lc, text string) {
	if loggerCfg.level > 0 {
		return
	}

	logManage("DEBUG", lc, text)
}

/*
Info - the category of logs used for information messages in the algorithm.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func Info(lc, text string) {
	if loggerCfg.level > 1 {
		return
	}

	logManage("INFO", lc, text)
}

/*
Warning - the category of logs used to display warnings of the program logic.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func Warning(lc, text string) {
	if loggerCfg.level > 2 {
		return
	}

	logManage("WARNING", lc, text)
}

/*
Error - a category of logs used to display errors in the program logic.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func Error(lc, text string) {
	if loggerCfg.level > 3 {
		return
	}

	logManage("ERROR", lc, text)
}

/*
Critical - a category of logs used to display fatal errors in the event of which the program terminates.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func Critical(lc, text string) {
	if loggerCfg.level > 4 {
		return
	}

	logManage("CRITICAL", lc, text)
}

/*
logManage - used to manage the display of logs.

- level <string> - level of logs

- lc <string> - logging category

- text <string> - text of error
*/
func logManage(level, lc, text string) {
	if bWriter == nil {
		fmt.Println(makeLogString(level, lc, text))
		return
	}

	fmt.Fprintf(bWriter, "%s\n", makeLogString(level, lc, text))
	bWriter.Flush()
}

/*
makeLogString - used to create a formatted log string.

- level <string> - level of logs

- lc <string> - logging category

- text <string> - text of error

*/
func makeLogString(level, lc, text string) string {
	now := time.Now().Format(timeFormat)
	logPos := getLogPosition()
	return fmt.Sprint(levelColors[level],
		"[", loggerCfg.appName, "]", "[", now, "]", "[", lc, "]", "[", logPos, "]", "[", level, "]  ▶  ", text)
}

/*
getLogPosition - used to determine where the log is called.
*/
func getLogPosition() string {
	_, file, line, _ := runtime.Caller(4)

	return fmt.Sprint(filepath.Base(file), ":", line)
}

/*
SetConfig - configures the logging module for further work.

- appName <string> - name of the program.

- logPath <string> - path of the file for writing logs to it.
If the variable is empty, then logs are not written to the file.

- level <int> - logger's level: 0 - DEBUG, 1 - INFO, 2 - WARNING, 3 - ERROR, 4 - CRITICAL
*/
func SetConfig(appName, logPath string, level int) {
	loggerCfg.appName = appName
	loggerCfg.logPath = logPath
	loggerCfg.level = level

	checkConfig()
}

/*
checkConfig - checks the config and reconfigures it if necessary.
*/
func checkConfig() {
	if loggerCfg.level < 0 || loggerCfg.level > 4 {
		loggerCfg.level = 0
	}

	if len(loggerCfg.logPath) > 0 {
		logs, err := os.OpenFile(loggerCfg.logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return
		}

		bWriter = bufio.NewWriter(logs)
	}
}
