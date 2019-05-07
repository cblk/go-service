package logy

type logConfig struct {
	AppID      string   `yaml:"app_id"`      // global unique application id, etc: ****.****.*******
	HostName   string   `yaml:"host_name"`   // host name
	Region     string   `yaml:"region"`      // application Region, etc: sh
	Identifier string   `yaml:"identifier"`  // application identifier etc: sh001
	PublishEnv string   `yaml:"publish_env"` // publish environment
	LogStdout  bool     `yaml:"log_stdout"`  //
	LogLevel   int      `yaml:"log_level"`   // 1=debug 2=info 3=notice 4=warning 5=error 6=critical
	LogFilter  []string `yaml:"log_filter"`  // Filter tell log handler which Field are sensitive message, use * instead.

	// The following CI environments are not required.
	LogDir         string `yaml:"log_dir"`           // todo
	LogMaxFileNum  int    `yaml:"log_max_file_num"`  // todo
	LogMaxFileSize int    `yaml:"log_max_file_size"` // todo
	LogSplit       bool   `yaml:"log_split"`         // todo
	LogSplitBy     string `yaml:"log_split_by"`      // todo
}

func GetAppID() string {
	return getLogConfig().AppID
}

func SetAppID(appID string) {
	getLogConfig().AppID = appID
}

func GetHostName() string {
	return getLogConfig().HostName
}

func GetRegion() string {
	return getLogConfig().Region
}

func SetRegion(region string) {
	getLogConfig().Region = region
}

func GetIdentifier() string {
	return getLogConfig().Identifier
}

func SetIdentifier(identifier string) {
	getLogConfig().Identifier = identifier
}

func GetPublishEnv() string {
	return getLogConfig().PublishEnv
}

func SetPublishEnv(publishEnv string) {
	getLogConfig().PublishEnv = publishEnv
}

func GetLogStdout() bool {
	return getLogConfig().LogStdout
}

func SetLogStdout(logStdout bool) {
	getLogConfig().LogStdout = logStdout
}

func GetLogLevel() int {
	return getLogConfig().LogLevel
}

func SetLogLevel(logLevel int) {
	getLogConfig().LogLevel = logLevel
}

func GetLogFilter() []string {
	return getLogConfig().LogFilter
}

func SetLogFilter(logFilter []string) {
	getLogConfig().LogFilter = logFilter
}

func GetLogDir() string {
	return getLogConfig().LogDir
}

func SetLogDir(logDir string) {
	getLogConfig().LogDir = logDir
}

func GetLogMaxFileNum() int {
	return getLogConfig().LogMaxFileNum
}

func SetLogMaxFileNum(logMaxFileNum int) {
	getLogConfig().LogMaxFileNum = logMaxFileNum
}

func GetLogMaxFileSize() int {
	return getLogConfig().LogMaxFileSize
}

func SetLogMaxFileSize(logMaxFileSize int) {
	getLogConfig().LogMaxFileSize = logMaxFileSize
}

func GetLogSplit() bool {
	return getLogConfig().LogSplit
}

func SetLogSplit(logSplit bool) {
	getLogConfig().LogSplit = logSplit
}

func GetLogSplitBy() string {
	return getLogConfig().LogSplitBy
}

func SetLogSplitBy(logSplitBy string) {
	getLogConfig().LogSplitBy = logSplitBy
}
