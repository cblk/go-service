package logy

const (
	_key_wrap_count = "wrap_count"
	_key_wrap       = "wrap"

	_timeFormat = "2006-01-02 15:04:05.999999"

	// in log_init.go, 0=all 1=debug 2=info 3=notice 4=warning 5=error 6=critical
	_key_logLevel     = "log_level"      // log level value : 0, 1, 2, 3, 4, 5, 6
	_key_logLevelName = "log_level_name" // log level name : all, debug, info, notice, warning, error, critical
	_key_logTime      = "log_time"
	_key_log          = "log"
	_key_publishEnv   = "publish_env"
	_key_appID        = "app_id"
	_key_region       = "region"
	_key_identifier   = "identifier"
	_key_color        = "color"
	_key_source       = "source"
)
