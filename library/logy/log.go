package logy

import (
	"context"

	"github.com/spf13/viper"
)

func getLogConfig() *logConfig {
	return _logConfig
}

func LoadLogConfig(cnf *viper.Viper) {
	lc := &logConfig{}

	if cnf != nil {
		lc.AppID = cnf.GetString("log.app_id")
		lc.LogFormat = cnf.GetString("log.format")
		lc.LogStdout = cnf.GetBool("log.stdout")
		lc.LogLevel = cnf.GetString("log.level")
		lc.LogFilter = cnf.GetStringSlice("log.filter")

		lc.HostName = cnf.GetString("host_name")
		lc.Region = cnf.GetString("region")
		lc.Identifier = cnf.GetString("identifier")
		lc.PublishEnv = cnf.GetString("publish_env")

		lc.LogDir = cnf.GetString("log.dir")
		lc.LogMaxFileNum = cnf.GetInt("log.max_file_num")
		lc.LogMaxFileSize = cnf.GetInt("log.max_file_size")
		lc.LogSplit = cnf.GetBool("log.split")
		lc.LogSplitBy = cnf.GetString("log.split_by")
	}

	if lc.AppID != "" {
		SetAppID(lc.AppID)
	}

	SetLogStdout(lc.LogStdout)

	if lc.LogLevel != "" {
		SetLogLevel(lc.LogLevel)
	}

	if lc.LogDir != "" {
		SetLogDir(lc.LogDir)
	}

	if lc.LogMaxFileNum > 0 {
		SetLogMaxFileNum(lc.LogMaxFileNum)
	}

	if lc.LogMaxFileSize > 0 {
		SetLogMaxFileSize(lc.LogMaxFileSize)
	}

	if lc.LogSplit {
		SetLogSplit(lc.LogSplit)
	}

	if lc.LogSplitBy != "" {
		SetLogSplitBy(lc.LogSplitBy)
	}

	SetLogFilter(lc.LogFilter)

	var hds []IHandler

	if getLogConfig().LogStdout || getLogConfig().PublishEnv == "" || getLogConfig().PublishEnv == PublishEnvDev {
		hds = append(hds, _defaultStdout)
	}

	if getLogConfig().LogDir != "" {
		// todo Log Dir
	}

	_logHandler = newHandlers(getLogConfig().LogFilter, hds...)

	if lc.LogFormat != "" {
		SetFormat(lc.LogFormat)
	}
}

// nil Error
func Nil() IError {
	return logWrap(context.Background(), LogLevelAll, "", nil, nil)
}

func IsNil(val IError) bool {
	if val == nil {
		return true
	}

	if loc, ok := val.(*logWithValue); ok {
		if loc.Err == nil {
			return true
		}
	}

	return false
}

// Debug
func Debug(message string, err error) {
	_ = logWrap(context.Background(), LogLevelDebug, message, err, nil).Error()
}

// Info
func Info(message string, err error) {
	_ = logWrap(context.Background(), LogLevelInfo, message, err, nil).Error()
}

// Notice
func Notice(message string, err error) {
	_ = logWrap(context.Background(), LogLevelNotice, message, err, nil).Error()
}

// Warn
func Warn(message string, err error) {
	_ = logWrap(context.Background(), LogLevelWarning, message, err, nil).Error()
}

// Error
func Error(message string, err error) {
	_ = logWrap(context.Background(), LogLevelError, message, err, nil).Error()
}

// Fatal
func Fatal(message string, err error) {
	_ = logWrap(context.Background(), LogLevelFatal, message, err, nil).Error()
}

/*
SetFormat only effective on stdout and file handler
%L log level e.g. DEBUG INFO NOTICE WARN ERROR FATAL
%T time format at "15:04:05.000"
%t time format at "15:04"
%D data format at "2006-01-02"
%d data format at "01-02"
%a appId
%e deploy env e.g. dev pre prod
%r region
%i identifier
%f function name and line number e.g. model.Get:121
%S full file name and line number: /a/b/c/d.go:23
%s final file name element and line number: d.go:23
%M log message and additional fields: key=value this is log message
*/
func SetFormat(format string) {
	_logHandler.SetFormat(format)
}

func Close() (err error) {
	err = _logHandler.Close()

	_logHandler = _defaultStdout
	return
}

func GetLogStdoutInstance() IHandler {
	return _logHandler
}
