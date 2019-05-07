package logy

import (
	"context"
	"testing"

	"go_service/config"

	"github.com/pkg/errors"
)

const (
	CONST_Config_Path = "../../config"
)

/*
test empty log config
*/
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

	if GetLogLevel() != int(LogLevelAll) {
		t.Fatal("LogLevel error")
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

func TestLoadLogConfig(t *testing.T) {

	err := config.InitConfig(CONST_Config_Path)
	if err != nil {
		t.Fatal(err.Error())
	}
	cnf := config.GetConfig()
	LoadLogConfig(cnf)

	if GetAppID() != "test.env.config.0111" {
		t.Fatal("AppID error")
	}

	if GetHostName() == "" {
		t.Fatal("HostName error")
	}

	if !GetLogStdout() {
		t.Fatal("LogStdout error")
	}

	if GetLogLevel() != int(LogLevelInfo) {
		t.Fatal("LogLevel error")
	}

	if GetLogDir() == "" {
		t.Fatal("LogDir error")
	}

	if GetLogMaxFileNum() != 100 {
		t.Fatal("LogMaxFileNum error")
	}

	if GetLogMaxFileSize() != 1000000000 {
		t.Fatal("LogMaxFileSize error")
	}

	if !GetLogSplit() {
		t.Fatal("LogSplit error")
	}

	if GetLogSplitBy() != LogSplitByWeek {
		t.Fatal("LogSplitBy error")
	}

	if len(GetLogFilter()) != 2 {
		t.Fatal("LogFilter error")
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

	SetFormat("%L %e %D-%T %a %f %S %M")

	Debug("Value %v %v", errors.New("a1"), errors.New("b1"))
	Info("Value %v", errors.New("a2"))
	Error("Value %v", errors.New("a3"))
	Notice("Value %v", errors.New("a4"))
	Warn("Value %v", errors.New("a5"))
	Fatal("Value %v", errors.New("a6"))

	_ = Close()
}

func TestLogStdoutByField(t *testing.T) {
	err := config.InitConfig(CONST_Config_Path)
	if err != nil {
		t.Fatal(err.Error())
	}
	cnf := config.GetConfig()
	LoadLogConfig(cnf)

	SetFormat("%L %e %D-%T %a %f %S %M")

	DebugF(context.Background(), Field{Key: "A1", Value: "B1"}, Field{Key: "A2", Value: "B2"})
	InfoF(context.Background(), Field{Key: "A2", Value: "B2"})
	ErrorF(context.Background(), Field{Key: "A3", Value: "B3"})
	NoticeF(context.Background(), Field{Key: "A4", Value: "B4"})
	WarnF(context.Background(), Field{Key: "A5", Value: "B5"})
	FatalF(context.Background(), Field{Key: "A6", Value: "B6"})

	_ = Close()
}

func TestLogStdoutByString(t *testing.T) {
	err := config.InitConfig(CONST_Config_Path)
	if err != nil {
		t.Fatal(err.Error())
	}
	cnf := config.GetConfig()
	LoadLogConfig(cnf)

	SetFormat("%L %e %D-%T %a %f %S %M")

	DebugS(context.Background(), "A1", "B1", "A2", "B2")
	InfoS(context.Background(), "A2", "B2")
	ErrorS(context.Background(), "A3", "B3")
	NoticeS(context.Background(), "A4", "B1")
	WarnS(context.Background(), "A5", "B5")
	FatalS(context.Background(), "A6", "B6")

	_ = Close()
}

func TestLogStdoutByWrap(t *testing.T) {
	err := config.InitConfig(CONST_Config_Path)
	if err != nil {
		t.Fatal(err.Error())
	}
	cnf := config.GetConfig()
	LoadLogConfig(cnf)

	// SetFormat("%L %e %D %T %a %f %S %M")
	SetFormat("%L %e %D %T %a %M")

	debugWrap := DebugW("A1", errors.New("B1")).
		DebugW("A11", errors.New("B11")).
		DebugW("A12", errors.New("B12"))
	debugWrap.Error()

	InfoW("A2", errors.New("B2")).Error()
	ErrorW("A3", errors.New("B3")).Error()
	NoticeW("A4", errors.New("B4")).Error()
	WarnW("A5", errors.New("B5")).Error()
	FatalW("A6", errors.New("B6")).Error()

	_ = Close()
}

func testNilFunc() IError {
	return Nil()
}

func TestLogStdoutByErrorC(t *testing.T) {
	err := config.InitConfig(CONST_Config_Path)
	if err != nil {
		t.Fatal(err.Error())
	}
	cnf := config.GetConfig()
	LoadLogConfig(cnf)

	// SetFormat("%L %e %D %T %a %f %S %M")
	SetFormat("%L %e %D %T %a %M")
	DebugC(context.Background(), "A1", errors.New("B1")).ErrorW("A2", errors.New("B2")).Error()

	// import
	if !IsNil(testNilFunc()) {
		t.Fatal("nil func return error")
	}

	// nil test, import
	nilErr := testNilFunc()
	if nilErr.GetContent() != context.Background() {
		t.Fatal("nilErr GetContent error")
	}
	if nilErr.GetError() != nil {
		t.Fatal("nilErr GetError error")
	}
	if nilErr.GetLevel() != LogLevelAll {
		t.Fatal("nilErr GetLevel error")
	}
	if nilErr.GetMessage() != "" {
		t.Fatal("nilErr GetMessage error")
	}
	if nilErr.GetPre() != nil {
		t.Fatal("nilErr GetPre error")
	}
	nilErr.Error()


	dErr := testNilFunc().ErrorW("D2", errors.New("D2"))
	if dErr.GetContent() != context.Background() {
		t.Fatal("dErr GetContent error")
	}
	if dErr.GetError() == nil {
		t.Fatal("dErr GetError error")
	}
	if dErr.GetLevel() != LogLevelError {
		t.Fatal("dErr GetLevel error")
	}
	if dErr.GetMessage() != "D2" {
		t.Fatal("dErr GetMessage error")
	}
	if !IsNil(dErr.GetPre())   {
		t.Fatal("dErr GetPre error")
	}
	if !dErr.GetPre().IsEqual(Nil())   {
		t.Fatal("dErr GetPre error")
	}
	dErr.Error()

	if dErr.IsEqual(nilErr) {
		t.Fatal("IsEqual error")
	}

	if !dErr.IsEqual(dErr) {
		t.Fatal("IsEqual error")
	}
}
