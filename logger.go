package logger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	s "strings"
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
	fgPurple      string = "\033[35m"
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
var loggerCfg = Cfg{
	AppName: "APP",
	LogPath: "",
	Level:   0,
}

// bWriter - used to output logs to a file
var bWriter *bufio.Writer

/*
Cfg - a structure that stores the settings of the logging module
	- Cfg - contains:

		- AppName <string> - name of the program.

		- LogPath <string> - path of the file for writing logs to it.
		If the variable is empty, then logs are not written to the file.

		- Level <int> - logger's level: 0 - DEBUG, 1 - INFO, 2 - WARNING, 3 - ERROR, 4 - CRITICAL
*/
type Cfg struct {
	AppName string
	LogPath string
	Level   int
}

/*
JSONLog - structure of a log message in JSON format
	- JSONLog - contains:

		- AppName <string> - name of the program.

		- Dttm <string> - log's date and time.

		- LC <string> - logging category

		- File <string> - file name where log is called.

		- Line <string> - line where log is called.

		- Level <string> - logger level

		- Text <string> - log's text
*/
type JSONLog struct {
	AppName string `json:"app_name"`
	Dttm    string `json:"dttm"`
	LC      string `json:"lc"`
	File    string `json:"file"`
	Line    int    `json:"line"`
	Level   string `json:"level"`
	Text    string `json:"text"`
}

/*
Debug - category of logs, used for debugging code.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func Debug(lc, text string) {
	if loggerCfg.Level > 0 {
		return
	}

	logManage("DEBUG", lc, text)
}

/*
DebugJ - category of JSON logs, used for debugging code.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func DebugJ(lc, text string) {
	if loggerCfg.Level > 0 {
		return
	}

	logJSONManage("DEBUG", lc, text)
}

/*
Info - the category of logs used for information messages in the algorithm.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func Info(lc, text string) {
	if loggerCfg.Level > 1 {
		return
	}

	logManage("INFO", lc, text)
}

/*
InfoJ - the category of JSON logs used for information messages in the algorithm.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func InfoJ(lc, text string) {
	if loggerCfg.Level > 1 {
		return
	}

	logJSONManage("INFO", lc, text)
}

/*
Warning - the category of logs used to display warnings of the program logic.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func Warning(lc, text string) {
	if loggerCfg.Level > 2 {
		return
	}

	logManage("WARNING", lc, text)
}

/*
WarningJ - the category of JSON logs used to display warnings of the program logic.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func WarningJ(lc, text string) {
	if loggerCfg.Level > 2 {
		return
	}

	logJSONManage("WARNING", lc, text)
}

/*
Error - a category of logs used to display errors in the program logic.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func Error(lc, text string) {
	if loggerCfg.Level > 3 {
		return
	}

	logManage("ERROR", lc, text)
}

/*
ErrorJ - a category of JSON logs used to display errors in the program logic.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func ErrorJ(lc, text string) {
	if loggerCfg.Level > 3 {
		return
	}

	logJSONManage("ERROR", lc, text)
}

/*
Critical - a category of logs used to display fatal errors in the event of which the program terminates.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func Critical(lc, text string) {
	if loggerCfg.Level > 4 {
		return
	}

	logManage("CRITICAL", lc, text)
}

/*
CriticalJ - a category of JSON logs used to display fatal errors in the event of which the program terminates.

- lc <string> - logging category

- text <string> - text of the log's message
*/
func CriticalJ(lc, text string) {
	if loggerCfg.Level > 4 {
		return
	}

	logJSONManage("CRITICAL", lc, text)
}

/*
logJSONManage - used to manage the display of JSON logs.

- level <string> - level of logs

- lc <string> - logging category

- text <string> - text of error
*/
func logJSONManage(level, lc, text string) {
	if bWriter == nil {
		fmt.Println(makeLogJSONString(level, lc, text))
		return
	}

	fmt.Fprintf(bWriter, "%s\n", makeLogJSONString(level, lc, text))
	bWriter.Flush()
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

	strLog := fmt.Sprint("[", loggerCfg.AppName, "]", "[", now, "]", "[", lc, "]", "[", logPos, "]", "[", level, "]  ▶  ", text)

	if bWriter == nil {
		return fmt.Sprint(levelColors[level], strLog, fgNormal)
	}

	return fmt.Sprint(strLog)
}

/*
makeLogJSONString - used to create a formatted JSON log string.

- level <string> - level of logs

- lc <string> - logging category

- text <string> - text of error
*/
func makeLogJSONString(level, lc, text string) string {
	logPos := s.Split(getLogPosition(), ":")
	line, _ := strconv.Atoi(logPos[1])

	JSONLog := JSONLog{
		AppName: loggerCfg.AppName,
		Dttm:    time.Now().Format(timeFormat),
		LC:      lc,
		File:    logPos[0],
		Line:    line,
		Level:   level,
		Text:    text,
	}

	bLog, err := json.Marshal(JSONLog)
	if err != nil {
		return ""
	}

	if bWriter == nil {
		return fmt.Sprint(levelColors[level], string(bLog), fgNormal)
	}

	return fmt.Sprint(string(bLog))
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
	- config Cfg - contains:

		- appName <string> - name of the program.

		- logPath <string> - path of the file for writing logs to it.
		If the variable is empty, then logs are not written to the file.

		- level <int> - logger's level: 0 - DEBUG, 1 - INFO, 2 - WARNING, 3 - ERROR, 4 - CRITICAL
*/
func SetConfig(config Cfg) {
	loggerCfg = config
	checkConfig()
}

/*
checkConfig - checks the config and reconfigures it if necessary.
*/
func checkConfig() {
	if loggerCfg.Level < 0 || loggerCfg.Level > 4 {
		loggerCfg.Level = 0
	}

	if len(loggerCfg.LogPath) > 0 {
		if err := os.MkdirAll(filepath.Dir(loggerCfg.LogPath), os.ModePerm); err != nil {
			return
		}

		logs, err := os.OpenFile(loggerCfg.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return
		}

		bWriter = bufio.NewWriter(logs)
	}
}
