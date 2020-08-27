# Project: easy-logger

Just a Logger written in Go

## Content

* [Tech stack](#tech_stack)
* [Functions](#description)

## <a name="tech_stack">Tech stack</a>
- Golang 1.15 [Doc](https://golang.org/doc/)

## <a name="description">Functions</a>

1) #### func Debug

    `func Debug(lc , text string)`

    `Debug` - category of logs, used for debugging code.

2) #### func DebugJ

    `func DebugJ(lc , text string)`

    `DebugJ` - category of JSON logs, used for debugging code.

3) #### func Info

    `func Info(lc , text string)`

    `Info` - the category of logs used for information messages in the algorithm.

4) #### func InfoJ

    `func InfoJ(lc , text string)`

    `InfoJ` - the category of JSON logs used for information messages in the algorithm.

5) #### func Warning

    `func Warning(lc , text string)`

    `Warning` - the category of logs used to display warnings of the program logic.

6) #### func WarningJ

    `func WarningJ(lc , text string)`

    `WarningJ` - the category of JSON logs used to display warnings of the program logic.

7) #### func Error

    `func Error(lc , text string)`

    `Error` - a category of logs used to display errors in the program logic.

8) #### func ErrorJ

    `func ErrorJ(lc , text string)`

    `ErrorJ` - a category of JSON logs used to display errors in the program logic.

9) #### func Critical

    `func Critical(lc , text string)`

    `Critical` - a category of logs used to display fatal errors in the event of which the program terminates.

10) #### func CriticalJ

    `func CriticalJ(lc , text string)`

    `CriticalJ` - a category of JSON logs used to display fatal errors in the event of which the program terminates.

11) #### func SetConfig

    `func SetConfig(lc , text string)`

    `SetConfig` - configures the logging module for further work.
    - `config Cfg` - contains:
		- `appName <string>` - name of the program.
		- `logPath <string>` - path of the file for writing logs to it.
		If the variable is empty, then logs are not written to the file.
		- `level <int>` - logger's level: 0 - `DEBUG`, 1 - `INFO`, 2 - `WARNING`, 3 - `ERROR`, 4 - `CRITICAL`
