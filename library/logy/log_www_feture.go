package logy

/*
// Debug
func Debug(message string, err error) IError {
	return logWrap(context.Background(), LogLevelDebug, message, err, nil).Error()
}

// Info
func Info(message string, err error) IError {
	return logWrap(context.Background(), LogLevelInfo, message, err, nil)
}

// Notice
func Notice(message string, err error) IError {
	return logWrap(context.Background(), LogLevelNotice, message, err, nil)
}

// Warn
func Warn(message string, err error) IError {
	return logWrap(context.Background(), LogLevelWarning, message, err, nil)
}

// Error
func Error(message string, err error) IError {
	return logWrap(context.Background(), LogLevelError, message, err, nil)
}

// Fatal
func Fatal(message string, err error) IError {
	return logWrap(context.Background(), LogLevelFatal, message, err, nil)
}

// Debug
func DebugF(format string, args ...interface{}) {
	_logHandler.Log(context.Background(), LogLevelDebug, newField(_key_log, fmt.Sprintf(format, args...)))
}

func DebugC(ctx context.Context, args ...Field) {
	_logHandler.Log(ctx, LogLevelDebug, args...)
}

func DebugS(ctx context.Context, args ...string) {
	_logHandler.Log(ctx, LogLevelDebug, logParams(args)...)
}

func DebugE(ctx context.Context, message string, err error) IError {
	return logWrap(ctx, LogLevelDebug, message, err, nil)
}


// Info
func InfoF(format string, args ...interface{}) {
	_logHandler.Log(context.Background(), LogLevelInfo, newField(_key_log, fmt.Sprintf(format, args...)))
}

func InfoC(ctx context.Context, args ...Field) {
	_logHandler.Log(ctx, LogLevelInfo, args...)
}

func InfoS(ctx context.Context, args ...string) {
	_logHandler.Log(ctx, LogLevelInfo, logParams(args)...)
}

func InfoE(ctx context.Context, message string, err error) IError {
	return logWrap(ctx, LogLevelInfo, message, err, nil)
}

// Notice
func NoticeF(format string, args ...interface{}) {
	_logHandler.Log(context.Background(), LogLevelNotice, newField(_key_log, fmt.Sprintf(format, args...)))
}

func NoticeC(ctx context.Context, args ...Field) {
	_logHandler.Log(ctx, LogLevelNotice, args...)
}

func NoticeS(ctx context.Context, args ...string) {
	_logHandler.Log(ctx, LogLevelNotice, logParams(args)...)
}

func NoticeM(ctx context.Context, message string, err error) IError {
	return logWrap(ctx, LogLevelNotice, message, err, nil)
}

// Warn
func WarnF(format string, args ...interface{}) {
	_logHandler.Log(context.Background(), LogLevelWarning, newField(_key_log, fmt.Sprintf(format, args...)))
}

func WarnC(ctx context.Context, args ...Field) {
	_logHandler.Log(ctx, LogLevelWarning, args...)
}

func WarnS(ctx context.Context, args ...string) {
	_logHandler.Log(ctx, LogLevelWarning, logParams(args)...)
}

func WarnM(ctx context.Context, message string, err error) IError {
	return logWrap(ctx, LogLevelWarning, message, err, nil)
}

// Error
func ErrorF(format string, args ...interface{}) {
	_logHandler.Log(context.Background(), LogLevelError, newField(_key_log, fmt.Sprintf(format, args...)))
}

func ErrorC(ctx context.Context, args ...Field) {
	_logHandler.Log(ctx, LogLevelError, args...)
}

func ErrorS(ctx context.Context, args ...string) {
	_logHandler.Log(ctx, LogLevelError, logParams(args)...)
}

func ErrorM(ctx context.Context, message string, err error) IError {
	return logWrap(ctx, LogLevelError, message, err, nil)
}

//Fatal
func FatalF(format string, args ...interface{}) {
	_logHandler.Log(context.Background(), LogLevelFatal, newField(_key_log, fmt.Sprintf(format, args...)))
}

func FatalC(ctx context.Context, args ...Field) {
	_logHandler.Log(ctx, LogLevelFatal, args...)
}

func FatalS(ctx context.Context, args ...string) {
	_logHandler.Log(ctx, LogLevelFatal, logParams(args)...)
}

func FatalM(ctx context.Context, message string, err error) IError {
	return logWrap(ctx, LogLevelFatal, message, err, nil)
}*/