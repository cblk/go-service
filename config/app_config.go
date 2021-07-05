package config

const (
	EnvProduction = "release"
	EnvDebug      = "debug"
	EnvTest       = "test"
)

type AppConfig struct {
	Environment string `mapstructure:"environment"`

	Db struct {
		Driver           string `mapstructure:"driver"`
		ConnectionString string `mapstructure:"connection"`
	} `mapstructure:"db"`

	Log struct {
		Level   string `mapstructure:"level"`
		GinMode string `mapstructure:"ginMode"`
	} `mapstructure:"log"`

	Http struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"http"`

	AllowOrigins []string `mapstructure:"allow_origins"`

	File struct {
		StorageType int    `mapstructure:"storage_type"`
		StorageUrl  string `mapstructure:"storage_url"`
	} `mapstructure:"file"`

	Oss struct {
		Endpoint              string `mapstructure:"endpoint"`
		AccessKeyID           string `mapstructure:"access_key_id"`
		AccessKeySecret       string `mapstructure:"access_key_secret"`
		Bucket                string `mapstructure:"bucket"`
		PathPrefix            string `mapstructure:"path_prefix"`
		CallBackUrl           string `mapstructure:"callback_url"`
		UploadFileCallbackUrl string `mapstructure:"upload_file_callback_url"`
	} `mapstructure:"oss"`
}
