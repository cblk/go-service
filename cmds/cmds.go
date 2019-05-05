package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"portal/internal/cnst"
	"portal/utils"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	HomeFlag  = "home"
	TraceFlag = "trace"
	LogFlag   = "log_level"
)

func PrepareBaseCmd(cmd *cobra.Command, envPrefix, defaultHome string) Executor {
	cobra.OnInitialize(func() { initEnv(envPrefix) })

	cmd.PersistentFlags().StringP(HomeFlag, "", defaultHome, "directory for config and data")
	cmd.PersistentFlags().Bool(TraceFlag, false, "print out full stack trace on errors.toml")
	cmd.PersistentFlags().StringP(LogFlag, "l", "debug", "log level")
	cmd.PersistentPreRunE = concatCobraCmdFuncs(bindFlagsLoadViper, cmd.PersistentPreRunE)
	return Executor{Command: cmd, Exit: os.Exit}
}

// initEnv sets to use ENV variables if set.
func initEnv(prefix string) {
	copyEnvVars(prefix)

	// env variables with TM prefix (eg. TM_ROOT)
	viper.SetEnvPrefix(prefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
}

// This copies all variables like TMROOT to TM_ROOT,
// so we can support both formats for the user
func copyEnvVars(prefix string) {
	prefix = strings.ToUpper(prefix)
	ps := prefix + "_"
	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "=", 2)
		if len(kv) == 2 {
			k, v := kv[0], kv[1]
			if strings.HasPrefix(k, prefix) && !strings.HasPrefix(k, ps) {
				k2 := strings.Replace(k, prefix, ps, 1)
				utils.PanicErr(os.Setenv(k2, v))
			}
		}
	}
}

// Executor wraps the cobra Command with a nicer Execute method
type Executor struct {
	*cobra.Command
	Exit func(int) // this is os.Exit by default, override in tests
}

type ExitCoder interface {
	ExitCode() int
}

// execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func (e Executor) Execute() error {
	e.SilenceUsage = true
	e.SilenceErrors = true

	return e.Command.Execute()
}

type cobraCmdFunc func(cmd *cobra.Command, args []string) error

// Returns a single function that calls each argument function in sequence
// RunE, PreRunE, PersistentPreRunE, etc. all have this same signature
func concatCobraCmdFuncs(fs ...cobraCmdFunc) cobraCmdFunc {
	return func(cmd *cobra.Command, args []string) error {
		for _, f := range fs {
			if f != nil {
				utils.PanicErr(f(cmd, args))
			}
		}
		return nil
	}
}

// Bind all flags and read the config into viper
func bindFlagsLoadViper(cmd *cobra.Command, args []string) error {
	return utils.Try(func() {
		// cmd.Flags() includes flags from this command and all persistent flags from the parent
		utils.PanicErr(viper.BindPFlags(cmd.Flags()))

		homeDir := viper.GetString(HomeFlag)

		dl := strings.Split(os.Getenv("DRONE_TAG"), ".")
		_env1 := ""
		if len(dl) != 0 {
			_env1 = dl[len(dl)-1]
		}

		_env2 := os.Getenv("env")
		utils.PanicBool(_env1 == "" && _env2 == "", "env(%s)(%s) is null", _env1, _env2)

		env := _env1
		if env == "" {
			env = _env2
		}

		utils.PanicBool(!(env != cnst.Env.Dev || env != cnst.Env.Stag || env != cnst.Env.Prod), "env error(%s)", env)

		viper.SetConfigType("yml")
		viper.Set(HomeFlag, homeDir)
		viper.SetConfigName(fmt.Sprintf("config.%s", env))

		viper.AddConfigPath("/etc/kdata")
		viper.AddConfigPath("$HOME/.kdata")
		viper.AddConfigPath(".")
		viper.AddConfigPath(homeDir)                          // search root directory
		viper.AddConfigPath(filepath.Join(homeDir, "config")) // search root directory /config

		// load config
		utils.PanicErr(viper.ReadInConfig())
	})
}
