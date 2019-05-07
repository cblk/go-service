package logy

import (
	"context"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var funcNames sync.Map

type Field struct {
	Key   string
	Value interface{}
}

func newField(key string, value interface{}) Field {
	return Field{
		Key:   key,
		Value: value,
	}
}

type IHandler interface {
	SetFormat(string)

	Log(context.Context, LogLevel, ...Field) string

	Close() error

	Write(p []byte) (n int, err error)
}

type handlers struct {
	filters  map[string]struct{}
	handlers []IHandler
}

func newHandlers(filters []string, hds ...IHandler) *handlers {
	locFilters := make(map[string]struct{})
	for _, item := range filters {
		locFilters[item] = struct{}{}
	}
	return &handlers{filters: locFilters, handlers: hds}
}

func (hs *handlers) Log(ctx context.Context, logLevel LogLevel, fields ...Field) string {
	var fn string
	for i := range fields {
		if _, ok := hs.filters[fields[i].Key]; ok {
			fields[i].Value = "***"
		}
	}
	if fn == "" {
		fields = append(fields, newField(_key_source, funcName(4)))
	}
	fields = append(fields,
		newField(_key_logTime, time.Now()),
		newField(_key_logLevel, int(logLevel)),
		newField(_key_logLevelName, logLevel.String()))

	if GetAppID() != "" {
		fields = append(fields, newField(_key_appID, GetAppID()))
	}

	val := ""
	for _, hItem := range hs.handlers {
		val = hItem.Log(ctx, logLevel, fields...)
	}

	return val
}

func (hs *handlers) Close() (err error) {
	for _, hItem := range hs.handlers {
		if err := hItem.Close(); err != nil {
			err = errors.WithStack(err)
		}
	}
	return
}

func (hs *handlers) Write(p []byte) (n int, err error) {
	for _, hItem := range hs.handlers {
		n, err = hItem.Write(p)
	}

	return
}

// SetFormat .
func (hs *handlers) SetFormat(format string) {
	for _, h := range hs.handlers {
		h.SetFormat(format)
	}
}

// funcName get func name.
func funcName(skip int) (name string) {
	if pc, _, lineNo, ok := runtime.Caller(skip); ok {
		if v, ok := funcNames.Load(pc); ok {
			name = v.(string)
		} else {
			name = runtime.FuncForPC(pc).Name() + ":" + strconv.FormatInt(int64(lineNo), 10)
			funcNames.Store(pc, name)
		}
	}
	return
}
