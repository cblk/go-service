package logy

import (
	"bytes"
	"io"
	"sync"
)

type iWriter interface {
	Write(io.Writer, LogLevel, map[string]interface{}, bool) (string, error)
	WriteString(LogLevel, map[string]interface{}) string
}

type writer struct {
	funcs []func(LogLevel, map[string]interface{}) string
	bufPool sync.Pool
}

func newWriter(format string) iWriter {
	wr := &writer{
		bufPool: sync.Pool{New: func() interface{} { return &bytes.Buffer{} }},
	}
	bs := make([]byte, 0, len(format))
	for i := 0; i < len(format); i++ {
		if format[i] != '%' {
			bs = append(bs, format[i])
			continue
		}
		if i+1 >= len(format) {
			bs = append(bs, format[i])
			continue
		}
		f, ok := BaseMap[string(format[i+1])]
		if !ok {
			bs = append(bs, format[i])
			continue
		}
		if len(bs) != 0 {
			wr.funcs = append(wr.funcs, textFormat(LogLevelAll, string(bs)))
			bs = bs[:0]
		}
		wr.funcs = append(wr.funcs, f)
		i++
	}
	if len(bs) != 0 {
		wr.funcs = append(wr.funcs, textFormat(LogLevelAll, string(bs)))
	}
	return wr
}

func (wr *writer) Write(w io.Writer, ll LogLevel, params map[string]interface{}, realWrite bool) (string, error) {
	buf := wr.bufPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		wr.bufPool.Put(buf)
	}()

	val := ""
	for _, fItem := range wr.funcs {
		local := fItem(ll, params)
		val = val + local
		buf.WriteString(local)
	}

	var err error
	if realWrite {
		_, err = buf.WriteTo(w)
	}

	return val, err
}

func (wr *writer) WriteString(ll LogLevel, params map[string]interface{}) string {
	buf := wr.bufPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		wr.bufPool.Put(buf)
	}()
	for _, fItem := range wr.funcs {
		buf.WriteString(fItem(ll, params))
	}

	return buf.String()
}
