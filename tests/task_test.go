package tests

import (
	"fmt"
	"go-service/utils"
	"testing"
)

func TestTask(t *testing.T) {
	t.Log("test")
}

func TestTry(t *testing.T) {
	fmt.Println(utils.Try(func() {
		utils.PanicWrap(utils.Try(func() {
			utils.PanicWrap(utils.Try(func() {
				utils.PanicWrap(utils.Error("test error"), "sss%s", "ddd")
				utils.PanicBool(true, "sss%s", "ddd")
			}), "test error try")
		}), "test 12345")
	}))
}

func TestSendTask(t *testing.T) {
}
