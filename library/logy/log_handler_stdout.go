package logy

import (
	"context"
	"io"
	"os"
	"strconv"
	"time"
)

const defaultPattern = "%L %e %D %T %a %f %s %M"

var _defaultStdout = newStdout()

// stdoutHandler stdout log iHandler
type stdoutHandler struct {
	out    io.Writer
	writer iWriter
}

func newStdout() *stdoutHandler {
	return &stdoutHandler{
		out:    os.Stderr,
		writer: newWriter(defaultPattern),
	}
}

func (sh *stdoutHandler) Log(ctx context.Context, ll LogLevel, args ...Field) string {
	fs := make(map[string]interface{}, 10+len(args))
	for _, arg := range args {
		fs[arg.Key] = arg.Value
	}

	// add extra fields
	addExtraFields(ctx, fs)

	fs[_key_logTime] = time.Now().Format(_timeFormat)

	val, _ := sh.writer.Write(sh.out, fs, int(ll) >= GetLogLevel())
	if int(ll) >= GetLogLevel() {
		_, _ = sh.out.Write([]byte("\n"))
	}

	return val
}

func (sh *stdoutHandler) Close() error {
	return nil
}

// SetFormat set stdout log output format
// %T time format at "15:04:05.999"
// %t time format at "15:04:05"
// %D data format at "2006/01/02"
// %d data format at "01/02"
// %L log level e.g. INFO WARN ERROR
// %f function name and line number e.g. model.Get:121
// %i instance id
// %e deploy env e.g. dev uat fat prod
// %z zone
// %S full file name and line number: /a/b/c/d.go:23
// %s final file name element and line number: d.go:23
// %M log message and additional fields: key=value this is log message
func (sh *stdoutHandler) SetFormat(format string) {
	sh.writer = newWriter(format)
}

func addExtraFields(ctx context.Context, fields map[string]interface{}) {
	count, ok := ctx.Value(_key_wrap_count).(int)
	if !ok {
		return
	}

	for i := 1; i <= count; count++ {
		wrapIndex := strconv.Itoa(i)
		if withValue, ok := ctx.Value(_key_wrap + wrapIndex).(logWithValue); ok {
			fields[withValue.Message] = withValue.Err
		}
	}
}
