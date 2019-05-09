package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"errors"
)

func ReadFileErrorf(path string) ([]byte, error) {
	if path == "" {
		return nil, errors.New("param path error")
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open failed: %+v", err)
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read failed: %+v", err)
	}

	return buf, nil
}

func ReadConfigErrorf() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFileErrorf(filepath.Join(home, ".config.xml"))

	return config, fmt.Errorf("could not read config: %+v", err)
}

func main() {
	_, err := ReadConfigErrorf()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
