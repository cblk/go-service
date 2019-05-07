package logy

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

var BaseMap = map[string]func(map[string]interface{}) string{
	"T": longTime,
	"t": shortTime,
	"D": longDate,
	"d": shortDate,
	"a": keyFormat(_key_appID),
	"L": keyFormat(_key_logLevelName),
	"f": keyFormat(_key_source),
	"e": keyFormat(_key_publishEnv),
	"r": keyFormat(_key_region),
	"i": keyFormat(_key_identifier),
	"S": longStack,
	"s": shortStack,
	"M": message,
}

func longTime(map[string]interface{}) string {
	return time.Now().Format("15:04:05.000")
}

func shortTime(map[string]interface{}) string {
	return time.Now().Format("15:04")
}

func longDate(map[string]interface{}) string {
	return time.Now().Format("2006-01-02")
}

func shortDate(map[string]interface{}) string {
	return time.Now().Format("01-02")
}

func longStack(map[string]interface{}) string {
	if _, file, lineNo, ok := runtime.Caller(6); ok {
		return fmt.Sprintf("stack:%s:%d", file, lineNo)
	}
	return "stack:0:0"
}

func shortStack(map[string]interface{}) string {
	if _, file, lineNo, ok := runtime.Caller(6); ok {
		return fmt.Sprintf("stack:%s:%d", path.Base(file), lineNo)
	}
	return "stack:0:0"
}

func isInternalKey(k string) bool {
	switch k {
	case _key_logLevel, _key_logLevelName, _key_logTime, _key_log,
		_key_publishEnv, _key_appID, _key_region, _key_identifier,
		_key_color, _key_source:
		return true
	}

	return false
}

func message(d map[string]interface{}) string {
	var m string
	var s []string
	for k, v := range d {
		if k == _key_log {
			m = fmt.Sprint(v)
			continue
		}
		if isInternalKey(k) {
			continue
		}

		s = append(s, fmt.Sprintf("%s=%v", k, v))
	}
	s = append(s, m)
	return strings.Join(s, " ")
}

func textFormat(text string) func(map[string]interface{}) string {
	return func(map[string]interface{}) string {
		return text
	}
}

func keyFormat(key string) func(map[string]interface{}) string {
	return func(params map[string]interface{}) string {
		if val, ok := params[key]; ok {
			if s, ok := val.(string); ok {
				return s
			}
			return fmt.Sprint(val)
		}
		return ""
	}
}
