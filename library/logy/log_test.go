package logy

import (
	"testing"

	"go_service/config"

	"github.com/pkg/errors"
)

func TestLogConfig(t *testing.T) {

	if GetAppID() != "" {
		t.Fatal("AppID error")
	}

	if GetHostName() == "" {
		t.Fatal("HostName error")
	}

	if !GetLogStdout() {
		t.Fatal("LogStdout error")
	}

	if GetLogLevel() != "all" {
		t.Fatal("logLevelIn error")
	}

	if GetLogDir() != "" {
		t.Fatal("LogDir error")
	}

	if GetLogMaxFileNum() != 0 {
		t.Fatal("LogMaxFileNum error")
	}

	if GetLogMaxFileSize() != 0 {
		t.Fatal("LogMaxFileSize error")
	}

	if GetLogSplit() {
		t.Fatal("LogSplit error")
	}

	if GetLogSplitBy() != LogSplitByDay {
		t.Fatal("LogSplitBy error")
	}

	_ = Close()
}

func TestLogStdout(t *testing.T) {
	err := config.InitConfig(CONST_Config_Path)
	if err != nil {
		t.Fatal(err.Error())
	}
	cnf := config.GetConfig()
	LoadLogConfig(cnf)

	// can set from config file
	// SetFormat("%L %e %D %T %a %f %M %S")

	Debug("A1", errors.New("a1"))
	Info("A2", errors.New("a2"))
	Notice("A3", errors.New("a3"))
	Warn("A4", errors.New("a4"))
	Error("A5", errors.New("a5"))
	Fatal("A6", errors.New("a6"))

	Debug("B1", nil)
	Info("B2", nil)
	Notice("B3", nil)
	Warn("B4", nil)
	Error("B5", nil)
	Fatal("B6", nil)

	Debug("", errors.New("a1"))
	Info("", errors.New("a2"))
	Notice("", errors.New("a3"))
	Warn("", errors.New("a4"))
	Error("", errors.New("a5"))
	Fatal("", errors.New("a6"))

	_ = Close()
}