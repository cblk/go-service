package logy

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
)

func getLogConfig() *logConfig {
	return _logConfig
}

func LoadLogConfig(cnf *viper.Viper) {
	lc := &logConfig{}

	if cnf != nil {
		lc.AppID = cnf.GetString("app_id")
		lc.HostName = cnf.GetString("host_name")
		lc.Region = cnf.GetString("region")
		lc.Identifier = cnf.GetString("identifier")
		lc.PublishEnv = cnf.GetString("publish_env")
		lc.LogStdout = cnf.GetBool("log_stdout")
		lc.LogLevel = cnf.GetInt("log_level")
		lc.LogFilter = cnf.GetStringSlice("log_filter")
	}

	if lc.AppID != "" {
		SetAppID(lc.AppID)
	}

	SetLogStdout(lc.LogStdout)

	if lc.LogLevel > 0 {
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

	var hds []iHandler

	if getLogConfig().LogStdout || getLogConfig().PublishEnv == "" || getLogConfig().PublishEnv == PublishEnvDev {
		hds = append(hds, newStdout())
	}

	if getLogConfig().LogDir != "" {
		// todo Log Dir
	}

	_logHandler = newHandlers(getLogConfig().LogFilter, hds...)
}

// Debug
func Debug(format string, args ...interface{}) {
	_logHandler.Log(context.Background(), LogLevelDebug, newField(_key_log, fmt.Sprintf(format, args...)))
}

func DebugF(ctx context.Context, args ...Field) {
	_logHandler.Log(ctx, LogLevelDebug, args...)
}

func DebugS(ctx context.Context, args ...string) {
	_logHandler.Log(ctx, LogLevelDebug, logParams(args)...)
}

func DebugW(message string, err error) IError {
	return logWrap(context.Background(), LogLevelDebug, message, err, nil)
}

func DebugC(ctx context.Context, message string, err error) IError {
	return logWrap(ctx, LogLevelDebug, message, err, nil)
}

// Info
func Info(format string, args ...interface{}) {
	_logHandler.Log(context.Background(), LogLevelInfo, newField(_key_log, fmt.Sprintf(format, args...)))
}

func InfoF(ctx context.Context, args ...Field) {
	_logHandler.Log(ctx, LogLevelInfo, args...)
}

func InfoS(ctx context.Context, args ...string) {
	_logHandler.Log(ctx, LogLevelInfo, logParams(args)...)
}

func InfoW(message string, err error) IError {
	return logWrap(context.Background(), LogLevelInfo, message, err, nil)
}

func InfoC(ctx context.Context, message string, err error) IError {
	return logWrap(ctx, LogLevelInfo, message, err, nil)
}

// Notice
func Notice(format string, args ...interface{}) {
	_logHandler.Log(context.Background(), LogLevelNotice, newField(_key_log, fmt.Sprintf(format, args...)))
}

func NoticeF(ctx context.Context, args ...Field) {
	_logHandler.Log(ctx, LogLevelNotice, args...)
}

func NoticeS(ctx context.Context, args ...string) {
	_logHandler.Log(ctx, LogLevelNotice, logParams(args)...)
}

func NoticeW(message string, err error) IError {
	return logWrap(context.Background(), LogLevelNotice, message, err, nil)
}

func NoticeC(ctx context.Context, message string, err error) IError {
	return logWrap(ctx, LogLevelNotice, message, err, nil)
}

// Warn
func Warn(format string, args ...interface{}) {
	_logHandler.Log(context.Background(), LogLevelWarning, newField(_key_log, fmt.Sprintf(format, args...)))
}

func WarnF(ctx context.Context, args ...Field) {
	_logHandler.Log(ctx, LogLevelWarning, args...)
}

func WarnS(ctx context.Context, args ...string) {
	_logHandler.Log(ctx, LogLevelWarning, logParams(args)...)
}

func WarnW(message string, err error) IError {
	return logWrap(context.Background(), LogLevelWarning, message, err, nil)
}

func WarnC(ctx context.Context, message string, err error) IError {
	return logWrap(ctx, LogLevelWarning, message, err, nil)
}

// Error
func Error(format string, args ...interface{}) {
	_logHandler.Log(context.Background(), LogLevelError, newField(_key_log, fmt.Sprintf(format, args...)))
}

func ErrorF(ctx context.Context, args ...Field) {
	_logHandler.Log(ctx, LogLevelError, args...)
}

func ErrorS(ctx context.Context, args ...string) {
	_logHandler.Log(ctx, LogLevelError, logParams(args)...)
}

func ErrorW(message string, err error) IError {
	return logWrap(context.Background(), LogLevelError, message, err, nil)
}

// Fatal
func Fatal(format string, args ...interface{}) {
	_logHandler.Log(context.Background(), LogLevelFatal, newField(_key_log, fmt.Sprintf(format, args...)))
}

func FatalF(ctx context.Context, args ...Field) {
	_logHandler.Log(ctx, LogLevelFatal, args...)
}

func FatalS(ctx context.Context, args ...string) {
	_logHandler.Log(ctx, LogLevelFatal, logParams(args)...)
}

func FatalW(ctx context.Context, message string, err error) IError {
	return logWrap(ctx, LogLevelFatal, message, err, nil)
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
