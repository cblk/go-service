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

	GetContent() context.Context
	GetMessage() string
	GetLevel() LogLevel
	GetError() error
	GetPre() IError
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

	if len(locFields) == 0 {
		return ""
	}

	return _logHandler.Log(lw.Context, lw.Level, locFields...)
}

func (lw *logWithValue) GetContent() context.Context {
	if lw == nil {
		return context.Background()
	}

	return lw.Context
}

func (lw *logWithValue) GetMessage() string {
	if lw == nil {
		return ""
	}

	return lw.Message
}

func (lw *logWithValue) GetLevel() LogLevel {
	if lw == nil {
		return LogLevelAll
	}

	return lw.Level
}

func (lw *logWithValue) GetError() error {
	if lw == nil {
		return nil
	}

	return lw.Err
}

func (lw *logWithValue) GetPre() IError {
	if lw == nil {
		return Nil()
	}

	return lw.Pre
}
