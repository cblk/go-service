package main

import (
	"go_service/library/logy"

	"github.com/pkg/errors"
)

var (
	ERR_CREATE_RESOURCE1_FAILED = errors.New("create resource1 failed")
	ERR_CREATE_RESOURCE2_FAILED = errors.New("create resource2 failed")
	ERR_CREATE_RESOURCE3_FAILED = errors.New("create resource3 failed")
	ERR_CREATE_RESOURCE4_FAILED = errors.New("create resource4 failed")
)

var (
	_mock_check_data string
)

func createResource1() error {
	// do something
	if _mock_check_data == "" {
		logy.Error("create resource1", ERR_CREATE_RESOURCE1_FAILED)
		return ERR_CREATE_RESOURCE1_FAILED
	}

	return nil
}

func createResource2() error {
	// do something
	if _mock_check_data == "" {
		return ERR_CREATE_RESOURCE2_FAILED
	}

	if err := createResource1(); err != nil {
		// logy.Error("create resource1", err)
		return err
	}


	return nil
}


func createResource3() error {
	// do something
	if _mock_check_data == "" {
		return ERR_CREATE_RESOURCE3_FAILED
	}

	if err := createResource1(); err != nil {
		// logy.Error("create resource1", err)
		return err
	}

	if err := createResource2(); err != nil {
		// logy.Error("create resource2", err)
		return err
	}

	return nil
}


func createResource4() error {
	// do something
	if _mock_check_data == "" {
		return ERR_CREATE_RESOURCE4_FAILED
	}

	if err := createResource1(); err != nil {
		// logy.Error("create resource1", err)
		return err
	}

	if err := createResource2(); err != nil {
		// logy.Error("create resource2", err)
		return err
	}

	if err := createResource3(); err != nil {
		// logy.Error("create createResource3", err)
		return err
	}

	return nil
}



func demo() error {
	err := createResource4()
	if err != nil {
		// logy.Error("create createResource4", err)
		return err
	}

	return nil
}

func main() {
	if err := demo(); err != nil {
		logy.Error("demo error", err)
		return
	}

	logy.Info("deferDemo success ", nil)
}
