package main

import (
	"net/http"

	"go_service/internal/catch"
	"go_service/library/logy"
)

func visitFunc() error {
	_, err := http.Get("www.baidu.com")
	if err != nil {
		// logy.Error("http.Get", err)
		return err
	}

	return nil
}

func goEvent() {
	defer func() {
		testTx := func(params ...interface{}) {
			logy.Info("scroll tx", nil)
		}

		catch.Finally(recover(), testTx, "")
	}()

	err := visitFunc()
	if err != nil {
		logy.Error("visitFunc", nil)
		return
	}

	logy.Info("go event", nil)
}

func main() {
	go goEvent()

	select {}
}
