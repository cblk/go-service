package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"
	"runtime"
)

type M struct {
	Msg        string `json:"msg,omitempty"`
	FuncCaller string `json:"funcCaller,omitempty"`
	Sub        error  `json:"sub,omitempty"`
}

func (t *M) Error() string {
	_dt, _ := json.Marshal(t)
	return string(_dt)
}

func Error(msg string, args ...interface{}) *M {
	return &M{
		FuncCaller: funcCaller(),
		Msg:        fmt.Sprintf(msg, args...),
	}
}

func Log(err interface{}) {
	P(err)
	fmt.Println(reflect.TypeOf(err).String(), funcCaller())
	fmt.Println("******************************")
}

func P(d ...interface{}) {
	for _, i := range d {
		dt, err := json.MarshalIndent(i, "", "\t")
		PanicErr(err)
		fmt.Println(reflect.ValueOf(i).String(), "->", string(dt))
	}
}

func PanicErr(err error) {
	if err == nil {
		return
	}

	var m = &M{}
	switch e := err.(type) {
	case *M:
		m = e
	case error:
		m.Msg = e.Error()
	}

	log.Println(funcCaller())
	panic(&M{
		Sub:        m,
		FuncCaller: funcCaller(),
	})
}

func funcCaller() string {
	_, file, line, _ := runtime.Caller(2)
	_f := fmt.Sprintf("%s:%d ", file, line)
	return _f
}

func PanicBool(b bool, msg string, args ...interface{}) {
	if !b {
		return
	}

	log.Println(funcCaller())
	panic(&M{
		FuncCaller: funcCaller(),
		Msg:        fmt.Sprintf(msg, args...),
		Sub:        Error(msg, args...),
	})
}

func PanicWrap(err error, msg string, args ...interface{}) {
	if err == nil {
		return
	}

	var m = &M{}
	switch e := err.(type) {
	case *M:
		m = e
	case error:
		m.Msg = e.Error()
	}

	log.Println(funcCaller())
	panic(&M{
		FuncCaller: funcCaller(),
		Msg:        fmt.Sprintf(msg, args...),
		Sub:        m,
	})
}

func IpAddress() string {
	addrs, err := net.InterfaceAddrs()
	PanicErr(err)

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func Try(fn func()) (err error) {
	PanicBool(fn == nil, "the func is nil")

	_v := reflect.TypeOf(fn)
	PanicBool(_v.Kind() != reflect.Func, "the params type(%s) is not func", _v.String())

	defer func() {
		defer func() {
			m := &M{}
			if r := recover(); r != nil {
				switch d := r.(type) {
				case *M:
					m = d
				case error:
					m.Sub = d
				case string:
					m.Sub = errors.New(d)
				}
			}

			if m.Sub == nil {
				err = nil
			} else {
				err = m
			}
		}()
		reflect.ValueOf(fn).Call([]reflect.Value{})
	}()
	return
}

func Fn(f interface{}, params ...interface{}) func() error {
	t := reflect.TypeOf(f)
	PanicBool(t.Kind() != reflect.Func, "the params is not func type")

	return func() error {
		return Try(func() {
			var vs []reflect.Value
			for i, p := range params {
				if p == nil {
					vs = append(vs, reflect.New(t.In(i)).Elem())
				} else {
					vs = append(vs, reflect.ValueOf(p))
				}
			}
			reflect.ValueOf(f).Call(vs)
		})
	}
}
