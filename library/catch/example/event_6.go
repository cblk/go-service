package main

import "errors"

var (
	ERR_Params         = errors.New("params error")
	ERR_CLOSED_PIPE    = errors.New("io: read/write on closed pipe")
	ERR_NO_PROGRESS    = errors.New("multiple Read calls return no data or error")
	ERR_SHORT_BUFFER   = errors.New("short buffer")
	ERR_SHORT_WRITE    = errors.New("short write")
	ERR_UNEXPECTED_EOF = errors.New("unexpected EOF")
)

var _mock_data string

func WriteData(val string) error {
	if val == "" {
		return ERR_Params
	}

	_mock_data = val
	// do something

	return nil
}

func ReadData() (string, error) {
	// do something

	// mock
	if _mock_data == "" {
		return "", ERR_UNEXPECTED_EOF
	}

	return _mock_data, nil
}