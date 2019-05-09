package logy

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

var BaseMap = map[string]func(LogLevel, map[string]interface{}) string{
	"T": longTime,
	"t": shortTime,
	"D": longDate,
	"d": shortDate,
	"a": keyFormat(LogLevelAll, _key_appID),
	"L": keyFormat(LogLevelAll, _key_logLevelName),
	"f": keyFormat(LogLevelAll, _key_source),
	"e": keyFormat(LogLevelAll, _key_publishEnv),
	"r": keyFormat(LogLevelAll, _key_region),
	"i": keyFormat(LogLevelAll, _key_identifier),
	"S": longStack,
	"s": shortStack,
	"M": message,
}

func longTime(ll LogLevel, ds map[string]interface{}) string {
	return time.Now().Format("15:04:05.000")
}

func shortTime(ll LogLevel, ds map[string]interface{}) string {
	return time.Now().Format("15:04")
}

func longDate(ll LogLevel, ds map[string]interface{}) string {
	return time.Now().Format("2006-01-02")
}

func shortDate(ll LogLevel, ds map[string]interface{}) string {
	return time.Now().Format("01-02")
}

func longStack(ll LogLevel, ds map[string]interface{}) string {
	if ll <= LogLevelWarning {
		return ""
	}

	if _, file, lineNo, ok := runtime.Caller(6); ok {
		val := fmt.Sprintf("fixed:%s:%d\n", file, lineNo)

		val = val + "stack:\n"
		buf := make([]byte, 4096)
		n := runtime.Stack(buf, true)
		if n < len(buf) {
			val = val + string(buf[:n])
		} else {
			val = val + string(buf)
		}

		return "\n" + val
	}
	return "\n" + "stack:0:0"
}

func shortStack(ll LogLevel, ds map[string]interface{}) string {
	if ll <= LogLevelWarning {
		return ""
	}

	if _, file, lineNo, ok := runtime.Caller(6); ok {
		val := fmt.Sprintf("fixed:%s:%d\n", path.Base(file), lineNo)

		val = val + "stack:\n"
		buf := make([]byte, 4096)
		n := runtime.Stack(buf, true)
		if n < len(buf) {
			val = val + string(buf[:n])
		} else {
			val = val + string(buf)
		}

		return "\n" + val
	}
	return "\n" + "stack:0:0"
}

func isInternalKey(ll LogLevel, k string) bool {
	switch k {
	case _key_logLevel, _key_logLevelName, _key_logTime, _key_log,
		_key_publishEnv, _key_appID, _key_region, _key_identifier,
		_key_color, _key_source:
		return true
	}

	return false
}

func message(ll LogLevel, ds map[string]interface{}) string {
	var m string
	var s []string
	for k, v := range ds {
		if k == _key_log {
			m = fmt.Sprint(v)
			continue
		}
		if isInternalKey(ll, k) {
			continue
		}

		if v != "" {
			s = append(s, fmt.Sprintf("%s=%v", k, v))
		} else {
			s = append(s, fmt.Sprintf("%s", k))
		}
	}
	s = append(s, m)
	return strings.Join(s, " ")
}

func textFormat(ll LogLevel, text string) func(LogLevel, map[string]interface{}) string {
	return func(LogLevel, map[string]interface{}) string {
		return text
	}
}

func keyFormat(ll LogLevel, key string) func(LogLevel, map[string]interface{}) string {
	return func(ll LogLevel, params map[string]interface{}) string {
		if val, ok := params[key]; ok {
			if s, ok := val.(string); ok {
				return s
			}
			return fmt.Sprint(val)
		}
		return ""
	}
}
