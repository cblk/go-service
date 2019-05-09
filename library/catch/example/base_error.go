package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"errors"
)

func ReadFileError(path string) ([]byte, error) {
	if path == "" {
		return nil, errors.New("param path error")
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func ReadConfigError() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFileError(filepath.Join(home, ".config.yml"))

	return config, err
}

func main() {
	_, err := ReadConfigError()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
