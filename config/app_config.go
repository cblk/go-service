package config

const (
	EnvRelease = "release"
	EnvDebug   = "debug"
	EnvTest    = "test"
)

type AppConfig struct {
	Environment string `mapstructure:"environment" json:"environment"`

	Db struct {
		Driver     string `mapstructure:"driver" json:"driver"`
		Connection string `mapstructure:"connection" json:"connection"`
	} `mapstructure:"db" json:"db"`

	Log struct {
		Level   string `mapstructure:"level" json:"level"`
		GinMode string `mapstructure:"ginMode" json:"ginMode"`
	} `mapstructure:"log" json:"log"`

	Http struct {
		Host string `mapstructure:"host" json:"host"`
		Port string `mapstructure:"port" json:"port"`
	} `mapstructure:"http" json:"http"`

	AllowOrigins []string `mapstructure:"allow_origins" json:"allow_origins"`

	File struct {
		StorageType int    `mapstructure:"storage_type" json:"storage_type"`
		StorageUrl  string `mapstructure:"storage_url" json:"storage_url"`
	} `mapstructure:"file" json:"file"`

	Oss struct {
		Endpoint              string `mapstructure:"endpoint" json:"endpoint"`
		AccessKeyID           string `mapstructure:"access_key_id" json:"access_key_id"`
		AccessKeySecret       string `mapstructure:"access_key_secret" json:"access_key_secret"`
		Bucket                string `mapstructure:"bucket" json:"bucket"`
		PathPrefix            string `mapstructure:"path_prefix" json:"path_prefix"`
		CallBackUrl           string `mapstructure:"callback_url" json:"callback_url"`
		UploadFileCallbackUrl string `mapstructure:"upload_file_callback_url" json:"upload_file_callback_url"`
	} `mapstructure:"oss" json:"oss"`

	Session struct {
		Driver             string `mapstructure:"driver" json:"driver"`
		Connection         string `mapstructure:"connection" json:"connection"`
		SecretKey          string `mapstructure:"secret_key" json:"secret_key"`
		LoginMaxAgeSeconds int    `mapstructure:"login_max_age_seconds" json:"login_max_age_seconds"`
		VaultMaxAgeSeconds int    `mapstructure:"vault_max_age_seconds" json:"vault_max_age_seconds"`
	} `mapstructure:"session" json:"session"`
}
