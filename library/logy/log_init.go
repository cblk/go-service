package logy

import "os"

type LogLevel int

const (
	LogSplitByHour  = "hour"
	LogSplitByDay   = "day"
	LogSplitByWeek  = "week"
	LogSplitByMonth = "month"

	// log level : 0= all, 1=debug 2=info 3=notice 4=warning 5=error 7=fatal
	LogLevelAll     LogLevel = 0
	LogLevelDebug   LogLevel = 1
	LogLevelInfo    LogLevel = 2
	LogLevelNotice  LogLevel = 3
	LogLevelWarning LogLevel = 4
	LogLevelError   LogLevel = 5
	LogLevelFatal   LogLevel = 6

	PublishEnvDev  = "dev"
	PublishEnvPre  = "pre"
	PublishEnvProd = "prod"
)

var (
	_level       int
	_stdout      bool
	_dir         string
	_maxFileNum  int
	_maxFileSize int
	_logSplit    bool
	_logSplitBy  string

	_logLevelNames = [...]string{
		LogLevelAll:     "[ALL  ]",
		LogLevelDebug:   "[DEBUG ]",
		LogLevelInfo:    "[INFO  ]",
		LogLevelNotice:  "[NOTICE]",
		LogLevelWarning: "[WARN  ]",
		LogLevelError:   "[ERROR ]",
		LogLevelFatal:   "[FATAL ]",
	}

	_logConfig  *logConfig
	_logHandler IHandler
)

func init() {
	_hostName, _ := os.Hostname()
	_logConfig = &logConfig{
		HostName:   _hostName,
		LogLevel:   "all",
		logLevelIn: LogLevelAll,
		LogStdout:  true,
		LogSplit:   false,
		LogSplitBy: LogSplitByDay,
	}

	_logHandler = newHandlers([]string{}, newStdout())
}

func (ll LogLevel) String() string {
	return _logLevelNames[ll]
}
