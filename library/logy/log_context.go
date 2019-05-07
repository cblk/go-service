package logy

import (
	"context"
)

type IError interface {
	DebugW(message string, err error) IError
	InfoW(message string, err error) IError
	NoticeW(message string, err error) IError
	WarnW(message string, err error) IError
	ErrorW(message string, err error) IError
	FatalW(message string, err error) IError

	Error() string
}

type logWithValue struct {
	Context context.Context
	Message string
	Level   LogLevel
	Err     error
	Pre     IError
}

func logWrap(ctx context.Context, logLevel LogLevel, message string, err error, pre IError) IError {
	return &logWithValue{Context: ctx, Message: message, Level: logLevel, Err: err, Pre: pre}
}

func (lw *logWithValue) DebugW(message string, err error) IError {
	if lw == nil {
		return lw
	}

	newLog := logWrap(lw.Context, LogLevelDebug, message, err, lw)

	return newLog
}

func (lw *logWithValue) InfoW(message string, err error) IError {
	if lw == nil {
		return lw
	}

	newLog := logWrap(lw.Context, LogLevelInfo, message, err, lw)

	return newLog
}

func (lw *logWithValue) NoticeW(message string, err error) IError {
	if lw == nil {
		return lw
	}

	newLog := logWrap(lw.Context, LogLevelNotice, message, err, lw)

	return newLog
}

func (lw *logWithValue) WarnW(message string, err error) IError {
	if lw == nil {
		return lw
	}

	newLog := logWrap(lw.Context, LogLevelWarning, message, err, lw)

	return newLog
}

func (lw *logWithValue) ErrorW(message string, err error) IError {
	if lw == nil {
		return lw
	}

	newLog := logWrap(lw.Context, LogLevelError, message, err, lw)

	return newLog
}

func (lw *logWithValue) FatalW(message string, err error) IError {
	if lw == nil {
		return lw
	}

	newLog := logWrap(lw.Context, LogLevelFatal, message, err, lw)

	return newLog
}

func (lw *logWithValue) Error() string {
	if lw == nil {
		return ""
	}

	locFields := make([]Field, 0)
	for {
		if lw.Err == nil {
			if lw.Pre == nil {
				break
			}

			continue
		}

		locFields = append(locFields, newField(lw.Message, lw.Err))

		if lw.Pre == nil {
			break
		}

		lw = lw.Pre.(*logWithValue)
	}

	return _logHandler.Log(lw.Context, lw.Level, locFields...)
}
