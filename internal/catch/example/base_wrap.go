package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func ReadFileWrap(path string) ([]byte, error) {
	if path == "" {
		return nil, errors.New("param path error")
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}

	return buf, nil
}

func ReadConfigWrap() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFileWrap(filepath.Join(home, ".config.xml"))

	return config, errors.Wrap(err, "could not read config")
}

func main() {
	_, err := ReadConfigWrap()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
